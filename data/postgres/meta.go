package postgres

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/data"
)

const driverName = "postgres"

type meta struct{}

func init() {
	m := new(meta)
	data.Register(driverName, m)
	data.Register("pgx", m)

}
func (m *meta) Concat(vals ...string) string {
	return fmt.Sprintf("CONCAT(%s)", strings.Join(vals, ","))
}
func (m *meta) Merge(destTable, srcDataSQL string, pks, columns []string) string {
	updateSet := []string{}
	pkMap := map[string]bool{}
	onConflict := "ON CONFLICT(" + strings.Join(pks, ",") + ")"
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
	if len(updateSet) > 0 {
		onConflict = onConflict + " DO UPDATE SET\n" + strings.Join(updateSet, ",\n")
	} else {
		onConflict = onConflict + " DO NOTHING"
	}
	return fmt.Sprintf("insert %s into %s(%s)select %s from (%s) merge_src %s",
		ignore, destTable, strings.Join(columns, ","), strings.Join(columns, ","),
		srcDataSQL, onConflict)

}
func (m *meta) Minus(table1, where1, table2, where2 string, primaryKeys, cols []string) string {
	strSql := ""
	if len(where1) > 0 {
		where1 = "where " + where1
	}
	if len(where2) > 0 {
		where2 = "where " + where2
	}

	strSql = fmt.Sprintf(
		"select %s from %s %s EXCEPT select %s from %s %s",
		strings.Join(cols, ","),
		table1,
		where1,
		strings.Join(cols, ","),
		table2,
		where2)

	return strSql
}
func (m *meta) Concat_ws(separator string, vals ...string) string {
	list := make([]string, len(vals))
	for i, v := range vals {
		list[i] = fmt.Sprintf("nullif(%s,'')", v)
	}
	return fmt.Sprintf("concat_ws(%s,%s)", signString(separator), strings.Join(list, ","))
}

// 返回单引号包括的字符串
func signString(str string) string {

	return "'" + strings.Replace(str, "'", "''", -1) + "'"
}
