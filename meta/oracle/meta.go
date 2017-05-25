package mysql

import (
	"dbweb/lib/safe"
	"fmt"
	"sort"
	"strings"

	"github.com/linlexing/dbx"
)

type meta struct {
}

func (m *meta) IsNull() string {
	return "nvl"
}

//执行create table as select语句
func (m *meta) CreateTableAs(db dbx.DB, tableName, strSql string, pks []string) error {
	s := fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSql)
	if _, err := db.Exec(s); err != nil {
		return dbx.NewSQLError(s, nil, err)
	}
	s = fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ","))
	if _, err := db.Exec(s); err != nil {
		return dbx.NewSQLError(s, nil, err)
	}
	return nil
}
func (m *meta) DropColumnIndex(db dbx.DB, tableName, indexName string) error {
	strSQL := fmt.Sprintf("drop index %s", indexName)
	if _, err := db.Exec(strSQL); err != nil {
		return dbx.NewSQLError(strSQL, nil, err)
	}
	return nil
}

//CreateColumnIndex 新增单字段索引
func (m *meta) CreateColumnIndex(db dbx.DB, tableName, colName string) error {
	ns := strings.Split(tableName, ".")
	schema := ""
	tname := ""
	if len(ns) > 1 {
		schema = ns[0] + "."
		tname = ns[1]
	} else {
		tname = tableName
	}
	//这里会有问题，如果表名和字段名比较长就会出错
	strSQL := fmt.Sprintf("create index %si%s%s on %s(%s)", schema, tname, colName, tableName, colName)
	if _, err := db.Exec(strSQL); err != nil {
		return dbx.NewSQLError(strSQL, nil, err)
	}
	return nil
}

func (m *meta) TableNames(db dbx.DB) (names []string, err error) {
	strSQL := "SELECT table_name FROM user_tables"
	names = []string{}
	if err = db.Select(&names, strSQL); err != nil {
		return
	}
	sort.Strings(names)
	return
}
func (m *meta) RemoveColumns(db dbx.DB, tabName string, cols []string) error {
	strSQL := fmt.Sprintf("ALTER table %s drop(%s)", tabName, strings.Join(cols, ","))
	if _, err := db.Exec(strSQL); err != nil {
		return dbx.NewSQLError(strSQL, nil, err)
	}
	return nil
}
func (m *meta) TableExists(db dbx.DB, tabName string) (bool, error) {
	schema := ""
	ns := strings.Split(tabName, ".")
	tname := ""
	if len(ns) > 1 {
		schema = ns[0]
		tname = ns[1]
	} else {
		tname = tabName
	}
	if len(schema) == 0 {
		schema = safe.String(dbx.MustGetSqlFun(db, "select user from dual", nil))
	}
	strSQL := fmt.Sprintf("SELECT count(*) FROM all_tables where owner='%s' and table_name=:tname", schema)
	var iCount int64
	p := map[string]interface{}{"tname": strings.ToUpper(tname)}
	if err := dbx.NameGet(db, &iCount, strSQL, p); err != nil {
		return false, dbx.NewSQLError(strSQL, p, err)
	}
	return iCount > 0, nil
}
func (m *meta) TableRename(db dbx.DB, oldName, newName string) error {
	strSQL := fmt.Sprintf("rename table %s TO %s", oldName, newName)
	if _, err := db.Exec(strSQL); err != nil {
		return dbx.NewSQLError(strSQL, nil, err)
	}
	return nil
}

func init() {
	dbx.RegisterMeta("oci8", new(meta))
}
