package common

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

// Execer 是大多数元数据操作函数用到的参数
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

// Queryer 是大多数元数据操作函数用到的参数
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// Preparer 是能够批量执行同一个语句的对象
type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// Scaner 是能扫描数据的类，一般是sql.Rows
type Scaner interface {
	Scan(dest ...interface{}) error
}

// DB 是结合了查询和操作的数据库类
// DB 一个能返回DriverName 的DB,并且所有的sql统一使用?作为参数占位符
type DB interface {
	Execer
	Queryer
	Preparer
	DriverName() string
	ConnectString() string
}

// Txer 是事务
// Txer 一个能返回DriverName 的Txer

type Tx interface {
	Execer
	Queryer
	Preparer
	DriverName() string
	ConnectString() string
	Conn() *sql.Conn
	Commit() error
	Rollback() error
}

// TxDB 一个能返回DriverName 的TxDB,并且所有的sql统一使用?作为参数占位符
type TxDB interface {
	DB
	Conn(ctx context.Context) (*sql.Conn, error)
	Close() error
	Beginx() (Tx, error)
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
	Driver() driver.Driver
	Ping() error
	SetConnMaxLifetime(d time.Duration)
	SetConnMaxIdleTime(d time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	ResetConnect() error
	Stats() sql.DBStats
}
