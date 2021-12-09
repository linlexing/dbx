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

	updateSet := []string{}
	pkMap := map[string]bool{}
	ignore := ""
	for _, v := range pks {
		pkMap[v] = true
	}
	for _, field := range columns {
		//非主键的才更新
		if _, ok := pkMap[field]; !ok {
			updateSet = append(updateSet, fmt.Sprintf("%s = excluded.%[1]s", field))
		}
	}
	//如果只有主键字段，则省略WHEN MATCHED THEN子句
	updateStr := ""
	if len(updateSet) > 0 {
		updateStr = fmt.Sprintf("ON CONFLICT(%s) DO UPDATE SET %s\n", strings.Join(pks, ","),
			strings.Join(updateSet, ",\n"))
	} else {
		ignore = "NOTHING"
	}
	strSQL := fmt.Sprintf("insert %s into %s(%s)select %s from %s where true %s",
		ignore, destTable, strings.Join(columns, ","), strings.Join(columns, ","),
		srcTable, updateStr)
	_, err := db.Exec(strSQL)
	return err
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
