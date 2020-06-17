package sqlite

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

//GetTempTableName 获取一个临时表名
func getTempTableName(db common.DB, prev string) (string, error) {
	if len(prev) == 0 {
		return "", fmt.Errorf("prev can't empty")
	}
	//确定名称
	tableName := ""
	rand.Seed(time.Now().UnixNano())
	bys := make([]byte, 4)
	icount := 0
	for {
		binary.BigEndian.PutUint32(bys, rand.Uint32())
		tableName = fmt.Sprintf("%s%X", prev, bys)
		if exists, err := tableExists(db, tableName); err != nil {
			return "", err
		} else if !exists {
			break
		}
		icount++
		if icount > 100 {
			return "", fmt.Errorf("find table name too much")
		}
	}
	return tableName, nil
}
func tableRenameSQL(oldName string, newName string) []string {
	return []string{fmt.Sprintf("ALTER table %s RENAME TO %s", oldName, newName)}
}
func tableExists(db common.DB, tabName string) (bool, error) {

	strSQL := "SELECT count(*) FROM sqlite_master WHERE type='table' AND name=:tname"
	var iCount int64

	if err := db.QueryRow(strSQL, tabName).Scan(&iCount); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return false, err
	}

	return iCount > 0, nil
}

//createColumnIndex 新增单字段索引
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

func dbType(dataType schema.DataType, maxLength int) string {
	switch dataType {

	case schema.TypeBytea:
		return "BLOB"
	case schema.TypeDatetime:
		return "DATE"
	case schema.TypeFloat:
		return "REAL"
	case schema.TypeInt:
		return "INTEGER"
	case schema.TypeString:
		if maxLength <= 0 {
			return "TEXT"
		}
		return fmt.Sprintf("TEXT(%d)", maxLength)

	}

	panic("not impl DBType")

}
func SqliteType(typeName string) schema.DataType {
	a, _ := sqliteType(typeName)
	return a
}
func sqliteType(typeName string) (schema.DataType, int) {
	/*
		<1> 如果声明类型包含”INT”字符串，那么这个列被赋予INTEGER近似
		<2> 如果这个列的声明类型包含”CHAR”，”CLOB”，或者”TEXT”中的任意一个，那么这个列就有了TEXT近似。注意类型VARCHAR包含了”CHAR”字符串，那么也就被赋予了TEXT近似
		<3> 如果列的声明类型中包含了字符串”BLOB”或者没有为其声明类型，这个列被赋予NONE近似
		<4> 其他的情况，列被赋予NUMERIC近似
	*/
	typeName = strings.ToUpper(typeName)
	if strings.Contains(typeName, "INT") {
		return schema.TypeInt, 0
	}
	if strings.Contains(typeName, "CHAR") ||
		strings.Contains(typeName, "CLOB") ||
		strings.Contains(typeName, "TEXT") {
		length := "-1"
		if ts := strings.Split(typeName, "("); len(ts) > 1 {
			length = ts[1]
			length = length[:len(length)-1]
		}
		i, err := strconv.ParseInt(length, 10, 64)
		if err != nil {
			panic(err)
		}
		return schema.TypeString, int(i)
	}
	if strings.Contains(typeName, "BLOB") || strings.Contains(typeName, "BYTEA") ||
		len(typeName) == 0 {
		return schema.TypeBytea, 0
	}
	if strings.Contains(typeName, "DATE") || strings.Contains(typeName, "TIME") {
		return schema.TypeDatetime, 0
	}
	return schema.TypeFloat, 0
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
