package common

import "log"

//BatchRunAndPrint 批量运行sql并打印
func BatchRunAndPrint(db Execer, list []string) error {
	for _, one := range list {
		log.Println(one)
		if _, err := db.Exec(one); err != nil {
			return err
		}
	}
	return nil
}
func BatchRun(db Execer, list []string) error {
	for _, one := range list {

		if _, err := db.Exec(one); err != nil {
			return err
		}
	}
	return nil
}
