package oracle

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/linlexing/dbx/scan"
	"github.com/sirupsen/logrus"

	ps "github.com/linlexing/dbx/pageselect"
	"github.com/linlexing/dbx/schema"
)

const driverName = "postgres"
const driverNamePgx = "pgx"

type meta struct{}

func init() {
	ps.Register(driverName, new(meta))
	ps.Register(driverNamePgx, new(meta))
}
func fromDBType(ty string) schema.DataType {
	switch ty {
	case "INT8", "INT4", "INT2":
		return schema.TypeInt
	case "_VARCHAR", "VARCHAR", "TEXT", "JSON", "JSONB", "NAME": //NAME 是information_schema.tables 用到的类型，等同字符串
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
	return `"` + strings.ReplaceAll(col, `"`, `""`) + `"`
}
func (m *meta) SortByAsc(field string, notNull bool) string {
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

			strSQL = fmt.Sprintf("%s ~ %s", left, valueExpress(dataType, right))

		}
	case ps.OperatorNotRegexp: //"!~" 非正则
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {

			strSQL = fmt.Sprintf("%s !~ %s", left, valueExpress(dataType, right))

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
	case schema.TypeFloat, schema.TypeInt:
		return value
	case schema.TypeString:
		return "'" + strings.Replace(value, "'", "''", -1) + "'"
	case schema.TypeDatetime:
		//如果是like，则直接返回不转换
		if strings.HasPrefix(value, "%") || strings.HasSuffix(value, "%") {
			return value
		}
		if len(value) == 10 {
			return fmt.Sprintf("TO_DATE('%s','YYYY-MM-DD')", value)
		} else if len(value) == 19 {
			return fmt.Sprintf("TO_DATE('%s','YYYY-MM-DD HH24:MI:SS')", value)
		} else {
			panic(fmt.Errorf("invalid datetime:%s", value))
		}
	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}
