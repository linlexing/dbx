package pageselect

import "github.com/linlexing/dbx/schema"

//PageSelecter 是pageselect类用来接入任意数据库系统的方法
type PageSelecter interface {
	IsNull() string

	GetOperatorExpress(ope Operator, dataType schema.DataType, left, right string) string
	SortByAsc(string) string
	SortByDesc(string) string
	LimitSQL(sel, strSQL, where, orderby string, limit int) string
}
