package sqlite

import (
	"fmt"
	"sort"
	"strings"

	"database/sql"

	"github.com/linlexing/dbx"
)

type tableColumn struct {
	CID       int
	Name      string
	Type      string
	NotNull   int
	DfltValue sql.NullString `db:"dflt_value"`
	PK        int
}
type meta struct {
}

func (m *meta) IsNull() string {
	return "ifnull"
}
func init() {
	dbx.RegisterMeta("sqlite3", new(meta))
}

//CreateTableAs 执行create table as select语句
func (m *meta) CreateTableAs(db dbx.DB, tableName, strSQL string, pks []string) error {
	tmpTableName, err := dbx.GetTempTableName(db, "t_")
	if err != nil {
		return err
	}
	//由于sqlite不支持alter table，所以需要先创建表，然后insert
	s := fmt.Sprintf("CREATE TABLE %s as select * from (%s) limit 0", tmpTableName, strSQL)
	if _, err := db.Exec(s); err != nil {
		return dbx.NewSQLError(s, nil, err)
	}
	s = "SELECT sql FROM sqlite_master WHERE type = 'table' AND name = :tname"
	param := map[string]interface{}{"tname": tmpTableName}
	var createSQL string
	if err := dbx.NameGet(db, &createSQL, s, param); err != nil {
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
		return dbx.NewSQLError(s, nil, err)
	}

	s = fmt.Sprintf("insert into %s %s", tableName, strSQL)
	if _, err := db.Exec(s); err != nil {
		return dbx.NewSQLError(s, nil, err)
	}
	s = fmt.Sprintf("drop table %s", tmpTableName)
	if _, err := db.Exec(s); err != nil {
		return dbx.NewSQLError(s, nil, err)
	}
	return nil
}

func (m *meta) TableNames(db dbx.DB) (names []string, err error) {
	strSQL := "SELECT name FROM sqlite_master WHERE type='table'"
	names = []string{}
	if err = db.Select(&names, strSQL); err != nil {
		return
	}
	sort.Strings(names)
	return
}
func (m *meta) RemoveColumns(db dbx.DB, tabName string, cols []string) error {
	strSQL := fmt.Sprintf("PRAGMA table_info(%s)", tabName)
	tabCols := []tableColumn{}
	if err := db.Select(&tabCols, strSQL); err != nil {
		println(err)
		return dbx.NewSQLError(strSQL, nil, err)
	}
	pkCols := []string{}
	copyCols := []string{}
	sort.Strings(cols)
	//整理出要复制的列和主键列
	for _, one := range tabCols {
		pk := false
		if one.PK > 0 {
			pkCols = append(pkCols, one.Name)
			pk = true
		}
		if idx := sort.SearchStrings(cols, one.Name); idx < len(cols) &&
			cols[idx] == one.Name {
			if pk {
				return fmt.Errorf("pk %s can't remove", one.Name)
			}
		} else { //找不到则复制
			copyCols = append(copyCols, one.Name)
		}
	}
	//复制新表
	newTable, err := dbx.GetTempTableName(db, "cp_")
	if err != nil {
		return err
	}
	if err := m.CreateTableAs(db, newTable,
		fmt.Sprintf("select %s from %s", strings.Join(copyCols, ","), tabName),
		pkCols); err != nil {
		return err
	}
	//删除旧表
	strSQL = "drop table " + tabName
	if _, err := db.Exec(strSQL); err != nil {
		return dbx.NewSQLError(strSQL, nil, err)
	}
	//新表改名
	return m.TableRename(db, newTable, tabName)
}
func (m *meta) TableExists(db dbx.DB, tabName string) (bool, error) {

	strSQL := "SELECT count(*) FROM sqlite_master WHERE type='table' AND name=:tname"
	var iCount int64
	p := map[string]interface{}{"tname": strings.ToUpper(tabName)}
	if err := dbx.NameGet(db, &iCount, strSQL, p); err != nil {
		return false, dbx.NewSQLError(strSQL, p, err)
	}
	return iCount > 0, nil
}

func (m *meta) TableRename(db dbx.DB, oldName, newName string) error {
	strSQL := fmt.Sprintf("ALTER table %s RENAME TO %s", oldName, newName)
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
func (m *meta) DropColumnIndex(db dbx.DB, tableName, indexName string) error {
	strSQL := fmt.Sprintf("drop index %s", indexName)
	if _, err := db.Exec(strSQL); err != nil {
		return dbx.NewSQLError(strSQL, nil, err)
	}
	return nil
}
