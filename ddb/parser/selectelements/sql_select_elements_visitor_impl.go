package selectelements

import (
	"regexp"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/linlexing/dbx/ddb/parser"
	"github.com/linlexing/dbx/ddb/parser/model"
)

var (
	//按.分隔，单双引号、括号内的不算
	//`(\([^)]*\)|'[^']*'|[^.()']+)+`
	regPoint = regexp.MustCompile(`(\([^)]*\)|'[^']*'|"[^"]*"|[^.()'"]+)+`)

	regComment = regexp.MustCompile(`(?:[^']|'[^']*')*?(/\*[^*]*\*+(?:[^/*][^*]*\*+)*/)`)
)

type SqlSelectelementsVisitorImpl struct {
	parser.SqlVisitor
	vars map[string]interface{}
}

func (s *SqlSelectelementsVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.SelectElementsContext:
		arr := val.Accept(s).([]interface{})
		//有*只返回*
		if val.GetStar() != nil {
			return &model.NodeSelectelements{NodeType: model.NodeStar, Elements: nil}
		}

		var eles []*model.Element
		for _, v := range arr {
			eles = append(eles, v.(*model.Element))
		}
		return &model.NodeSelectelements{NodeType: model.NodeElements, Elements: eles}
	default:
		panic("not impl")
	}
}
func (s *SqlSelectelementsVisitorImpl) VisitSelectElements(ctx *parser.SelectElementsContext) interface{} {
	var res []interface{}
	for k := range ctx.AllSelectElement() {
		res = append(res, ctx.SelectElement(k).Accept(s))
	}
	return res
}

func (s *SqlSelectelementsVisitorImpl) VisitSelectElement(ctx *parser.SelectElementContext) interface{} {
	var tableAlias, columnName, exprStr, asStr, aliaStr string
	expr, as, alias := ctx.Expr(), ctx.AS(), ctx.Alias()
	if expr != nil {
		//分隔a.xx，识别表别名
		if expr.ColumnName() != nil {
			col := expr.ColumnName().GetText()
			result := regPoint.FindAllString(col, -1)
			if len(result) == 1 {
				columnName = result[0]
			}
			if len(result) == 2 {
				tableAlias = result[0]
				columnName = result[1]
			}
		} else {
			exprStr = expr.GetText()
		}
	}
	if as != nil {
		asStr = as.GetText()
	}
	if alias != nil {
		aliaStr = alias.GetText()
	}
	return NewElement(tableAlias, columnName, exprStr, asStr, aliaStr)
}
