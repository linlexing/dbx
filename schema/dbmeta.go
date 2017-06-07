package schema

import (
	"fmt"

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
	CreateTableAs(db common.DB, tableName, strSQL string, pks []string) error
	TableExists(db common.DB, tableName string) (bool, error)
	CreateTable(db common.DB, table *Table) error
	OpenTable(db common.DB, tableName string) (*Table, error)
	ChangeTable(db common.DB, change *TableSchemaChange) error
	TableNames(db common.DB) (names []string, err error)
	DropIndexIfExists(db common.DB, indexName,tableName string) error
	CreateIndexIfNotExists(db common.DB, indexName, tableName, express string) error
}

//Register 注册一个新元数据操作驱动。
func Register(driverName string, meta Meta) {
	metas[driverName] = meta
}

//Find 根据实际的数据库连接返回一个元数据操纵类，缓存
func Find(dname string) Meta {
	if one, ok := metas[dname]; ok {
		return one
	}
	panic(fmt.Errorf("not found %s meta", dname))

}
