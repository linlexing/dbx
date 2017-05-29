package mysql

import (
	"fmt"
	"log"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

func removeColumns(db common.DB, tabName string, cols []string) error {
	var strSQL string
	strList := []string{}
	for _, v := range cols {
		strList = append(strList, "DROP COLUMN "+v)
	}
	strSQL = fmt.Sprintf("ALTER table %s %s", tabName, strings.Join(strList, ","))
	if _, err := db.Exec(strSQL); err != nil {
		log.Println(strSQL)
		return common.NewSQLError(err, strSQL)
	}
	return nil
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
	strSQL := "SELECT count(*) FROM information_schema.tables WHERE table_schema = ? and UPPER(table_name)=?"
	var iCount int64
	row := db.QueryRow(strSQL, schemaName, tname)
	if err := row.Scan(&iCount); err != nil {
		return false, err
	}

	return iCount > 0, nil
}

//tableRename 处理表改名，旧名可以是多个，任意一个对上就改名，如果都没有存在，则不处理，也不返回出错
func tableRename(db common.DB, oldName string, newName string) error {

	strSQL := fmt.Sprintf("rename table %s TO %s", oldName, newName)
	if _, err := db.Exec(strSQL); err != nil {
		log.Println(strSQL)
		return common.NewSQLError(err, strSQL)
	}

	return nil
}

//CreateColumnIndex 新增单字段索引
func createColumnIndex(db common.DB, tableName, colName string) error {
	ns := strings.Split(tableName, ".")
	schema := ""
	tname := ""
	if len(ns) > 1 {
		schema = ns[0] + "."
		tname = ns[1]
	} else {
		tname = tableName
	}
	//这里会有问题，如果表名和字段名比较长就会出错
	strSQL := fmt.Sprintf("create index %si%s%s on %s(%s)", schema, tname, colName, tableName, colName)
	if _, err := db.Exec(strSQL); err != nil {
		log.Println(strSQL)
		return common.NewSQLError(err, strSQL)
	}
	return nil
}
func dropColumnIndex(db common.DB, tableName, indexName string) error {
	strSQL := fmt.Sprintf("drop index %s on %s", indexName, tableName)
	if _, err := db.Exec(strSQL); err != nil {
		log.Println(strSQL)
		return common.NewSQLError(err, strSQL)
	}
	return nil
}

//dropTablePrimaryKey 删除主键
func dropTablePrimaryKey(db common.DB, tableName string) error {
	strSQL := fmt.Sprintf("ALTER TABLE %s DROP PRIMARY KEY", tableName)
	if _, err := db.Exec(strSQL); err != nil {
		log.Println(strSQL)
		return common.NewSQLError(err, strSQL)
	}
	return nil
}

//addTablePrimaryKey 新增主键
func addTablePrimaryKey(db common.DB, tableName string, pks []string) error {
	strSQL := fmt.Sprintf("alter table %s add primary key(%s)", tableName, strings.Join(pks, ","))
	if _, err := db.Exec(strSQL); err != nil {
		log.Println(strSQL)
		return common.NewSQLError(err, strSQL)
	}
	return nil
}

func dbDefine(c *schema.Column) string {
	nullStr := ""
	typeStr := ""
	if !c.Null {
		nullStr = " NOT NULL"
	}
	if c.FetchDriver == driverName && len(c.TrueType) > 0 {
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
func valueExpress(dataType schema.DataType, value string) string {
	switch dataType {
	case schema.TypeFloat, schema.TypeInt:
		return value
	case schema.TypeString:
		return "'" + strings.Replace(value, "'", "''", -1) + "'"
	case schema.TypeDatetime:
		if len(value) == 10 {
			return "STR_TO_DATE('" + value + "','%Y-%m-%d')"
		} else if len(value) == 19 {
			return "STR_TO_DATE('" + value + "','%Y-%m-%d %h:%i:%s')"
		} else {
			panic(fmt.Errorf("invalid datetime:%s", value))
		}
	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}
