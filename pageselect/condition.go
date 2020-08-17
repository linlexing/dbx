package pageselect

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/schema"
)

const (
	//AND 逻辑
	AND = "AND"
	//OR 逻辑
	OR = "OR"
)

//ConditionLine 条件一行
type ConditionLine struct {
	LeftBrackets string
	ColumnName   string

	Operators     Operator
	Value         string
	RightBrackets string
	Logic         string
	PlainText     string //与上面的条件成and关系
}

//SQLCondition 模板条件,多行并且可以带一段高级条件
type SQLCondition struct {
	Name      string
	Lines     []*ConditionLine
	PlainText string
}

//GetExpress 根据条件返回一个SQL条件
func (c *ConditionLine) GetExpress(driver string, dataType schema.DataType) string {
	//加上括号
	rev := ""
	if len(c.ColumnName) > 0 {
		rev = fmt.Sprintf("%s%s%s", c.LeftBrackets,
			Find(driver).GetOperatorExpress(c.Operators, dataType,
				Find(driver).QuotedIdentifier(c.ColumnName), c.Value),
			c.RightBrackets)
	}
	if len(c.PlainText) > 0 {
		if len(rev) > 0 {
			rev += " " + AND + " "
		}
		rev += "(" + c.PlainText + ")"
	}
	return rev
}

//BuildWhere 构造where条件，可选传入一个schema.Table来更准确地界定每列的数据类型
func (c *SQLCondition) BuildWhere(driver string, cols ColumnTypes) string {
	strLines := []string{}

	if len(c.Lines) > 0 {
		//最后一行的逻辑设置为空
		c.Lines[len(c.Lines)-1].Logic = ""
		for i, v := range c.Lines {
			dataType := schema.TypeString
			if len(cols) > 0 {
				if field := cols.byName(v.ColumnName); field != nil {
					dataType = field.Type
				}
			}
			exp := v.GetExpress(driver, dataType)
			//最后一行不需要加逻辑
			if i < len(c.Lines)-1 {
				strLines = append(strLines, exp+" "+v.Logic)
			} else {
				strLines = append(strLines, exp)
			}

		}
	}
	if len(c.PlainText) > 0 {
		if len(strLines) > 0 {
			return fmt.Sprintf("(\n%s\n) %s (\n%s\n)", strings.Join(strLines, "\n"), AND, c.PlainText)
		}
		return c.PlainText

	}
	if len(strLines) > 0 {
		return strings.Join(strLines, "\n")
	}
	return ""

}
