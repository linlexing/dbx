package data

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/linlexing/dbx/common"
	"github.com/patrickmn/go-cache"
)

var (
	bindCache = cache.New(5*time.Minute, 10*time.Minute)
)

//Bind 将?占位符替换成实际驱动的占位符
//发现sqlx.Rebind分配大量内存，加上缓存
func Bind(driver, strSQL string) string {
	key := driver + ":" + strSQL
	if x, found := bindCache.Get(key); found {
		return x.(string)
	}
	x := sqlx.Rebind(sqlx.BindType(driver), strSQL)
	bindCache.Set(key, x, cache.DefaultExpiration)
	return x
}

//In 将一个参数类型是数组的参数对应的?，扩展成多个?，并把参数也扁平化
func In(query string, args ...interface{}) (string, []interface{}, error) {
	return sqlx.In(query, args...)
}

//AsInt 返回整形
func AsInt(db common.Queryer, strSQL string, args ...interface{}) (r int64, err error) {
	c := []byte{}
	row := db.QueryRow(strSQL, args...)
	if err := row.Scan(&c); err != nil {
		log.Println(err)
		err = common.NewSQLError(err, strSQL)
	}
	str := string(c)
	if strings.Contains(str, ".") {
		var f float64
		f, err = strconv.ParseFloat(str, 64)
		if err != nil {
			err = common.NewSQLError(err, strSQL)
		}
		r = int64(f)
	} else {
		r, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			err = common.NewSQLError(err, strSQL)
		}
	}
	return
}

//RunAtTx 在一个事务中运行，自动处理commit 和rollback
func RunAtTx(db common.TxDB, callback func(common.Txer) error) (err error) {
	var tx common.Txer
	if tx, err = db.Begin(); err != nil {
		return err
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
