package data

import (
	"context"
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

// Bind 将?占位符替换成实际驱动的占位符
// 发现sqlx.Rebind分配大量内存，加上缓存
func Bind(driver, strSQL string) string {
	key := driver + ":" + strSQL
	if x, found := bindCache.Get(key); found {
		return x.(string)
	}
	x := Rebind(BindType(driver), strSQL)
	bindCache.Set(key, x, cache.DefaultExpiration)
	return x
}

// In 将一个参数类型是数组的参数对应的?，扩展成多个?，并把参数也扁平化
func In(query string, args ...interface{}) (string, []interface{}, error) {
	return sqlx.In(query, args...)
}

// AsInt 返回整形
func AsInt(db common.Queryer, strSQL string, args ...interface{}) (r int64, err error) {
	return AsIntContext(db, context.Background(), strSQL, args...)
}

// AsIntContext 返回整形
func AsIntContext(db common.Queryer, ctx context.Context, strSQL string, args ...interface{}) (r int64, err error) {
	c := []byte{}
	row := db.QueryRowContext(ctx, strSQL, args...)
	if err = row.Scan(&c); err != nil {
		log.Println(err)
		err = common.NewSQLError(err, strSQL)
		return
	}
	str := string(c)
	if strings.Contains(str, "e") || strings.Contains(str, ".") {
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
