package sqlite

import (
	"fmt"
	"log"

	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/ddb"
	"github.com/linlexing/dbx/schema"
)

func (m *meta) ChangeTableSQL(db common.DB, change *schema.TableSchemaChange) (rev []string, err error) {
	tabName := change.NewName
	//处理表更名,处理过后，所有后续操作都在新表名上进行
	if !strings.EqualFold(change.OldName, change.NewName) {
		rev = append(rev, tableRenameSQL(change.OldName, change.NewName)...)
	}
	//needCopy 指明需用复制表的手段
	needCopy := change.PKChange || len(change.RemoveFields) > 0
	if !needCopy {
		for _, one := range change.ChangeFields {
			if one.OldField != nil {
				needCopy = true
				break
			}
		}
	}
	if !needCopy {
		//到这，说明只有字段新增
		for _, one := range change.ChangeFields {
			rev = append(rev, addColumnSQL(tabName, one.NewField)...)
		}
		return
	}
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
func addColumnSQL(tabName string, newCol *schema.Column) []string {
	rev := []string{fmt.Sprintf("alter table %s add %s", tabName, dbDefine(newCol))}
	//处理索引
	if newCol.Index == schema.Index {
		rev = append(rev, createColumnIndexSQL(tabName, false, newCol.Name)...)
	} else if newCol.Index == schema.UniqueIndex {
		rev = append(rev, createColumnIndexSQL(tabName, true, newCol.Name)...)
	}
	return rev

}
