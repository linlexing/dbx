package mysql

import (
	"github.com/linlexing/dbx/data"
)

const driverName = "mysql"

type meta struct{}

func init() {
	data.Register(driverName, new(meta))
}
func (m *meta) Merge(tabName string, skipUpdateColumns ...string) error {
	panic("not impl")
}
