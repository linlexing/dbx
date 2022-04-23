package mysql

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

const driverName = "mysql"

type meta struct {
}

func init() {
	schema.Register(driverName, new(meta))
}

//CreateTableAsSQL 生成create table as select语句
func (m *meta) CreateTableAsSQL(db common.DB, tableName, strSQL string, param []interface{},
	pks []string) ([]string, error) {
	return []string{
		fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSQL),
		fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ",")),
	}, nil
}
func (m *meta) TableEmpty(db common.DB, tableName string) (bool, error) {
	var a int
	if err := db.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s)",
		tableName)).Scan(&a); err != nil {
		return true, err
	}
	return a == 0, nil
}
func (m *meta) TableNames(db common.DB) (names []string, err error) {
	strSQL := "SELECT table_name FROM information_schema.tables WHERE table_schema = schema()"
	names = []string{}
	rows, err := db.Query(strSQL)
	if err != nil {
		log.Println(strSQL)
		return nil, common.NewSQLError(err, strSQL)
	}
	defer rows.Close()
	var name string
	for rows.Next() {
		if err = rows.Scan(&name); err != nil {
			return
		}
		names = append(names, name)
	}
	err = rows.Err()
	sort.Strings(names)
	return
}
func (m *meta) TableExists(db common.DB, tabName string) (bool, error) {
	return tableExists(db, tabName)
}

func (m *meta) CreateTableSQL(db common.DB, tab *schema.Table) ([]string, error) {
	rev := []string{}
	cols := []string{}
	for _, v := range tab.Columns {
		cols = append(cols, dbDefine(v))
	}
	if len(tab.PrimaryKeys) > 0 {
		rev = append(rev, fmt.Sprintf(
			"CREATE TABLE %s(\n%s,\nCONSTRAINT %s_pkey PRIMARY KEY(%s)\n)",
			tab.FullName(), strings.Join(cols, ",\n"), tab.Name, strings.Join(tab.PrimaryKeys, ",")))
	} else {
		rev = append(rev, fmt.Sprintf(
			"CREATE TABLE %s(\n%s\n)",
			tab.FullName(), strings.Join(cols, ",\n")))
	}

	//最后处理索引
	for _, col := range tab.Columns {
		if col.Index == schema.Index {
			rev = append(rev, createColumnIndexSQL(tab.FullName(), false, col.Name)...)
		} else if col.Index == schema.UniqueIndex {
			rev = append(rev, createColumnIndexSQL(tab.FullName(), true, col.Name)...)
		}
	}
	return rev, nil
}
func (m *meta) DropIndexIfExistsSQL(db common.DB, indexName, tableName string) ([]string, error) {
	panic(errors.New("not impl"))
}
func (m *meta) CreateIndexIfNotExistsSQL(db common.DB, unique bool, indexName, tableName, express string) ([]string, error) {
	panic(errors.New("not impl"))
}

// CreateSchemaSQL 建库
func (m *meta) CreateSchemaSQL(db common.DB, dbInfo schema.DataBaseInfo) ([]string, error) {
	// createDB := fmt.Sprintf("CREATE DATABASE %s", dbInfo.DBName)
	// createUser := fmt.Sprintf(`GRANT ALL PRIVILEGES ON %s.* TO %s@"%s" IDENTIFIED BY "%s"`,
	// 	dbInfo.DBName, dbInfo.UserName, "%", dbInfo.PassWord)
	// flush := fmt.Sprintf("FLUSH PRIVILEGES")
	// return []string{createDB, createUser, flush}, nil
	return []string{}, nil
}

// DropSchemaSQL 删库
func (m *meta) DropSchemaSQL(db common.DB, dbInfo schema.DataBaseInfo) ([]string, error) {

	return []string{}, nil
}
