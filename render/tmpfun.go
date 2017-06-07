package render

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

var tempFunc = template.FuncMap{

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
}
