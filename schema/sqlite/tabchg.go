package sqlite

import (
	"fmt"
	"log"

	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

func (m *meta) ChangeTable(db common.DB, change *schema.TableSchemaChange) error {
	tabName := change.NewName
	//处理表更名,处理过后，所有后续操作都在新表名上进行
	if change.OldName != change.NewName {
		if err := tableRename(db, change.OldName, change.NewName); err != nil {
			return err
		}
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
			if err := addColumn(db, tabName, one.NewField); err != nil {
				return err
			}
		}
		return nil
	}
	//先创建新表
	tmpTableName, err := getTempTableName(db, "tu__")
	if err != nil {
		log.Println(err)
		return err
	}
	tab := schema.NewTable(tmpTableName)
	cols := []*schema.Column{}
	selFields := []string{} //select 的列名
outLoop:
	for _, one := range change.OriginFields {
		//如果是删除，忽略这列
		for _, rmFieldName := range change.RemoveFields {
			if rmFieldName == one.Name {
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
		cols = append(cols, col)
	}
	tab.Columns = cols
	tab.PrimaryKeys = change.PK
	intoFields := []string{}
	for _, one := range tab.Columns {
		intoFields = append(intoFields, one.Name)
	}
	if err := m.CreateTable(db, tab); err != nil {
		return err
	}
	//再复制数据

	strSQL := fmt.Sprintf("insert into %s(%s)select %s from %s",
		change.NewName, strings.Join(intoFields, ","),
		strings.Join(selFields, ","), tabName)
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	//然后drop 旧表
	strSQL = "drop table " + tabName
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	//最后改回名称
	return tableRename(db, tmpTableName, tabName)
}
func addColumn(db common.DB, tabName string, newCol *schema.Column) error {
	strSQL := fmt.Sprintf("alter table %s add %s", tabName, dbDefine(newCol))
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	log.Println(strSQL)
	//处理索引
	if newCol.Index {
		if err := createColumnIndex(db, tabName, newCol.Name); err != nil {
			return err
		}
		log.Printf("table %s add column index %s\n", tabName, newCol.Name)
	}
	return nil

}
