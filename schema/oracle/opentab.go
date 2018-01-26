package oracle

import (
	"log"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

//获取主键字段
func getPk(db common.DB, tableName string) ([]string, error) {
	ns := strings.Split(tableName, ".")
	var schemaName string
	var table string
	if len(ns) == 1 {
		table = tableName
	} else {
		schemaName = ns[0]
		table = ns[1]
	}
	result := []string{}
	if len(schemaName) == 0 {
		strSQL := "select user from dual"
		row := db.QueryRow(strSQL)
		if err := row.Scan(&schemaName); err != nil {
			err = common.NewSQLError(err, strSQL)
			log.Println(err)
			return nil, err
		}

	}
	strSQL :=
		`SELECT cols.column_name
			FROM all_constraints cons,all_cons_columns cols
			WHERE cons.owner=:1
			and cons.OWNER=cols.owner
			and cols.table_name = :2
			AND cons.constraint_type = 'P'
			AND cons.constraint_name = cols.constraint_name
			AND cons.owner = cols.owner
			ORDER BY cols.table_name, cols.position`
	rows, err := db.Query(strSQL, strings.ToUpper(schemaName), strings.ToUpper(table))
	if err != nil {
		err = common.NewSQLError(err, strSQL, schemaName, table)
		log.Println(err)
		return nil, err
	}
	var col string
	for rows.Next() {
		if err := rows.Scan(&col); err != nil {
			return nil, err
		}
		result = append(result, col)
	}
	return result, nil
}
func getColumns(db common.DB, schemaName, table string) ([]*schema.Column, error) {
	if len(schemaName) == 0 {
		strSQL := "select user from dual"
		row := db.QueryRow(strSQL)
		if err := row.Scan(&schemaName); err != nil {
			err = common.NewSQLError(err, strSQL)
			log.Println(err)
			return nil, err
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
		Owner  string `db:"INDEXOWNER"`
		Name   string `db:"INDEXNAME"`
		Column string `db:"COLUMNNAME"`
	}
	columns := []columnType{}
	if err := func() error {
		strSQL := `select column_name as "DBNAME",
					decode(nullable,'Y',1,0) as "DBNULL",
					(case when data_type in ('CLOB','VARCHAR', 'VARCHAR2','CHAR','NVARCHAR2')
						then 'STR'
						when  data_type ='NUMBER' AND DATA_PRECISION IS NULL AND DATA_SCALE = 0 
						then 'INT'
						when data_type ='DATE'
						then 'DATE'
						when data_type in('NUMBER','BINARY_DOUBLE')
						then 'FLOAT'
						when data_type in ('BLOB','RAW')
						then 'BYTEA'
						else data_type
					end) as "DBTYPE",
					(case when data_type = 'CLOB' then 0 when CHAR_LENGTH =0 then DATA_LENGTH else CHAR_LENGTH end) as "DBMAXLENGTH",
					data_type||
						case
						when data_precision is not null and nvl(data_scale,0)>0 then '('||data_precision||','||data_scale||')'
						when data_precision is not null and nvl(data_scale,0)=0 then '('||data_precision||')'
						when data_precision is null and data_scale is not null then '(*,'||data_scale||')'
						when char_length>0 then '('||char_length|| case char_used 
						                                                         when 'B' then ' Byte'
						                                                         when 'C' then ' Char'
						                                                         else null 
						                                           end||')'
						end as "TRUETYPE"
				from ALL_TAB_COLUMNS 
				where owner=:1 and table_name=:2
				order by column_id`
		rows, err := db.Query(strSQL, strings.ToUpper(schemaName), strings.ToUpper(table))
		if err != nil {
			err = common.NewSQLError(err, strSQL, schemaName, table)
			log.Println(err)
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var col columnType
			if err := rows.Scan(
				&col.Name,
				&col.Null,
				&col.Type,
				&col.MaxLength,
				&col.TrueType); err != nil {
				return err
			}
			columns = append(columns, col)
		}
		return nil
	}(); err != nil {
		return nil, err
	}
	indexColumns := []indexType{}
	if err := func() error {

		strSQL := `SELECT min(index_owner) as "INDEXOWNER",
					index_name as "INDEXNAME",min(column_name) as "COLUMNNAME"
				from all_ind_columns a
				where table_owner=:1 and table_name = :2 and
					exists(select 1 from all_indexes b where 
						a.index_owner=b.owner and a.index_name =b.index_name and 
						UNIQUENESS ='NONUNIQUE')
				group by index_name having count(*)=1`
		rows, err := db.Query(strSQL, strings.ToUpper(schemaName), strings.ToUpper(table))
		if err != nil {
			err = common.NewSQLError(err, strSQL, schemaName, table)
			log.Println(err)
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var idx indexType
			if err := rows.Scan(
				&idx.Owner,
				&idx.Name,
				&idx.Column); err != nil {
				return err
			}
			indexColumns = append(indexColumns, idx)
		}
		return nil
	}(); err != nil {
		return nil, err
	}
	//注意indexColumns中可能含有非表字段的名称，例如oracle中的function index
	indexColumnsMap := map[string]indexType{}
	for _, s := range indexColumns {
		indexColumnsMap[s.Column] = s
	}
	rev := []*schema.Column{}
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
			col.Index = true
			col.IndexName = s.Name
			if len(schemaName) > 0 || //如果是其他schema的表，则必定带上schema
				s.Owner != schemaName { //如果index不和表在同一个schema中，也带上schema
				col.IndexName = s.Owner + "." + col.IndexName
			}
		}
		rev = append(rev, col)
	}
	return rev, nil
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
	return t, nil
}
