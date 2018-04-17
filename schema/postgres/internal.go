package postgres

import (
	"fmt"
	"log"
	"strings"

	"database/sql"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

//createColumnIndex 新增单字段索引
func createColumnIndex(db common.DB, tableName, colName string) error {
	//这里会有问题，如果表名和字段名比较长就会出错
	strSQL := fmt.Sprintf("create index on %s(%s)", tableName, colName)
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	return nil
}

func removeColumns(db common.DB, tabName string, cols []string) error {
	var strSQL string
	strList := []string{}
	for _, v := range cols {
		strList = append(strList, "DROP COLUMN "+v)
	}
	strSQL = fmt.Sprintf("ALTER table %s %s", tabName, strings.Join(strList, ","))
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	return nil
}

func tableRename(db common.DB, oldName string, newName string) error {

	strSQL := fmt.Sprintf("ALTER table %s RENAME TO %s", oldName, newName)
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
func dbType(dataType schema.DataType, maxLength int) string {
	switch dataType {

	case schema.TypeBytea:
		return "bytea"
	case schema.TypeDatetime:
		return "timestamp without time zone"
	case schema.TypeFloat:
		return "double precision"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeString:
		if maxLength <= 0 {
			return "text"
		}
		return fmt.Sprintf("character varying(%d)", maxLength)

	}
	panic("not impl DBType")

}

func dropTablePrimaryKey(db common.DB, tableName string) error {
	strSQL := fmt.Sprintf(
		"select b.relname from  pg_index a inner join pg_class b on a.indexrelid =b.oid where indisprimary and indrelid='%s'::regclass",
		tableName)
	var pkCons string
	if err := db.QueryRow(strSQL).Scan(&pkCons); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		err = common.NewSQLError(err, strSQL)
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
//只有当表不存在主键时，才可以新增主键
func addTablePrimaryKey(db common.DB, tableName string, pks []string) error {
	strSQL := fmt.Sprintf("alter table %s add primary key(%s)", tableName, strings.Join(pks, ","))
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
