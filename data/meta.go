package data

import (
	"strings"
)

// 目标与来源字段的映射
type ColMap struct {
	Dest string
	Src  string
}

// Accesser 不同数据库驱动需要实现的数据访问类
type Accesser interface {
	//merge操作，目标和源数据中的字段名要一致
	Merge(destTable, srcDataSQL string, pks, columns []string) string
	//关联更新，目标和源数据中的字段名要一致
	UpdateFrom(destTable, srcDataSQL, additionSet string, pk, columns []ColMap) string
	Minus(table1, where1, table2, where2 string, primaryKeys, cols []string) string
	//Concat 串联字符串，null被忽略，而不是返回null
	Concat(vals ...string) string
	Concat_ws(separator string, vals ...string) string
}

var (
	metas = map[string]Accesser{}
)

// Register 注册一个数据库接口，其实现了指定的方法
func Register(driver string, da Accesser) {
	metas[driver] = da
}

// Find 根据一个驱动找到正确的Ps
func Find(driver string) Accesser {
	//sqlite3可能会有不同的变种
	if strings.HasPrefix(driver, "sqlite3") {
		driver = "sqlite3"
	}
	if v, ok := metas[driver]; !ok {
		panic(driver + " not registe dataaccess")
	} else {
		return v
	}

}
