package common

import (
	"errors"
	"log"
)

// BatchRunAndPrint 批量运行sql并打印
func BatchRunAndPrint(db Execer, list []string, param ...interface{}) error {
	for _, one := range list {
		log.Println(one)
		if _, err := db.Exec(one, param...); err != nil {
			return err
		}
	}
	return nil
}
func BatchRun(db Execer, list []string, param ...[]interface{}) error {
	if len(param) == 0 {
		param = make([][]interface{}, len(list))
	} else if len(param) != len(list) {
		return errors.New("sql与param数量不符")
	}
	for i, one := range list {
		if _, err := db.Exec(one, param[i]...); err != nil {
			return err
		}
	}
	return nil
}
