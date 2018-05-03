package mysql

import (
	"fmt"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

func (m *meta) ChangeTableSQL(db common.DB, change *schema.TableSchemaChange) ([]string, error) {
	rev := []string{}
	tabName := change.NewName
	//处理表更名,处理过后，所有后续操作都在新表名上进行
	if change.OldName != change.NewName {
		rev = append(rev, tableRenameSQL(change.OldName, change.NewName)...)
	}

	//如果主键变更，则需要先除去主键
	if change.PKChange {
		rev = append(rev, dropTablePrimaryKeySQL(tabName)...)
	}
	//逐个处理字段，每处理一个字段，旧表字段就标上标记，最后删除没有标记的字段
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
	return rev, nil
}

func processColumnSQL(tabName string, oldCol, newCol *schema.Column) []string {
	rev := []string{}
	//如果是新增字段
	if oldCol == nil {
		rev := append(rev, fmt.Sprintf("alter table %s add %s", tabName, dbDefine(newCol)))
		//处理索引
		if newCol.Index {
			rev = append(rev, createColumnIndexSQL(tabName, newCol.Name)...)
		}
		return nil
	}
	//如果是更名，需要先处理
	if oldCol.Name != newCol.Name {
		rev = append(rev, fmt.Sprintf("alter table %s CHANGE column %s %s", tabName, oldCol.Name, dbDefine(newCol)))
	}
	if !oldCol.EqueNoIndex(newCol) {
		rev = append(rev, fmt.Sprintf("alter table %s MODIFY %s", tabName, dbDefine(newCol)))
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
	return rev
}
