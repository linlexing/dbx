package oracle

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

const driverName = "oci8"

type meta struct{}

func init() {
	ps.Register(driverName, new(meta))
}
func fromDBType(ty string) schema.DataType {
	switch ty {
	case "SQLT_INT", "SQLT_UIN": //, "SQLT_NUM":
		return schema.TypeInt
	case "SQLT_CHR", "SQLT_STR", "SQLT_CLOB", "SQLT_VCS", "SQLT_LVC", "SQLT_AFC",
		"SQLT_AVC", "SQLT_VST", "SQLT_LNG", "SQLT_VBI", "SQLT_BIN", "SQLT_LBI", "SQLT_LVB":
		return schema.TypeString
	case "SQLT_FLT", "SQLT_BDOUBLE", "SQLT_BFLOAT", "SQLT_VNU",
		"SQLT_NUM", /*number,int都是这个类型，目前驱动不支持精度查询，如果放int，解析小数就出错，所以放这里*/
		"":
		return schema.TypeFloat
	case "SQLT_DAT", "SQLT_DATE", "SQLT_TIMESTAMP", "SQLT_TIMESTAMP_TZ", "SQLT_TIMESTAMP_LTZ":
		return schema.TypeDatetime
	case "SQLT_BLOB":
		return schema.TypeBytea
	default:
		logrus.WithFields(logrus.Fields{
			"type": ty,
		}).Panic("invalid type")
	}
	return 0
}
func (m *meta) ColumnTypes(rows *sql.Rows) ([]*scan.ColumnType, error) {
	cols, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	rev := []*scan.ColumnType{}
	for _, one := range cols {
		rev = append(rev,
			&scan.ColumnType{
				Name: one.Name(),
				Type: fromDBType(one.DatabaseTypeName()),
			})
	}
	return rev, nil
}
func (m *meta) SortByAsc(field string, notNull bool) string {
	//非空字段（一般是主键）不加NULLs first性能会有很大提升
	if notNull {
		return field
	}
	return field + " NULLS FIRST"
}
func (m *meta) SortByDesc(field string, notNull bool) string {
	if notNull {
		return field + " DESC"
	}
	return field + " DESC NULLS LAST"
}
func (m *meta) QuotedIdentifier(col string) string {
	return "\"" + col + "\""
}
func (m *meta) LimitSQL(sel, strSQL, where, orderby string, limit int) string {
	return fmt.Sprintf(
		"select * from (select %s from (\n%s\n) wholesql %s%s) outer_wsql where rownum<=%d",
		sel, strSQL, where, orderby, limit)
}
func (m *meta) Sum(col string) string {
	return fmt.Sprintf("sum(cast(nvl(%s,0) as decimal(29,6)))", col)
}
func (m *meta) Avg(col string) string {
	return fmt.Sprintf("avg(cast(nvl(%s,0) as decimal(29,6)))", col)
}
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
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd')", value)
		} else if len(value) == 19 {
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd hh24:mi:ss')", value)
		} else if len(value) > 19 {
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd hh24:mi:ss')", value[:19])
		} else {
			panic(fmt.Errorf("invalid datetime:%s", value))
		}
	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}
func valueExpressNoQuotes(dataType schema.DataType, value string) string {
	if len(value) >= 2 && value[0] == '`' && value[len(value)-1] == '`' {
		return value[1 : len(value)-1]
	}
	switch dataType {
	case schema.TypeFloat, schema.TypeInt:
		return value
	case schema.TypeString:
		return strings.Replace(value, "'", "''", -1)
	case schema.TypeDatetime:
		if len(value) == 10 {
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd')", value)
		} else if len(value) == 19 {
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd hh24:mi:ss')", value)
		} else if len(value) > 19 {
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd hh24:mi:ss')", value[:19])
		} else {
			panic(fmt.Errorf("invalid datetime:%s", value))
		}
	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}

func (m *meta) GetOperatorExpress(ope ps.Operator, dataType schema.DataType, column, value, value2 string) (strSQL string) {
	//需要考虑到null的情况
	switch ope {
	case ps.OperatorEqu: // "=" 等于
		if value == "" {
			strSQL = fmt.Sprintf("%s is null", column)
		} else {
			strSQL = fmt.Sprintf("%s = %s", column, valueExpress(dataType, value))
		}
	case ps.OperatorNotEqu: // "!=" 不等于
		if value == "" {
			strSQL = fmt.Sprintf("%s is not null", column)
		} else {
			strSQL = fmt.Sprintf("(%s <> %s or %[1]s is null)", column, valueExpress(dataType, value))
		}
	case ps.OperatorGreaterThan: // ">" 大于
		if value == "" {
			strSQL = fmt.Sprintf("%s is not null", column)
		} else {
			strSQL = fmt.Sprintf("%s > %s", column, valueExpress(dataType, value))
		}
	case ps.OperatorGreaterThanOrEqu: // ">=" 大于等于
		if value == "" {
			strSQL = "1=1"
		} else {
			strSQL = fmt.Sprintf("%s >= %s", column, valueExpress(dataType, value))
		}
	case ps.OperatorLessThan: //"<" 小于
		if value == "" {
			strSQL = "1=2"
		} else {
			strSQL = fmt.Sprintf("(%s < %s or %[1]s is null)", column, valueExpress(dataType, value))
		}
	case ps.OperatorLessThanOrEqu: // "<=" 小于等于
		if value == "" {
			strSQL = fmt.Sprintf("%s is null", column)

		} else {
			strSQL = fmt.Sprintf("(%s <= %s or %[1]s is null)", column, valueExpress(dataType, value))
		}
	case ps.OperatorLike: //"?" 包含
		if value == "" {
			strSQL = fmt.Sprintf("%s is null", column)
		} else {
			strSQL = fmt.Sprintf("%s like %s", column, valueExpress(dataType, "%"+value+"%"))
		}
	case ps.OperatorNotLike: //"!?" 不包含
		if value == "" {
			strSQL = fmt.Sprintf("%s is not null", column)
		} else {
			strSQL = fmt.Sprintf("%s not like %s", column, valueExpress(dataType, "%"+value+"%"))
		}
	case ps.OperatorPrefix: // "?>" 前缀
		if value == "" {
			strSQL = fmt.Sprintf("%s is null", column)
		} else {
			strSQL = fmt.Sprintf("%s like %s", column, valueExpress(dataType, value+"%"))
		}
	case ps.OperatorNotPrefix: //"!?>" 非前缀
		if value == "" {
			strSQL = fmt.Sprintf("%s is not null", column)
		} else {
			strSQL = fmt.Sprintf("%s not like %s", column, valueExpress(dataType, value+"%"))
		}
	case ps.OperatorSuffix: // "<?" 后缀
		if value == "" {
			strSQL = fmt.Sprintf("%s is null", column)
		} else {
			strSQL = fmt.Sprintf("%s like %s", column, valueExpress(dataType, "%"+value))
		}
	case ps.OperatorNotSuffix: // "!<?" 非后缀
		if value == "" {
			strSQL = fmt.Sprintf("%s is not null", column)
		} else {
			strSQL = fmt.Sprintf("%s not like %s", column, valueExpress(dataType, "%"+value))
		}
	case ps.OperatorIn: //"in" 在列表
		if value == "" {
			strSQL = fmt.Sprintf("%s is null", column)
		} else {

			if array, err := csv.NewReader(strings.NewReader(value)).Read(); err != nil {
				log.Panic(err)
			} else {
				list := []string{}
				for _, v := range array {
					list = append(list, valueExpress(dataType, v))
				}
				strSQL = fmt.Sprintf("%s in (%s)", column, strings.Join(list, ",\n"))
			}

		}
	case ps.OperatorNotIn: //"!in" 不在列表
		if value == "" {
			strSQL = fmt.Sprintf("%s is not null", column)
		} else {

			if array, err := csv.NewReader(strings.NewReader(value)).Read(); err != nil {
				log.Panic(err)
			} else {
				list := []string{}
				for _, v := range array {
					list = append(list, valueExpress(dataType, v))
				}
				strSQL = fmt.Sprintf("%s not in (%s)", column, strings.Join(list, ",\n"))
			}
		}
	case ps.OperatorLikeArray:
		if value == "" {
			strSQL = fmt.Sprintf("%s is null", column)
		} else {

			if array, err := csv.NewReader(strings.NewReader(value)).Read(); err != nil {
				log.Panic(err)
			} else {
				rList := []string{}
				var matchItem string
				for _, v := range array {
					//正则长度限制
					if len(v) > 256 { //单个就超长就跳过
						continue
					}
					if len(matchItem)+len(v) > 256 {
						rList = append(rList, fmt.Sprintf("regexp_like(%s,'%s')", column, matchItem))
						matchItem = ""
					}
					if len(matchItem) > 0 {
						matchItem = matchItem + fmt.Sprintf("|%s", valueExpressNoQuotes(dataType, v))
					} else {
						matchItem = valueExpressNoQuotes(dataType, v)
					}
				}
				rList = append(rList, fmt.Sprintf("regexp_like(%s,'%s')", column, matchItem))
				strSQL = fmt.Sprintf("(%s)", strings.Join(rList, " or "))
			}

		}
	case ps.OperatorNotLikeArray:
		if value == "" {
			strSQL = fmt.Sprintf("%s is not null", column)
		} else {

			if array, err := csv.NewReader(strings.NewReader(value)).Read(); err != nil {
				log.Panic(err)
			} else {
				rList := []string{}
				var matchItem string
				for _, v := range array {
					//正则长度限制
					if len(v) > 256 { //单个就超长就跳过
						continue
					}
					if len(matchItem)+len(v) > 256 {
						rList = append(rList, fmt.Sprintf("regexp_like(%s,'%s')", column, matchItem))
						matchItem = ""
					}
					if len(matchItem) > 0 {
						matchItem = matchItem + fmt.Sprintf("|%s", valueExpressNoQuotes(dataType, v))
					} else {
						matchItem = valueExpressNoQuotes(dataType, v)
					}
				}
				rList = append(rList, fmt.Sprintf("regexp_like(%s,'%s')", column, matchItem))
				strSQL = fmt.Sprintf("not (%s)", strings.Join(rList, " or "))
			}

		}
	case ps.OperatorRegexp: // "~" 正则
		if value == "" {
			strSQL = fmt.Sprintf("%s is null", column)
		} else {

			strSQL = fmt.Sprintf("regexp_like(%s,%s)", column, valueExpress(dataType, value))

		}
	case ps.OperatorNotRegexp: //"!~" 非正则
		if value == "" {
			strSQL = fmt.Sprintf("%s is not null", column)
		} else {

			strSQL = fmt.Sprintf("not regexp_like(%s,%s)", column, valueExpress(dataType, value))

		}
	case ps.OperatorIsNull: // "e" 为空
		strSQL = fmt.Sprintf("%s is null", column)
	case ps.OperatorIsNotNull: //"!e" 不为空
		strSQL = fmt.Sprintf("%s is not null", column)
	case ps.OperatorLengthEqu: // "_" 长度等于

		strSQL = fmt.Sprintf("length(%s) = %s", column, value)

	case ps.OperatorLengthNotEqu: // "!_" 长度不等于

		strSQL = fmt.Sprintf("length(%s) <> %s", column, value)

	case ps.OperatorLengthGreaterThan: // "_>" 长度大于

		strSQL = fmt.Sprintf("length(%s) > %s", column, value)

	case ps.OperatorLengthGreaterThanOrEqu: // "_>=" 长度大于等于

		strSQL = fmt.Sprintf("length(%s) >= %s", column, value)

	case ps.OperatorLengthLessThan: //"_<" 长度小于

		strSQL = fmt.Sprintf("length(%s) < %s", column, value)

	case ps.OperatorLengthLessThanOrEqu: //"_<=" 长度小于

		strSQL = fmt.Sprintf("length(%s) <= %s", column, value)
	case ps.OperatorBetween:
		strSQL = fmt.Sprintf("%s between %s and %s", column, valueExpress(dataType, value), valueExpress(dataType, value2))
	case ps.OperatorNotBetween:
		strSQL = fmt.Sprintf("%s not between %s and %s", column, valueExpress(dataType, value), valueExpress(dataType, value2))
	default:
		log.Panic(fmt.Errorf("the opt:%s not impl", ope))
	}
	return
}
