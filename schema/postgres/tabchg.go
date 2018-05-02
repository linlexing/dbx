package postgres

import (
	"fmt"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

func (m *meta) ChangeTableSQL(db common.DB, change *schema.TableSchemaChange) (rev []string, err error) {
	tabName := change.NewName
	//处理表更名,处理过后，所有后续操作都在新表名上进行
	if change.OldName != change.NewName {
		rev = append(rev, tableRenameSQL(change.OldName, change.NewName)...)
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
		rev = append(rev, fmt.Sprintf("alter table %s add %s", tabName, dbDefine(newCol)))
		//处理索引
		if newCol.Index {
			rev = append(rev, createColumnIndexSQL(tabName, newCol.Name)...)
		}
		return
	}
	//如果是更名，需要先处理
	if oldCol.Name != newCol.Name {
		rev = append(rev, fmt.Sprintf("alter table %s rename %s to %s", tabName, oldCol.Name, newCol.Name))
	}
	if !oldCol.Eque(newCol) {

		//去掉最后的notnull
		//去掉定义中的字段名，因为中间多了个type字样
		rev = append(rev, fmt.Sprintf(
			"alter table %s ALTER COLUMN %s type %s",
			tabName, newCol.Name, colDBType(newCol)))

		//再改not null
		if oldCol.Null && !newCol.Null {
			rev = append(rev, fmt.Sprintf(
				"alter table %s alter column %s set not null",
				tabName, newCol.Name))
		}
		if !oldCol.Null && newCol.Null {
			rev = append(rev, fmt.Sprintf(
				"alter table %s alter column %s drop not null",
				tabName, newCol.Name))
		}
	}
	//处理索引,字段更名的操作，oracle、postgres、mysql都是安全的，所以不需处理
	//ref:http://stackoverflow.com/questions/6732896/does-rename-column-take-care-of-indexes
	if oldCol.Index && !newCol.Index {
		//删除索引
		rev = append(rev, dropColumnIndexSQL(tabName, oldCol.IndexName)...)
	} else if !oldCol.Index && newCol.Index {
		//新增索引
		rev = append(rev, createColumnIndexSQL(tabName, oldCol.Name)...)
	}
	return
}
