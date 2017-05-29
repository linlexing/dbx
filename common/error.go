package common

import (
	"fmt"
	"strings"
	"time"
)

//SQLError 表示一个sql语句执行出错
type SQLError struct {
	SQL    string
	Params interface{}
	Err    error
}

//NewSQLError 构造
func NewSQLError(err error, sql string, params ...interface{}) SQLError {
	var p interface{}
	if len(params) == 1 {
		p = params[0]
	} else {
		p = params
	}
	return SQLError{
		SQL:    sql,
		Params: p,
		Err:    err,
	}
}
func (e SQLError) Error() string {
	l := 0
	content := fmt.Sprintf("%#v", e.Params)
	switch tv := e.Params.(type) {
	case []interface{}:
		l = len(tv)
	case map[string]interface{}:
		l = len(tv)
		list := []string{}
		for k, v := range tv {
			switch subtv := v.(type) {
			case time.Time:
				list = append(list, fmt.Sprintf("%s(time):%s", k, subtv.Format(time.RFC3339)))
			default:
				list = append(list, fmt.Sprintf("%s(%T):%#v", k, v, v))
			}

		}
		content = strings.Join(list, "\n")
	}
	return fmt.Sprintf("%s\n%s\nparams len is %d,content is:\n%s", e.Err, e.SQL, l, content)
}
