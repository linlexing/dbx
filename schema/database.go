package schema

import (
	"github.com/linlexing/dbx/common"
)

// DataBaseInfo 库结构
type DataBaseInfo struct {
	UserName string // 库用户
	PassWord string // 用户密码
	DBName   string // 库名称
	DBSize   int64  // 库大小 单位M
	DataFile string // 物理文件路径
}

// CreateSchema 创建
func CreateSchema(driver string, db common.DB, dbInfo DataBaseInfo) error {
	mt := Find(driver)
	rev, err := mt.CreateSchemaSQL(db, dbInfo)
	if err != nil {
		return err
	}
	return common.BatchRunAndPrint(db, rev)
}

// DropSchema 删除
func DropSchema(driver string, db common.DB, dbInfo DataBaseInfo) error {
	mt := Find(driver)
	list, err := mt.DropSchemaSQL(db, dbInfo)
	if err != nil {
		return err
	}
	return common.BatchRunAndPrint(db, list)
}
