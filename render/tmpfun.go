package render

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/linlexing/dbx/suid"
	"github.com/pborman/uuid"
)

var tempFunc = template.FuncMap{
	"guid64": func() string {
		return base64.RawURLEncoding.EncodeToString(uuid.NewUUID())
	},
	"guid": func() string {
		return strings.ToUpper(hex.EncodeToString(uuid.NewUUID()))
	},
	"suid": func() string {
		id, err := suid.Next()
		if err != nil {
			return err.Error()
		}
		return id
	},
	//P 函数将具体的参数值转换成文字量，这里不用绑定，会有一些性能损失
	"P": func(val interface{}) string {
		switch tv := val.(type) {
		case string:
			return "'" + strings.Replace(tv, "'", "''", -1) + "'"
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			return fmt.Sprintf("%d", tv)
		case float32:
			return strconv.FormatFloat(float64(tv), 'f', -1, 32)
		case float64:
			return strconv.FormatFloat(tv, 'f', -1, 64)
		case []string: //字符串数组一定是用在in语句中
			list := []string{}
			for _, v := range tv {
				list = append(list, "'"+strings.Replace(v, "'", "''", -1)+"'")
			}
			return strings.Join(list, ",")
		default:
			panic(fmt.Errorf("not impl P,data type :%T", val))
		}
	},
	"cell": newCell,
	"in": func(val string, s ...string) bool {
		for _, one := range s {
			if one == val {
				return true
			}
		}
		return false
	},
	"subtract": func(y, x int) int {
		return y - x
	},
	"fields": strings.Fields,
	"split":  strings.Split,
	"dec": func(i int) int {
		return i - 1
	},
	"int": strconv.Atoi,
	"N": func(start, end int) (stream chan int) {
		stream = make(chan int)
		go func() {
			for i := start; i <= end; i++ {
				stream <- i
			}
			close(stream)
		}()
		return
	},
}

// AddFunc 增加模板函数，如果前面有同名的函数，将被覆盖
func AddFunc(fs template.FuncMap) {
	for k, v := range fs {
		tempFunc[k] = v
	}
}

// Cell是用来模拟可变变量的，模板系统不允许修改变量的值，因此用struct来模拟
type cell struct{ v interface{} }

func newCell(v ...interface{}) (*cell, error) {
	switch len(v) {
	case 0:
		return new(cell), nil
	case 1:
		return &cell{v[0]}, nil
	default:
		return nil, fmt.Errorf("wrong number of args: want 0 or 1, got %v", len(v))
	}
}

func (c *cell) Set(v interface{}) *cell { c.v = v; return c }
func (c *cell) Get() interface{}        { return c.v }
