package render

import (
	"bytes"
	"fmt"
	"text/template"
)

//SQLRenderError  表示一个SQL语句渲染错误
type SQLRenderError struct {
	template string

	renderArgs interface{}
	err        error
}

func (r *SQLRenderError) Error() string {
	return fmt.Sprintf("Template:\n%s\nRenderArgs:\n%#v\nError:\n%s", r.template, r.renderArgs, r.err)
}

//NewSQLRenderError 返回一个渲染错误
func NewSQLRenderError(err error, temp string, args interface{}) *SQLRenderError {
	rev := &SQLRenderError{
		template:   temp,
		err:        err,
		renderArgs: args,
	}

	return rev
}

//RenderSQL 渲染一个sql,后期如果速度慢，可以加入一个模板缓存
func RenderSQL(strSQL string, renderArgs interface{}) (string, error) {
	return RenderSQLCustom(strSQL, "{{", "}}", renderArgs)
}

func RenderSQLCustom(strSQL, delimLeft, delimRight string, renderArgs interface{}) (string, error) {

	if len(strSQL) == 0 {
		return strSQL, nil
	}
	var err error
	var t *template.Template
	if t, err = template.New("sql").Delims(delimLeft, delimRight).Funcs(tempFunc).Parse(strSQL); err != nil {
		return "", NewSQLRenderError(err, strSQL, renderArgs)
	}

	out := bytes.NewBuffer(nil)
	if err = t.Execute(out, renderArgs); err != nil {
		return "", NewSQLRenderError(err, strSQL, renderArgs)
	}
	strSQL = out.String()
	return strSQL, nil
}
