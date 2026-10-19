package scan

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

type AnyString struct {
	String string
	Valid  bool
}

func (a *AnyString) Scan(src interface{}) error {
	if src == nil {
		a.Valid = false
		a.String = ""
		return nil
	}
	a.Valid = true
	switch v := src.(type) {
	case string:
		a.String = v
	case []byte:
		a.String = string(v)
	case []interface{}:
		b, _ := json.Marshal(v)
		a.String = string(b)
	default:
		a.String = fmt.Sprintf("%v", v)
	}
	return nil
}

// ColumnType 表示一个列，带有类型，主要是go1.8的规范，还很少有driver实现，后期可去掉类型
type ColumnType struct {
	Name string
	Type schema.DataType
}

// TypeScan 根据一个字段类型清单，扫描出正确的数据，null值返回nil,类型对照如下：
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
			newV = new(AnyString)
		case schema.TypeInt:
			newV = new(sql.NullInt64)
		case schema.TypeDatetime:
			newV = new(NullTime)
		case schema.TypeFloat:
			newV = new(sql.NullFloat64)
		case schema.TypeBytea:
			newV = &[]byte{}
		default:
			str, _ := v.Type.String()
			return nil, errors.New("invalid type " + str)
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
			tv := outList[i].(*AnyString)
			if tv.Valid {
				outV = tv.String
			}
		case schema.TypeInt:
			tv := outList[i].(*sql.NullInt64)
			if tv.Valid {
				outV = tv.Int64
			}
		case schema.TypeDatetime:
			tv := outList[i].(*NullTime)
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
			str, _ := v.Type.String()
			return nil, errors.New("invalid type " + str)
		}
		rev = append(rev, outV)
	}
	return rev, nil
}
