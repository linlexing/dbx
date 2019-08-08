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

func renderManualPageSQL(driver string, strSQL string, columnList, whereList, orderbyList []string, limit int) (string, error) {

	var where string
	var columns string
	var orderby string
	if len(whereList) > 0 {
		where = "(" + strings.Join(whereList, " "+AND+" ") + ")"
	}
	if len(columnList) > 0 {
		columns = strings.Join(columnList, ",")
	}
	if len(orderbyList) > 0 {
		orderby = strings.Join(orderbyList, ",")
	}
	return render.RenderSQLCustom(strSQL, "<<", ">>", map[string]interface{}{
		"Driver":  driver,
		"Columns": columns,
		"Where":   where,
		"OrderBy": orderby,
		"Limit":   limit,
	}, nil)
}
