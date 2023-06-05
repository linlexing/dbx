package postgres

import (
	"fmt"
	"log"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/ddb"
	"github.com/linlexing/dbx/schema"
)

func (m *meta) changeTableSQLGauss(db common.DB, change *schema.TableSchemaChange) (rev []string, err error) {
	tabName := change.NewName

	//先创建新表
	tmpTableName, err := ddb.GetTempTableName("tu__")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	tab := schema.NewTable(tmpTableName)
	cols := []*schema.Column{}
	selFields := []string{} //select 的列名
	intoFields := []string{}
outLoop:
	for _, one := range change.OriginFields {
		//如果是删除，忽略这列
		for _, rmFieldName := range change.RemoveFields {
			if strings.EqualFold(rmFieldName, one.Name) {
				continue outLoop
			}
		}
		col := one
		selFields = append(selFields, one.Name)
		//找出修改后的新列，如有
		for _, upField := range change.ChangeFields {
			if upField.OldField == one {
				col = upField.NewField
				break
			}
		}
		intoFields = append(intoFields, col.Name)
		cols = append(cols, col)
	}
	//加上新增的列
	for _, insField := range change.ChangeFields {
		if insField.OldField == nil {
			cols = append(cols, insField.NewField)
		}
	}
	tab.Columns = cols
	tab.PrimaryKeys = change.PK

	list, err := m.CreateTableSQL(db, tab)
	if err != nil {
		return
	}
	rev = append(rev, list...)
	//如果有intoFields，则再复制数据
	if len(intoFields) > 0 {
		rev = append(rev, fmt.Sprintf("insert into %s(%s)select %s from %s",
			tmpTableName, strings.Join(intoFields, ","),
			strings.Join(selFields, ","), tabName))
	}
	//然后drop 旧表
	rev = append(rev, "drop table "+tabName)
	rev = append(rev, tableRenameSQL(tmpTableName, tabName)...)
	return

}
func (m *meta) ChangeTableSQL(db common.DB, change *schema.TableSchemaChange) (rev []string, err error) {
	tabName := change.NewName
	//处理表更名,处理过后，所有后续操作都在新表名上进行
	if !strings.EqualFold(change.OldName, change.NewName) {
		rev = append(rev, tableRenameSQL(change.OldName, change.NewName)...)
	}
	//如果是华为高斯，不能直接修改主键，则需要用临时表来迁移数据
	if m.isGauss() {
		needCopy := change.PKChange
		if !needCopy {
			//判断主键类型有没有变化
		out:
			for _, one := range change.ChangeFields {
				if one.OldField != nil {
					for _, pk := range change.PK {
						if strings.EqualFold(pk, one.OldField.Name) {
							needCopy = true
							break out
						}
					}
				}
			}
		}
		if needCopy {
			list, er := m.changeTableSQLGauss(db, change)
			if er != nil {
				err = er
				return
			}
			//需要加上改名的语句
			rev = append(rev, list...)
			return
		}
	}

	//如果主键变更，则需要先除去主键
	if change.PKChange {
		list, err := dropTablePrimaryKeySQL(db, tabName)
		if err != nil {
			return nil, err
		}
		rev = append(rev, list...)
	}
	//逐个处理字段
	for _, cf := range change.ChangeFields {

		rev = append(rev, processColumnSQL(tabName, cf.OldField, cf.NewField)...)
	}
	if len(change.RemoveFields) > 0 {
		rev = append(rev, removeColumnsSQL(tabName, change.RemoveFields)...)
	}
	//如果主键变过，则新增主键
	if change.PKChange {
		rev = append(rev, addTablePrimaryKeySQL(tabName, change.PK)...)
	}
	return
}
func processColumnSQL(tabName string, oldCol, newCol *schema.Column) (rev []string) {
	//如果是新增字段
	if oldCol == nil {
		rev = append(rev, fmt.Sprintf("ALTER TABLE %s ADD %s", tabName, dbDefine(newCol)))
		//处理索引
		if newCol.Index == schema.Index {
			rev = append(rev, createColumnIndexSQL(tabName, false, newCol.Name)...)
		} else if newCol.Index == schema.UniqueIndex {
			rev = append(rev, createColumnIndexSQL(tabName, true, newCol.Name)...)
		}
		return
	}
	//如果是更名，需要先处理
	if strings.ToUpper(oldCol.Name) != strings.ToUpper(newCol.Name) {
		rev = append(rev, fmt.Sprintf("ALTER TABLE %s RENAME %s to %s", tabName,
			oldCol.Name, strings.ToLower(newCol.Name)))
	}
	if !oldCol.EqueNoIndexAndName(newCol) {

		//去掉最后的notnull
		//去掉定义中的字段名，因为中间多了个type字样
		if !oldCol.EqueType(newCol) {
			//如果类型变过，现在类型是数值或者日期，则需要加using子句
			if newCol.Type == schema.TypeFloat ||
				newCol.Type == schema.TypeInt ||
				newCol.Type == schema.TypeDatetime {

				if oldCol.Type == schema.TypeString {
					rev = append(rev, fmt.Sprintf(
						"ALTER TABLE %s ALTER COLUMN %s TYPE %s using(trim(%[2]s)::%[3]s)",
						tabName, newCol.Name, colDBType(newCol)))
				} else {
					rev = append(rev, fmt.Sprintf(
						"ALTER TABLE %s ALTER COLUMN %s TYPE %s using(%[2]s::%[3]s)",
						tabName, newCol.Name, colDBType(newCol)))
				}
			} else {
				rev = append(rev, fmt.Sprintf(
					"ALTER TABLE %s ALTER COLUMN %s TYPE %s",
					tabName, newCol.Name, colDBType(newCol)))
			}
		}
		//再改not null
		if oldCol.Null && !newCol.Null {
			rev = append(rev, fmt.Sprintf(
				"ALTER TABLE %s ALTER COLUMN %s SET NOT NULL",
				tabName, newCol.Name))
		}
		if !oldCol.Null && newCol.Null {
			rev = append(rev, fmt.Sprintf(
				"ALTER TABLE %s ALTER COLUMN %s DROP NOT NULL",
				tabName, newCol.Name))
		}
	}
	//处理索引,字段更名的操作，oracle、postgres、mysql都是安全的，所以不需处理
	//ref:http://stackoverflow.com/questions/6732896/does-rename-column-take-care-of-indexes
	if (oldCol.Index == schema.Index || oldCol.Index == schema.UniqueIndex) &&
		newCol.Index == schema.NoIndex {
		//删除索引
		rev = append(rev, dropColumnIndexSQL(tabName, oldCol.IndexName)...)
	} else if oldCol.Index == schema.NoIndex && (newCol.Index == schema.Index || newCol.Index == schema.UniqueIndex) {
		//新增索引
		if newCol.Index == schema.Index {
			rev = append(rev, createColumnIndexSQL(tabName, false, newCol.Name)...)
		} else {
			rev = append(rev, createColumnIndexSQL(tabName, true, newCol.Name)...)
		}
	}
	return
}
