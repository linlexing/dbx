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
func createColumnIndexSQL(tableName, colName string) []string {
	//这里会有问题，如果表名和字段名比较长就会出错
	return []string{fmt.Sprintf("create index on %s(%s)", tableName, colName)}
}

func removeColumnsSQL(tabName string, cols []string) []string {
	strList := []string{}
	for _, v := range cols {
		strList = append(strList, "DROP COLUMN "+v)
	}
	return []string{fmt.Sprintf("ALTER table %s %s", tabName, strings.Join(strList, ","))}
}

func tableRenameSQL(oldName string, newName string) []string {
	return []string{fmt.Sprintf("ALTER table %s RENAME TO %s", oldName, newName)}
}
func dropColumnIndexSQL(tableName, indexName string) []string {
	return []string{fmt.Sprintf("drop index %s", indexName)}
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

func dropTablePrimaryKeySQL(db common.DB, tableName string) ([]string, error) {
	strSQL := fmt.Sprintf(
		"select b.relname from  pg_index a inner join pg_class b on a.indexrelid =b.oid where indisprimary and indrelid='%s'::regclass",
		tableName)
	var pkCons string
	if err := db.QueryRow(strSQL).Scan(&pkCons); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	return []string{fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, pkCons)}, nil
}

//addTablePrimaryKey 新增主键
//只有当表不存在主键时，才可以新增主键
func addTablePrimaryKeySQL(tableName string, pks []string) []string {
	return []string{fmt.Sprintf("alter table %s add primary key(%s)", tableName, strings.Join(pks, ","))}
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
	return fmt.Sprintf("%s\t%s%s", c.Name, colDBType(c), nullStr)
}

func dbDefineNull(c *schema.Column) string {
	nullStr := " NULL"
	if !c.Null {
		nullStr = " NOT NULL"
	}
	return fmt.Sprintf("%s\t%s%s", c.Name, colDBType(c), nullStr)
}
