package common

import (
	"fmt"

	"github.com/linlexing/dbx/suid"
)

// GetTempTableName 获取一个临时表名
func GetTempTableName(prev string) (string, error) {

	if len(prev) == 0 {
		return "", fmt.Errorf("prev can't empty")
	}
	id, err := suid.Next()
	if err != nil {
		return "", err
	}
	return prev + id, nil
}
