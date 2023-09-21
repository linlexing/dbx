package oracle

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/linlexing/dbx/scan"

	ps "github.com/linlexing/dbx/pageselect"
	"github.com/linlexing/dbx/schema"
	"github.com/linlexing/dbx/schema/sqlite"
	// sqlite3 "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

type meta struct{}

func init() {
	ps.Register(driverName, new(meta))
	// regex := func(re, s string) (bool, error) {
	// 	return regexp.MatchString(re, s)
	// }
	// sql.Register("sqlite3_with_go_func",
	// 	&sqlite3.SQLiteDriver{

	// 		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
	// 			return conn.RegisterFunc("regexp", regex, true)
	// 		},
	// 	})
}
func (m *meta) QuotedIdentifier(col string) string {
	return "`" + col + "`"
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
				Type: sqlite.SqliteType(one.Name(), one.DatabaseTypeName()),
			})
	}
	return rev, nil
}

func (m *meta) SortByAsc(field string, _ bool) string {
	return field
}
func (m *meta) SortByDesc(field string, _ bool) string {
	return field + " DESC"
}
func (m *meta) LimitSQL(sel, strSQL, where, orderby string, limit int) string {
	return fmt.Sprintf("select %s from (\n%s\n) wholesql %s%s limit %d",
		sel, strSQL, where, orderby, limit)
}
func (m *meta) Sum(col string) string {
	return fmt.Sprintf("sum(cast(ifnull(%s,0) as decimal(29,6)))", col)
}
func (m *meta) Avg(col string) string {
	return fmt.Sprintf("avg(cast(ifnull(%s,0) as decimal(29,6)))", col)
}

func (m *meta) GetOperatorExpress(ope ps.Operator, dataType schema.DataType, column, value, value2 string) (strSQL string) {
	//需要考虑到null的情况
	switch ope {
	case ps.OperatorEqu: // "=" 等于
		if value == "" {
			strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)
		} else {
			strSQL = fmt.Sprintf("%s = %s", column, valueExpress(dataType, value))
		}
	case ps.OperatorNotEqu: // "!=" 不等于
		if value == "" {
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
		} else {
			strSQL = fmt.Sprintf("(%s <> %s or %[1]s is null or %[1]s = '')", column, valueExpress(dataType, value))
		}
	case ps.OperatorGreaterThan: // ">" 大于
		if value == "" {
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
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
			strSQL = fmt.Sprintf("(%s < %s or %[1]s is null or %[1]s = '')", column, valueExpress(dataType, value))
		}
	case ps.OperatorLessThanOrEqu: // "<=" 小于等于
		if value == "" {
			strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)

		} else {
			strSQL = fmt.Sprintf("(%s <= %s or %[1]s is null or %[1]s = '')", column, valueExpress(dataType, value))
		}
	case ps.OperatorLike: //"?" 包含
		if value == "" {
			strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)
		} else {
			strSQL = fmt.Sprintf("%s like %s", column, valueExpress(dataType, "%"+value+"%"))
		}
	case ps.OperatorNotLike: //"!?" 不包含
		if value == "" {
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
		} else {
			strSQL = fmt.Sprintf("%s not like %s", column, valueExpress(dataType, "%"+value+"%"))
		}
	case ps.OperatorPrefix: // "?>" 前缀
		if value == "" {
			strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)
		} else {
			strSQL = fmt.Sprintf("%s like %s", column, valueExpress(dataType, value+"%"))
		}
	case ps.OperatorNotPrefix: //"!?>" 非前缀
		if value == "" {
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
		} else {
			strSQL = fmt.Sprintf("%s not like %s", column, valueExpress(dataType, value+"%"))
		}
	case ps.OperatorSuffix: // "<?" 后缀
		if value == "" {
			strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)
		} else {
			strSQL = fmt.Sprintf("%s like %s", column, valueExpress(dataType, "%"+value))
		}
	case ps.OperatorNotSuffix: // "!<?" 非后缀
		if value == "" {
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
		} else {
			strSQL = fmt.Sprintf("%s not like %s", column, valueExpress(dataType, "%"+value))
		}
	case ps.OperatorIn: //"in" 在列表
		if value == "" {
			strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)
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
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
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
			strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)
		} else {

			if array, err := csv.NewReader(strings.NewReader(value)).Read(); err != nil {
				log.Panic(err)
			} else {
				rList := []string{}
				var matchItems string
				for _, v := range array {
					//正则长度限制
					if len(v) > 256 { //单个就超长就跳过
						continue
					}
					if len(matchItems)+len(v) > 256 {
						rList = append(rList, fmt.Sprintf("%s regexp '%s'", column, matchItems))
						matchItems = ""
					}
					if len(matchItems) > 0 {
						matchItems = matchItems + fmt.Sprintf("|%s", valueExpressNoQuotes(dataType, v))
					} else {
						matchItems = valueExpressNoQuotes(dataType, v)
					}
				}
				rList = append(rList, fmt.Sprintf("%s regexp '%s'", column, matchItems))
				strSQL = fmt.Sprintf("(%s)", strings.Join(rList, " or "))
			}

		}
	case ps.OperatorNotLikeArray:
		if value == "" {
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
		} else {

			if array, err := csv.NewReader(strings.NewReader(value)).Read(); err != nil {
				log.Panic(err)
			} else {
				rList := []string{}
				var matchItems string
				for _, v := range array {
					//正则长度限制
					if len(v) > 256 { //单个就超长就跳过
						continue
					}
					if len(matchItems)+len(v) > 256 {
						rList = append(rList, fmt.Sprintf("%s regexp '%s'", column, matchItems))
						matchItems = ""
					}
					if len(matchItems) > 0 {
						matchItems = matchItems + fmt.Sprintf("|%s", valueExpressNoQuotes(dataType, v))
					} else {
						matchItems = valueExpressNoQuotes(dataType, v)
					}
				}
				rList = append(rList, fmt.Sprintf("%s regexp '%s'", column, matchItems))
				strSQL = fmt.Sprintf("not (%s)", strings.Join(rList, " or "))
			}

		}
	case ps.OperatorRegexp: // "~" 正则
		if value == "" {
			strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)
		} else {

			strSQL = fmt.Sprintf("%s REGEXP %s", column, valueExpress(dataType, value))

		}
	case ps.OperatorNotRegexp: //"!~" 非正则
		if value == "" {
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
		} else {

			strSQL = fmt.Sprintf("not (%s REGEXP %s)", column, valueExpress(dataType, value))

		}
	case ps.OperatorIsNull: // "e" 为空
		strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)
	case ps.OperatorIsNotNull: //"!e" 不为空
		strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
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

	case schema.TypeDatetime:
		return "'" + strings.Replace(value, "'", "''", -1) + "'"
	case schema.TypeString:
		return "'" + strings.Replace(value, "'", "''", -1) + "'"

	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}
func valueExpressNoQuotes(dataType schema.DataType, value string) string {
	if len(value) >= 2 && value[0] == '`' && value[len(value)-1] == '`' {
		return value[1 : len(value)-1]
	}
	switch dataType {
	case schema.TypeFloat, schema.TypeInt, schema.TypeDatetime:
		return value
	case schema.TypeString:
		return strings.Replace(value, "'", "''", -1)

	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}
