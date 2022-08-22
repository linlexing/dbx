package data

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"errors"
	"fmt"
	"log"

	"strconv"
	"strings"
	"time"

	"github.com/linlexing/mapfun"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/scan"
	"github.com/linlexing/dbx/schema"
)

//Table 表现一个数据库表,扩展了schema.Table ，提供了数据访问，
//注意该表实例后，不能再去修改其结构
type Table struct {
	Driver string
	DB     common.DB
	*schema.Table

	ColumnTypes    []*scan.ColumnType
	notnullColumns []string
	ColumnNames    []string
	columnsMap     map[string]*schema.Column //用于快速查询，名称用大写的
}

//NewTable 用schema.Table构造一个Table,没有数据库操作发生
func NewTable(driver string, db common.DB, st *schema.Table) *Table {
	rev := &Table{
		Driver:         driver,
		DB:             db,
		ColumnNames:    []string{},
		Table:          st,
		ColumnTypes:    []*scan.ColumnType{},
		notnullColumns: []string{},
		columnsMap:     map[string]*schema.Column{},
	}
	rev.BuildColumnIndex()
	return rev
}

//OpenTable 从数据库取出结构构造Table,表名用 schema.tablename的方式,
//如果不读取数据，仅定义结构，应当使用schema.Table
//注意：这里返回的所有字段名都转换成了大写，如果要获取实际的字段名，使用OpenTableCase
func OpenTable(driver string, db common.DB, tabName string) (*Table, error) {
	if len(tabName) == 0 {
		return nil, errors.New("table name is empty")
	}
	st, err := schema.Find(driver).OpenTable(db, tabName)
	if err != nil {
		return nil, err
	}
	st.ToUpper()
	return NewTable(driver, db, st), nil
}
func OpenTableCase(driver string, db common.DB, tabName string) (*Table, error) {
	if len(tabName) == 0 {
		return nil, errors.New("table name is empty")
	}
	st, err := schema.Find(driver).OpenTable(db, tabName)
	if err != nil {
		return nil, err
	}
	return NewTable(driver, db, st), nil
}
func (t *Table) bind(strSQL string) string {
	return Bind(t.Driver, strSQL)
}

//BuildColumnIndex 在变动表结构后调用，一般是自动调用，只有在NewTable后，
//又去手工变动过schema.Table.Columns,才需要去手动调用
func (t *Table) BuildColumnIndex() {
	//构造索引
	t.ColumnNames = []string{}
	t.ColumnTypes = []*scan.ColumnType{}
	t.notnullColumns = []string{}
	t.columnsMap = map[string]*schema.Column{}
	for _, col := range t.Columns {
		t.ColumnNames = append(t.ColumnNames, col.Name)
		if !col.Null {
			t.notnullColumns = append(t.notnullColumns, col.Name)
		}
		t.ColumnTypes = append(t.ColumnTypes, &scan.ColumnType{
			Name: col.Name,
			Type: col.Type,
		})
		t.columnsMap[strings.ToUpper(col.Name)] = col
	}

}
func (t *Table) findColumn(name string) (*schema.Column, bool) {
	rev, ok := t.columnsMap[strings.ToUpper(name)]
	return rev, ok
}

//RefreshSchema 从数据库重新检索表结构
func (t *Table) RefreshSchema() error {
	tab, err := schema.Find(t.Driver).OpenTable(t.DB, t.FullName())
	if err != nil {
		return err
	}
	t.Table = tab
	t.BuildColumnIndex()
	return nil
}

//Row 根据一个主键值返回一个记录,如果没有找到返回nil
func (t *Table) Row(pks ...interface{}) (map[string]interface{}, error) {
	whereList := []string{}
	if len(t.PrimaryKeys) != len(pks) {
		return nil, fmt.Errorf("the table %s pk values number error.table pk:%#v,pkvalues:%#v", t.FullName(), t.PrimaryKeys, pks)
	}
	for i, onePk := range t.PrimaryKeys {
		whereList = append(whereList, fmt.Sprintf("%s=?", onePk))
		if pks[i] == nil {
			return nil, fmt.Errorf("the pk field %s value is nil", onePk)
		}
	}
	rows, err := t.Query(strings.Join(whereList, " and\n"), pks...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, rows.Err()
	}
	return t.ScanMap(rows)
}

//ToJSON 将一行数据转换成json,日期、二进制、int64数据转换成文本
//注意传入的字段不一定是全字段
func (t *Table) ToJSON(row map[string]interface{}) (map[string]interface{}, error) {
	transRecord := map[string]interface{}{}
	for k, v := range row {
		col, ok := t.findColumn(k)
		if !ok {
			return nil, fmt.Errorf("not found column %s", k)
		}
		var err error
		if transRecord[k], err = col.Type.ToJSON(v); err != nil {
			return nil, err
		}
	}
	return transRecord, nil
}

//FromJSON 将一个json数据转换回row
//注意传入的字段不一定是全字段
//字段名忽略大小写
func (t *Table) FromJSON(row map[string]interface{}) (map[string]interface{}, error) {
	transRecord := map[string]interface{}{}
	for k, v := range row {
		col, ok := t.findColumn(k)
		if !ok {
			return nil, fmt.Errorf("not found column %s", k)
		}
		var err error
		if transRecord[k], err = col.Type.ParseJSON(v); err != nil {
			return nil, err
		}
	}
	return transRecord, nil
}

//SafeFromJSON 将一个json数据转换回row,忽略不能转换的字段
//注意传入的字段不一定是全字段
func (t *Table) SafeFromJSON(row map[string]interface{}) map[string]interface{} {
	transRecord := map[string]interface{}{}
	for k, v := range row {
		col, ok := t.findColumn(k)
		if !ok {
			continue
		}
		if tv, err := col.Type.ParseJSON(v); err != nil {
			continue
		} else {
			transRecord[k] = tv
		}
	}
	return transRecord
}

//ScanSlice 根据一个Scaner，再根据Table的字段数据类型，扫描出一个slice
func (t *Table) ScanSlice(s common.Scaner) (result []interface{}, err error) {
	result, err = scan.TypeScan(s, t.ColumnTypes)
	return
}

//ScanMap 根据一个Scaner，再根据Table的字段数据类型，扫描出一个map
func (t *Table) ScanMap(s common.Scaner) (result map[string]interface{}, err error) {
	outList, err := t.ScanSlice(s)
	if err != nil {
		return
	}
	result = map[string]interface{}{}
	for i, one := range outList {
		result[t.Columns[i].Name] = one
	}
	return result, nil
}

//QueryOrderRows 直接返回所有记录
func (t *Table) QueryOrderRows(orderby []string, where string,
	param ...interface{}) (rows []map[string]interface{}, err error) {
	rs, err := t.QueryOrder(orderby, where, param...)
	if err != nil {
		return
	}
	defer rs.Close()
	rows = []map[string]interface{}{}
	for rs.Next() {
		one, err := t.ScanMap(rs)
		if err != nil {
			return nil, err
		}
		rows = append(rows, one)
	}
	return
}

//QueryOrder 查询返回排序记录，where如有参数，必须用 ?
func (t *Table) QueryOrder(orderby []string, where string,
	param ...interface{}) (rows *sql.Rows, err error) {

	if len(where) > 0 {
		where = " where " + where
	}
	columnsStr := strings.Join(t.ColumnNames, ",")

	strOrderby := ""
	if len(orderby) > 0 {
		strOrderby = " order by " + strings.Join(orderby, ",")
	}
	strSQL := t.bind(fmt.Sprintf("select %s from %s%s%s", columnsStr, t.FullName(), where, strOrderby))

	if rows, err = t.DB.Query(strSQL, param...); err != nil {
		err = common.NewSQLError(err, strSQL, param)
	}
	return
}

//Query 查询返回记录，无排序
func (t *Table) Query(where string, param ...interface{}) (*sql.Rows, error) {
	return t.QueryOrder(nil, where, param...)
}

//QueryRows 直接返回记录
func (t *Table) QueryRows(where string, param ...interface{}) ([]map[string]interface{}, error) {
	return t.QueryOrderRows(nil, where, param...)
}

//KeyExists 检查一个主键是否存在
func (t *Table) KeyExists(pks ...interface{}) (result bool, err error) {
	whereList := []string{}

	for _, v := range t.PrimaryKeys {
		whereList = append(whereList, v+"=?")
	}
	var strWhere string
	if len(whereList) > 0 {
		strWhere = " where " + strings.Join(whereList, " and ")
	}

	var rows *sql.Rows
	strSQL := fmt.Sprintf("select 1 from %s %s", t.FullName(), strWhere)
	if rows, err = t.DB.Query(t.bind(strSQL), pks...); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return
	}
	defer rows.Close()
	result = rows.Next()
	err = rows.Err()
	return

}

//KeyValues 返回一个记录的主键值,row中的字段名会忽略大小写和主键进行比较
func (t *Table) KeyValues(row map[string]interface{}) []interface{} {
	rev := []interface{}{}
	nameMap := map[string]string{}
	for k := range row {
		nameMap[strings.ToLower(k)] = k
	}
	for _, v := range t.PrimaryKeys {
		rev = append(rev, row[nameMap[strings.ToLower(v)]])
	}
	return rev
}

//MustCount 统计记录数
//参数可以传入string 代表where,map[string]interface{} 代表字段条件
func (t *Table) MustCount(where string, params ...interface{}) int64 {
	i, err := t.Count(where, params...)
	if err != nil {
		log.Panic(err)
		return -1

	}
	return i

}

//Exists 检测指定条件的记录是否存在，只要用于Bill.Remove方法，目前没有用到数据的exists，
//后期优化性能，可以考虑改成select 1 from dual where exists()
func (t *Table) Exists(where string, params ...interface{}) (rev bool, err error) {
	strSQL := "select 1 from " + t.FullName()
	if len(where) > 0 {

		strSQL += " where " + where
	}
	rows, err := t.DB.Query(t.bind(strSQL), params...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, params...)
		log.Println(err)
		return
	}
	defer rows.Close()
	rev = rows.Next()
	err = rows.Err()
	return
}

//Count 统计表中记录数，其实没什么逻辑，就是省了组合一个sql语句
func (t *Table) Count(where string, params ...interface{}) (rev int64, err error) {
	strSQL := "select count(*) from " + t.FullName()
	if len(where) > 0 {

		strSQL += " where " + where
	}
	rev, err = AsInt(t.DB, t.bind(strSQL), params...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, params...)
	}

	return
}

func (t *Table) checkNotNull(row map[string]interface{}) error {
	for _, v := range t.notnullColumns {
		if val, ok := row[v]; ok && val == nil {
			return fmt.Errorf("the not null column:%s is null", v)
		}
	}
	return nil
}

func truncateTimeZone(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), 0, time.UTC)
}

//检查row中是否含有非空字段的值，以及去掉多余的字段值
//如果是oracle，则需要去除时间中的时区，以免触发ORA-01878错误
func (t *Table) checkAndConvertRow(row map[string]interface{}) error {
	if t.Driver == "oci8" {
		for k, v := range row {
			if tm, ok := v.(time.Time); ok {
				row[k] = truncateTimeZone(tm)
			}
		}
	}
	return t.checkNotNull(row)
}

//ImportFromTable 从另一个表中导入数据，表中列数量、名称、类型必须一致
func (t *Table) ImportFromTable(srcTable *Table, progressFunc func(string, interface{}), where string,
	args ...interface{}) (iCount int64, err error) {
	if len(t.ColumnNames) != len(srcTable.ColumnNames) {
		return -1, errors.New("column number not equ")
	}
	for _, col := range t.ColumnNames {
		if srcTable.ColumnByName(col) == nil {
			return -1, errors.New("column:" + col + " not exists")
		}
	}
	var whereStr string
	if len(where) > 0 {
		whereStr = " where " + where
	}
	query := fmt.Sprintf("select %s from %s%s", strings.Join(t.ColumnNames, ","),
		srcTable.FullName(), whereStr)
	return t.ImportFrom(srcTable.DB, progressFunc, query, args...)
}

//ImportFrom 从一个查询中导入数据,其列必须与表中数量一致,且序号类型一致，因为可能是异构数据库
//所以不能用直接的CreateTableAs,由于数据可能比较多，采用5秒钟提交一次事务，
//所以Table.DB必须是TxDB
func (t *Table) ImportFrom(db common.Queryer, progressFunc func(string, interface{}), query string,
	args ...interface{}) (iCount int64, err error) {
	var rowCount int64
	strSQL := fmt.Sprintf("select count(*) from (%s) out_count", query)
	rowCount, err = AsInt(db, t.bind(strSQL), args...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, args...)
		return
	}
	progressFunc(fmt.Sprintf("start import table %s,total %d records", t.FullName(), rowCount), nil)

	rows, err := db.Query(query, args...)
	if err != nil {
		err = common.NewSQLError(err, query, args...)
		return
	}
	defer rows.Close()
	var batCount int64

	//再构造insert语句
	insertSQL := t.InsertSQL()
	//再开始事务

	startTime := time.Now()
	beginTime := startTime
	finished := false
	tx, err := t.DB.(common.TxDB).Begin()
	if err != nil {
		return
	}
	defer func() {
		if !finished && tx != nil {
			tx.Rollback()
		}
	}()
	insertStmt, err := tx.Prepare(insertSQL)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}
	iCount = 0
	batCount = 0
	for rows.Next() {
		outList, err := scan.TypeScan(rows, t.ColumnTypes)
		if err != nil {
			return 0, err
		}

		if _, err = insertStmt.Exec(outList...); err != nil {
			log.Printf("error:%s,values:\n", err)
			for ei, ev := range outList {
				log.Printf("\t%s=%#v", t.Columns[ei].Name, ev)
			}
			return 0, err
		}
		iCount++
		batCount++
		totalSec := time.Since(startTime).Seconds()
		//5秒提交一次
		if totalSec >= 5 {
			if err = insertStmt.Close(); err != nil {
				log.Println(err)
				return 0, err
			}
			if err = tx.Commit(); err != nil {
				log.Println(err)
				return 0, err
			}
			batCount = 0
			tx = nil
			tx, err = t.DB.(common.TxDB).Begin()
			if err != nil {
				log.Println(err)
				return 0, err
			}
			insertStmt, err = tx.Prepare(insertSQL)
			if err != nil {
				log.Println(err)
				return 0, err
			}
			progressFunc(fmt.Sprintf("\t%.2f%%\t%d/%d\t%.2fs", 100.0*float64(iCount)/float64(rowCount), iCount, rowCount, totalSec), nil)
			startTime = time.Now()
		}
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return
	}
	//如果最后一批有数据
	if batCount > 0 {
		if err = insertStmt.Close(); err != nil {
			log.Println(err)
			return
		}
		if err = tx.Commit(); err != nil {
			tx = nil
			return
		}
	}
	finished = true
	progressFunc(fmt.Sprintf("%s,total %d records imported %.2fs", t.FullName(), iCount, time.Since(beginTime).Seconds()), nil)
	return
}

//InsertSQL 生成一个InsertSQL
func (t *Table) InsertSQL() string {
	insertSQL := fmt.Sprintf(
		"insert into %s(%s)values(%s)",
		t.FullName(), strings.Join(t.ColumnNames, ","),
		strings.Join(strings.Split(strings.Repeat("?", len(t.Columns)), ""), ","))
	insertSQL = t.bind(insertSQL)
	return insertSQL

}

//仅非空字段生成语句
func (t *Table) insertAsPack(row map[string]interface{}) (err error) {
	columns := []string{}
	data := []interface{}{}

	for k, v := range row {
		if v != nil {
			data = append(data, v)
			columns = append(columns, k)
		}
	}

	strSQL := fmt.Sprintf(
		"insert into %s(%s)values(%s)",
		t.FullName(), strings.Join(columns, ","),
		strings.Join(strings.Split(strings.Repeat("?", len(columns)), ""), ","))
	if _, err = t.DB.Exec(strSQL, data...); err != nil {
		err = common.NewSQLError(err, strSQL, data...)
	}
	return
}

//编码key值，如果是复合主键，则用gob序列化
func (t *Table) encodeKey(keys ...interface{}) []byte {
	if len(keys) == 1 {
		switch t.ColumnByName(t.PrimaryKeys[0]).Type {
		case schema.TypeString:
			return []byte(keys[0].(string))
		case schema.TypeInt:
			return []byte(fmt.Sprintf("%d", keys[0]))
		case schema.TypeBytea:
			return keys[0].([]byte)
		default:
			panic("invalid primary key datatype")
		}
	}
	out := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(out).Encode(keys); err != nil {
		log.Panic(err)
	}
	return out.Bytes()

}

//解开主键
func (t *Table) decodeKey(key []byte) []interface{} {

	if len(t.PrimaryKeys) == 1 {
		switch t.ColumnByName(t.PrimaryKeys[0]).Type {
		case schema.TypeBytea:
			return []interface{}{key}
		case schema.TypeString:
			return []interface{}{string(key)}
		case schema.TypeInt:
			i, err := strconv.ParseInt(string(key), 10, 64)
			if err != nil {
				panic("invalid int data" + string(key))
			}
			return []interface{}{i}
		}
	}
	in := bytes.NewBuffer(key)
	rev := []interface{}{}
	if err := gob.NewDecoder(in).Decode(&rev); err != nil {
		log.Panic(err)
	}
	return rev
}

//Insert 插入一批记录,使用第一行数据中的字段，并没有使用表中的字段,因此可以插入部分字段
func (t *Table) Insert(rows []map[string]interface{}) (err error) {
	if len(rows) == 0 {
		return nil
	}
	if len(rows) == 1 {
		if err = t.checkAndConvertRow(rows[0]); err != nil {

			return
		}
		return t.insertAsPack(rows[0])

	}
	cols := []string{}
	colMaps := map[string]struct{}{}
	for k := range rows[0] {
		colMaps[k] = struct{}{}
	}
	//按照现有字段顺序进行排序
	for _, one := range t.Columns {
		if _, ok := colMaps[one.Name]; ok {
			cols = append(cols, one.Name)
		}
	}
	strSQL := fmt.Sprintf("insert into %s(%s)values(%s)",
		t.FullName(), strings.Join(cols, ","),
		strings.Join(strings.Split(strings.Repeat("?", len(cols)), ""), ","))
	stmt, err := t.DB.Prepare(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return
	}
	defer stmt.Close()
	//先检查并插入数据

	for _, row := range rows {
		if err = t.checkAndConvertRow(row); err != nil {
			return
		}
		data := []interface{}{}
		for _, col := range cols {
			data = append(data, row[col])
		}
		if _, err = stmt.Exec(data...); err != nil {
			return common.NewSQLError(err, strSQL, data...)
		}
	}

	return
}

//Delete 删除记录，全部字段值将被生成where字句(text、bytea除外)
func (t *Table) Delete(rows []map[string]interface{}) (delCount int64, err error) {
	//考虑到null值，所有的行不能用一个语句，必须单独删除
	for _, v := range rows {
		var i int64
		if i, err = t.Remove(v); err != nil {
			return
		}
		delCount += i
	}
	return
}

//RemoveByKeyValues 根据完整或者部分主键值，删除记录，返回删除的行数
func (t *Table) RemoveByKeyValues(keyValues ...interface{}) (delCount int64, err error) {
	whereList := []string{}
	for i := range keyValues {
		whereList = append(whereList, fmt.Sprintf("%s=?", t.PrimaryKeys[i]))
	}
	strSQL := t.bind(fmt.Sprintf("delete from %s where %s", t.FullName(),
		strings.Join(whereList, " and\n")))

	sr, err := t.DB.Exec(strSQL, keyValues...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, keyValues...)
		log.Println(err)
		return
	}
	delCount, err = sr.RowsAffected()
	return

}
func numberOfDigits(f float64) int {
	ls := strings.Split(strconv.FormatFloat(f, 'f', -1, 64), ".")
	if len(ls) > 1 {
		return len(ls[1])
	}
	return 0
}
func (t *Table) buildWhere(row map[string]interface{}) (string, []interface{}) {
	pkMap := map[string]struct{}{}
	for _, one := range t.PrimaryKeys {
		pkMap[one] = struct{}{}
	}

	strWhere := []string{}
	newRow := []interface{}{}
	for k, v := range row {
		//如果是没有长度的string，即text，以及bytea、datetime不参与where条件
		//float因为精度问题，超过六位小数不能用来做where
		//如果是主键，则一定参与条件
		fld := t.ColumnByName(k)
		if v == nil {
			strWhere = append(strWhere, fmt.Sprintf("%s is null", k))
			continue
		}
		_, isPK := pkMap[k]
		if !isPK && (fld.Type == schema.TypeBytea ||
			fld.Type == schema.TypeDatetime ||
			(fld.Type == schema.TypeFloat && numberOfDigits(v.(float64)) > 6) ||
			(fld.Type == schema.TypeString && fld.MaxLength <= 0)) {
			continue
		}
		if v == nil {
			strWhere = append(strWhere, fmt.Sprintf("%s is null", k))
		} else {
			strWhere = append(strWhere, fmt.Sprintf("%s=?", k))
			newRow = append(newRow, v)
		}
	}
	return strings.Join(strWhere, " and\n"), newRow

}

//Remove 删除一个记录，必须是全指标的记录
func (t *Table) Remove(row map[string]interface{}) (delCount int64, err error) {
	err = t.checkAndConvertRow(row)
	if err != nil {
		return
	}
	strWhere, newRow := t.buildWhere(row)
	strSQL := t.bind(fmt.Sprintf(
		"delete from %s where %s", t.FullName(), strWhere))
	var sqlr sql.Result
	if sqlr, err = t.DB.Exec(strSQL, newRow...); err != nil {
		err = common.NewSQLError(err, strSQL, newRow...)
		log.Println(err)
		return
	}

	delCount, err = sqlr.RowsAffected()

	return
}

//UpdateByKey 通过一个key更新记录
func (t *Table) UpdateByKey(key []interface{}, row map[string]interface{}) (upCount int64, err error) {

	whereList := []string{}
	for _, v := range t.PrimaryKeys {
		whereList = append(whereList, fmt.Sprintf("%s=?", v))
	}

	return t.UpdateByWhere(row, strings.Join(whereList, " and\n"), key...)
}

//UpdateByWhere 通过一个条件更新指定的字段值
func (t *Table) UpdateByWhere(row map[string]interface{}, where string, params ...interface{}) (upCount int64, err error) {
	if len(row) == 0 {
		err = fmt.Errorf("data is null,row:%v,where:%v,params:%#v", row, where, params)
		return
	}

	if err = t.checkAndConvertRow(row); err != nil {
		return 0, err
	}
	set := []string{}
	setVals := []interface{}{}
	for k, v := range row {
		set = append(set, fmt.Sprintf("%s=?", k))
		setVals = append(setVals, v)
	}
	whereStr := ""
	if len(where) > 0 {
		whereStr = "where " + where
	}
	setVals = append(setVals, params...)
	strSQL := t.bind(fmt.Sprintf("update %s set %s %s",
		t.FullName(), strings.Join(set, ","), whereStr))
	var sqlr sql.Result

	if sqlr, err = t.DB.Exec(strSQL, setVals...); err != nil {
		err = common.NewSQLError(err, strSQL, setVals...)
		return
	}
	upCount, err = sqlr.RowsAffected()
	//去除调试的信息
	// if upCount == 0 {
	// 	if _, er := os.Stdout.WriteString(fmt.Sprintf(
	// 		"---------[%v] update nothing,maybe is insert-------\n%s\n%s",
	// 		time.Now(), strSQL, spew.Sdump(setVals))); er != nil {
	// 		panic(er)
	// 	}
	// }
	return
}

//Update 只有修改过的字段才被更新，where采用全部旧值判断（没有长度的string将不参与，因为oracle会出错）
func (t *Table) Update(oldData, newData map[string]interface{}) (upCount int64, err error) {
	if oldData == nil || len(oldData) == 0 || newData == nil || len(newData) == 0 {
		err = errors.New("data is empty")
		return
	}
	// 为允许部分更新，以下代码注释
	// if len(oldData) != len(newData) {
	// 	err = errors.New("the old and new record,field number not same")
	// 	return
	// }
	if err = t.checkAndConvertRow(oldData); err != nil {
		return
	}
	if err = t.checkAndConvertRow(newData); err != nil {
		return
	}
	whereStr, whereVals := t.buildWhere(oldData)
	//仅修改差异部分
	chgs := mapfun.Changes(oldData, newData)
	if len(chgs) == 0 {
		return
	}
	return t.UpdateByWhere(chgs, whereStr, whereVals...)
}

//Save 保存一个记录，先尝试用keyvalue去update，如果更新到记录为0再insert，
//逻辑上是正确的，同时，速度也会有保障
func (t *Table) Save(row map[string]interface{}) error {
	if len(t.PrimaryKeys) == 0 {
		return errors.New("no pk")
	}
	i, err := t.UpdateByKey(t.KeyValues(row), row)
	if err != nil {
		return err
	}
	if i > 0 {
		return nil
	}
	return t.insertAsPack(row)
}

//BatchSave 批量保存记录，返回插入和更新记录数,注意性能，update采用bykey方式
func (t *Table) BatchSave(rows []map[string]interface{}) (insNum int64, updNum int64, err error) {
	//得到一个插入用的sql.Stmt
	insertStmt, err := t.DB.Prepare(t.InsertSQL())
	if err != nil {
		return
	}
	defer insertStmt.Close()
	var setList []string
	var whereList []string
	for _, col := range t.ColumnNames {
		setList = append(setList, fmt.Sprintf("%s=?", col))
	}
	for _, key := range t.PrimaryKeys {
		whereList = append(whereList, fmt.Sprintf("%s=?", key))
	}
	//得到一个更新用的sql.Stmt
	updateStmt, err := t.DB.Prepare(fmt.Sprintf("update %s set %s where %s", t.Name, strings.Join(setList, ","),
		strings.Join(whereList, " and ")))
	if err != nil {
		return
	}
	defer updateStmt.Close()
	for _, row := range rows {
		param := []interface{}{}
		for _, col := range t.ColumnNames {
			param = append(param, row[col])
		}
		//先插入数据
		if _, err = insertStmt.Exec(param...); err != nil {
			//插入失败则更新数据
			if _, err = updateStmt.Exec(append(mapfun.ValuesByKeys(row, t.ColumnNames...), t.KeyValues(row)...)...); err != nil {
				return
			}
			updNum++
			continue
		}
		insNum++
	}
	return
}

//Replace 将一批记录替换成另一批记录，自动删除旧在新中不存在，插入新在旧中不存在的，更新主键相同的
func (t *Table) Replace(oldRows, newRows []map[string]interface{}) (insCount, upCount, delCount int64, err error) {

	if delCount, err = t.Delete(mapfun.Difference(oldRows, newRows, t.PrimaryKeys)); err != nil {
		return
	}
	updateRowsOld, updateRowsNew := mapfun.Intersection(oldRows, newRows, t.PrimaryKeys)
	for i, v := range updateRowsOld {
		var up int64
		if up, err = t.Update(v, updateRowsNew[i]); err != nil {
			return
		}
		upCount += up
	}
	insertRows := mapfun.Difference(newRows, oldRows, t.PrimaryKeys)
	insCount = int64(len(insertRows))
	err = t.Insert(insertRows)
	return
}
func (t *Table) pgMergeForNotNull(tabName string, cols ...string) error {
	updateCols := []string{}
	linkCols := []string{}
	insertNotExistsLink := []string{}
	pkMap := map[string]struct{}{}
	for _, c := range t.PrimaryKeys {
		pkMap[c] = struct{}{}
		linkCols = append(linkCols, fmt.Sprintf("%s.%s=t.%[2]s", t.Name, c))
		insertNotExistsLink = append(insertNotExistsLink, fmt.Sprintf("dest.%s=src.%[1]s", c))
	}
	for _, col := range cols {
		if _, ok := pkMap[col]; !ok {
			updateCols = append(updateCols, fmt.Sprintf("%s=t.%[1]s", col))
		}
	}
	updateSQL := fmt.Sprintf("update %s set %s from %s t where %s",
		t.Name, strings.Join(updateCols, ","), tabName, strings.Join(linkCols, " and "))
	if _, err := t.DB.Exec(updateSQL); err != nil {
		return err
	}
	//再insert
	insertSQL := fmt.Sprintf(
		"insert into %s(%s)select %[2]s from %[3]s src where not exists(select 1 from %[1]s dest where %[4]s)",
		t.Name, strings.Join(cols, ","), tabName, strings.Join(insertNotExistsLink, " and "))
	_, err := t.DB.Exec(insertSQL)
	return err

}

//Merge 将另一个表中的数据合并进本表，要求两个表的主键相同,相同主键的被覆盖
func (t *Table) Merge(tabName string, cols ...string) error {
	colMap := map[string]struct{}{}
	for _, c := range cols {
		colMap[c] = struct{}{}
	}
	//判断是否是postgres
	if IsPostgres(t.Driver) {
		//是否有非空字段而且不在字段范围内
		hasNotNull := false
		for _, col := range t.Columns {
			if _, ok := colMap[col.Name]; !ok && !col.Null {
				hasNotNull = true
				break
			}
		}
		if hasNotNull {
			//改用分离的update和insert
			return t.pgMergeForNotNull(tabName, cols...)
		}
	}
	strSQL := Find(t.Driver).Merge(t.FullName(), "select * from "+tabName, t.PrimaryKeys, cols)
	_, err := t.DB.Exec(strSQL)
	return err
}
