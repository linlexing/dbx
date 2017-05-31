package data

import (
	"bytes"
	"database/sql"
"errors"

	"encoding/gob"
	"fmt"


	"strconv"
	"strings"
	"time"


	log "github.com/Sirupsen/logrus"

	"github.com/linlexing/mapfun"

	"github.com/jmoiron/sqlx"
	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/scan"
	"github.com/linlexing/dbx/schema"
)

type txDB interface {
	Begin() (txer, error)
}
type txer interface {
	Prepare(query string) (*sql.Stmt, error)
	Commit() error
	Rollback() error
}

//Table 表现一个数据库表,扩展了schema.Table ，提供了数据访问
type Table struct {
	driver string
	DB     common.DB
	*schema.Table
	columnTypes    []*scan.ColumnType
	notnullColumns []string
	ColumnNames    []string
	columnsMap     map[string]*schema.Column //用于快速查询
}

//NewTable 返回一个数据表，表名用 schema.tablename的方式，立即打开表，获取其结构
//因此，如果不读取数据，仅定义结构，应当使用schema.Table
func NewTable(driver string, db common.DB, tabName string) *Table {
	if len(tabName) == 0 {
		log.Panic("table name is empty")
	}
	st, err := schema.Find(driver).OpenTable(db, tabName)
	if err != nil {
		log.Panic(err)
	}
	rev := &Table{
		driver:         driver,
		DB:             db,
		ColumnNames:    []string{},
		Table:          st,
		columnTypes:    []*scan.ColumnType{},
		notnullColumns: []string{},
		columnsMap:     map[string]*schema.Column{},
	}
	//构造索引

	for _, col := range st.Columns {
		rev.ColumnNames = append(rev.ColumnNames, col.Name)
		if !col.Null {
			rev.notnullColumns = append(rev.notnullColumns, col.Name)
		}
		rev.columnTypes = append(rev.columnTypes, &scan.ColumnType{
			Name: col.Name,
			Type: col.Type,
		})
		rev.columnsMap[col.Name] = col
	}
	return rev
}
func (t *Table) bind(strSQL string) string {
	return sqlx.Rebind(sqlx.BindType(t.driver), strSQL)
}

//Row 根据一个主键值返回一个记录,如果没有找到返回一个错误
func (t *Table) Row(pks ...interface{}) (map[string]interface{}, error) {
whereList :=[]string{}
	if len(t.PrimaryKeys) != len(pks) {
		return nil, fmt.Errorf("the table %s pk values number error.table pk:%#v,pkvalues:%#v", t.FullName(), t.PrimaryKeys, pks)
	}
	for i, v := range pks {
		whereList = append(whereList, fmt.Sprintf("%s=?",t.PrimaryKeys[i]))
	}
	rows, err := t.QueryRows(strings.Join(whereList," and\n"),pks...)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, errors.New("the record not found")
	}
	return rows[0], nil
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

//QueryRowsOrder 查询返回排序记录，where如有参数，必须用 ?
func (t *Table) QueryRowsOrder(orderby []string, where string, param ...interface{}) (record []map[string]interface{}, err error) {

	if len(where) > 0 {
		where = " where " + where
	}
	columnsStr := strings.Join(t.ColumnNames, ",")

	strOrderby := ""
	if len(orderby) > 0 {
		strOrderby = " order by " + strings.Join(orderby, ",")
	}
	strSQL := t.bind(fmt.Sprintf("select %s from %s%s%s", columnsStr, t.FullName(), where, strOrderby))

	var rows *sql.Rows
	if rows, err = t.DB.Query(strSQL, param...); err != nil {
		err = common.NewSQLError(err, strSQL, param)

		return
	}
	record = []map[string]interface{}{}
	defer rows.Close()
	for rows.Next() {
		oneRecord := map[string]interface{}{}
		var outList []interface{}
		if outList, err = scan.TypeScan(rows, t.columnTypes); err != nil {
			return
		}
		for i, one := range outList {
			oneRecord[t.Columns[i].Name] = one
		}
		record = append(record, oneRecord)
	}
	return
}

//QueryRows 查询返回记录，无排序
func (t *Table) QueryRows(where string, param  ...interface{}) (record []map[string]interface{}, err error) {
	return t.QueryRowsOrder(nil,where, param...)
}

//KeyExists 检查一个主键是否存在
func (t *Table) KeyExists(pks ...interface{}) (result bool, err error) {
	strWhere := []string{}

	for _, v := range t.PrimaryKeys {
		strWhere = append(strWhere, v+"=?")
	}

	var rows *sql.Rows
	strSQL := fmt.Sprintf("select 1 from %s%s", t.FullName(), strings.Join(strWhere, " and "))
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

//Count 统计表中记录数，其实没什么逻辑，就是省了组合一个sql语句
func (t *Table) Count(where string, params ...interface{}) (rev int64, err error) {

	var pam map[string]interface{}
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

//ImportFrom 从一个sql语句导入数据,sql语句返回的列必须与表中数量一致,因为可能是异构数据库
//所以不能用直接的CreateTableAs,由于数据可能比较多，采用5秒钟提交一次事务，所以Table.DB不能
//是事务Tx
func (t *Table) ImportFrom(dataDB common.DB, strSQL string, progressFunc func(string)) (err error) {
	rowCount, err := t.Count(dataDB, strSql, nil)
	if err != nil {
		log.Println(err)
		return
	}

	progressFunc(fmt.Sprintf("start CreateAs table %s,total %d records", t.Name(), rowCount))
	//创建表

	rows, err := dataDB.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return
	}
	defer rows.Close()

	var icount, batCount int64 = 0, 0

	//再构造insert语句
	insertSQL := t.InsertSQL()
	//再开始事务

	startTime := time.Now()
	beginTime := startTime
	finished := false
	tx, err := t.DB.(txDB).Begin()
	if err != nil {
		return err
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
		return err
	}
	icount = 0
	batCount = 0
	for rows.Next() {
		var outList []interface{}

		if outList, err = scan.TypeScan(rows, t.columnTypes); err != nil {

			return err
		}

		if _, err = insertStmt.Exec(outList...); err != nil {
			log.Printf("error:%s,values:\n", err)
			for ei, ev := range outList {
				log.Printf("\t%s=%#v", t.ColumnByName(ei).Name, ev)
			}

			return err
		}
		icount++
		batCount++
		totalSec := time.Since(startTime).Seconds()
		if totalSec >= 5 {
			if err = insertStmt.Close(); err != nil {
				log.Println(err)

				return err
			}
			if err = tx.Commit(); err != nil {
				log.Println(err)
				return err
			}
			batCount = 0
			tx = nil
			tx, err = t.DB.Begin()
			if err != nil {
				log.Println(err)
				return err
			}
			insertStmt, err = tx.Prepare(insertSql)
			if err != nil {
				log.Println(err)

				return err
			}
			progressFunc(fmt.Sprintf("\t%.2f%%\t%d/%d\t%.2fs", 100.0*float64(icount)/float64(rowCount), icount, rowCount, totalSec))
			startTime = time.Now()
		}
	}

	if err = rows.Err(); err != nil {
		log.Println(err)

		return err
	}
	if batCount > 0 {
		if err = insertStmt.Close(); err != nil {
			log.Println(err)

			return err
		}
		if err = tx.Commit(); err != nil {
			tx = nil
			return err
		}
		finished = true
	}

	progressFunc(fmt.Sprintf("%s,total %d records imported %.2fs", t.Name(), icount, time.Since(beginTime).Seconds()))

	return

}

//InsertSQL 生成一个InsertSQL
func (t *Table) InsertSQL() string {
	insertSQL := fmt.Sprintf(
		"insert into %s(%s)values(%s)",
		t.FullName(), strings.Join(t.ColumnNames, ","),
		strings.Join(strings.Split(strings.Repeat("?", len(t.Columns)), ""), ","))
	insertSql = t.bind(insertSql)
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
		err = common.NewSQLError(err, strSql, data...)
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
			return []byte(fmt.Sprintf("%d",keys[0]))
		case []byte:
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
			i,err := strconv.ParseInt(string(key),10,64)
			if err !=nil{
				panic("invalid int data"+ string(key))
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
	if len(rows)==0{
		return nil
	}
	if len(rows) == 1 {
		if  err= t.checkAndConvertRow(rows[0]); err != nil {
			
			return
		} 
			return t.insertAsPack(rows[0])
		
	}
	cols :=[]string{}
	for k,_:=range rows[0]{
		cols = append(cols,k)
	}
	strSQL := fmt.Sprintf("insert into %s(%s)values(%s)",
	t.FullName(),strings.Join(cols,","),
	strings.Join(strings.Split(strings.Repeat("?",len(cols),""),",")))
	stmt,err := t.DB.Prepare(strSQL)
	if err !=nil{
		err = common.NewSQLError(err,strSQL)
		log.Println(err)
		return
	}
	defer stmt.Close()
	//先检查并插入数据
	
	for _, row := range rows {
		if err= t.checkAndConvertRow(row); err != nil {
			
			return
		}
		data :=[]interface{}{}
		for _,col :=range cols{
			data = append(data,row[col])
		}
		if err = stmt.Exec(data...);err !=nil{
			return
		}
	
	}
	
	return
}

//Delete 删除记录，全部字段值将被生成where字句(text、bytea除外)
func (t *Table) Delete(rows []map[string]interface{}) (delCount int64,err error) {
	//考虑到null值，所有的行不能用一个语句，必须单独删除
	for _, v := range rows {
		var i int64
		if i,err = t.Remove(v); err != nil {
			return
		}
		delCount += i
	}
	return
}
//RemoveByKeyValues 根据一个主键值，删除记录，如果没有记录被删除，则返回一个错误
func (t *Table) RemoveByKeyValues(keyValues ...interface{}) (delCount int64,err error) {
	if len(keyValues) != len(t.PrimaryKeys){
		return errors.New("the key number not equ primary keys")
	}
	whereList :=[]string{}
	for i,v :=range keyValues{
		whereList = append(whereList,fmt.Sprintf("%s=?",t.PrimaryKeys[i]))
	}
	strSQL :=t.bind( fmt.Sprintf("delete from %s where %s",t.FullName(),
	strings.Join(whereList," and\n")))


	sr,err :=t.DB.Exec(strSQL,keyValues...)
	if err !=nil{
		err = common.NewSQLError(err,strSQL,keyValues...)
		log.Println(err)
		return
	}
	delCount,err = sr.RowsAffected()
return
			

}
func (t *Table)buildWhere(row map[string]interface{})(string,[]interface{}){
		strWhere := []string{}
	newRow := []interface{}{}
	for k, v := range row {
		//如果是没有长度的string，即text，以及bytea、datetime不参与where条件
		fld := t.ColumnByName(k)
		if fld.Type == schema.TypeBytea ||
		fld.Type == schema.TypeDatetime||	
		 (fld.Type == schema.TypeString && fld.MaxLength <=0) {
			 continue
		 }
			if v == nil {
				strWhere = append(strWhere, fmt.Sprintf("%s is null", k))
			} else {
				strWhere = append(strWhere, fmt.Sprintf("%s=?", k))
			}
		newRow = append(newRow,v)
	}
	return strWhere,newRow

}
//Remove 删除一个记录，必须是全指标的记录
func (t *Table) Remove(row map[string]interface{}) (delCount int64,err error) {
	row, err = t.checkAndConvertRow(row)
	if err != nil {
		return
	}
	strWhere,newRow :=t.buildWhere(row)
	strSql :=t.bind( fmt.Sprintf(
		"delete from %s where %s", t.FullName(), strings.Join(strWhere, " and\n")))
	var sqlr sql.Result
	if sqlr, err = t.DB.Exec(strSql, newRow...); err != nil {
		err = NewSQLError(err,strSql, newRow...)
		log.Println(err)
		return
	}
	
	 delCount, err = sqlr.RowsAffected()

	return
}

//UpdateByKey 通过一个key更新记录
func (t *Table) UpdateByKey(key []interface{}, row map[string]interface{}) (upCount int64,err error) {

whereList :=[]string{}
	for i,v :=range key{
		whereList = append(whereList,fmt.Sprintf("%s=?",t.PrimaryKeys[i]))
	}
		
	return t.UpdateByWhere(row,strings.Join(whereList," and\n"),key...)
}

//UpdateByWhere 通过一个条件更新指定的字段值
func (t *Table) UpdateByWhere(row map[string]interface{},where string,params ...interface{}) (upCount int64,err error) {
	if len(row) == 0 {
		return fmt.Errorf("data is null,row:%v,where:%v,params:%#v", row, where,params)
	}

	if err = t.checkAndConvertRow(row); err != nil {
		return err
	}
	set := []string{}
	setVals :=[]interface{}{}
	for k,v :=range row{
		set = append(set,fmt.Sprintf("%s=?",k))
		setVals=append(setVals,v)
	}
	whereStr:=""
	if len(where)>0{
whereStr = "where "+where
	}
	setVals = append(setVals,params...)
	strSql :=t.bind( fmt.Sprintf("update %s set %s %s",
		t.FullName(), strings.Join(set, ","), whereStr))
	var sqlr sql.Result

	if sqlr, err = t.DB.Exec(strSql, setVals...); err != nil {
		err = common.NewSQLError(err,strSql,setVals... )
		return
	}
	 upCount, err = sqlr.RowsAffected()
	return
}

//Update 只有修改过的字段才被更新，where采用全部旧值判断（没有长度的string将不参与，因为oracle会出错）
//如果old、new中有多余字段，则会自动剔除，如果主键缺失，则会出错
func (t *Table) Update(oldData, newData map[string]interface{}) (upCount int64,err error) {
	if oldData == nil || len(oldData) == 0 || newData == nil || len(newData) == 0 {
		err= errors.New("data is empty")
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
	whereStr,whereVals := t.buildWhere(oldData)
	return	t.UpdateByWhere(newData,whereStr,whereVals...)
}

//Save 保存一个记录，先尝试用keyvalue去update，如果更新到记录为0再insert，
//逻辑上是正确的，同时，速度也会有保障
func (t *Table) Save(row map[string]interface{}) error {
	i,err := t.UpdateByKey(t.KeyValues(row),row)
	if err !=nil{
		return err
	}
	if i >0 {
		return nil
	}
	return t.insertAsPack(row)
}

//Replace 将一批记录替换成另一批记录，自动删除旧在新中不存在，插入新在旧中不存在的，更新主键相同的
func (t *Table) Replace(oldRows, newRows []map[string]interface{}) (insCount,upCounterr,delCount int64,err  error) {
	pkNames := t.PrimaryKeys()
	updateRowsOld, updateRowsNew := mapfun.Intersection(oldRows, newRows, pkNames)
	if delCount, err = t.Delete(mapfun.Difference(oldRows, newRows, pkNames)); err != nil {
		return
	}
	for i, v := range updateRowsOld {
		var up int64
		if up,err = t.Update(v, updateRowsNew[i]); err != nil {
			return
		}
		upCount += up
	}
	insertRows := mapfun.Difference(newRows, oldRows, pkNames)
	insCount = len(insertRows)
	err = t.Insert(insertRows)
	return
}

//Merge 将另一个表中的数据合并进本表，要求两个表的主键相同,相同主键的被覆盖
//skipColumns指定跳过update的字段清单
func (t *Table) Merge(tabName string, skipUpdateColumns ...string) error {
	cols := mapfun.WithoutStr( t.ColumnNames,skipUpdateColumns...)
	return Find(t.driver).Merge(t.DB,t.FullName(),tabName,t.PrimaryKeys,cols)
}
