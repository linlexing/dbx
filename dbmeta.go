package dbx

import "fmt"

var (
	metas = map[string]DBMeta{}
)

//ChangedField 存贮字段变更信息
type ChangedField struct {
	OldField *DBTableColumn
	NewField *DBTableColumn
}

//TableSchemaChange 存贮一个表结构变化的数据
type TableSchemaChange struct {
	PKChange     bool
	PK           []string
	ChangeFields []*ChangedField
	RemoveFields []string
}

//DBMeta 是数据库操作元数据的接口，任意go标准的sql驱动，实现了这个接口就可以使用dbx
type DBMeta interface {
	IsNull() string
	CreateTableAs(db DB, tableName, strSQL string, pks []string) error
	TableExists(db DB, tableName string) (bool, error)
	TableRename(db DB, oldName, newName string) error
	RemoveColumns(db DB, tabName string, cols []string) error
	TableNames(db DB) (names []string, err error)
	CreateColumnIndex(db DB, tableName, colName string) error
	DropColumnIndex(db DB, tableName, indexName string) error
}

//RegisterMeta 注册一个新的驱动
func RegisterMeta(driverName string, meta DBMeta) {
	metas[driverName] = meta
}

//Meta 根据实际的数据库连接返回一个元数据操纵类，缓存
func Meta(db DB) DBMeta {
	if one, ok := metas[db.DriverName()]; ok {
		return one
	}
	panic(fmt.Errorf("not found %s meta", db.DriverName()))

}
