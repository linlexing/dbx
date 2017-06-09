package schema

import (
	"errors"
	"reflect"

	"strings"
	"time"
)

type structField struct {
	name    string  //结构属性名称
	define  *colDef //字段定义
	st      reflect.Type
	valPath []int //属性路径，即每层属性序号
}

func (s *structField) get(obj reflect.Value) interface{} {
	p := obj
	for _, i := range s.valPath {
		p = p.Field(i)
	}

	//空字符串返回nil
	switch s.st.Kind() {
	case reflect.String:
		if s.define.Null && len(p.String()) == 0 {
			return nil
		}
	}
	//0时间返回nil
	switch tv := p.Interface().(type) {
	case time.Time:
		if s.define.Null && tv.IsZero() {
			return nil
		}
	}

	return p.Interface()
}
func (s *structField) set(obj reflect.Value, val interface{}) {
	p := obj
	for _, i := range s.valPath {
		p = p.Field(i)
	}
	//nil特殊处理
	if val == nil {
		//字符串设置成空串
		switch s.st.Kind() {
		case reflect.String:
			p.SetString("")
			return
		}
		//时间设置成0值
		switch p.Interface().(type) {
		case time.Time:
			p.Set(reflect.ValueOf(time.Time{}))
			return
		}
	}
	p.Set(reflect.ValueOf(val))
}

func fieldsFromStruct(vtype reflect.Type, parentPath []int) (rev []*structField, err error) {
	rev = []*structField{}
	if vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}
	if vtype.Kind() != reflect.Struct {
		err = errors.New("val type not is struct")
		return
	}
	for i := 0; i < vtype.NumField(); i++ {
		field := vtype.Field(i)
		newPath := []int{}
		copy(newPath, parentPath)
		newPath = append(newPath, i)
		if field.Anonymous {
			list, e := fieldsFromStruct(field.Type, newPath)
			if e != nil {
				err = e
				return
			}
			rev = append(rev, list...)
			continue
		}
		sf := &structField{
			name:    field.Name,
			valPath: newPath,
			st:      field.Type,
		}
		name := strings.ToUpper(field.Name)
		tag, ok := field.Tag.Lookup("dbx")
		//没有定义，则只有名称
		if !ok || len(tag) == 0 {
			sf.define, err = columnDefine(name)
			if err != nil {
				return
			}
			rev = append(rev, sf)
			continue
		}
		tag = strings.ToUpper(tag)
		//如果有定义，则必定是类型名称或者字段名+类型
		tags := strings.Fields(tag)

		//如果不是任意一种类型名，则说明是完整的定义，即名称开始
		if !strings.HasPrefix(tags[0], TypeString.String()) &&
			!strings.HasPrefix(tags[0], TypeBytea.String()) &&
			!strings.HasPrefix(tags[0], TypeDatetime.String()) &&
			!strings.HasPrefix(tags[0], TypeFloat.String()) &&
			!strings.HasPrefix(tags[0], TypeInt.String()) {

			sf.define, err = columnDefine(tag)
			if err != nil {
				return
			}
			rev = append(rev, sf)
			continue
		}
		//到这就说明，是省略名称的定义，补上大写的名称
		sf.define, err = columnDefine(name + " " + tag)
		if err != nil {
			return
		}
		rev = append(rev, sf)
	}
	return rev, nil
}

//TableFromStruct 将一个struct转换成table
func TableFromStruct(tableName string, meta interface{}) (*Table, error) {
	vtype := reflect.TypeOf(meta)
	if vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}
	list, err := fieldsFromStruct(vtype, nil)
	if err != nil {
		return nil, err
	}
	//剥离出其中的定义
	coldefs := []*colDef{}
	for _, one := range list {
		coldefs = append(coldefs, one.define)
	}
	tab := NewTable(tableName)
	tab.Columns, tab.PrimaryKeys = columnsDefine(coldefs)
	return tab, nil
}

//Struct2Row 结构体的值转换成map
func Struct2Row(meta interface{}) (map[string]interface{}, error) {
	vtype := reflect.TypeOf(meta)
	if vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}
	types, err := fieldsFromStruct(vtype, nil)
	if err != nil {
		return nil, err
	}
	vval := reflect.ValueOf(meta)
	if vval.Kind() == reflect.Ptr {
		vval = vval.Elem()
	}
	result := map[string]interface{}{}
	for _, v := range types {
		result[v.define.Name] = v.get(vval)
	}
	return result, nil
}

//Row2Struct map转换成结构体的值
func Row2Struct(row map[string]interface{}, meta interface{}) error {
	vtype := reflect.TypeOf(meta)
	if vtype.Kind() != reflect.Ptr {
		return errors.New("the meta must is ptr")
	}
	vtype = vtype.Elem()
	types, err := fieldsFromStruct(vtype, nil)
	if err != nil {
		return err
	}
	vval := reflect.ValueOf(meta)
	if vval.Kind() == reflect.Ptr {
		vval = vval.Elem()
	}
	for _, v := range types {
		if tv, ok := row[v.define.Name]; ok {
			v.set(vval, tv)
		}
	}
	return nil
}
