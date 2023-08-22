package ddb

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"time"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/data"
	"github.com/sirupsen/logrus"
)

type Txer = common.Tx
type TxDB = common.TxDB
type Tx = common.Tx
type DB = common.DB
type db struct {
	db              *sql.DB
	driverName      string
	connectString   string
	connMaxLifetime time.Duration
	connMaxIdleTime time.Duration
	maxIdleConns    int
	maxOpenConns    int
}

func (d *db) Driver() driver.Driver {
	return d.db.Driver()
}
func (d *db) Conn(ctx context.Context) (*sql.Conn, error) {
	return d.db.Conn(ctx)
}
func (d *db) SetConnMaxLifetime(t time.Duration) {
	d.connMaxLifetime = t
	d.db.SetConnMaxLifetime(t)
}
func (d *db) SetConnMaxIdleTime(t time.Duration) {
	d.connMaxIdleTime = t
	d.db.SetConnMaxIdleTime(t)
}
func (d *db) SetMaxIdleConns(n int) {
	d.maxIdleConns = n
	d.db.SetMaxIdleConns(n)
}
func (d *db) SetMaxOpenConns(n int) {
	d.maxOpenConns = n
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

//	func (d *db) Begin() (*sql.Tx, error) {
//		return d.db.Begin()
//	}
func (d *db) Beginx() (Txer, error) {
	return d.BeginTxx(context.Background(), nil)
}

//	func (d *db) BeginTx(ctx context.Context, opts *sql.TxOptions) (common.Txer, error) {
//		return d.db.BeginTx(ctx, opts)
//	}
func (d *db) BeginTxx(ctx context.Context, opts *sql.TxOptions) (Txer, error) {
	conn, err := d.db.Conn(ctx)
	if err != nil {
		return nil, err
	}

	t, err := conn.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	rev := &tx{
		driverName:    d.DriverName(),
		tx:            t,
		conn:          conn, //commit或者rollback时会自动调用conn.Close
		connectString: d.ConnectString(),
	}
	return rev, nil
}
func (d *db) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	done := make(chan struct{})
	tm := time.NewTimer(120 * time.Second)
	go func() {
		select {
		case <-tm.C:
			stats := d.db.Stats()
			logrus.WithFields(logrus.Fields{
				"sql":     query,
				"param":   fmt.Sprintf("%v", args),
				"maxconn": stats.MaxOpenConnections,
				"inuse":   stats.InUse,
			}).Info("slow-sql-execcontext")
		case <-done:
		}
	}()
	defer func() {
		tm.Stop()
		close(done)
	}()
	r, err := d.db.ExecContext(ctx, data.Bind(d.driverName, query), args...)
	if err != nil {
		err = common.NewSQLError(err, query, args...)
	}
	return r, err
}
func (d *db) Exec(query string, args ...interface{}) (sql.Result, error) {
	done := make(chan struct{})
	tm := time.NewTimer(120 * time.Second)
	go func() {
		select {
		case <-tm.C:
			stats := d.db.Stats()
			logrus.WithFields(logrus.Fields{
				"sql":     query,
				"param":   fmt.Sprintf("%v", args),
				"maxconn": stats.MaxOpenConnections,
				"inuse":   stats.InUse,
			}).Info("slow-sql-exec")
		case <-done:
		}
	}()
	defer func() {
		tm.Stop()
		close(done)
	}()
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
func (d *db) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	done := make(chan struct{})
	tm := time.NewTimer(120 * time.Second)
	go func() {
		select {
		case <-tm.C:
			stats := d.db.Stats()
			logrus.WithFields(logrus.Fields{
				"sql":     query,
				"param":   fmt.Sprintf("%v", args),
				"maxconn": stats.MaxOpenConnections,
				"inuse":   stats.InUse,
			}).Info("slow-sql-querycontext")
		case <-done:
		}
	}()
	defer func() {
		tm.Stop()
		close(done)
	}()
	r, err := d.db.QueryContext(ctx, data.Bind(d.driverName, query), args...)
	if err != nil {
		err = common.NewSQLError(err, query, args...)
	}
	return r, err
}
func (d *db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	done := make(chan struct{})
	tm := time.NewTimer(120 * time.Second)
	go func() {
		select {
		case <-tm.C:
			stats := d.db.Stats()
			logrus.WithFields(logrus.Fields{
				"sql":     query,
				"param":   fmt.Sprintf("%v", args),
				"maxconn": stats.MaxOpenConnections,
				"inuse":   stats.InUse,
			}).Info("slow-sql-query")
		case <-done:
		}
	}()
	defer func() {
		tm.Stop()
		close(done)
	}()
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
	done := make(chan struct{})
	tm := time.NewTimer(120 * time.Second)
	go func() {
		select {
		case <-tm.C:
			stats := d.db.Stats()
			logrus.WithFields(logrus.Fields{
				"sql":     query,
				"param":   fmt.Sprintf("%v", args),
				"maxconn": stats.MaxOpenConnections,
				"inuse":   stats.InUse,
			}).Info("slow-sql-queryrow")
		case <-done:
		}
	}()
	defer func() {
		tm.Stop()
		close(done)
	}()
	if len(args) > 0 {
		query = data.Bind(d.driverName, query)
	}
	return d.db.QueryRow(query, args...)
}
func (d *db) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	done := make(chan struct{})
	tm := time.NewTimer(120 * time.Second)
	go func() {
		select {
		case <-tm.C:
			stats := d.db.Stats()
			logrus.WithFields(logrus.Fields{
				"sql":     query,
				"param":   fmt.Sprintf("%v", args),
				"maxconn": stats.MaxOpenConnections,
				"inuse":   stats.InUse,
			}).Info("slow-sql-queryrowcontext")
		case <-done:
		}
	}()
	defer func() {
		tm.Stop()
		close(done)
	}()
	return d.db.QueryRowContext(ctx, data.Bind(d.driverName, query), args...)
}
func (d *db) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return d.db.PrepareContext(ctx, data.Bind(d.driverName, query))
}
func (d *db) Prepare(query string) (*sql.Stmt, error) {
	return d.db.Prepare(data.Bind(d.driverName, query))
}
func (d *db) Close() error {
	return d.db.Close()
}

// 释放所有的连接
func (d *db) ResetConnect() error {
	db := d.db
	//先开启新的连接
	if err := d.connect(); err != nil {
		return err
	}
	d.db.SetMaxIdleConns(d.maxIdleConns)
	d.db.SetMaxOpenConns(d.maxOpenConns)
	d.db.SetConnMaxIdleTime(d.connMaxIdleTime)
	d.db.SetConnMaxLifetime(d.connMaxLifetime)
	go func() {
		if err := db.Close(); err != nil {
			logrus.Error(fmt.Errorf("ResetConnect\n%w", err))
		}
	}()
	return nil

}
func (d *db) connect() (err error) {
	d.db, err = sql.Open(d.driverName, d.connectString)
	return
}
