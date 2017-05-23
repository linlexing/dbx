package mysql

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx"
)

type meta struct {
}

func (m *meta) IsNull() string {
	return "ifnull"
}

//执行create table as select语句
func (m *meta) CreateTableAs(db dbx.DB, tableName, strSql string, pks []string) error {
	tmpTableName := "t_"+tableName
	//由于sqlite不支持alter table，所以需要先创建表，然后insert
	s := fmt.Sprintf("CREATE TABLE %s as select * from (%s) limit 0", tmpTableName, strSql)
	if _, err := db.Exec(s); err != nil {
		return dbx.NewSQLError(s, nil, err)
	}

	strSQL := "SELECT sql FROM sqlite_master WHERE type = 'table' AND name = :tname"
	param := map[string]interface{}{"tname": tmpTableName}
		rev, err := dbx.GetSqlFun(db, strSQL, param)
	if err != nil {
		return err
	}
	createScript := strings.Replace(string(rev),"CREATE TABLE "+ tableName,
	"CREATE TABLE "+ tableName
	if _, err := db.Exec(s); err != nil {
		return dbx.NewSQLError(s, nil, err)
	}
	s = fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ","))
	if _, err := db.Exec(s); err != nil {
		return dbx.NewSQLError(s, nil, err)
	}
	return nil
}

func init() {
	dbx.RegisterMeta("sqlite3", new(meta))
}
