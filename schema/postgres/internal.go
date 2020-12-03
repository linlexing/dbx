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
func createColumnIndexSQL(tableName string, unique bool, colName string) []string {
	ustr := ""
	if unique {
		ustr = "UNIQUE "
	}
	//这里会有问题，如果表名和字段名比较长就会出错
	return []string{fmt.Sprintf("CREATE %sINDEX ON %s(%s)", ustr, tableName, colName)}
}

func removeColumnsSQL(tabName string, cols []string) []string {
	strList := []string{}
	for _, v := range cols {
		strList = append(strList, "DROP COLUMN "+v)
	}
	return []string{fmt.Sprintf("ALTER TABLE %s %s", tabName, strings.Join(strList, ","))}
}

func tableRenameSQL(oldName string, newName string) []string {
	return []string{fmt.Sprintf("ALTER TABLE %s RENAME TO %s", oldName, strings.ToLower(newName))}
}
func dropColumnIndexSQL(tableName, indexName string) []string {
	return []string{fmt.Sprintf("DROP INDEX %s", indexName)}
}
func dbType(dataType schema.DataType, maxLength int) string {
	switch dataType {

	case schema.TypeBytea:
		return "bytea"
	case schema.TypeDatetime:
		//为了支持多国语言，必须要考虑时区的问题，时间需要保存时区
		return "timestamp with time zone"
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
	return []string{fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ","))}
}
func colDBType(c *schema.Column) string {
	if (len(c.FetchDriver) == 0 || strings.ToLower(c.FetchDriver) == strings.ToLower(driverName)) && len(c.TrueType) > 0 {
		return c.TrueType
	}
	return dbType(c.Type, c.MaxLength)
}

func dbDefine(c *schema.Column) string {
	nullStr := ""
	if !c.Null {
		nullStr = "\tNOT NULL"
	}
	//postgresql 中，中英文混合的字段名，默认会转换成小写字母，不加转换会出现找不到列的错误
	return fmt.Sprintf("%s\t%s%s", strings.ToLower(c.Name), colDBType(c), nullStr)
}

func dbDefineNull(c *schema.Column) string {
	nullStr := "\tNULL"
	if !c.Null {
		nullStr = "\tNOT NULL"
	}
	//postgresql 中，中英文混合的字段名，默认会转换成小写字母，不加转换会出现找不到列的错误
	return fmt.Sprintf("%s\t%s%s", strings.ToLower(c.Name), colDBType(c), nullStr)
}
