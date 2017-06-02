package common

import "database/sql"

//Execer 是大多数元数据操作函数用到的参数
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

//Queryer 是大多数元数据操作函数用到的参数
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

//Preparer 是能够批量执行同一个语句的对象
type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

//Scaner 是能扫描数据的类，一般是sql.Rows
type Scaner interface {
	Scan(dest ...interface{}) error
}

//DB 是结合了查询和操作的数据库类
type DB interface {
	Execer
	Queryer
	Preparer
}

//TxDB 是带事务操作的DB
type TxDB interface {
	DB
	Begin() (*sql.Tx, error)
}

//Txer 是事务
type Txer interface {
	DB
	Commit() error
	Rollback() error
}
