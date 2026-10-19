package duckdb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

type tableColumn struct {
	CID       int
	Name      string
	Type      string
	NotNull   bool
	DfltValue sql.NullString `db:"dflt_value"`
	PK        bool
}
type tableIndex struct {
	Name    string
	Unique  bool
	Columns []string
}

func getParquetTableInfo(db common.DB, fileName string) ([]tableColumn, error) {
	strSQL := fmt.Sprintf(
		"select column_name,column_type from(DESCRIBE SELECT * FROM '%s')", fileName)
	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	rev := []tableColumn{}
	for rows.Next() {
		var colName, colType string
		if err = rows.Scan(&colName, &colType); err != nil {
			return nil, err
		}
		rev = append(rev, tableColumn{
			Name: colName,
			Type: colType,
		})
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return rev, nil
}
func getTableInfo(db common.DB, tableName string) ([]tableColumn, error) {
	strSQL := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	tabCols := []tableColumn{}
	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var col tableColumn
		if err = rows.Scan(
			&col.CID,
			&col.Name,
			&col.Type,
			&col.NotNull,
			&col.DfltValue,
			&col.PK); err != nil {
			return nil, err
		}
		tabCols = append(tabCols, col)
	}
	return tabCols, rows.Err()
}

// 只能获取到显式声明的索引 create index
func getTableIndex(db common.DB, tableName string) ([]tableIndex, error) {
	rev := []tableIndex{}
	strSQL := fmt.Sprintf("select index_name,is_unique,expressions from duckdb_indexes() where table_name = '%s'", tableName)
	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var idx tableIndex
		var expressions sql.NullString //格式为 [c,b] 或 ['"name"']
		if err = rows.Scan(
			&idx.Name,
			&idx.Unique,
			&expressions,
		); err != nil {
			return nil, err
		}
		if len(expressions.String) > 0 {
			//去掉[和 ]
			expressions.String = expressions.String[1 : len(expressions.String)-1]
			expressions.String = strings.ReplaceAll(expressions.String, "'", "")
			expressions.String = strings.ReplaceAll(expressions.String, "\"", "")
			expressions.String = strings.ReplaceAll(expressions.String, " ", "")
			idx.Columns = strings.Split(expressions.String, ",")
		}
		rev = append(rev, idx)
	}
	return rev, rows.Err()
}
func (m *meta) TablePK(db common.DB, tableName string) ([]string, error) {
	tab, err := m.OpenTable(db, tableName)
	if err != nil {
		return nil, err
	}
	return tab.PrimaryKeys, nil
}
func (m *meta) openTableParquet(db common.DB, fileName string) (*schema.Table, error) {
	tabCols, err := getParquetTableInfo(db, fileName)
	if err != nil {
		return nil, err
	}
	columns := []*schema.Column{}
	for _, col := range tabCols {
		c := &schema.Column{
			Name:        col.Name,
			FetchDriver: driverName,
			TrueType:    col.Type,
			Null:        !col.NotNull,
			Extended:    map[string]any{},
			MaxLength:   0,
		}
		c.Type = DuckDBType(col.Type)
		columns = append(columns, c)
	}
	bname := filepath.Base(fileName)
	bname = strings.TrimSuffix(bname, filepath.Ext(bname))
	t := schema.NewTable(bname)
	t.Columns = columns
	//获取主键
	var keys string
	strSQL := fmt.Sprintf(
		"SELECT value FROM parquet_kv_metadata('%s') where key='primary_keys'",
		fileName)
	if err := db.QueryRow(strSQL).Scan(&keys); err != nil {
		//没有主键，也可以返回
		if errors.Is(err, sql.ErrNoRows) {
			return t, nil
		}
		return nil, err
	}
	t.PrimaryKeys = strings.Split(keys, ",")
	return t, nil
}
func (m *meta) OpenTable(db common.DB, tableName string) (*schema.Table, error) {
	if strings.HasSuffix(tableName, ".parquet") {
		return m.openTableParquet(db, tableName)
	}
	type pkInfo struct {
		Name  string
		Order int
	}
	columns := []*schema.Column{}
	tabCols, err := getTableInfo(db, tableName)
	if err != nil {
		return nil, err
	}
	for _, col := range tabCols {
		c := &schema.Column{
			Name:        col.Name,
			FetchDriver: driverName,
			TrueType:    col.Type,
			Null:        !col.NotNull,
			Extended:    map[string]any{},
			MaxLength:   0,
		}
		c.Type = DuckDBType(col.Type)
		columns = append(columns, c)
	}
	var pks []pkInfo
	var constraintColumnNames []any
	if err = db.QueryRow(`select constraint_column_names from duckdb_constraints() where 
			table_name = ? and constraint_type = 'PRIMARY KEY'`, tableName).Scan(&constraintColumnNames); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		} else {
			return nil, err
		}
	}
	if len(constraintColumnNames) > 0 {
		//已经排过序了
		for i := range constraintColumnNames {
			pks = append(pks, pkInfo{
				Name:  constraintColumnNames[i].(string),
				Order: i,
			})
		}
	}
	pksStr := []string{}
	for _, one := range pks {
		pksStr = append(pksStr, one.Name)
	}
	tabIdxs, err := getTableIndex(db, tableName)
	if err != nil {
		return nil, err
	}
	//indexColumn 是字段名-->索引的map
	indexColumn := map[string]tableIndex{}
	for _, idx := range tabIdxs {
		//只找出一个字段的索引
		if len(idx.Columns) == 1 {
			indexColumn[idx.Columns[0]] = idx
		}
	}
	for _, col := range columns {
		if idx, ok := indexColumn[col.Name]; ok {
			col.IndexName = idx.Name
			if idx.Unique {
				col.Index = schema.UniqueIndex
			} else {
				col.Index = schema.Index
			}
		}
	}
	t := schema.NewTable(tableName)
	t.Columns = columns

	t.PrimaryKeys = pksStr

	return t, nil
}
