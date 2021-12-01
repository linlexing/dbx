package ddb

import (
	"github.com/linlexing/dbx/data"
	"github.com/linlexing/dbx/schema"
)

//Table 是对data.Table 的包装，主要为了table.DB的类型
type Table struct {
	db DB
	*data.Table
}

//DB 返回当前的DB
func (t *Table) DB() DB {
	return t.db
}

//SetDB 设置当前表的DB
func (t *Table) SetDB(db DB) {
	t.db = db
	t.Table.DB = db
}

//OpenTable 是data.OpenTable 的包装
func OpenTable(db DB, tableName string) (*Table, error) {
	tab, err := data.OpenTable(db.DriverName(), db, tableName)
	if err != nil {
		return nil, err
	}
	return &Table{
		Table: tab,
		db:    db,
	}, nil
}

//NewTable 是data.NewTable 的包装
func NewTable(db DB, tab *schema.Table) *Table {
	return &Table{
		Table: data.NewTable(db.DriverName(), db, tab),
		db:    db,
	}
}
