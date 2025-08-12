package duckdb

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/schema"
)

func tableRenameSQL(oldName string, newName string) []string {
	return []string{fmt.Sprintf("ALTER table %s RENAME TO %s", oldName, newName)}
}

// createColumnIndex 新增单字段索引
func createColumnIndexSQL(tableName string, unique bool, colName string) []string {
	ns := strings.Split(tableName, ".")
	schemaName := ""
	tname := ""
	if len(ns) > 1 {
		schemaName = ns[0] + "."
		tname = ns[1]
	} else {
		tname = tableName
	}
	ustr := ""
	if unique {
		ustr = "unique "
	}
	//这里会有问题，如果表名和字段名比较长就会出错
	return []string{fmt.Sprintf("create %sindex %si%s%s on \"%s\"(\"%s\")", ustr, schemaName, tname, colName, tableName, colName)}
}
func dropColumnIndexSQL(tableName, indexName string) []string {
	return []string{fmt.Sprintf("drop index %s on \"%s\"", indexName, tableName)}
}

// dropTablePrimaryKeySQL 删除主键
func dropTablePrimaryKeySQL(tableName string) []string {
	return []string{fmt.Sprintf("ALTER TABLE \"%s\" DROP PRIMARY KEY", tableName)}
}

// addTablePrimaryKeySQL 新增主键
func addTablePrimaryKeySQL(tableName string, pks []string) []string {
	return []string{fmt.Sprintf("alter table \"%s\" add primary key(\"%s\")", tableName, strings.Join(pks, ","))}
}

func dbType(dataType schema.DataType) string {
	switch dataType {
	case schema.TypeBytea:
		return "BLOB"
	case schema.TypeDatetime:
		return "timestamp with time zone"
	case schema.TypeFloat:
		return "DOUBLE"
	case schema.TypeInt:
		return "BIGINT"
	case schema.TypeString:
		return "TEXT"
	}
	panic("not impl DBType")
}

func colDBType(c *schema.Column) string {
	if (len(c.FetchDriver) == 0 || strings.ToLower(c.FetchDriver) == strings.ToLower(driverName)) && len(c.TrueType) > 0 {
		return c.TrueType
	}
	return dbType(c.Type)
}

func dbDefine(c *schema.Column) string {
	nullStr := ""
	if !c.Null {
		nullStr = " NOT NULL"
	}
	return fmt.Sprintf("\"%s\" %s%s", c.Name, colDBType(c), nullStr)
}

func removeColumnsSQL(tabName string, cols []string) []string {
	strList := []string{}
	for _, v := range cols {
		strList = append(strList, fmt.Sprintf("DROP COLUMN \"%s\"", v))
	}
	return []string{fmt.Sprintf("ALTER table \"%s\" %s", tabName, strings.Join(strList, ","))}
}

func DuckDBType(typeName string) schema.DataType {
	typeName = strings.ToUpper(typeName)
	switch typeName {
	case "TINYINT", "SMALLINT", "INTEGER", "BIGINT", "HUGEINT", "UTINYINT", "USMALLINT", "UINTEGER", "UBIGINT", "UHUGEINT":
		return schema.TypeInt
	case "VARCHAR":
		return schema.TypeString
	case "BLOB":
		return schema.TypeBytea
	case "FLOAT", "DOUBLE":
		return schema.TypeFloat
	case "DATE", "TIME", "TIME WITH TIME ZONE", "TIMESTAMP_NS", "TIMESTAMP", "TIMESTAMP_MS", "TIMESTAMP_S", "TIMESTAMP WITH TIME ZONE":
		return schema.TypeDatetime
	}
	return schema.TypeString
}
