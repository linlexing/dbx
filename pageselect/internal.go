package pageselect

import (
	"strings"

	"github.com/linlexing/dbx/render"
)

func buildCondition(order, divide []string) []*ConditionLine {
	result := []*ConditionLine{}
	//a=:a and b=:b and c>:c or
	//a=:a and b>:b or
	//a>:a
	for i := len(divide) - 1; i >= 0; i-- {
		lines := []*ConditionLine{}
		for j := i; j >= 0; j-- {
			colName := order[j]
			isDesc := strings.HasPrefix(colName, "-")
			if isDesc {
				colName = colName[1:]
			}
			//尾部指标是用大于或者小于（倒序）
			if j == i {
				opt := OperatorGreaterThan // ">"
				if isDesc {
					opt = OperatorLessThan //"<"
				}
				lines = append(lines, &ConditionLine{
					ColumnName: colName,
					Operators:  opt,
					Value:      divide[j],
					Logic:      AND,
				})
			} else {
				lines = append(lines, &ConditionLine{
					ColumnName: colName,
					Operators:  OperatorEqu, // "=",
					Value:      divide[j],
					Logic:      AND,
				})
			}
		}
		if len(lines) > 1 {
			lines[0].LeftBrackets = "("
			lines[len(lines)-1].RightBrackets = ")"
		}
		lines[len(lines)-1].Logic = OR
		result = append(result, lines...)
	}
	return result
}

func renderManualPageSQL(driver string, strSQL string, columnList []string,
	columnListIsExpress bool, whereList, orderbyList []string, limit int,
	autoQuoted bool) (string, error) {

	var where string
	var columns string
	if len(whereList) > 0 {
		where = "(" + strings.Join(whereList, " "+AND+" ") + ")"
	}
	if len(columnList) > 0 {

		list := []string{}
		for _, c := range columnList {
			if columnListIsExpress {
				list = append(list, c)
			} else {
				colName := c
				if autoQuoted {
					colName = Find(driver).QuotedIdentifier(colName)
				}
				list = append(list, colName)
			}
		}
		columns = strings.Join(list, ",")
	}

	return render.RenderSQLCustom(strSQL, "<<", ">>", map[string]interface{}{
		"Driver":  driver,
		"Columns": columns,
		"Where":   where,
		"OrderBy": strings.Join(orderbyList, ","),
		"Limit":   limit,
	}, nil)
}
