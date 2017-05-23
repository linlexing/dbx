package dbx

import (
	"fmt"
)

var (
	metas map[string]DBMeta
)

//DBMeta 是数据库操作元数据的接口，任意go标准的sql驱动，实现了这个接口就可以使用dbx
type DBMeta interface {
	IsNull() string
	CreateTableAs(db DB, tableName, strSql string, pks []string) error
}

func RegisterMeta(driverName string, meta DBMeta) {
	metas[driverName] = meta
}
func Meta(db DB) DBMeta {
	if one, ok := metas[db.DriverName()]; ok {
		return one
	} else {
		panic(fmt.Errorf("not found %s meta", db.DriverName()))
	}
}
