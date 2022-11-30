package ddb

import (
	"context"
	"database/sql"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/data"
)

// Txer 一个能返回DriverName 的Txer
type Txer interface {
	common.Txer
	DriverName() string
	ConnectString() string
}

// Tx 一个能存贮driverName的tx
type tx struct {
	tx            *sql.Tx
	driverName    string
	connectString string
}

func (t *tx) DriverName() string {
	return t.driverName
}
func (t *tx) ConnectString() string {
	return t.connectString
}
func (t *tx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	r, err := t.tx.ExecContext(ctx, data.Bind(t.driverName, query), args...)
	if err != nil {
		err = common.NewSQLError(err, query, args...)
	}
	return r, err
}
func (t *tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	r, err := t.tx.Exec(data.Bind(t.driverName, query), args...)
	if err != nil {
		err = common.NewSQLError(err, query, args...)
	}
	return r, err
}
func (t *tx) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	r, err := t.tx.QueryContext(ctx, data.Bind(t.driverName, query), args...)
	if err != nil {
		err = common.NewSQLError(err, query, args...)
	}
	return r, err
}

func (t *tx) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return t.tx.QueryRowContext(ctx, data.Bind(t.driverName, query), args...)
}

func (t *tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	r, err := t.tx.Query(data.Bind(t.driverName, query), args...)
	if err != nil {
		err = common.NewSQLError(err, query, args...)
	}
	return r, err
}
func (t *tx) QueryRow(query string, args ...interface{}) *sql.Row {
	return t.tx.QueryRow(data.Bind(t.driverName, query), args...)
}

func (t *tx) Prepare(query string) (*sql.Stmt, error) {
	return t.tx.Prepare(data.Bind(t.driverName, query))
}
func (t *tx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return t.tx.PrepareContext(ctx, data.Bind(t.driverName, query))
}

func (t *tx) Commit() error {
	return t.tx.Commit()
}
func (t *tx) Rollback() error {
	return t.tx.Rollback()
}
