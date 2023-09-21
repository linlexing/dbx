package oracle

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/linlexing/dbx/scan"
	"github.com/sirupsen/logrus"

	ps "github.com/linlexing/dbx/pageselect"
	"github.com/linlexing/dbx/schema"
)

var (
	normalStart = regexp.MustCompile("^[a-zA-Z_\u4e00-\u9fa5]")
	spec        = regexp.MustCompile("[^0-9a-zA-Z_\u4e00-\u9fa5]")
)

const driverName = "postgres"
const driverNamePgx = "pgx"
const driverNameGauss = "opengauss"
const driverNamePgxGauss = "pgx-opengauss"

type meta struct{}

func init() {
	m := new(meta)
	ps.Register(driverName, m)
	ps.Register(driverNamePgx, m)
	ps.Register(driverNameGauss, m)
	ps.Register(driverNamePgxGauss, m)
}
func fromDBType(ty string) schema.DataType {
	switch ty {
	case "INT8", "INT4", "INT2":
		return schema.TypeInt
	case "_VARCHAR", "VARCHAR", "TEXT", "JSON", "JSONB", "NAME", "BOOL": //NAME 是information_schema.tables 用到的类型，等同字符串
		return schema.TypeString
	case "NUMERIC", "FLOAT8", "FLOAT4":
		return schema.TypeFloat
	case "DATE", "TIME", "TIMETZ", "TIMESTAMP", "TIMESTAMPTZ":
		return schema.TypeDatetime
	case "BYTEA":
		return schema.TypeBytea
	case "": //USER-DEFINED 返回的是nil
		return schema.TypeString
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
				//postgres默认是小写，如果是中英文混排，含有大写字母，则不能转换成大写，
				//因为select使用时，pg会自动转换成小写
				Name: one.Name(),
				Type: fromDBType(one.DatabaseTypeName()),
			})
	}
	return rev, nil
}
func (m *meta) QuotedIdentifier(col string) string {
	isNormal := true
	//不是字母下划线汉字开头的，要加包括引号
	if !normalStart.MatchString(col) {
		isNormal = false
	} else if strings.ToLower(col) != col && strings.ToUpper(col) != col {
		//不全是大小写
		isNormal = false
	} else if spec.MatchString(col) {
		isNormal = false
	}
	if isNormal {
		return col
	}

	return `"` + strings.ReplaceAll(col, `"`, `""`) + `"`
}
func (m *meta) SortByAsc(field string, notNull bool) string {
	if notNull {
		return field + " ASC"
	}
	return field + " ASC NULLS FIRST"
}
func (m *meta) SortByDesc(field string, notNull bool) string {
	if notNull {
		return field + " DESC"
	}

	return field + " DESC NULLS LAST"
}
func (m *meta) LimitSQL(sel, strSQL, where, orderby string, limit int) string {
	return fmt.Sprintf("select %s from (\n%s\n) wholesql %s%s limit %d",
		sel, strSQL, where, orderby, limit)
}
func (m *meta) Sum(col string) string {
	return fmt.Sprintf("sum(cast(COALESCE(%s,0) as decimal(29,6)))", col)
}
func (m *meta) Avg(col string) string {
	return fmt.Sprintf("avg(cast(COALESCE(%s,0) as decimal(29,6)))", col)
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
				list := []string{}
				for _, v := range array {
					list = append(list, valueExpress(dataType, "%"+v+"%"))
				}
				strSQL = fmt.Sprintf("%s like any (array[%s])", column, strings.Join(list, ",\n"))
			}

		}
	case ps.OperatorNotLikeArray:
		if value == "" {
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
		} else {

			if array, err := csv.NewReader(strings.NewReader(value)).Read(); err != nil {
				log.Panic(err)
			} else {
				list := []string{}
				for _, v := range array {
					list = append(list, valueExpress(dataType, "%"+v+"%"))
				}
				strSQL = fmt.Sprintf("%s not like all (array[%s])", column, strings.Join(list, ",\n"))
			}

		}
	case ps.OperatorRegexp: // "~" 正则
		if value == "" {
			strSQL = fmt.Sprintf("(%s is null or %[1]s ='')", column)
		} else {

			strSQL = fmt.Sprintf("%s ~ %s", column, valueExpress(dataType, value))

		}
	case ps.OperatorNotRegexp: //"!~" 非正则
		if value == "" {
			strSQL = fmt.Sprintf("(%s is not null and %[1]s <>'')", column)
		} else {

			strSQL = fmt.Sprintf("%s !~ %s", column, valueExpress(dataType, value))

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
	case ps.OperatorQuery:
		strSQL = fmt.Sprintf("to_tsvector('ngram',%s) @@ to_tsquery('ngram',%s)", column, signString(value))
	case ps.OperatorNotQuery:
		strSQL = fmt.Sprintf("not(to_tsvector('ngram',%s) @@ to_tsquery('ngram',%s))", column, signString(value))
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

	case schema.TypeString:
		return "'" + strings.Replace(value, "'", "''", -1) + "'"
	case schema.TypeDatetime:
		//如果是like，则直接返回不转换
		if strings.HasPrefix(value, "%") || strings.HasSuffix(value, "%") {
			return "'" + strings.Replace(value, "'", "''", -1) + "'"
		}
		if len(value) == 10 {
			return fmt.Sprintf("TO_DATE('%s','YYYY-MM-DD')", value)
		} else if len(value) == 19 {
			return fmt.Sprintf("to_timestamp('%s','YYYY-MM-DD HH24:MI:SS')", value)
		} else if len(value) > 19 {
			return fmt.Sprintf("to_timestamp('%s','YYYY-MM-DD HH24:MI:SS.US')", value)
		} else {
			panic(fmt.Errorf("invalid datetime:%s", value))
		}
	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}

// 返回单引号包括的字符串
func signString(str string) string {

	return "'" + strings.Replace(str, "'", "''", -1) + "'"
}
