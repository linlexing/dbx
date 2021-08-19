package ddb

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/data"
)

//DB 一个能返回DriverName 的DB,并且所有的sql统一使用?作为参数占位符
type DB interface {
	common.DB
	DriverName() string
	ConnectString() string
}

//TxDB 一个能返回DriverName 的TxDB,并且所有的sql统一使用?作为参数占位符
type TxDB interface {
	common.TxDB
	Beginx() (Txer, error)
	DriverName() string
	ConnectString() string
	Ping() error
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	ResetConnect() error
	Stats() sql.DBStats
}

type db struct {
	db            *sql.DB
	driverName    string
	connectString string
}

func (d *db) Conn(ctx context.Context) (*sql.Conn, error) {
	return d.db.Conn(ctx)
}
func (d *db) SetConnMaxLifetime(t time.Duration) {
	d.db.SetConnMaxLifetime(t)
}
func (d *db) SetMaxIdleConns(n int) {
	d.db.SetMaxIdleConns(n)
}
func (d *db) SetMaxOpenConns(n int) {
	d.db.SetMaxOpenConns(n)
}
func (d *db) Stats() sql.DBStats {
	return d.db.Stats()
}

func (d *db) Ping() error {
	return d.db.Ping()
}

func (d *db) DriverName() string {
	return d.driverName
}
func (d *db) ConnectString() string {
	return d.connectString
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
		driverName:    d.DriverName(),
		tx:            t,
		connectString: d.ConnectString(),
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
func (d *db) ResetConnect() error {
	if err := d.db.Close(); err != nil {
		return err
	}
	return d.connect()
}
func (d *db) connect() (err error) {
	d.db, err = sql.Open(d.driverName, d.connectString)
	return
}
