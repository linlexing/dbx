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
//skipColumns指定跳过update的字段清单
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
	for _, field := range t.columns {
		//非主键的才更新
		if _, ok := pkMap[field.Name]; !ok {
			updateSet = append(updateSet, fmt.Sprintf("dest.%s = src.%[1]s", field))
		}
		insertColumns = append(insertColumns, fmt.Sprintf("dest.%s", field))
		insertValues = append(insertValues, fmt.Sprintf("src.%s", field))
	}
	strSQL := fmt.Sprintf(`
MERGE INTO %s dest
USING(select * from %s) src 
ON(%s)
WHEN MATCHED THEN UPDATE SET
	%s
WHEN NOT MATCHED THEN INSERT
	(%s)
	values
	(%s)`, destTable, srcTable,
		strings.Join(join, " and "),
		strings.Join(updateSet, ",\n"),
		strings.Join(insertColumns, ","),
		strings.Join(insertValues, ","))
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSql)
		return err
	}

	return nil

}
