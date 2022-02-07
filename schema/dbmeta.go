package schema

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/common"
)

var (
	metas = map[string]Meta{}
)

//ChangedField 存贮字段变更信息
type ChangedField struct {
	OldField *Column
	NewField *Column
}

//TableSchemaChange 存贮一个表结构变化的数据
type TableSchemaChange struct {
	OldName      string
	NewName      string
	PKChange     bool
	PK           []string
	OriginFields []*Column
	ChangeFields []*ChangedField
	RemoveFields []string
}

//Meta 是数据库操作元数据的接口，任意go标准的sql驱动，实现了这个接口就可以使用dbx
type Meta interface {
	CreateTableAsSQL(db common.DB, tableName, strSQL string, pks []string) ([]string, error)
	TableExists(db common.DB, tableName string) (bool, error)
	CreateTableSQL(db common.DB, table *Table) ([]string, error)
	OpenTable(db common.DB, tableName string) (*Table, error)
	ChangeTableSQL(db common.DB, change *TableSchemaChange) ([]string, error)
	TableNames(db common.DB) (names []string, err error)
	DropIndexIfExistsSQL(db common.DB, indexName, tableName string) ([]string, error)
	CreateIndexIfNotExistsSQL(db common.DB, unique bool, indexName, tableName, express string) ([]string, error)
	CreateSchemaSQL(db common.DB, dbInfo DataBaseInfo) ([]string, error)
	DropSchemaSQL(db common.DB, dbInfo DataBaseInfo) ([]string, error)
}

//Register 注册一个新元数据操作驱动。
func Register(driverName string, meta Meta) {
	metas[driverName] = meta
}

//Find 根据实际的数据库连接返回一个元数据操纵类，缓存
func Find(drivername string) Meta {
	//sqlite3可能会有不同的变种
	if strings.HasPrefix(drivername, "sqlite3") {
		drivername = "sqlite3"
	}
	if one, ok := metas[drivername]; ok {
		return one
	}
	panic(fmt.Errorf("not found %s meta", drivername))

}
