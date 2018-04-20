package mysql

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/data"
)

const driverName = "mysql"

type meta struct{}

func init() {
	data.Register(driverName, new(meta))
}
func (m *meta) Concat(vals ...string) string {
	return fmt.Sprintf("CONCAT_WS('',%s)", strings.Join(vals, ","))
}
func (m *meta) Merge(db common.DB, destTable, srcTable string, pks,
	columns []string) error {
	updateSet := []string{}
	pkMap := map[string]bool{}
	ignore := ""
	for _, v := range pks {
		pkMap[v] = true
	}
	for _, field := range columns {
		//非主键的才更新
		if _, ok := pkMap[field]; !ok {
			updateSet = append(updateSet, fmt.Sprintf("%s = values(%[1]s)", field))
		}
	}
	//如果只有主键字段，则省略WHEN MATCHED THEN子句
	updateStr := ""
	if len(updateSet) > 0 {
		updateStr = "ON DUPLICATE KEY UPDATE\n" + strings.Join(updateSet, ",\n")
	} else {
		ignore = "IGNORE"
	}
	strSQL := fmt.Sprintf("insert %s into %s(%s)select %s from %s %s",
		ignore, destTable, strings.Join(columns, ","), strings.Join(columns, ","),
		srcTable, updateStr)
	_, err := db.Exec(strSQL)
	return err

}
func (m *meta) Minus(table1, where1, table2, where2 string,
	primaryKeys, cols []string) string {
	strSQL := ""
	if len(where1) > 0 {
		where1 = "where " + where1
	}
	if len(where2) > 0 {
		where2 = "where " + where2
	}
	keyMap := map[string]bool{}
	for _, v := range primaryKeys {
		keyMap[v] = true
	}
	join := []string{}
	cols_l := []string{}
	for _, str := range cols {
		cols_l = append(cols_l, fmt.Sprintf("l_a.%s", str))
		//如果是主键，则不需要检查null
		if _, ok := keyMap[str]; ok {
			join = append(join, fmt.Sprintf("l_a.%[1]s=l_b.%[1]s", str))
		} else {
			join = append(join, fmt.Sprintf("(l_a.%s is null and l_b.%[1]s is null or l_a.%[1]s=l_b.%[1]s)", str))
		}
	}
	var from1 string
	var from2 string
	if len(where1) > 0 {
		from1 = fmt.Sprintf("(select * from %s %s)", table1, where1)
	} else {
		from1 = table1
	}
	if len(where2) > 0 {
		from2 = fmt.Sprintf("(select * from %s %s)", table2, where2)
	} else {
		from2 = table2
	}
	strSQL = fmt.Sprintf(
		"select %s from %s l_a left join %s l_b on %s where %s",
		strings.Join(cols_l, ",\n"),
		from1,
		from2,
		strings.Join(join, " and\n"),
		fmt.Sprintf("l_b.%s is null", primaryKeys[0]))
	return strSQL
}
