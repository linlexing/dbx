package pageselect

import (
	"database/sql"

	"github.com/linlexing/dbx/scan"
	"github.com/linlexing/dbx/schema"
)

//PageSelecter 是pageselect类用来接入任意数据库系统的方法
type PageSelecter interface {
	Sum(col string) string
	Avg(col string) string
	ColumnTypes(rows *sql.Rows) ([]*scan.ColumnType, error)
	//标识符加引号，mysql是`
	QuotedIdentifier(col string) string
	GetOperatorExpress(ope Operator, dataType schema.DataType, left, right string) string
	SortByAsc(string, bool) string
	SortByDesc(string, bool) string
	LimitSQL(sel, strSQL, where, orderby string, limit int) string
}
