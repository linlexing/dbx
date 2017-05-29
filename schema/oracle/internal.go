package oracle

import (
	"fmt"
	"log"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

func dbType(dataType schema.DataType, maxLength int) string {
	switch dataType {
	case schema.TypeBytea:
		return "BLOB"
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
func dropTablePrimaryKey(db common.DB, tableName string) error {
	ns := strings.Split(tableName, ".")
	var strSQL string
	var params []interface{}
	if len(ns) > 1 {
		strSQL =
			"select constraint_name from ALL_CONSTRAINTS where owner = :schema and table_name =:table and constraint_type='P'"
		params = []interface{}{
			strings.ToUpper(ns[0]),
			strings.ToUpper(ns[1])}
	} else {
		strSQL =
			"select constraint_name from user_CONSTRAINTS where table_name =:table and constraint_type='P'"
		params = []interface{}{
			strings.ToUpper(tableName)}
	}
	var pkCons string
	row := db.QueryRow(strSQL, params...)
	if err := row.Scan(&pkCons); err != nil {
		err = common.NewSQLError(err, strSQL, params...)
		log.Println(err)
		return err
	}
	strSQL = fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, pkCons)
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}

	return nil
}

//addTablePrimaryKey 新增主键
func addTablePrimaryKey(db common.DB, tableName string, pks []string) error {
	ns := strings.Split(tableName, ".")
	var clearTableName string
	if len(ns) > 1 {
		clearTableName = ns[1]
	} else {
		clearTableName = tableName
	}
	strSQL := fmt.Sprintf("alter table %s add constraint %s_pk primary key(%s)", tableName, clearTableName, strings.Join(pks, ","))
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}

	return nil
}
func colDBType(c *schema.Column) string {
	if c.FetchDriver == driverName && len(c.TrueType) > 0 {
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

//tableRename 处理表改名，旧名可以是多个，任意一个对上就改名，如果都没有存在，则不处理，也不返回出错
func tableRename(db common.DB, oldName string, newName string) error {

	strSQL := fmt.Sprintf("rename table %s TO %s", oldName, newName)
	if _, err := db.Exec(strSQL); err != nil {
		log.Println(strSQL)
		return common.NewSQLError(err, strSQL)
	}

	return nil
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
	strSQL := "SELECT count(*) FROM all_tables where owner=:schema and table_name=:tname"
	var iCount int64
	row := db.QueryRow(strSQL, schemaName, tname)
	if err := row.Scan(&iCount); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return false, err
	}
	return iCount > 0, nil
}
func removeColumns(db common.DB, tabName string, cols []string) error {
	strSQL := fmt.Sprintf("ALTER table %s drop(%s)", tabName, strings.Join(cols, ","))
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	return nil
}

//createColumnIndex 新增单字段索引
func createColumnIndex(db common.DB, tableName, colName string) error {
	ns := strings.Split(tableName, ".")
	schemaName := ""
	tname := ""
	if len(ns) > 1 {
		schemaName = ns[0] + "."
		tname = ns[1]
	} else {
		tname = tableName
	}
	//这里会有问题，如果表名和字段名比较长就会出错
	strSQL := fmt.Sprintf("create index %si%s%s on %s(%s)", schemaName, tname, colName, tableName, colName)
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	return nil
}

func dropColumnIndex(db common.DB, tableName, indexName string) error {
	strSQL := fmt.Sprintf("drop index %s", indexName)
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	return nil
}
