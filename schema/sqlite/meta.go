package sqlite

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

const driverName = "sqlite3"

type meta struct {
}

func init() {
	schema.Register(driverName, new(meta))
}

//CreateTableAs 执行create table as select语句
func (m *meta) CreateTableAsSQL(db common.DB, tableName, strSQL string, pks []string) (rev []string, err error) {
	tmpTableName, err := getTempTableName(db, "t_")
	if err != nil {
		return
	}
	//由于sqlite不支持alter table，所以需要先创建表，然后insert
	s := fmt.Sprintf("CREATE TABLE %s as select * from (%s) limit 0", tmpTableName, strSQL)
	if _, err := db.Exec(s); err != nil {
		err = common.NewSQLError(err, s)
		log.Println(err)
		return nil, err
	}
	s = "SELECT sql FROM sqlite_master WHERE type = 'table' AND name = :tname"

	var createSQL string

	if err := db.QueryRow(s, tmpTableName).Scan(&createSQL); err != nil {
		err = common.NewSQLError(err, s, tmpTableName)
		log.Println(err)
		return nil, err
	}
	s = strings.Replace(createSQL,
		"CREATE TABLE "+tmpTableName,
		"CREATE TABLE "+tableName, 1)
	//去除尾部括号和换行
	s = strings.TrimSpace(s[:len(s)-1])
	//加上主键定义
	s = s + fmt.Sprintf(",\n  CONSTRAINT %s_pkey primary key(%s)\n)",
		tableName, strings.Join(pks, ","))
	rev = []string{s, fmt.Sprintf("insert into %s %s", tableName, strSQL)}
	return
}

func (m *meta) TableNames(db common.DB) (names []string, err error) {
	strSQL := "SELECT name FROM sqlite_master WHERE type='table'"
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
	return tableExists(db, tabName)
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
		if col.Index {
			rev = append(rev, createColumnIndexSQL(tab.FullName(), col.Name)...)
		}
	}
	return
}

func (m *meta) DropIndexIfExistsSQL(db common.DB, indexName, tableName string) ([]string, error) {
	return []string{fmt.Sprintf("drop index if exists %s", indexName)}, nil
}

func (m *meta) CreateIndexIfNotExistsSQL(db common.DB, indexName, tableName, express string) ([]string, error) {
	return []string{fmt.Sprintf("create index if not exists %s on %s(%s)", indexName, tableName, express)}, nil
}
