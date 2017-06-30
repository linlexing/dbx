package ddb

import (
	"database/sql"
	"log"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/data"
)

//DB 一个能返回DriverName 的DB,并且所有的sql统一使用?作为参数占位符
type DB interface {
	common.DB
	DriverName() string
}

//TxDB 一个能返回DriverName 的TxDB,并且所有的sql统一使用?作为参数占位符
type TxDB interface {
	common.TxDB
	Beginx() (Txer, error)
	DriverName() string
}

type db struct {
	db         *sql.DB
	driverName string
}

func (d *db) DriverName() string {
	return d.driverName
}
func (d *db) Begin() (*sql.Tx, error) {
	return d.db.Begin()
}
func (d *db) Beginx() (Txer, error) {
	t, err := d.db.Begin()
	if err != nil {
		return nil, err
	}
	rev := &tx{
		driverName: d.driverName,
		tx:         t,
	}
	return rev, nil
}
func (d *db) Exec(query string, args ...interface{}) (sql.Result, error) {
	//无参数不用转换
	if len(args) > 0 {
		query = data.Bind(d.driverName, query)
	}
	r, err := d.db.Exec(query, args...)
	if err != nil {
		err = common.NewSQLError(err, query, args...)
		log.Println(err)
	}
	return r, err
}

func (d *db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	//无参数不用转换
	if len(args) > 0 {
		query = data.Bind(d.driverName, query)
	}
	r, err := d.db.Query(query, args...)
	if err != nil {
		err = common.NewSQLError(err, query, args...)
		log.Println(err)
	}
	return r, err
}
func (d *db) QueryRow(query string, args ...interface{}) *sql.Row {
	if len(args) > 0 {
		query = data.Bind(d.driverName, query)
	}
	return d.db.QueryRow(query, args...)
}

func (d *db) Prepare(query string) (*sql.Stmt, error) {
	return d.db.Prepare(data.Bind(d.driverName, query))
}
func (d *db) Close() error {
	return d.db.Close()
}
