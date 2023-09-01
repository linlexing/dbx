package oracle

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/data"
)

var (
	driverName = []string{"oci8", "oracle"}
)

type meta struct{}

func init() {
	for _, one := range driverName {
		data.Register(one, new(meta))
	}
}
func (m *meta) Concat(vals ...string) string {
	return strings.Join(vals, "||")
}
func (m *meta) UpdateFrom(destTable, srcDataSQL, additionSet string, pks, columns []data.ColMap) string {
	dataAligs := "datasrc_"
	// pkMap := map[string]struct{}{}
	links := make([]string, len(pks))
	for i, v := range pks {
		// pkMap[v] = struct{}{}
		links[i] = fmt.Sprintf("%s.%s=%s.%s", destTable, v.Dest, dataAligs, v.Src)
	}

	if len(additionSet) > 0 {
		additionSet = "," + additionSet
	}
	sets := make([]string, len(columns))
	for i, v := range columns {
		sets[i] = v.Dest
	}
	return fmt.Sprintf(
		"update %s set (%s)=(select %[2]s from (%s) %s where %s)%s where exists(select 1 from (%[3]s) %[4]s where %[6]s)",
		destTable, strings.Join(sets, ","), srcDataSQL, dataAligs, strings.Join(links, " and "),
		additionSet)
}

// Merge 将另一个表中的数据合并进本表，要求两个表的主键相同,相同主键的被覆盖
// columns指定字段清单,不在清单内的字段不会被update
func (m *meta) Merge(destTable, srcDataSQL string, pks, columns []string) string {
	join := []string{}
	updateSet := []string{}
	insertColumns := []string{}
	insertValues := []string{}
	pkMap := map[string]bool{}
	for _, v := range pks {
		pkMap[v] = true
		join = append(join, fmt.Sprintf("dest.%s = src.%s", v, v))
	}
	for _, field := range columns {
		//非主键的才更新
		if _, ok := pkMap[field]; !ok {
			updateSet = append(updateSet, fmt.Sprintf("dest.%s = src.%[1]s", field))
		}
		insertColumns = append(insertColumns, fmt.Sprintf("dest.%s", field))
		insertValues = append(insertValues, fmt.Sprintf("src.%s", field))
	}
	//如果只有主键字段，则省略WHEN MATCHED THEN子句
	updateStr := ""
	if len(updateSet) > 0 {
		updateStr = "WHEN MATCHED THEN UPDATE SET\n" + strings.Join(updateSet, ",\n")
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
	strSQL := ""
	if len(where1) > 0 {
		where1 = "where " + where1
	}
	if len(where2) > 0 {
		where2 = "where " + where2
	}

	strSQL = fmt.Sprintf(
		"select %s from %s %s minus select %s from %s %s",
		strings.Join(cols, ","),
		table1,
		where1,
		strings.Join(cols, ","),
		table2,
		where2)

	return strSQL
}
func (m *meta) Concat_ws(separator string, vals ...string) string {
	// 	SELECT TRIM(LEADING ';'
	//                FROM dx89||NVL2(dx90,';'||dx90,dx90)||
	//                           NVL2(dx91,';'||dx91,dx91)||
	//                           NVL2(dx92,';'||dx92,dx92)) AS "Concatenated String"
	//   FROM t
	if len(vals) == 0 {
		return "null"
	}
	secList := []string{vals[0]}
	for i := 1; i < len(vals); i++ {
		secList = append(secList, fmt.Sprintf("NVL2(%s,%s||%[1]s,%[1]s)", vals[i], signString(separator)))
	}
	return fmt.Sprintf("TRIM(LEADING %s FROM %s)", signString(separator), strings.Join(secList, "||"))
}

// 返回单引号包括的字符串
func signString(str string) string {

	return "'" + strings.Replace(str, "'", "''", -1) + "'"
}
