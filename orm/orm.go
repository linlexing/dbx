package orm

import (
	"errors"
	"reflect"

	"strings"

	"github.com/linlexing/dbx/schema"
)

//TableNamer 用来返回一个dbtablename
type TableNamer interface {
	TableName() string
}

func scriptFromStruct(vtype reflect.Type) (rev []string, err error) {
	rev = []string{}
	if vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}
	if vtype.Kind() != reflect.Struct {
		err = errors.New("val type not is struct")
		return
	}
	for i := 0; i < vtype.NumField(); i++ {
		field := vtype.Field(i)
		if field.Anonymous {
			list, e := scriptFromStruct(field.Type)
			if e != nil {
				err = e
				return
			}
			rev = append(rev, list...)
			continue
		}
		name := strings.ToUpper(field.Name)
		tag, ok := field.Tag.Lookup("dbx")
		//没有定义，则只有名称
		if !ok || len(tag) == 0 {
			rev = append(rev, name)
			continue
		}
		tag = strings.ToUpper(tag)
		//如果有定义，则必定是类型名称或者字段名+类型
		tags := strings.Fields(tag)
		//如果不是任意一种类型名，则说明是完整的定义，即名称开始
		if !strings.HasPrefix(tags[0], schema.TypeString.String()) &&
			!strings.HasPrefix(tags[0], schema.TypeBytea.String()) &&
			!strings.HasPrefix(tags[0], schema.TypeDatetime.String()) &&
			!strings.HasPrefix(tags[0], schema.TypeFloat.String()) &&
			!strings.HasPrefix(tags[0], schema.TypeInt.String()) {
			rev = append(rev, tag)
			continue
		}
		//到这就说明，是省略名称的定义，补上大写的名称
		rev = append(rev, name+" "+tag)
	}
	return rev, nil
}

//TableFromStruct 将一个struct转换成table
func TableFromStruct(meta interface{}) (*schema.Table, error) {
	vtype := reflect.TypeOf(meta)
	if vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}
	tableName := strings.ToUpper(vtype.Name())
	if tnr, ok := meta.(TableNamer); ok {
		tableName = tnr.TableName()
	}
	list, err := scriptFromStruct(vtype)
	if err != nil {
		return nil, err
	}
	tab := schema.NewTable(tableName)
	if err := tab.DefineScript(strings.Join(list, "\n")); err != nil {
		return nil, err
	}
	return tab, nil
}
