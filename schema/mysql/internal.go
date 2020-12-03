package mysql

import (
	"fmt"
	"log"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

func removeColumnsSQL(tabName string, cols []string) []string {
	strList := []string{}
	for _, v := range cols {
		strList = append(strList, "DROP COLUMN "+v)
	}
	return []string{fmt.Sprintf("ALTER table %s %s", tabName, strings.Join(strList, ","))}
}
func tableExists(db common.DB, tabName string) (bool, error) {
	schemaName := ""
	ns := strings.Split(tabName, ".")
	tname := ""
	if len(ns) > 1 {
		schemaName = ns[0]
		tname = ns[1]
	} else {
		tname = tabName
	}
	if len(schemaName) == 0 {
		strSQL := "select schema()"

		row := db.QueryRow(strSQL)
		if err := row.Scan(&schemaName); err != nil {
			log.Println(strSQL)
			return false, common.NewSQLError(err, strSQL)
		}
	}
	strSQL := "SELECT count(*) FROM information_schema.tables WHERE table_schema = ? and table_name=?"
	var iCount int64
	row := db.QueryRow(strSQL, schemaName, tname)
	if err := row.Scan(&iCount); err != nil {
		return false, err
	}

	return iCount > 0, nil
}

//tableRenameSQL 处理表改名，旧名可以是多个，任意一个对上就改名，如果都没有存在，则不处理，也不返回出错
func tableRenameSQL(oldName string, newName string) []string {
	return []string{fmt.Sprintf("rename table %s TO %s", oldName, newName)}
}

//createColumnIndexSQL 新增单字段索引
func createColumnIndexSQL(tableName string, unique bool, colName string) []string {
	ns := strings.Split(tableName, ".")
	schema := ""
	tname := ""
	ustr := ""
	if unique {
		ustr = "unique "
	}
	if len(ns) > 1 {
		schema = ns[0] + "."
		tname = ns[1]
	} else {
		tname = tableName
	}
	//这里会有问题，如果表名和字段名比较长就会出错
	return []string{fmt.Sprintf("create %sindex %si%s%s on %s(%s)", ustr, schema, tname, colName, tableName, colName)}
}

func dropColumnIndexSQL(tableName, indexName string) []string {
	return []string{fmt.Sprintf("drop index %s on %s", indexName, tableName)}
}

//dropTablePrimaryKeySQL 删除主键
func dropTablePrimaryKeySQL(tableName string) []string {
	return []string{fmt.Sprintf("ALTER TABLE %s DROP PRIMARY KEY", tableName)}
}

//addTablePrimaryKeySQL 新增主键
func addTablePrimaryKeySQL(tableName string, pks []string) []string {
	return []string{fmt.Sprintf("alter table %s add primary key(%s)", tableName, strings.Join(pks, ","))}
}
func dbDefine(c *schema.Column) string {
	nullStr := ""
	typeStr := ""
	if !c.Null {
		nullStr = " NOT NULL"
	}
	if (len(c.FetchDriver) == 0 ||
		strings.ToLower(c.FetchDriver) == strings.ToLower(driverName)) && len(c.TrueType) > 0 {
		typeStr = c.TrueType
	}
	if len(typeStr) == 0 {
		switch c.Type {
		case schema.TypeBytea:
			typeStr = "BLOB"
		case schema.TypeDatetime:
			typeStr = "DATETIME"
		case schema.TypeFloat:
			typeStr = "DOUBLE PRECISION"
		case schema.TypeInt:
			typeStr = "BIGINT"
		case schema.TypeString:
			if c.MaxLength <= 0 {
				typeStr = "TEXT"
			} else {
				typeStr = fmt.Sprintf("VARCHAR(%d)", c.MaxLength)
			}
		default:
			panic("not impl DBType")
		}
	}
	return fmt.Sprintf("%s %s%s", c.Name, typeStr, nullStr)
}
