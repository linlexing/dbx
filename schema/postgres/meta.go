package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

const driverName = "postgres"
const driverNamePgx = "pgx"
const driverNameGauss = "opengauss"
const driverNamePgxGauss = "pgx-opengauss"
const driverNamePgxUnify = "pgx-un"

type meta struct {
	driverName string
}

// 是否高斯数据库
func (m *meta) isGauss() bool {
	switch m.driverName {
	case driverNameGauss, driverNamePgxGauss:
		return true
	}
	return false
}

func init() {
	schema.Register(driverName, &meta{driverName})
	schema.Register(driverNamePgx, &meta{driverNamePgx})
	schema.Register(driverNameGauss, &meta{driverNameGauss})
	schema.Register(driverNamePgxGauss, &meta{driverNamePgxGauss})
	schema.Register(driverNamePgxUnify, &meta{driverNamePgxUnify})
}

// 执行create table as select语句
func (m *meta) CreateTableAsSQL(db common.DB, tableName, strSQL string,
	param []interface{}, pks []string) ([]string, error) {
	if m.isGauss() {
		return []string{fmt.Sprintf("CREATE TABLE %s DISTRIBUTE BY HASH(%s) as %s",
			tableName, strings.Join(pks, ","), strSQL),
			fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ",")),
		}, nil

	} else {
		return []string{fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSQL),
			fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ",")),
		}, nil
	}
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
	err = rows.Err()
	sort.Strings(names)
	return
}

/*
// 性能提升一倍，用处不大,！！！不能用，因为在事务中，只要发生一次error，整个事务就被污染，必须rollback
func (m *meta) TableExists(db common.DB, tabName string) (bool, error) {
	var str sql.NullString
	if err := db.QueryRow("select '" + tabName + "'::regclass").Scan(&str); err != nil {
		// if err := db.QueryRow("select to_regclass('" + tabName + "')").Scan(&str); err != nil {
		// if err := db.QueryRow("select '1' from " + tabName + " where 1=2").Scan(&str); err != nil {
		return false, nil
	}
	return true, nil
}*/

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
func (m *meta) CreateTableSQL(db common.DB, tab *schema.Table) (rev []string, _ error) {
	cols := []string{}
	for _, v := range tab.Columns {
		cols = append(cols, dbDefine(v))
	}
	outBuf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(outBuf, 0, 0, 1, ' ', 0)
	if len(tab.PrimaryKeys) > 0 {
		if m.isGauss() {
			fmt.Fprintf(w, "CREATE TABLE %s(\n  %s,\n  CONSTRAINT %s_PKEY PRIMARY KEY(%s)\n) DISTRIBUTE BY HASH(%s)",
				tab.FullName(), strings.Join(cols, ",\n  "), tab.Name, strings.Join(tab.PrimaryKeys, ","),
				strings.Join(tab.PrimaryKeys, ","))
		} else {
			fmt.Fprintf(w, "CREATE TABLE %s(\n  %s,\n  CONSTRAINT %s_PKEY PRIMARY KEY(%s)\n)",
				tab.FullName(), strings.Join(cols, ",\n  "), tab.Name, strings.Join(tab.PrimaryKeys, ","))
		}
		w.Flush()
		rev = append(rev, outBuf.String())
	} else {
		fmt.Fprintf(w, "CREATE TABLE %s(\n  %s\n)",
			tab.FullName(), strings.Join(cols, ",\n  "))
		w.Flush()
		rev = append(rev, outBuf.String())
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
	return []string{fmt.Sprintf("DROP INDEX IF EXISTS %s", indexName)}, nil

}

func (m *meta) CreateIndexIfNotExistsSQL(db common.DB, unique bool, indexName, tableName, express string) ([]string, error) {
	idx := "INDEX"
	if unique {
		idx = "UNIQUE INDEX"
	}
	str := fmt.Sprintf("CREATE %s %s ON %s(%s)", idx, strings.ToLower(indexName), tableName, express)
	return []string{fmt.Sprintf(
		`DO $BLOCK$
	BEGIN
		BEGIN
			%s;
		EXCEPTION
			WHEN duplicate_table THEN
				RAISE NOTICE 'index ''index_name '' on table_name already exists, skipping';
			WHEN OTHERS THEN 
				IF SQLERRM LIKE '%%already exists%%' THEN
					RAISE NOTICE 'index ''index_name '' on table_name already exists, skipping';
				ELSE
					RAISE;
				END IF;
			END;
		END;
$BLOCK$;`, str)}, nil
}

func (m *meta) CreateSchemaSQL(db common.DB, dbInfo schema.DataBaseInfo) ([]string, error) {
	return []string{
		fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s'", dbInfo.UserName, dbInfo.PassWord),
		fmt.Sprintf("CREATE SCHEMA %s AUTHORIZATION %s", dbInfo.DBName, dbInfo.UserName)}, nil
}

func (m *meta) DropSchemaSQL(db common.DB, dbInfo schema.DataBaseInfo) ([]string, error) {
	return []string{
		fmt.Sprintf("DROP SCHEMA %s CASCADE", dbInfo.DBName),
		fmt.Sprintf("DROP USER %s", dbInfo.UserName)}, nil
}
