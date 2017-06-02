package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/linlexing/dbx/common"
)

//Bind 将?占位符替换成实际驱动的占位符
func Bind(driver, strSQL string) string {
	return sqlx.Rebind(sqlx.BindType(driver), strSQL)
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
