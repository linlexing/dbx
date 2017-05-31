package data

//在一个事务中运行，自动处理commit 和rollback
func runAtTx(db txDB, callback func(txer) error) (err error) {
	var tx txer
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
