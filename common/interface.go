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

//DB 是结合了查询和操作的数据库类
type DB interface {
	Execer
	Queryer
}
