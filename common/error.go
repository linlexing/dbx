package common

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

//SQLError 表示一个sql语句执行出错
type SQLError struct {
	SQL    string
	Params interface{}
	Err    error
}

//NewSQLError 构造
func NewSQLError(err error, sql string, params ...interface{}) SQLError {
	//如果已经是sqlerr，则不进行再次的包装，因为有ddb的存在
	if v, ok := err.(SQLError); ok {
		return v
	}
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
	if yes, table := isDuplicatePKErrorPostgres(e.Err); yes {
		return "DuplicatePKError:" + table
	}
	// l := 0
	// content := fmt.Sprintf("%#v", e.Params)
	// switch tv := e.Params.(type) {
	// case []interface{}:
	// 	l = len(tv)
	// case map[string]interface{}:
	// 	l = len(tv)
	// 	list := []string{}
	// 	for k, v := range tv {
	// 		switch subtv := v.(type) {
	// 		case time.Time:
	// 			list = append(list, fmt.Sprintf("%s(time):%s", k, subtv.Format(time.RFC3339)))
	// 		default:
	// 			list = append(list, fmt.Sprintf("%s(%T):%#v", k, v, v))
	// 		}

	// 	}
	// 	content = strings.Join(list, "\n")
	// }
	// return fmt.Sprintf("%s\n%s\nparams len is %d,content is:\n%s", e.Err, e.SQL, l, content)
	return fmt.Sprintf("%s\n%s\nparams:\n%s", e.Err, e.SQL, spew.Sdump(e.Params))
}

//isDuplicatePKErrorPostgres 是否主键重复错误，以后可以优化成用reflect来读取结构值
func isDuplicatePKErrorPostgres(err error) (yes bool, table string) {
	v := reflect.ValueOf(err)
	if v.Kind() != reflect.Ptr {
		return
	}
	v = v.Elem()
	if vt := v.Type(); vt.PkgPath() == "github.com/lib/pq" && vt.Name() == "Error" {
		if v.FieldByName("Code").String() == "23505" {
			return true, v.FieldByName("Table").String()
		}
	}
	return false, ""
}
