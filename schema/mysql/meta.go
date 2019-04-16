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
func (m *meta) CreateTableAsSQL(db common.DB, tableName, strSQL string, pks []string) ([]string, error) {
	return []string{
		fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSQL),
		fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ",")),
	}, nil
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
		if col.Index {
			rev = append(rev, createColumnIndexSQL(tab.FullName(), col.Name)...)
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
