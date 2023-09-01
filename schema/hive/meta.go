package mysql

import (
	"errors"
	"log"
	"sort"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

const driverName = "hive"

type meta struct {
}

func init() {
	schema.Register(driverName, new(meta))
}
func (m *meta) TablePK(db common.DB, tableName string) ([]string, error) {
	tab, err := m.OpenTable(db, tableName)
	if err != nil {
		return nil, err
	}
	return tab.PrimaryKeys, nil
}

// CreateTableAsSQL 生成create table as select语句
func (m *meta) CreateTableAsSQL(db common.DB, tableName, strSQL string, param []interface{},
	pks []string) ([]string, error) {
	panic("not impl")
}
func (m *meta) ChangeTableSQL(db common.DB, change *schema.TableSchemaChange) ([]string, error) {
	panic("not impl")
}
func (m *meta) TableEmpty(db common.DB, tableName string) (bool, error) {
	panic("not impl")
}
func (m *meta) TableNames(db common.DB) (names []string, err error) {
	strSQL := "show tables"
	names = []string{}
	rows, err := db.Query(strSQL)
	if err != nil {
		log.Println(strSQL)
		return nil, common.NewSQLError(err, strSQL)
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		log.Println(strSQL)
		return nil, common.NewSQLError(err, strSQL)
	}
	var name string
	var dbname string
	var isTemp string
	for rows.Next() {
		if len(cols) == 1 {
			if err = rows.Scan(&name); err != nil {
				log.Println(strSQL)
				return nil, common.NewSQLError(err, strSQL)
			}
			names = append(names, name)
		} else {
			if err = rows.Scan(&dbname, &name, &isTemp); err != nil {
				log.Println(strSQL)
				return nil, common.NewSQLError(err, strSQL)
			}
			names = append(names, name)
		}
	}
	err = rows.Err()
	sort.Strings(names)
	return
}
func signString(str string) string {
	return "'" + strings.ReplaceAll(str, "'", "''") + "'"
}
func (m *meta) TableExists(db common.DB, tabName string) (bool, error) {
	strSQL := "show tables like " + signString(tabName)
	rows, err := db.Query(strSQL)
	if err != nil {
		log.Println(strSQL)
		return false, common.NewSQLError(err, strSQL)
	}
	defer rows.Close()
	return rows.Next(), nil

}

func (m *meta) CreateTableSQL(db common.DB, tab *schema.Table) ([]string, error) {
	panic("not impl")
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
