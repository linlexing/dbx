package postgres

import (
	"fmt"
	"log"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

type columnType struct {
	Name string `db:"DBNAME"`
	Null int    `db:"DBNULL"`
	Type string `db:"DBTYPE"`

	MaxLength int    `db:"DBMAXLENGTH"`
	TrueType  string `db:"TRUETYPE"`
}
type indexType struct {
	Owner  string `db:"INDEXOWNER"`
	Name   string `db:"INDEXNAME"`
	Column string `db:"COLUMNNAME"`
	Unique bool   `db:"INDISUNIQUE"`
}
type idOrder struct {
	ID   int
	Name string
}

// 获取主键字段编号的列表
func getPKIDS(db common.DB, tableName string) ([]int, error) {
	strSQL := fmt.Sprintf(
		`SELECT unnest(i.indkey) as pkid
		FROM   pg_index i
		WHERE  i.indrelid = '%s'::regclass
		AND    i.indisprimary`, tableName)
	result := []int{}
	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var oneCol int
		if err = rows.Scan(&oneCol); err != nil {
			return nil, err
		}
		result = append(result, oneCol)
	}

	return result, rows.Err()
}

// 获取主键字段
// tablename需要加单引号才能被sql语句识别
func getPk(db common.DB, tableName string) ([]string, error) {
	//为适应华为高斯，改成两步获取
	ids, err := getPKIDS(db, tableName)
	if err != nil {
		return nil, err
	}

	names := []*idOrder{}
	strSQL := fmt.Sprintf(
		`SELECT a.attname,a.attnum
			FROM   pg_index i
			JOIN   pg_attribute a ON a.attrelid = i.indrelid
				AND a.attnum = ANY(i.indkey)
			WHERE  i.indrelid = '%s'::regclass
			AND    i.indisprimary`, tableName)

	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&name, &id); err != nil {
			return nil, err
		}
		names = append(names, &idOrder{id, name})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	result := []string{}
	//按照顺序组合字段
	for _, one := range ids {
		for _, col := range names {
			if one == col.ID {
				result = append(result, col.Name)
			}
		}
	}
	return result, nil

}

// columnType中DBNULL被定义为int类型
// 0代表false 1代表true
func getTableColumns(db common.DB, schemaName, tableName string) ([]columnType, error) {
	columns := []columnType{}
	strSQL := `select column_name as "DBNAME",
					(case when is_nullable='YES' then 1 else 0 end) as "DBNULL",
					(case when data_type in ('text', 'varchar','character varying','jsob','jsonb','ARRAY','USER-DEFINED',
						'uuid','boolean','daterange','int8range','numrange','tsrange')
						then 'STR'
						when  data_type in ('smallint','integer','bigint')
						then 'INT'
						when data_type in ('timestamp(6) with time zone','timestamp with time zone','timestamp(6) without time zone', 'timestamp without time zone')
						then 'DATE'
						when data_type in('numeric','double precision','real')
						then 'FLOAT'
						when data_type ='bytea'
						then 'BYTEA'
						else data_type
					end) as "DBTYPE",
					(case when character_maximum_length is null then 0 else character_maximum_length end) as "DBMAXLENGTH",
					(SELECT CASE WHEN a.atttypid = ANY ('{int,int8,int2}'::regtype[])
						AND EXISTS (
							SELECT 1 FROM pg_attrdef ad
							WHERE  ad.adrelid = a.attrelid
							AND    ad.adnum   = a.attnum
							AND    pg_get_expr(ad.adbin, ad.adrelid)
								= 'nextval('''
								|| (pg_get_serial_sequence (a.attrelid::regclass::text, a.attname))::regclass
								|| '''::regclass)'
							) THEN
								CASE a.atttypid
									WHEN 'int'::regtype  THEN 'serial'
									WHEN 'int8'::regtype THEN 'bigserial'
									WHEN 'int2'::regtype THEN 'smallserial'
					   			END
				  		ELSE format_type(a.atttypid, a.atttypmod)
						end
					FROM pg_attribute a
						JOIN pg_class b ON (a.attrelid = b.oid)
						JOIN pg_namespace c ON (c.oid = b.relnamespace)
					WHERE
						b.relname = outa.table_name AND
						c.nspname = outa.table_schema AND
						a.attname = outa.column_name) as "TRUETYPE"
				from information_schema.columns outa
				where table_schema ilike $1 and table_name ilike $2
				order by ordinal_position`

	rows, err := db.Query(strSQL, schemaName, tableName)
	if err != nil {
		err = common.NewSQLError(err, strSQL, schemaName, tableName)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var col columnType
		if err = rows.Scan(
			&col.Name,
			&col.Null,
			&col.Type,
			&col.MaxLength,
			&col.TrueType); err != nil {
			return nil, err
		}
		columns = append(columns, col)
	}
	return columns, rows.Err()

}

// SQL语句查询结果为null
func getTableIndexes(db common.DB, schemaName, tableName string) ([]indexType, error) {
	indexes := []indexType{}
	strSQL := `select
					(select nspname from pg_namespace where oid=i.relnamespace) as "INDEXOWNER",
					i.relname as "INDEXNAME",
					min(a.attname) as "COLUMNNAME",
					ix.indisunique as "INDISUNIQUE"
				from
				    pg_class t,
				    pg_class i,
				    pg_index ix,
				    pg_attribute a,
				    pg_namespace tn,
					pg_am am
				where
				    t.oid = ix.indrelid
				    and i.oid = ix.indexrelid
				    and a.attrelid = t.oid
				    and t.relnamespace=tn.oid
				    and tn.nspname ilike $1
				    and a.attnum = ANY(ix.indkey)
				    and t.relkind = 'r'
					and t.relname ilike $2
					and not ix.indisprimary
					and i.relam=am.oid
					and am.amname='btree'
				group by
				   t.relname,
				   i.relnamespace,
				   i.relname,
				   ix.indisunique
				having count(*)=1
				order by
				    t.relname,
				    i.relname;`

	rows, err := db.Query(strSQL, schemaName, tableName)
	if err != nil {
		err = common.NewSQLError(err, strSQL, schemaName, tableName)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var idx indexType
		if err = rows.Scan(
			&idx.Owner,
			&idx.Name,
			&idx.Column,
			&idx.Unique); err != nil {
			return nil, err
		}
		indexes = append(indexes, idx)
	}
	return indexes, rows.Err()

}
func getColumns(db common.DB, schemaName, table string) ([]*schema.Column, error) {
	if len(schemaName) == 0 {
		strSQL := "select current_schema()"

		if err := db.QueryRow(strSQL).Scan(&schemaName); err != nil {
			err = common.NewSQLError(err, strSQL)
			log.Println(err)
			return nil, err
		}
	}
	columns, err := getTableColumns(db, schemaName, table)
	if err != nil {
		return nil, err
	}

	indexColumns, err := getTableIndexes(db, schemaName, table)
	if err != nil {
		return nil, err
	}

	//注意indexColumns中可能含有非表字段的名称，例如oracle中的function index
	indexColumnsMap := map[string]indexType{}
	for _, s := range indexColumns {
		indexColumnsMap[s.Column] = s
	}
	rev := []*schema.Column{}

	for _, v := range columns {
		vtype, err := schema.ParseDataType(v.Type)
		if err != nil {
			return nil, err
		}
		col := &schema.Column{
			Name:        v.Name,
			Type:        vtype,
			MaxLength:   v.MaxLength,
			Null:        v.Null > 0,
			TrueType:    v.TrueType,
			FetchDriver: driverName,
		}

		//组合主键，有时需要单字段索引
		if s, ok := indexColumnsMap[v.Name]; ok {
			if s.Unique {
				col.Index = schema.UniqueIndex
			} else {
				col.Index = schema.Index
			}
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
	//获取主键前需要先判断表是否存在，防止出现表不存在抛出异常
	tabExists, err := m.TableExists(db, tableName)
	if err != nil {
		return nil, err
	}
	pks := []string{}
	if tabExists {
		pks, err = getPk(db, tableName)
		if err != nil {
			return nil, err
		}
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
