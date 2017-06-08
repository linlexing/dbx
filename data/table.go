package data

import (
	"bytes"
	"database/sql"
	"errors"
	"log"

	"encoding/gob"
	"fmt"

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
	driver string
	DB     common.DB
	*schema.Table

	ColumnTypes    []*scan.ColumnType
	notnullColumns []string
	ColumnNames    []string
	columnsMap     map[string]*schema.Column //用于快速查询
}

//NewTable 用schema.Table构造一个Table,没有数据库操作发生
func NewTable(driver string, db common.DB, st *schema.Table) *Table {
	rev := &Table{
		driver:         driver,
		DB:             db,
		ColumnNames:    []string{},
		Table:          st,
		ColumnTypes:    []*scan.ColumnType{},
		notnullColumns: []string{},
		columnsMap:     map[string]*schema.Column{},
	}
	//构造索引

	for _, col := range st.Columns {
		rev.ColumnNames = append(rev.ColumnNames, col.Name)
		if !col.Null {
			rev.notnullColumns = append(rev.notnullColumns, col.Name)
		}
		rev.ColumnTypes = append(rev.ColumnTypes, &scan.ColumnType{
			Name: col.Name,
			Type: col.Type,
		})
		rev.columnsMap[col.Name] = col
	}
	return rev
}

//OpenTable 从数据库取出结构构造Table,表名用 schema.tablename的方式,
//如果不读取数据，仅定义结构，应当使用schema.Table
func OpenTable(driver string, db common.DB, tabName string) (*Table, error) {
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
	return Bind(t.driver, strSQL)
}

//Row 根据一个主键值返回一个记录,如果没有找到返回一个错误
func (t *Table) Row(pks ...interface{}) (map[string]interface{}, error) {
	whereList := []string{}
	if len(t.PrimaryKeys) != len(pks) {
		return nil, fmt.Errorf("the table %s pk values number error.table pk:%#v,pkvalues:%#v", t.FullName(), t.PrimaryKeys, pks)
	}
	for _, onePk := range t.PrimaryKeys {
		whereList = append(whereList, fmt.Sprintf("%s=?", onePk))
	}
	rows, err := t.Query(strings.Join(whereList, " and\n"), pks...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.New("the record not found")
	}
	return t.ScanMap(rows)
}

//ToJSON 将一行数据转换成json,日期、二进制、int64数据转换成文本
func (t *Table) ToJSON(row map[string]interface{}) (map[string]interface{}, error) {
	transRecord := map[string]interface{}{}
	for _, col := range t.Columns {
		v, err := col.Type.ToJSON(row[col.Name])
		if err != nil {
			return nil, err
		}
		transRecord[col.Name] = v
	}

	return transRecord, nil
}

//FromJSON 将一个json数据转换回row
func (t *Table) FromJSON(row map[string]interface{}) (map[string]interface{}, error) {
	transRecord := map[string]interface{}{}
	for _, col := range t.Columns {
		v, err := col.Type.ParseJSON(row[col.Name])
		if err != nil {
			return nil, err
		}
		transRecord[col.Name] = v
	}

	return transRecord, nil
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
	return

}

//KeyValues 返回一个记录的主键值
func (t *Table) KeyValues(row map[string]interface{}) []interface{} {
	rev := []interface{}{}
	for _, v := range t.PrimaryKeys {
		rev = append(rev, row[v])
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

		strSQL += "where " + where
	}
	rows, err := t.DB.Query(t.bind(strSQL), params...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, params...)
		log.Println(err)
		return
	}
	defer rows.Close()
	rev = rows.Next()
	return
}

//Count 统计表中记录数，其实没什么逻辑，就是省了组合一个sql语句
func (t *Table) Count(where string, params ...interface{}) (rev int64, err error) {
	strSQL := "select count(*) from " + t.FullName()
	if len(where) > 0 {

		strSQL += "where " + where
	}
	err = t.DB.QueryRow(t.bind(strSQL), params...).Scan(&rev)
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
	if t.driver == "oci8" {
		for k, v := range row {
			if tm, ok := v.(time.Time); ok {
				row[k] = truncateTimeZone(tm)
			}
		}
	}
	if err := t.checkNotNull(row); err != nil {
		return err
	}
	return nil
}

//ImportFromTable 从另一个表中导入数据，表中列数量、名称、类型必须一致
func (t *Table) ImportFromTable(srcTable *Table, progressFunc func(string), where string,
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
func (t *Table) ImportFrom(db common.Queryer, progressFunc func(string), query string,
	args ...interface{}) (iCount int64, err error) {
	var rowCount int64
	strSQL := fmt.Sprintf("select count(*) from (%s) out_count", query)
	if err = db.QueryRow(strSQL, args...).Scan(&rowCount); err != nil {
		err = common.NewSQLError(err, strSQL, args...)
		return
	}
	progressFunc(fmt.Sprintf("start import table %s,total %d records", t.FullName(), rowCount))

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
			progressFunc(fmt.Sprintf("\t%.2f%%\t%d/%d\t%.2fs", 100.0*float64(iCount)/float64(rowCount), iCount, rowCount, totalSec))
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
	progressFunc(fmt.Sprintf("%s,total %d records imported %.2fs", t.FullName(), iCount, time.Since(beginTime).Seconds()))
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
	for k := range rows[0] {
		cols = append(cols, k)
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
			return
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

//RemoveByKeyValues 根据一个主键值，删除记录，如果没有记录被删除，则返回一个错误
func (t *Table) RemoveByKeyValues(keyValues ...interface{}) (delCount int64, err error) {
	if len(keyValues) != len(t.PrimaryKeys) {
		err = errors.New("the key number not equ primary keys")
		return
	}
	whereList := []string{}
	for _, v := range t.PrimaryKeys {
		whereList = append(whereList, fmt.Sprintf("%s=?", v))
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
func (t *Table) buildWhere(row map[string]interface{}) (string, []interface{}) {
	strWhere := []string{}
	newRow := []interface{}{}
	for k, v := range row {
		//如果是没有长度的string，即text，以及bytea、datetime不参与where条件
		fld := t.ColumnByName(k)
		if fld.Type == schema.TypeBytea ||
			fld.Type == schema.TypeDatetime ||
			(fld.Type == schema.TypeString && fld.MaxLength <= 0) {
			continue
		}
		if v == nil {
			strWhere = append(strWhere, fmt.Sprintf("%s is null", k))
		} else {
			strWhere = append(strWhere, fmt.Sprintf("%s=?", k))
		}
		newRow = append(newRow, v)
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
	return
}

//Update 只有修改过的字段才被更新，where采用全部旧值判断（没有长度的string将不参与，因为oracle会出错）
func (t *Table) Update(oldData, newData map[string]interface{}) (upCount int64, err error) {
	if oldData == nil || len(oldData) == 0 || newData == nil || len(newData) == 0 {
		err = errors.New("data is empty")
		return
	}
	if len(oldData) != len(newData) {
		err = errors.New("the old and new record,field number not same")
		return
	}
	if err = t.checkAndConvertRow(oldData); err != nil {
		return
	}
	if err = t.checkAndConvertRow(newData); err != nil {
		return
	}
	whereStr, whereVals := t.buildWhere(oldData)
	return t.UpdateByWhere(newData, whereStr, whereVals...)
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

//Merge 将另一个表中的数据合并进本表，要求两个表的主键相同,相同主键的被覆盖
func (t *Table) Merge(tabName string, cols ...string) error {
	return Find(t.driver).Merge(t.DB, t.FullName(), tabName, t.PrimaryKeys, cols)
}
