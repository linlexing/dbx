package oracle

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/Sirupsen/logrus"
	ps "github.com/linlexing/dbx/pageselect"
	"github.com/linlexing/dbx/scan"
	"github.com/linlexing/dbx/schema"
)

const driverName = "oci8"

type meta struct{}

func init() {
	ps.Register(driverName, new(meta))
}
func fromDBType(ty string) schema.DataType {
	switch ty {
	case "SQLT_INT", "SQLT_NUM", "SQLT_UIN":
		return schema.TypeInt
	case "SQLT_CHR", "SQLT_STR", "SQLT_CLOB", "SQLT_VCS", "SQLT_LVC", "SQLT_AFC",
		"SQLT_AVC", "SQLT_VST", "SQLT_LNG", "SQLT_VBI", "SQLT_BIN", "SQLT_LBI", "SQLT_LVB":
		return schema.TypeString
	case "SQLT_FLT", "SQLT_BDOUBLE", "SQLT_BFLOAT", "SQLT_VNU", "":
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
func (m *meta) LimitSQL(sel, strSQL, where, orderby string, limit int) string {
	return fmt.Sprintf(
		"select * from (select %s from (%s) wholesql %s%s) where rownum<=%d",
		sel, strSQL, where, orderby, limit)
}
func (m *meta) IsNull() string {
	return "nvl"
}

func valueExpress(dataType schema.DataType, value string) string {
	switch dataType {
	case schema.TypeFloat, schema.TypeInt:
		return value
	case schema.TypeString:
		return "'" + strings.Replace(value, "'", "''", -1) + "'"
	case schema.TypeDatetime:
		if len(value) == 10 {
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd')", value)
		} else if len(value) == 19 {
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd hh24:mi:ss')", value)
		} else {
			panic(fmt.Errorf("invalid datetime:%s", value))
		}
	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
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

			strSQL = fmt.Sprintf("regexp_like(%s,%s)", left, valueExpress(dataType, right))

		}
	case ps.OperatorNotRegexp: //"!~" 非正则
		if right == "" {
			strSQL = fmt.Sprintf("%s is not null", left)
		} else {

			strSQL = fmt.Sprintf("not regexp_like(%s,%s)", left, valueExpress(dataType, right))

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
