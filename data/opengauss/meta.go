package postgres

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/data"
)

const driverName = "opengauss"

type meta struct{}

func init() {
	m := new(meta)
	data.Register(driverName, m)
	data.Register("pgx-opengauss", m)
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

// Merge 将另一个表中的数据合并进本表，要求两个表的主键相同,相同主键的被覆盖
// columns指定字段清单,不在清单内的字段不会被update
func (m *meta) Merge(destTable, srcDataSQL string, pks, columns []string, skipCheckCols ...string) string {
	join := []string{}
	updateSet := []string{}
	insertColumns := []string{}
	insertValues := []string{}
	pkMap := map[string]bool{}
	for _, v := range pks {
		pkMap[v] = true
		join = append(join, fmt.Sprintf("dest.%s = src.%s", v, v))
	}
	valNotEquWhereList := []string{}
	skipCheckColsMap := map[string]struct{}{}
	for _, v := range skipCheckCols {
		skipCheckColsMap[strings.ToUpper(v)] = struct{}{}
	}
	for _, field := range columns {
		//非主键的才更新
		if _, ok := pkMap[field]; !ok {
			updateSet = append(updateSet, fmt.Sprintf("dest.%s = src.%[1]s", field))
			//非跳过检查的，才进行值判断，一般是最后合并时间这类字段不需要检查，防止生成多余的日志
			if _, skip := skipCheckColsMap[strings.ToUpper(field)]; !skip {
				valNotEquWhereList = append(valNotEquWhereList, fmt.Sprintf("dest.%s is distinct from src.%[1]s", field))
			}
		}
		insertColumns = append(insertColumns, fmt.Sprintf("dest.%s", field))
		insertValues = append(insertValues, fmt.Sprintf("src.%s", field))
	}
	//如果只有主键字段，则省略WHEN MATCHED THEN子句
	updateStr := ""
	if len(updateSet) > 0 {
		updateStr = "WHEN MATCHED THEN UPDATE SET\n" + strings.Join(updateSet, ",\n") +
			"\nwhere\n" + strings.Join(valNotEquWhereList, " or ")
	}
	return fmt.Sprintf(`
MERGE INTO %s dest
USING(%s) src
ON(%s)
%s
WHEN NOT MATCHED THEN INSERT
	(%s)
	values
	(%s)`, destTable, srcDataSQL,
		strings.Join(join, " and "),
		updateStr,
		strings.Join(insertColumns, ","),
		strings.Join(insertValues, ","))

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
