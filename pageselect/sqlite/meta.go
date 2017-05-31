package oracle

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"regexp"
	"strings"

	ps "github.com/linlexing/dbx/pageselect"
	"github.com/linlexing/dbx/schema"
	sqlite3 "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

type meta struct{}

func init() {
	ps.Register(driverName, new(meta))
	regex := func(re, s string) (bool, error) {
		return regexp.MatchString(re, s)
	}
	sql.Register("sqlite3_with_go_func",
		&sqlite3.SQLiteDriver{

			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				return conn.RegisterFunc("regexp", regex, true)
			},
		})
}
func (m *meta) SortByAsc(field string) string {
	return field
}
func (m *meta) SortByDesc(field string) string {
	return field + " DESC"
}
func (m *meta) LimitSQL(sel, strSQL, where, orderby string, limit int) string {
	return fmt.Sprintf("select %s from (%s) wholesql %s%s limit %d",
		sel, strSQL, where, orderby, limit)
}

func (m *meta) IsNull() string {
	return "ifnull"
}

func (m *meta) GetOperatorExpress(ope ps.Operator, dataType schema.DataType, left, right string) (strSQL string) {
	//需要考虑到null的情况
	switch ope {
	case ps.OperatorEqu: // "=" 等于
		if right == "" {
			strSQL = fmt.Sprintf("%s is null", left)
		} else {
			strSQL = fmt.Sprintf("%s = %s", left, valueExpress(dataType, right))
		}
	case ps.OperatorNotEqu: // "!=" 不等于
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {
			strSQL = fmt.Sprintf("(%s <> %s or %[1]s is null)", left, valueExpress(dataType, right))
		}
	case ps.OperatorGreaterThan: // ">" 大于
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {
			strSQL = fmt.Sprintf("%s > %s", left, valueExpress(dataType, right))
		}
	case ps.OperatorGreaterThanOrEqu: // ">=" 大于等于
		if right == "" {
			strSQL = "1=1"
		} else {
			strSQL = fmt.Sprintf("%s >= %s", left, valueExpress(dataType, right))
		}
	case ps.OperatorLessThan: //"<" 小于
		if right == "" {
			strSQL = "1=2"
		} else {
			strSQL = fmt.Sprintf("(%s < %s or %[1]s is null)", left, valueExpress(dataType, right))
		}
	case ps.OperatorLessThanOrEqu: // "<=" 小于等于
		if right == "" {
			strSQL = fmt.Sprintf("%s is null", left)

		} else {
			strSQL = fmt.Sprintf("(%s <= %s or %[1]s is null)", left, valueExpress(dataType, right))
		}
	case ps.OperatorLike: //"?" 包含
		if right == "" {
			strSQL = fmt.Sprintf("%s is null", left)
		} else {
			strSQL = fmt.Sprintf("%s like %s", left, valueExpress(dataType, "%"+right+"%"))
		}
	case ps.OperatorNotLike: //"!?" 不包含
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {
			strSQL = fmt.Sprintf("%s not like %s", left, valueExpress(dataType, "%"+right+"%"))
		}
	case ps.OperatorPrefix: // "?>" 前缀
		if right == "" {
			strSQL = fmt.Sprintf("%s is null", left)
		} else {
			strSQL = fmt.Sprintf("%s like %s", left, valueExpress(dataType, right+"%"))
		}
	case ps.OperatorNotPrefix: //"!?>" 非前缀
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {
			strSQL = fmt.Sprintf("%s not like %s", left, valueExpress(dataType, right+"%"))
		}
	case ps.OperatorSuffix: // "<?" 后缀
		if right == "" {
			strSQL = fmt.Sprintf("%s is null", left)
		} else {
			strSQL = fmt.Sprintf("%s like %s", left, valueExpress(dataType, "%"+right))
		}
	case ps.OperatorNotSuffix: // "!<?" 非后缀
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {
			strSQL = fmt.Sprintf("%s not like %s", left, valueExpress(dataType, "%"+right))
		}
	case ps.OperatorIn: //"in" 在列表
		if right == "" {
			strSQL = fmt.Sprintf("%s is null", left)
		} else {

			if array, err := csv.NewReader(strings.NewReader(right)).Read(); err != nil {
				log.Panic(err)
			} else {
				list := []string{}
				for _, v := range array {
					list = append(list, valueExpress(dataType, v))
				}
				strSQL = fmt.Sprintf("%s in (%s)", left, strings.Join(list, ",\n"))
			}

		}
	case ps.OperatorNotIn: //"!in" 不在列表
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {

			if array, err := csv.NewReader(strings.NewReader(right)).Read(); err != nil {
				log.Panic(err)
			} else {
				list := []string{}
				for _, v := range array {
					list = append(list, valueExpress(dataType, v))
				}
				strSQL = fmt.Sprintf("%s not in (%s)", left, strings.Join(list, ",\n"))
			}
		}
	case ps.OperatorRegexp: // "~" 正则
		if right == "" {
			strSQL = fmt.Sprintf("%s is null", left)
		} else {

			strSQL = fmt.Sprintf("%s REGEXP %s", left, valueExpress(dataType, right))

		}
	case ps.OperatorNotRegexp: //"!~" 非正则
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {

			strSQL = fmt.Sprintf("not (%s REGEXP %s)", left, valueExpress(dataType, right))

		}
	case ps.OperatorIsNull: // "e" 为空
		strSQL = fmt.Sprintf("%s is null", left)
	case ps.OperatorIsNotNull: //"!e" 不为空
		strSQL = fmt.Sprintf("%s is not null", left)
	case ps.OperatorLengthEqu: // "_" 长度等于

		strSQL = fmt.Sprintf("length(%s) = %s", left, right)

	case ps.OperatorLengthNotEqu: // "!_" 长度不等于

		strSQL = fmt.Sprintf("length(%s) <> %s", left, right)

	case ps.OperatorLengthGreaterThan: // "_>" 长度大于

		strSQL = fmt.Sprintf("length(%s) > %s", left, right)

	case ps.OperatorLengthGreaterThanOrEqu: // "_>=" 长度大于等于

		strSQL = fmt.Sprintf("length(%s) >= %s", left, right)

	case ps.OperatorLengthLessThan: //"_<" 长度小于

		strSQL = fmt.Sprintf("length(%s) < %s", left, right)

	case ps.OperatorLengthLessThanOrEqu: //"_<=" 长度小于

		strSQL = fmt.Sprintf("length(%s) <= %s", left, right)

	default:
		log.Panic(fmt.Errorf("the opt:%s not impl", ope))
	}
	return
}

func valueExpress(dataType schema.DataType, value string) string {
	switch dataType {
	case schema.TypeFloat, schema.TypeInt, schema.TypeDatetime:
		return value
	case schema.TypeString:
		return "'" + strings.Replace(value, "'", "''", -1) + "'"

	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}
