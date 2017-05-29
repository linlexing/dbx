package mysql

import (
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

func (m *meta) IsNull() string {
	return "ifnull"
}
func init() {
	schema.Register(driverName, new(meta))
}

//CreateTableAs 执行create table as select语句
func (m *meta) CreateTableAs(db common.DB, tableName, strSQL string, pks []string) error {
	s := fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSQL)
	if _, err := db.Exec(s); err != nil {
		log.Println(s)
		return common.NewSQLError(err, s)

	}
	s = fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ","))
	if _, err := db.Exec(s); err != nil {
		log.Println(s)
		return common.NewSQLError(err, s)

	}
	return nil
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
		log.Println(strSQL)
		return common.NewSQLError(err, strSQL)
	}
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
