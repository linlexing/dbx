package oracle

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/data"
)

const driverName = "oci8"

type meta struct{}

func init() {
	data.Register(driverName, new(meta))
}

//Merge 将另一个表中的数据合并进本表，要求两个表的主键相同,相同主键的被覆盖
//columns指定字段清单,不在清单内的字段不会被update
func (m *meta) Merge(db common.DB, destTable, srcTable string, pks, columns []string) error {
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
	strSQL := fmt.Sprintf(`
MERGE INTO %s dest
USING(select * from %s) src 
ON(%s)
%s
WHEN NOT MATCHED THEN INSERT
	(%s)
	values
	(%s)`, destTable, srcTable,
		strings.Join(join, " and "),
		updateStr,
		strings.Join(insertColumns, ","),
		strings.Join(insertValues, ","))
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		return err
	}

	return nil

}
func (m *meta) Minus(db common.DB, table1, where1, table2, where2 string, primaryKeys, cols []string) string {
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
