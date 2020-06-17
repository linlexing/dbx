package mysql

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

//获取主键字段
func getPk(db common.DB, tableName string) ([]string, error) {
	pks := []string{}

	strSQL := fmt.Sprintf("SHOW KEYS FROM %s WHERE Key_name = 'PRIMARY'", tableName)

	rows, err := db.Query(strSQL)

	if err != nil {
		log.Println(strSQL)
		return nil, common.NewSQLError(err, strSQL)
	}
	// 判断取数是13列还是14列
	var mysqlVer = true
	colums, err := rows.Columns()
	if err != nil {
		return nil, common.NewSQLError(err, strSQL)
	}
	if len(colums) == 14 {
		mysqlVer = false
	}
	defer rows.Close()

	for rows.Next() {
		var Table sql.RawBytes
		var NonUnique sql.RawBytes
		var KeyName sql.RawBytes
		var SeqInIndex sql.RawBytes
		var ColumnName string
		var Collation sql.RawBytes
		var Cardinality sql.RawBytes
		var SubPart sql.RawBytes
		var Packed sql.RawBytes
		var Null sql.RawBytes
		var IndexType sql.RawBytes
		var Comment sql.RawBytes
		var IndexComment sql.RawBytes
		var Visible sql.RawBytes

		if mysqlVer {
			if err := rows.Scan(
				&Table,
				&NonUnique,
				&KeyName,
				&SeqInIndex,
				&ColumnName,
				&Collation,
				&Cardinality,
				&SubPart,
				&Packed,
				&Null,
				&IndexType,
				&Comment,
				&IndexComment); err != nil {
				log.Println(err)
				return nil, err
			}
		} else {
			if err := rows.Scan(
				&Table,
				&NonUnique,
				&KeyName,
				&SeqInIndex,
				&ColumnName,
				&Collation,
				&Cardinality,
				&SubPart,
				&Packed,
				&Null,
				&IndexType,
				&Comment,
				&IndexComment,
				&Visible); err != nil {
				log.Println(err)
				return nil, err
			}
		}

		pks = append(pks, ColumnName)
	}
	return pks, rows.Err()
}
func getColumns(db common.DB, schemaName, tableName string) ([]*schema.Column, error) {
	if len(schemaName) == 0 {
		strSQL := "select SCHEMA()"
		row := db.QueryRow(strSQL)
		if err := row.Scan(&schemaName); err != nil {
			log.Println(strSQL)
			return nil, common.NewSQLError(err, strSQL)
		}

	}

	type columnType struct {
		Name      string `db:"DBNAME"`
		Null      int    `db:"DBNULL"`
		Type      string `db:"DBTYPE"`
		MaxLength int    `db:"DBMAXLENGTH"`
		TrueType  string `db:"TRUETYPE"`
	}
	type indexType struct {
		Owner     string `db:"INDEXOWNER"`
		Name      string `db:"INDEXNAME"`
		Column    string `db:"COLUMNNAME"`
		NonUnique int    `db:"NON_UNIQUE"`
	}
	columns := []columnType{}
	if err := func() error {
		strSQL := `select
					column_name as DBNAME,
				    (case when is_nullable='YES' then 1 else 0 end) as DBNULL,
				    (case when data_type in('varchar','text','char','varbinary') then 'STR'
						  when data_type in('bigint','int','smallint','tinyint','mediumint') then 'INT'
						  when data_type in('decimal','double','float') then 'FLOAT'
				          when data_type ='blob' then 'BYTEA'
				          when data_type in('date','datetime','timestamp') then 'DATE'
						  else data_type
				    end) as DBTYPE,
				    (case when data_type in('text') then 0 else ifnull(CHARACTER_MAXIMUM_LENGTH,0) end) as DBMAXLENGTH,
					column_type as TRUETYPE
				from information_schema.columns
				where table_name=? and table_schema= ?
				order by ORDINAL_POSITION`
		rows, err := db.Query(strSQL, tableName, schemaName)
		if err != nil {
			log.Println(strSQL, tableName, schemaName)
			return common.NewSQLError(err, strSQL, tableName, schemaName)
		}
		defer rows.Close()
		for rows.Next() {
			row := columnType{}
			if err := rows.Scan(
				&row.Name,
				&row.Null,
				&row.Type,
				&row.MaxLength,
				&row.TrueType); err != nil {
				return err
			}
			columns = append(columns, row)
		}
		return rows.Err()
	}(); err != nil {
		return nil, err
	}
	indexColumns := []indexType{}
	if err := func() error {
		strSQL := `SELECT max(INDEX_SCHEMA) AS INDEXOWNER,
					INDEX_NAME as INDEXNAME,
					max(COLUMN_NAME) AS COLUMNNAME,
					max(NON_UNIQUE) as NON_UNIQUE
				FROM INFORMATION_SCHEMA.STATISTICS
				WHERE table_schema = ? and table_name=?
				group by index_name having count(*)=1
				ORDER BY table_name, index_name, seq_in_index`
		rows, err := db.Query(strSQL, schemaName, tableName)
		if err != nil {
			log.Println(strSQL, tableName, schemaName)
			return common.NewSQLError(err, strSQL, tableName, schemaName)
		}
		defer rows.Close()
		for rows.Next() {
			row := indexType{}
			if err = rows.Scan(
				&row.Owner,
				&row.Name,
				&row.Column,
				&row.NonUnique); err != nil {
				return err
			}
		}
		return rows.Err()
	}(); err != nil {
		return nil, err
	}
	//注意indexColumns中可能含有非表字段的名称，例如oracle中的function index
	indexColumnsMap := map[string]indexType{}
	for _, s := range indexColumns {
		indexColumnsMap[s.Column] = s
	}

	revColumns := []*schema.Column{}
	for _, v := range columns {
		col := &schema.Column{
			Name:        v.Name,
			Type:        schema.ParseDataType(v.Type),
			MaxLength:   v.MaxLength,
			Null:        v.Null > 0,
			TrueType:    v.TrueType,
			FetchDriver: driverName,
		}

		//组合主键，有时需要单字段索引
		if s, ok := indexColumnsMap[v.Name]; ok {
			if s.NonUnique == 1 {
				col.Index = schema.Index
			} else {
				col.Index = schema.UniqueIndex
			}
			col.IndexName = s.Name
			if len(schemaName) > 0 || //如果是其他schema的表，则必定带上schema
				s.Owner != schemaName { //如果index不和表在同一个schema中，也带上schema
				col.IndexName = s.Owner + "." + col.IndexName
			}
		}
		revColumns = append(revColumns, col)
	}
	return revColumns, nil

}
func (m *meta) OpenTable(db common.DB, tableName string) (*schema.Table, error) {
	t := schema.NewTable(tableName)

	pks, err := getPk(db, tableName)

	if err != nil {
		return nil, err
	}

	cols, err := getColumns(db, t.Schema, t.Name)

	if err != nil {
		return nil, err
	}

	t.Columns = cols
	t.PrimaryKeys = pks
	//最后去除单主键的主键字段的索引，因为主键自动加索引，不需要体现在字段定义中
	if len(t.PrimaryKeys) == 1 {
		for _, one := range t.Columns {
			if one.Name == t.PrimaryKeys[0] && one.Index != schema.NoIndex {
				one.Index = schema.NoIndex
				one.IndexName = ""
			}
		}
	}
	return t, nil
}
