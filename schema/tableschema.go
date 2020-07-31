package schema

import (
	"errors"
	"fmt"
	"strings"

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
func (t *tableSchema) pkIsChanged() bool {
	if len(t.oldTable.PrimaryKeys) != len(t.newTable.PrimaryKeys) {
		return true
	}
	//判断是否是改名
	beforeChangeName := []string{}
	for _, v := range t.newTable.PrimaryKeys {
		if fld := t.oldTable.findColumnAnyName(append([]string{v},
			t.newTable.ColumnByName(v).FormerName...)...); fld != nil {
			beforeChangeName = append(beforeChangeName, fld.Name)
		}
	}
	for i, one := range beforeChangeName {
		if strings.ToUpper(one) != strings.ToUpper(t.oldTable.PrimaryKeys[i]) {
			return true
		}
	}
	return false
}

//extract 提取修改数据库结构的sql语句
func (t *tableSchema) extract() ([]string, error) {
	//先检查字段名称是否合法
	if err := t.checkTableColumns(t.newTable); err != nil {
		return nil, err
	}
	//如果没有旧表，则是新增表
	if t.oldTable == nil {
		return t.mt.CreateTableSQL(t.db, t.newTable)
	}
	//处理表更名,处理过后，所有后续操作都在新表名上进行
	schg := &TableSchemaChange{
		OldName:      t.oldTable.FullName(),
		NewName:      t.newTable.FullName(),
		PK:           t.newTable.PrimaryKeys,
		OriginFields: t.oldTable.Columns,
	}

	//如果主键变更，则需要先除去主键
	//主键改名不属于主键变动
	schg.PKChange = t.pkIsChanged()

	//逐个处理字段，每处理一个字段，旧表字段就标上标记，最后删除没有标记的字段
	oldColumnProcesses := map[string]bool{}
	for _, v := range t.oldTable.Columns {
		oldColumnProcesses[strings.ToUpper(v.Name)] = false
	}
	for _, col := range t.newTable.Columns {
		//用曾用名+现有名称去找旧字段
		oldCol := t.oldTable.findColumnAnyName(append(col.FormerName, col.Name)...)
		if oldCol != nil {
			oldColumnProcesses[strings.ToUpper(oldCol.Name)] = true
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
	return t.mt.ChangeTableSQL(t.db, schg)
}
