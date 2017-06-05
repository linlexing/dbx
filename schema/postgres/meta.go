package postgres

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

const driverName = "postgres"

type meta struct {
}

func init() {
	schema.Register(driverName, new(meta))
}

//执行create table as select语句
func (m *meta) CreateTableAs(db common.DB, tableName, strSQL string, pks []string) error {
	s := fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSQL)
	if _, err := db.Exec(s); err != nil {
		err = common.NewSQLError(err, s)
		log.Println(err)
		return err
	}
	s = fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ","))
	if _, err := db.Exec(s); err != nil {
		err = common.NewSQLError(err, s)
		log.Println(err)
		return err
	}
	return nil
}

func (m *meta) TableNames(db common.DB) (names []string, err error) {
	strSQL := "SELECT table_name FROM information_schema.tables WHERE table_schema = current_schema()"
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
		if err = rows.Scan(&name); err != nil { //从rows中取值是用&name
			return nil, err
		}
		names = append(names, name)
	}
	sort.Strings(names)
	return
}

func (m *meta) TableExists(db common.DB, tabName string) (bool, error) {
	schemaName := ""
	ns := strings.Split(tabName, ".")
	tname := ""
	if len(ns) > 1 {
		schemaName = ns[0]
		tname = ns[1]
	} else {
		tname = tabName
	}
	if len(schemaName) == 0 {
		strSQL := "select current_schema()"
		if err := db.QueryRow(strSQL).Scan(&schemaName); err != nil {
			err = common.NewSQLError(err, strSQL)
			log.Println(err)
			return false, err
		}
	}

	strSQL :=
		"SELECT count(*) FROM information_schema.tables WHERE table_schema ilike $1 and table_name ilike $2"
	var iCount int64
	if err := db.QueryRow(strSQL, schemaName, tname).Scan(&iCount); err != nil {
		err = common.NewSQLError(err, strSQL, schemaName, tname)
		log.Println(err)
		return false, err
	}

	return iCount > 0, nil
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
func (m *meta) DropIndexIfExists(db common.DB, indexName, tableName string) error {
	strSQL := fmt.Sprintf("drop index if exists %s", indexName)
	_, err := db.Exec(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
	}
	return err

}

func (m *meta) CreateIndexIfNotExists(db common.DB, indexName, tableName, express string) error {
	var strSQL string
	strSQL = fmt.Sprintf("create index if not exists %s on %s(%s)", indexName, tableName, express)
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	return nil
}
