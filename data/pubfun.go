package data

import (
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
