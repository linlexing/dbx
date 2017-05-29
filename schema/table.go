package schema

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/common"
)

//Table 代表数据库中一个物理表
type Table struct {
	Schema      string
	Name        string
	Columns     []*Column
	FormerName  []string
	PrimaryKeys []string
}

//NewTable 返回一个新表，名称会自动依据句点拆分为shcema和name
func NewTable(name string) *Table {
	rev := new(Table)
	ns := strings.Split(name, ".")
	if len(ns) > 1 {
		rev.Schema = ns[0]
		rev.Name = ns[1]
	} else {
		rev.Name = name
	}
	return rev
}

//FullName 返回全名称，包括schema
func (t *Table) FullName() string {
	if len(t.Schema) > 0 {
		return t.Schema + "." + t.Name
	}
	return t.Name
}

//Update 更新一个表的结构至数据库中，会自动处理表改名、字段改名以及字段修改、索引修改等操作,
//先自动去数据库取出旧表结构
func (t *Table) Update(db common.DB, mt Meta) error {
	sch := &tableSchema{
		newTable: t,
		mt:       mt,
		db:       db,
	}
	if len(t.FormerName) > 0 {
		//如果有曾用名，则需验证曾用名不能和现有名称重复
		uname := map[string]bool{
			t.FullName(): true,
		}
		for _, v := range t.FormerName {
			if _, ok := uname[v]; ok {
				return fmt.Errorf("FormerName:%s dup", v)
			}
		}
		//并根据曾用名去获取之前的表结构
		for _, v := range t.FormerName {
			if b, _ := mt.TableExists(db, v); b {
				oldTable, err := mt.OpenTable(db, v)
				if err != nil {
					return nil
				}
				sch.oldTable = oldTable
				break
			}
		}
	}
	//如果曾用名的表找不到，则就用本来的名称，说明不需改名
	if sch.oldTable == nil {
		b, err := mt.TableExists(db, t.FullName())
		if err != nil {
			return err
		}
		if b {
			sch.oldTable, err = mt.OpenTable(db, t.FullName())
			if err != nil {
				return err
			}
		}
	}
	return sch.update()
}
func (t *Table) findColumnAnyName(names ...string) *Column {
	//用map作为检索索引
	idx := map[string]bool{}
	for _, oneName := range names {
		idx[oneName] = true
	}
	for _, col := range t.Columns {
		if _, ok := idx[col.Name]; ok {
			return col
		}
	}
	return nil
}
