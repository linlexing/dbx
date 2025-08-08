package duckdb

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

func (m *meta) ChangeTableSQL(db common.DB, change *schema.TableSchemaChange) ([]string, error) {
	rev := []string{}
	tabName := change.NewName
	//处理表更名,处理过后，所有后续操作都在新表名上进行
	if !strings.EqualFold(change.OldName, change.NewName) {
		rev = append(rev, tableRenameSQL(change.OldName, change.NewName)...)
	}

	//如果主键变更，则需要先除去主键
	if change.PKChange {
		rev = append(rev, dropTablePrimaryKeySQL(tabName)...)
	}
	//逐个处理字段，每处理一个字段，旧表字段就标上标记，最后删除没有标记的字段
	for _, cf := range change.ChangeFields {
		rev = append(rev, processColumnSQL(tabName, cf.OldField, cf.NewField, change.PK)...)
	}
	if len(change.RemoveFields) > 0 {
		rev = append(rev, removeColumnsSQL(tabName, change.RemoveFields)...)
	}
	//如果主键变过，则新增主键
	if change.PKChange {
		rev = append(rev, addTablePrimaryKeySQL(tabName, change.PK)...)
	}
	return rev, nil
}

func processColumnSQL(tabName string, oldCol, newCol *schema.Column, newPKS []string) []string {
	pksMap := make(map[string]any, len(newPKS))
	for i := range newPKS {
		pksMap[newPKS[i]] = nil
	}
	rev := []string{}
	//如果是新增字段
	if oldCol == nil {
		rev := append(rev, fmt.Sprintf("alter table \"%s\" add %s", tabName, dbDefine(newCol)))
		//处理索引
		if newCol.Index == schema.Index {
			rev = append(rev, createColumnIndexSQL(tabName, false, newCol.Name)...)
		} else if newCol.Index == schema.UniqueIndex {
			rev = append(rev, createColumnIndexSQL(tabName, true, newCol.Name)...)
		}
		return nil
	}
	//如果是更名，需要先处理
	if strings.ToUpper(oldCol.Name) != strings.ToUpper(newCol.Name) {
		rev = append(rev, fmt.Sprintf("alter table \"%s\" rename \"%s\" to \"%s\"", tabName, oldCol.Name, newCol.Name))
	}
	//改类型和设置非空一条语句无法完成
	if oldCol.Type != newCol.Type {
		rev = append(rev, fmt.Sprintf("alter table \"%s\" alter \"%s\" set type %s", tabName, newCol.Name, colDBType(newCol)))
	}
	if oldCol.Null && !newCol.Null {
		rev = append(rev, fmt.Sprintf("alter table \"%s\" alter column \"%s\" set not null", tabName, newCol.Name))
	}
	if newCol.Null && !oldCol.Null {
		rev = append(rev, fmt.Sprintf("alter table \"%s\" alter column \"%s\" drop not null", tabName, newCol.Name))
	}
	//处理索引,字段更名的操作
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
	return rev
}
