package oracle

import (
	"fmt"
	"log"
	"strings"

	"database/sql"

	"slices"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

func dbType(dataType schema.DataType, maxLength int) string {
	switch dataType {
	case schema.TypeBytea:
		if maxLength <= 0 || maxLength > 4000 {
			return "BLOB"
		}
		return "RAW"

	case schema.TypeDatetime:
		return "DATE"
	case schema.TypeFloat:
		return "BINARY_DOUBLE"
	case schema.TypeInt:
		return "INT"
	case schema.TypeString:
		if maxLength <= 0 {
			return "CLOB"
		}
		if maxLength > 4000 {
			return "VARCHAR2(4000)"
		}
		return fmt.Sprintf("VARCHAR2(%d CHAR)", maxLength)

	}

	panic("not impl DBType")

}
func dropTablePrimaryKeySQL(db common.DB, tableName string) ([]string, error) {
	ns := strings.Split(tableName, ".")
	var strSQL string
	var params []interface{}
	if len(ns) > 1 {
		strSQL =
			"select constraint_name from ALL_CONSTRAINTS where owner = :1 and table_name =:2 and constraint_type='P'"
		params = []interface{}{
			strings.ToUpper(ns[0]),
			strings.ToUpper(ns[1])}
	} else {
		strSQL =
			"select constraint_name from user_CONSTRAINTS where table_name =:1 and constraint_type='P'"
		params = []interface{}{
			strings.ToUpper(tableName)}
	}
	var pkCons string
	if err := db.QueryRow(strSQL, params...).Scan(&pkCons); err != nil {
		//如果找不到主键，则不需删除
		if err == sql.ErrNoRows {
			return nil, nil
		}
		err = common.NewSQLError(err, strSQL, params...)
		log.Println(err)
		return nil, err
	}
	return []string{fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, pkCons)}, nil
}

// addTablePrimaryKeySQL 新增主键
func addTablePrimaryKeySQL(tableName string, pks []string) []string {
	ns := strings.Split(tableName, ".")
	var clearTableName string
	if len(ns) > 1 {
		clearTableName = ns[1]
	} else {
		clearTableName = tableName
	}
	return []string{fmt.Sprintf("alter table %s add constraint %s_pk primary key(%s)", tableName, clearTableName, strings.Join(pks, ","))}
}
func colDBType(c *schema.Column) string {

	if (len(c.FetchDriver) == 0 || //内存中直接定义
		slices.Index(driverName, strings.ToLower(c.FetchDriver)) >= 0) && //从oracle数据库中取回
		len(c.TrueType) > 0 {
		return c.TrueType
	}
	return dbType(c.Type, c.MaxLength)
}

func dbDefine(c *schema.Column) string {
	nullStr := ""
	if !c.Null {
		nullStr = " NOT NULL"
	}
	return fmt.Sprintf("%s %s%s", c.Name, colDBType(c), nullStr)
}

func dbDefineNull(c *schema.Column) string {
	nullStr := " NULL"
	if !c.Null {
		nullStr = " NOT NULL"
	}
	return fmt.Sprintf("%s %s%s", c.Name, colDBType(c), nullStr)
}

// tableRenameSQL 处理表改名，旧名可以是多个，任意一个对上就改名，如果都没有存在，则不处理，也不返回出错
func tableRenameSQL(oldName string, newName string) []string {
	return []string{fmt.Sprintf("rename %s TO %s", oldName, newName)}
}
func tableExists(db common.DB, tabName string) (bool, error) {
	var schemaName string
	var tname string
	ns := strings.Split(tabName, ".")

	if len(ns) > 1 {
		schemaName = ns[0]
		tname = ns[1]
	} else {
		tname = tabName
	}
	if len(schemaName) == 0 {
		strSQL := "select user from dual"
		row := db.QueryRow(strSQL)
		if err := row.Scan(&schemaName); err != nil {
			err = common.NewSQLError(err, strSQL)
			log.Println(err)
			return false, err
		}

	}
	strSQL := "SELECT count(*) FROM all_tables where owner=:1 and table_name=:2"
	var iCount int64
	row := db.QueryRow(strSQL, strings.ToUpper(schemaName), strings.ToUpper(tname))
	if err := row.Scan(&iCount); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return false, err
	}
	return iCount > 0, nil
}
func removeColumnsSQL(tabName string, cols []string) []string {
	return []string{fmt.Sprintf("ALTER table %s drop(%s)", tabName, strings.Join(cols, ","))}
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
	return []string{fmt.Sprintf("create %sindex %si%s%s on %s(%s)", ustr, schemaName, tname, colName, tableName, colName)}
}

func dropColumnIndexSQL(tableName, indexName string) []string {
	return []string{fmt.Sprintf("drop index %s", indexName)}
}
