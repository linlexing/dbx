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
func (m *meta) CreateTableAs(db common.DB, tableName, strSQL string, pks []string) error {
	tmpTableName, err := getTempTableName(db, "t_")
	if err != nil {
		return err
	}
	//由于sqlite不支持alter table，所以需要先创建表，然后insert
	s := fmt.Sprintf("CREATE TABLE %s as select * from (%s) limit 0", tmpTableName, strSQL)
	if _, err := db.Exec(s); err != nil {
		err = common.NewSQLError(err, s)
		log.Println(err)
		return err
	}
	s = "SELECT sql FROM sqlite_master WHERE type = 'table' AND name = :tname"

	var createSQL string

	if err := db.QueryRow(s, tmpTableName).Scan(&createSQL); err != nil {
		err = common.NewSQLError(err, s, tmpTableName)
		log.Println(err)
		return err
	}
	s = strings.Replace(createSQL,
		"CREATE TABLE "+tmpTableName,
		"CREATE TABLE "+tableName, 1)
	//去除尾部括号和换行
	s = strings.TrimSpace(s[:len(s)-1])
	//加上主键定义
	s = s + fmt.Sprintf(",\n  CONSTRAINT %s_pkey primary key(%s)\n)",
		tableName, strings.Join(pks, ","))
	if _, err := db.Exec(s); err != nil {
		err = common.NewSQLError(err, s)
		log.Println(err)
		return err
	}

	s = fmt.Sprintf("insert into %s %s", tableName, strSQL)
	if _, err := db.Exec(s); err != nil {
		err = common.NewSQLError(err, s)
		log.Println(err)
		return err
	}
	s = fmt.Sprintf("drop table %s", tmpTableName)
	if _, err := db.Exec(s); err != nil {
		err = common.NewSQLError(err, s)
		log.Println(err)
		return err
	}
	return nil
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

	sort.Strings(names)
	return
}

func (m *meta) TableExists(db common.DB, tabName string) (bool, error) {
	return tableExists(db, tabName)
}

func (m *meta) CreateTable(db common.DB, tab *schema.Table) error {
	cols := []string{}
	for _, v := range tab.Columns {
		cols = append(cols, dbDefine(v))
	}
	var strSQL string
	if len(tab.PrimaryKeys) > 0 {
		strSQL = fmt.Sprintf(
			"CREATE TABLE %s(\n%s,\nCONSTRAINT %s_pkey PRIMARY KEY(%s)\n)",
			tab.FullName(), strings.Join(cols, ",\n"), tab.Name, strings.Join(tab.PrimaryKeys, ","))
	} else {
		strSQL = fmt.Sprintf(
			"CREATE TABLE %s(\n%s\n)",
			tab.FullName(), strings.Join(cols, ",\n"))
	}
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	log.Println(strSQL)
	//最后处理索引
	for _, col := range tab.Columns {
		if col.Index {
			if err := createColumnIndex(db, tab.FullName(), col.Name); err != nil {
				return err
			}
		}
	}
	return nil
}
