package schema

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/linlexing/dbx/common"
)

//tableSchema 该结构完成表结构的调整，自动处理
//1.表改名
//2.主键更改
//3.字段改名
//4.字段调整
//5.单字段的索引调整
type tableSchema struct {
	oldTable *Table
	newTable *Table
	mt       Meta
	db       common.DB
}

//CheckTableColumns 检查新表的字段定义是否合法：
//1.字段名（含曾用名）不能重复
//2.必须有主键
func (t *tableSchema) checkTableColumns(tab *Table) error {
	if len(tab.PrimaryKeys) == 0 {
		return errors.New("primary key is null")
	}
	uname := map[string]bool{}
	for _, c := range tab.Columns {
		if _, ok := uname[c.Name]; ok {
			return fmt.Errorf("column:%s dup", c.Name)
		}
		uname[c.Name] = true
		//如果该字段有曾用名，则需检查曾用名是否和现有的名称重复
		if len(c.FormerName) > 0 {
			for _, fc := range c.FormerName {
				if _, ok := uname[fc]; ok {
					return fmt.Errorf("column:%s former name:%s dup", c.Name, fc)
				}
				uname[fc] = true
			}
		}
	}
	return nil
}

//update 将定义更新到数据库中
func (t *tableSchema) update() error {
	//先检查字段名称是否合法
	if err := t.checkTableColumns(t.newTable); err != nil {
		return err
	}
	//如果没有旧表，则是新增表
	if t.oldTable == nil {
		return t.mt.CreateTable(t.db, t.newTable)
	}
	//处理表更名,处理过后，所有后续操作都在新表名上进行
	schg := &TableSchemaChange{
		OldName:      t.oldTable.FullName(),
		NewName:      t.newTable.FullName(),
		PK:           t.newTable.PrimaryKeys,
		OriginFields: t.oldTable.Columns,
	}

	//如果主键变更，则需要先除去主键
	if !reflect.DeepEqual(t.oldTable.PrimaryKeys, t.newTable.PrimaryKeys) {
		schg.PKChange = true

	}
	//逐个处理字段，每处理一个字段，旧表字段就标上标记，最后删除没有标记的字段
	oldColumnProcesses := map[string]bool{}
	for _, v := range t.oldTable.Columns {
		oldColumnProcesses[v.Name] = false
	}
	for _, col := range t.newTable.Columns {
		//用曾用名+现有名称去找旧字段
		oldCol := t.oldTable.findColumnAnyName(append(col.FormerName, col.Name)...)
		if oldCol != nil {
			oldColumnProcesses[oldCol.Name] = true
		}
		//只有新增（oldCol为nil）或不相等的字段才纳入修改
		if oldCol == nil || !oldCol.Eque(col) {
			schg.ChangeFields = append(schg.ChangeFields, &ChangedField{
				OldField: oldCol,
				NewField: col,
			})
		}
	}
	//最后删除没有处理过的旧字段

	for k, prc := range oldColumnProcesses {
		if !prc {
			schg.RemoveFields = append(schg.RemoveFields, k)

		}
	}
	return t.mt.ChangeTable(t.db, schg)

}
