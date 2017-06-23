package pageselect

import (
	"github.com/linlexing/dbx/scan"
)

//ColumnTypes 是对[]ColumnType的包装，主要增加搜索方法
type ColumnTypes []*scan.ColumnType

func (c ColumnTypes) byName(name string) *scan.ColumnType {
	for _, v := range c {
		if v.Name == name {
			return v
		}
	}
	return nil
}
