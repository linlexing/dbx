package sqlite

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/data"
)

const driverName = "sqlite3"

type meta struct{}

func init() {
	data.Register(driverName, new(meta))

}
func (m *meta) Concat(vals ...string) string {
	return strings.Join(vals, "||")
}
func (m *meta) Merge(db common.DB, destTable, srcTable string, pks, columns []string) error {
	panic("not impl")
}
func (m *meta) Minus(table1, where1, table2, where2 string, primaryKeys, cols []string) string {
	strSQL := ""
	if len(where1) > 0 {
		where1 = "where " + where1
	}
	if len(where2) > 0 {
		where2 = "where " + where2
	}

	strSQL = fmt.Sprintf(
		"select %s from %s %s EXCEPT select %s from %s %s",
		strings.Join(cols, ","),
		table1,
		where1,
		strings.Join(cols, ","),
		table2,
		where2)

	return strSQL
}
