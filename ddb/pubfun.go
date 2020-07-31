package ddb

import (
	"database/sql"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

//RunAtTx 在一个事务中运行，自动处理commit 和rollback
func RunAtTx(db TxDB, callback func(Txer) error) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return
	}

	finish := false
	defer func() {
		//如果没有设置，说明是中途跳出，发生了异常
		//这里不捕获异常是要保留现场
		if !finish {
			tx.Rollback()
		}
	}()
	if err = callback(tx); err != nil {
		tx.Rollback()
	} else {
		err = tx.Commit()
	}
	finish = true
	return
}

//Openx 打开一个数据库连接，返回一个包装过的DB对象，其能返回DriverName
func Openx(driverName, cnt string) (TxDB, error) {
	rev := &db{
		driverName:    driverName,
		connectString: cnt,
	}
	err := rev.connect()
	return rev, err
}

//ScanStrings 扫描一个单列的查询，并返回一个字符串数组
func ScanStrings(db common.Queryer, strSQL string, args ...interface{}) (strs []string, err error) {

	rows, err := db.Query(strSQL, args...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, args...)
		return
	}
	defer rows.Close()

	strs = []string{}
	for rows.Next() {
		var str sql.NullString
		if err = rows.Scan(&str); err != nil {
			return
		}
		strs = append(strs, str.String)
	}
	err = rows.Err()
	return
}

//Columns 根据一个sql返回列名
func Columns(db common.Queryer, strSQL string, args ...interface{}) (strs []string, err error) {

	rows, err := db.Query(strSQL, args...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, args...)
		return
	}
	defer rows.Close()
	strs, err = rows.Columns()
	return
}

//QueryMaps 获取一个查询的所有记录，智能识别其类型
func QueryMaps(db common.Queryer, strSQL string,
	args ...interface{}) (rev []map[string]interface{}, cols []string, err error) {

	rows, err := db.Query(strSQL, args...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, args...)
		return
	}
	defer rows.Close()
	cols, err = rows.Columns()
	if err != nil {
		return
	}
	rev = []map[string]interface{}{}
	for rows.Next() {
		outList := []interface{}{}
		for range cols {
			outList = append(outList, new(interface{}))
		}
		if err = rows.Scan(outList...); err != nil {
			return
		}
		row := map[string]interface{}{}
		for i, col := range cols {
			//将[]byte 转换成string
			v := *(outList[i].(*interface{}))
			switch tv := v.(type) {
			case []byte:
				v = string(tv)
			}
			row[col] = v
		}
		rev = append(rev, row)
	}
	err = rows.Err()
	return
}

//GetTempTableName 获取一个临时表名
func GetTempTableName(db DB, prev string) (string, error) {

	if len(prev) == 0 {
		return "", fmt.Errorf("prev can't empty")
	}
	//确定名称
	tableName := ""
	rand.Seed(time.Now().UnixNano())
	bys := make([]byte, 4)
	icount := 0
	for {
		binary.BigEndian.PutUint32(bys, rand.Uint32())
		tableName = fmt.Sprintf("%s%X", prev, bys)
		if exists, err := schema.Find(db.DriverName()).
			TableExists(db, tableName); err != nil {
			return "", err
		} else if !exists {
			break
		}
		icount++
		if icount > 100 {
			return "", fmt.Errorf("find table name too much")
		}
	}
	return tableName, nil
}
