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
	data.Register("pgx-un", m)
}
func (m *meta) Concat(vals ...string) string {
	return fmt.Sprintf("CONCAT(%s)", strings.Join(vals, ","))
}
func (m *meta) UpdateFrom(destTable, srcDataSQL, additionSet string, pks, columns []data.ColMap) string {
	dataAligs := "datasrc_"
	// pkMap := map[string]struct{}{}
	links := make([]string, len(pks))
	for i, v := range pks {
		// pkMap[v] = struct{}{}
		links[i] = fmt.Sprintf("%s.%s=%s.%s", destTable, v.Dest, dataAligs, v.Src)
	}

	sets := []string{}
	for _, col := range columns {
		sets = append(sets, fmt.Sprintf("%s=%s.%s", col.Dest, dataAligs, col.Src))
	}

	if len(additionSet) > 0 {
		sets = append(sets, additionSet)
	}

	setStr := strings.Join(sets, ",")
	return fmt.Sprintf("update %s set %s from (%s) %s where %s",
		destTable, setStr, srcDataSQL, dataAligs, strings.Join(links, " and "))
}
func (m *meta) Merge(destTable, srcDataSQL string, pks, columns []string, skipCheckCols ...string) string {
	updateSet := []string{}
	pkMap := map[string]bool{}
	onConflict := "ON CONFLICT(" + strings.Join(pks, ",") + ")"
	ignore := ""
	for _, v := range pks {
		pkMap[v] = true
	}
	valNotEquWhereList := []string{}
	skipCheckColsMap := map[string]struct{}{}
	for _, v := range skipCheckCols {
		skipCheckColsMap[strings.ToUpper(v)] = struct{}{}
	}
	for _, field := range columns {
		//非主键的才更新
		if _, ok := pkMap[field]; !ok {
			updateSet = append(updateSet, fmt.Sprintf("%s = excluded.%[1]s", field))
			//非跳过检查的，才进行值判断，一般是最后合并时间这类字段不需要检查，防止生成多余的日志
			if _, skip := skipCheckColsMap[strings.ToUpper(field)]; !skip {
				valNotEquWhereList = append(valNotEquWhereList, fmt.Sprintf("dest.%s is distinct from excluded.%[1]s", field))
			}
		}
	}
	//如果只有主键字段，则省略WHEN MATCHED THEN子句
	if len(updateSet) > 0 {
		onConflict = onConflict + " DO UPDATE SET\n" + strings.Join(updateSet, ",\n") +
			"\nwhere\n" + strings.Join(valNotEquWhereList, " or ")
	} else {
		onConflict = onConflict + " DO NOTHING"
	}
	return fmt.Sprintf("insert %s into %s as dest(%s)select %s from (%s) merge_src %s",
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
