package data

import "github.com/linlexing/dbx/common"

//Accesser 不同数据库驱动需要实现的数据访问类
type Accesser interface {
	Merge(db common.DB, destTable, srcTable string, pks, columns []string) error
	Minus(db common.DB, table1, where1, table2, where2 string, primaryKeys, cols []string) string
}

var (
	metas = map[string]Accesser{}
)

//Register 注册一个数据库接口，其实现了指定的方法
func Register(driver string, da Accesser) {
	metas[driver] = da
}

//Find 根据一个驱动找到正确的Ps
func Find(driver string) Accesser {
	if v, ok := metas[driver]; !ok {
		panic(driver + " not registe dataaccess")
	} else {
		return v
	}

}
