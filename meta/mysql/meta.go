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
func init() {
	dbx.RegisterMeta("mysql", new(meta))
}
