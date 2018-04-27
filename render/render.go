package render

import (
	"bytes"
	"fmt"
	"text/template"
)

//SQLRenderError  表示一个SQL语句渲染错误
type SQLRenderError struct {
	template   string
	funcs      template.FuncMap
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
func RenderSQL(strSQL string, renderArgs interface{}, funcs ...template.FuncMap) (string, error) {
	var fs template.FuncMap
	if len(funcs) > 1 {
		return "", fmt.Errorf("funcs can't >1")
	}
	if len(funcs) > 0 {
		fs = funcs[0]
	}
	return RenderSQLCustom(strSQL, "{{", "}}", renderArgs, fs)
}

func RenderSQLCustom(strSQL, delimLeft, delimRight string, renderArgs interface{}, funcs template.FuncMap) (string, error) {

	if len(strSQL) == 0 {
		return strSQL, nil
	}
	var err error
	var t *template.Template = template.New("sql").Delims(delimLeft, delimRight).Funcs(tempFunc)
	if funcs != nil {
		t = t.Funcs(funcs)
	}
	if t, err = t.Parse(strSQL); err != nil {
		return "", NewSQLRenderError(err, strSQL, renderArgs)
	}

	out := bytes.NewBuffer(nil)
	if err = t.Execute(out, renderArgs); err != nil {
		return "", NewSQLRenderError(err, strSQL, renderArgs)
	}
	strSQL = out.String()
	return strSQL, nil
}
