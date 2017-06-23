package pageselect

import (
	"bytes"
	"strings"
	"text/template"
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
					opt = OperatorGreaterThan //"<"
				}
				lines = append(lines, &ConditionLine{
					ColumnName: colName,
					Operators:  opt,
					Value:      divide[j],
					Logic:      "AND",
				})
			} else {
				lines = append(lines, &ConditionLine{
					ColumnName: colName,
					Operators:  OperatorEqu, // "=",
					Value:      divide[j],
					Logic:      "AND",
				})
			}
		}
		if len(lines) > 1 {
			lines[0].LeftBrackets = "("
			lines[len(lines)-1].RightBrackets = ")"
		}
		lines[len(lines)-1].Logic = "OR"
		result = append(result, lines...)
	}
	return result
}

func renderManualPageSQL(driver string, strSQL string, columnList, whereList, orderbyList []string, limit int) (string, error) {
	tmpl, err := template.New("ManualPage").Delims("<<", ">>").Parse(strSQL)
	if err != nil {
		return "", err
	}
	var where string
	var columns string
	var orderby string
	bys := bytes.NewBuffer(nil)
	if len(whereList) > 0 {
		where = "(" + strings.Join(whereList, " and ") + ")"
	}
	if len(columnList) > 0 {
		columns = strings.Join(columnList, ",")
	}
	if len(orderbyList) > 0 {
		orderby = strings.Join(orderbyList, ",")
	}
	if err = tmpl.Execute(bys, map[string]interface{}{
		"Driver":  driver,
		"Columns": columns,
		"Where":   where,
		"OrderBy": orderby,
		"Limit":   limit,
	}); err != nil {
		return "", err
	}
	return bys.String(), nil
}
