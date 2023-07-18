package mysql

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"strings"

	ps "github.com/linlexing/dbx/pageselect"
	"github.com/linlexing/dbx/scan"
	"github.com/linlexing/dbx/schema"
	"github.com/sirupsen/logrus"
)

const driverName = "elastic"

type meta struct{}

func init() {
	ps.Register(driverName, new(meta))
}
func fromDBType(ty string) schema.DataType {
	switch ty {
	case "byte", "short", "integer", "long":
		return schema.TypeInt
	case "keyword", "text":
		return schema.TypeString
	case "float", "double", "half_float", "scaled_float":
		return schema.TypeFloat
	case "date":
		return schema.TypeDatetime
	case "binary":
		return schema.TypeBytea
	default:
		logrus.WithFields(logrus.Fields{
			"type": ty,
		}).Panic("invalid type")
	}
	return 0
}
func (m *meta) QuotedIdentifier(col string) string {
	return "\"" + col + "\""
}
func (m *meta) ColumnTypes(rows *sql.Rows) ([]*scan.ColumnType, error) {
	return nil, fmt.Errorf("not impl")
}
func (m *meta) SortByAsc(field string, _ bool) string {
	return field
}
func (m *meta) Sum(col string) string {
	return fmt.Sprintf("sum(ifnull(%s,0))", col)
}
func (m *meta) Avg(col string) string {
	return fmt.Sprintf("avg(ifnull(%s,0))", col)
}
func (m *meta) SortByDesc(field string, _ bool) string {
	return field + " DESC"
}
func (m *meta) LimitSQL(sel, strSQL, where, orderby string, limit int) string {
	return fmt.Sprintf("select %s from (\n%s\n) wholesql %s%s limit %d",
		sel, strSQL, where, orderby, limit)
}

// 返回一个字段值的字符串表达式
func valueExpress(dataType schema.DataType, value string) string {
	if len(value) >= 2 && value[0] == '`' && value[len(value)-1] == '`' {
		return value[1 : len(value)-1]
	}
	switch dataType {
	case schema.TypeFloat:
		//防止注入攻击
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			return "'" + strings.Replace(value, "'", "''", -1) + "'"
		}
		return value
	case schema.TypeInt:
		//防止注入攻击
		if _, err := strconv.ParseInt(value, 10, 64); err != nil {
			return "'" + strings.Replace(value, "'", "''", -1) + "'"
		}
		return value

	case schema.TypeString:
		return "'" + strings.Replace(value, "'", "''", -1) + "'"
	case schema.TypeDatetime:
		if len(value) == 10 {
			return "cast('" + value + "' as TIMESTAMP)"
		} else if len(value) == 19 {
			vals := strings.Fields(value)
			return fmt.Sprintf("cast('%sT%sZ' as TIMESTAMP)", vals[0], vals[1])
		} else {
			panic(fmt.Errorf("invalid datetime:%s", value))
		}
	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}

func (m *meta) GetOperatorExpress(ope ps.Operator, dataType schema.DataType, left, right, value2 string) (strSQL string) {
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
			//在列表简化起见，不再类型化
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

			strSQL = fmt.Sprintf("%s rlike %s", left, valueExpress(dataType, right))

		}
	case ps.OperatorNotRegexp: //"!~" 非正则
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {

			strSQL = fmt.Sprintf("%s not rlike %s", left, valueExpress(dataType, right))

		}
	case ps.OperatorIsNull: // "e" 为空
		strSQL = fmt.Sprintf("%s is null", left)
	case ps.OperatorIsNotNull: //"!e" 不为空
		strSQL = fmt.Sprintf("%s is not null", left)
	case ps.OperatorLengthEqu: // "_" 长度等于

		strSQL = fmt.Sprintf("%s RLIKE '.{%s}'", left, right)

	case ps.OperatorLengthNotEqu: // "!_" 长度不等于
		ilen, err := strconv.Atoi(right)
		if err != nil {
			log.Panic(fmt.Errorf("the length:%s invalid", right))
		}
		strSQL = fmt.Sprintf("%s RLIKE '.{0,%d}|.{%d,}", left, ilen-1, ilen+1)

	case ps.OperatorLengthGreaterThan: // "_>" 长度大于
		ilen, err := strconv.Atoi(right)
		if err != nil {
			log.Panic(fmt.Errorf("the length:%s invalid", right))
		}

		strSQL = fmt.Sprintf("%s RLIKE '.{%d,}", left, ilen+1)

	case ps.OperatorLengthGreaterThanOrEqu: // "_>=" 长度大于等于

		strSQL = fmt.Sprintf("%s RLIKE '.{%s,}", left, right)

	case ps.OperatorLengthLessThan: //"_<" 长度小于
		ilen, err := strconv.Atoi(right)
		if err != nil {
			log.Panic(fmt.Errorf("the length:%s invalid", right))
		}
		strSQL = fmt.Sprintf("%s RLIKE '{0,%d}", left, ilen-1)

	case ps.OperatorLengthLessThanOrEqu: //"_<=" 长度小于

		strSQL = fmt.Sprintf("%s RLIKE '.{0,%s}", left, right)
	case ps.OperatorBetween:
		strSQL = fmt.Sprintf("%s between %s and %s", left, valueExpress(dataType, right), valueExpress(dataType, value2))
	case ps.OperatorNotBetween:
		strSQL = fmt.Sprintf("%s not between %s and %s", left, valueExpress(dataType, right), valueExpress(dataType, value2))
	default:
		log.Panic(fmt.Errorf("the opt:%s not impl", ope))
	}
	return
}
