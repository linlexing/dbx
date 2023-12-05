package selectelements

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/model"
)

func ParserNode(val string) *model.NodeSelectelements {
	if len(val) == 0 {
		return nil
	}
	//selectelements里没注释
	// var vars map[string]interface{}
	// val, _ = condition.ProcessComment(val)
	stream := antlr.NewInputStream(val)
	lexer := parser.NewSqlLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSqlParser(cs)
	p.BuildParseTrees = true
	tree := p.SelectElements()
	visitor := new(sqlSelectelementsVisitorImpl)
	// visitor.varssqlSelectelementsVisitorImpl
	return visitor.Visit(tree).(*model.NodeSelectelements)
}

func ParseByContext(ctx parser.ISelectElementsContext) *model.NodeSelectelements {
	visitor := new(sqlSelectelementsVisitorImpl)
	return visitor.Visit(ctx).(*model.NodeSelectelements)
}

func NewElement(tableAlias, columnName, express, as, alias string) *model.Element {
	return &model.Element{
		TableAlias: tableAlias,
		ColumnName: columnName,
		Express:    express,
		As:         as,
		Alias:      alias,
	}
}
func SelectElementsString(node *model.NodeSelectelements) string {
	if node.NodeType == model.NodeStar {
		return "*"
	}
	var elements []string
	for _, v := range node.Elements {
		col := v.Express
		var as, alias string
		if len(v.ColumnName) > 0 {
			if len(v.TableAlias) > 0 {
				col = v.TableAlias + "." + v.ColumnName
			} else {
				col = v.ColumnName
			}
		}
		if len(v.As) > 0 {
			as = " " + v.As
		}
		if len(v.Alias) > 0 {
			alias = " " + v.Alias
		}
		elements = append(elements, fmt.Sprintf("%s%s%s", col, as, alias))
	}
	return strings.Join(elements, ",")
}
