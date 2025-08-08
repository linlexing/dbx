package duckdb

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

const driverName = "duckdb"

type meta struct {
}

func init() {
	schema.Register(driverName, new(meta))
}

// CreateTableAs 执行create table as select语句
// todo:不支持date类型，会转换成num类型，需要整改
func (m *meta) CreateTableAsSQL(db common.DB, tableName, strSQL string,
	param []interface{}, pks []string) (rev []string, err error) {
	return []string{
		fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSQL),
		fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ",")),
	}, nil
}
func (m *meta) TableEmpty(db common.DB, tableName string) (bool, error) {
	var a int
	if err := db.QueryRow(fmt.Sprintf("select 1 where exists (select * from %s)",
		tableName)).Scan(&a); err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
func (m *meta) TableNames(db common.DB) (names []string, err error) {
	strSQL := "SELECT table_name FROM information_schema.tables where table_schema='main' "
	names = []string{}
	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	err = rows.Err()
	sort.Strings(names)
	return
}

func (m *meta) TableExists(db common.DB, tabName string) (bool, error) {
	strSQL := `SELECT count(*) FROM information_schema.tables 
                  where table_schema='main' and table_name = $1;`
	var iCount int64
	if err := db.QueryRow(strSQL, tabName).Scan(&iCount); err != nil {
		err = common.NewSQLError(err, strSQL, tabName)
		log.Println(err)
		return false, err
	}
	return iCount > 0, nil
}

func (m *meta) CreateTableSQL(db common.DB, tab *schema.Table) (rev []string, err error) {
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
	return
}

func (m *meta) DropIndexIfExistsSQL(db common.DB, indexName, tableName string) ([]string, error) {
	return []string{fmt.Sprintf("drop index if exists %s", indexName)}, nil
}

func (m *meta) CreateIndexIfNotExistsSQL(db common.DB, unique bool, indexName, tableName, express string) ([]string, error) {
	idx := "index"
	if unique {
		idx = "unique index"
	}
	return []string{fmt.Sprintf("create %s if not exists %s on %s(%s)", idx, indexName, tableName, express)}, nil
}

/*
open数据库时 若数据库不存在则会自动创建
*/
func (m *meta) CreateSchemaSQL(db common.DB, dbInfo schema.DataBaseInfo) ([]string, error) {
	return []string{}, nil
}
func (m *meta) DropSchemaSQL(db common.DB, dbInfo schema.DataBaseInfo) ([]string, error) {
	return []string{}, nil
}
