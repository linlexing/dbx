package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"sort"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

type tableColumn struct {
	CID       int
	Name      string
	Type      string
	NotNull   int
	DfltValue sql.NullString `db:"dflt_value"`
	PK        int
}
type tableIndex struct {
	Seq     int
	Name    string
	Unique  int
	Origin  string
	Partial int
}
type indexInfo struct {
	SeqNO int
	CID   int
	Name  string
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
func getTableIndex(db common.DB, tableName string) ([]tableIndex, error) {
	rev := []tableIndex{}
	strSQL := fmt.Sprintf("PRAGMA index_list(%s)", tableName)
	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var idx tableIndex
		if err = rows.Scan(
			&idx.Seq,
			&idx.Name,
			&idx.Unique,
			&idx.Origin,
			&idx.Partial); err != nil {
			return nil, err
		}
		rev = append(rev, idx)

	}
	return rev, rows.Err()
}
func getIndexInfo(db common.DB, indexName string) ([]indexInfo, error) {
	//每个索引再去找定义
	rev := []indexInfo{}
	strSQL := fmt.Sprintf("PRAGMA index_info(%s)", indexName)

	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var idxInfo indexInfo
		if err := rows.Scan(
			&idxInfo.SeqNO,
			&idxInfo.CID,
			&idxInfo.Name); err != nil {
			return nil, err
		}
		rev = append(rev, idxInfo)
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
func (m *meta) OpenTable(db common.DB, tableName string) (*schema.Table, error) {
	type pkInfo struct {
		Name  string
		Order int
	}
	pks := []pkInfo{}
	columns := []*schema.Column{}

	tabCols, err := getTableInfo(db, tableName)
	if err != nil {
		return nil, err
	}
	for _, col := range tabCols {
		if col.PK > 0 {
			pks = append(pks, pkInfo{Name: col.Name, Order: col.PK})
		}
		c := &schema.Column{
			Name:        col.Name,
			FetchDriver: driverName,
			TrueType:    col.Type,
			Null:        col.NotNull != 1,
			Extended:    map[string]any{},
		}
		c.Type, c.MaxLength = sqliteType(col.Name, col.Type)
		columns = append(columns, c)
	}
	//pks必须要排序,暂时不用slices包
	//slices.SortFunc(pks, func(a, b pkInfo) int {
	//	return a.Order - b.Order
	//})
	sort.Slice(pks, func(i, j int) bool {
		return pks[i].Order < pks[j].Order
	})
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
		idxInfo, err := getIndexInfo(db, idx.Name)
		if err != nil {
			return nil, err
		}
		//只找出一个字段的索引,并且不是主键索引
		if len(idxInfo) == 1 && (len(pks) > 1 ||
			idxInfo[0].Name != pks[0].Name) {
			indexColumn[idxInfo[0].Name] = idx
		}
	}
	for _, col := range columns {
		if idx, ok := indexColumn[col.Name]; ok {
			col.IndexName = idx.Name
			if idx.Unique == 1 {
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
