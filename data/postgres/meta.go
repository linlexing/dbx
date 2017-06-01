package postgres

import (
	"fmt"
	"log"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/data"
)

const driverName = "postgres"

type meta struct{}

func init() {
	data.Register(driverName, new(meta))
}
func (m *meta) Merge(db common.DB, destTable, srcTable string, pks, columns []string) error {
	panic("not impl")
}
func (m *meta) Minus(db common.DB, table1, where1, table2, strSQL string, primaryKeys, cols []string) string {
	strSql := ""
	if len(where1) > 0 {
		where1 = "where " + where1
	}
	if len(strSQL) > 0 {
		strSQL = "where " + strSQL
	}

	strSql = fmt.Sprintf(
		"select %s from %s %s EXCEPT select %s from %s %s",
		strings.Join(cols, ","),
		table1,
		where1,
		strings.Join(cols, ","),
		table2,
		strSQL)

	return strSql
}
