package data

import "github.com/linlexing/dbx/common"

//不同数据库驱动需要实现的数据访问类
type DataAccesser interface {
	Merge(db common.DB, destTable, srcTable string, pks, columns []string) error
}

var (
	metas = map[string]DataAccesser{}
)

//Register 注册一个数据库接口，其实现了指定的方法
func Register(driver string, da DataAccesser) {
	metas[driver] = da
}

//Find 根据一个驱动找到正确的Ps
func Find(driver string) DataAccesser {
	if v, ok := metas[driver]; !ok {
		panic(driver + " not registe dataaccess")
	} else {
		return v
	}

}
