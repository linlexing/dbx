package mysql

import (
	"fmt"
	"strings"

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
func (m *meta) UpdateFrom(destTable, srcDataSQL, additionSet string, pks, columns []data.ColMap) string {
	dataAligs := "datasrc_"

	links := make([]string, len(pks))
	for i, v := range pks {

		links[i] = fmt.Sprintf("%s.%s=%s.%s", destTable, v.Dest, dataAligs, v.Src)
	}

	sets := []string{}
	for _, col := range columns {
		sets = append(sets, fmt.Sprintf("%s.%s=%s.%s", destTable, col.Dest, dataAligs, col.Src))
	}

	if len(additionSet) > 0 {
		sets = append(sets, additionSet)
	}

	setStr := strings.Join(sets, ",")
	return fmt.Sprintf("update %s inner join (%s) %s on %s set %s",
		destTable, srcDataSQL, dataAligs, strings.Join(links, " and "), setStr)
}
func (m *meta) Merge(destTable, srcDataSQL string, pks, columns []string) string {
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
	return fmt.Sprintf("insert %s into %s(%s)select %s from (%s) merge_src %s",
		ignore, destTable, strings.Join(columns, ","), strings.Join(columns, ","),
		srcDataSQL, updateStr)

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
