package scan

import (
	"database/sql"
	"errors"
	"fmt"

	"time"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

type nullTime struct {
	Valid bool
	Time  time.Time
}

func (n *nullTime) Scan(src interface{}) error {
	switch tv := src.(type) {
	case nil:
		n.Valid = false
		println("nil")
	case time.Time:
		n.Valid = true
		n.Time = tv
	default:
		return fmt.Errorf("invalid type:%T", src)

	}
	return nil
}

//ColumnType 表示一个列，带有类型，主要是go1.8的规范，还很少有driver实现，后期可去掉类型
type ColumnType struct {
	Name string
	Type schema.DataType
}

// Scan 根据一个字段类型清单，扫描出正确的数据，null值返回nil,类型对照如下：
// TypeString	string
// TypeInt 	int64
// TypeDatetime time.Time
// TypeBytea 	[]byte
// TypeFloat 	float64
func TypeScan(s common.Scaner, cols []*ColumnType) ([]interface{}, error) {
	outList := []interface{}{}
	for _, v := range cols {
		var newV interface{}
		switch v.Type {
		case schema.TypeString:
			newV = new(sql.NullString)
		case schema.TypeInt:
			newV = new(sql.NullInt64)
		case schema.TypeDatetime:
			newV = new(nullTime)
		case schema.TypeFloat:
			newV = new(sql.NullFloat64)
		case schema.TypeBytea:
			newV = &[]byte{}
		default:
			return nil, errors.New("invalid type " + v.Type.String())
		}
		outList = append(outList, newV)
	}
	if err := s.Scan(outList...); err != nil {
		return nil, err
	}
	rev := []interface{}{}
	for i, v := range cols {
		var outV interface{}
		switch v.Type {
		case schema.TypeString:
			tv := outList[i].(*sql.NullString)
			if tv.Valid {
				outV = tv.String
			}
		case schema.TypeInt:
			tv := outList[i].(*sql.NullInt64)
			if tv.Valid {
				outV = tv.Int64
			}
		case schema.TypeDatetime:
			tv := outList[i].(*nullTime)
			if tv.Valid {
				outV = tv.Time
			}
		case schema.TypeFloat:
			tv := outList[i].(*sql.NullFloat64)
			if tv.Valid {
				outV = tv.Float64
			}
		case schema.TypeBytea:
			tv := outList[i].(*[]byte)
			if tv != nil {
				outV = *tv
			}
		default:
			return nil, errors.New("invalid type " + v.Type.String())
		}
		rev = append(rev, outV)
	}
	return rev, nil
}
