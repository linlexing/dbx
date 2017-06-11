package schema

import (
	"errors"
	"reflect"

	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type structField struct {
	fieldName string       //结构属性名称
	define    *colDef      //字段定义
	st        reflect.Type //Go类型
	child     bool         //是否明细表
	childName string       //如是明细表，这里是表名称
	valPath   []int        //属性路径，即每层属性序号
}

//checkType 检查数据库类型和实际类型是否相容
//数据库类型 Go类型
//String string struct slice map
//Bytea  []byte struct slice map
//Datetime time.Time
//Float  float64
//Int    int64
//child []struct
func (s *structField) checkType() error {
	dt := s.define.Type
	st := s.st
	if s.child {
		if st.Kind() == reflect.Slice &&
			st.Elem().Kind() == reflect.Struct {
			return nil
		}
		return errors.New("child table must is slice of struct")
	}
	switch dt {
	case TypeString:
		switch st.Kind() {
		case
			reflect.String,
			reflect.Struct,
			reflect.Slice,
			reflect.Map:
			return nil
		}
		return fmt.Errorf("string type must be one of string、Struct、Slice、Map")
	case TypeBytea:
		//[]byte 已经在下面的slice的范围中
		//if st.Kind() == reflect.Slice && st.Elem().Kind() == reflect.Uint8 {
		//	return nil
		//}
		switch st.Kind() {
		case
			reflect.Struct,
			reflect.Slice,
			reflect.Map:
			return nil
		}
		return fmt.Errorf("typea type must be one of []byte、Struct、Slice、Map")
	case TypeDatetime:
		if st.ConvertibleTo(reflect.TypeOf(time.Time{})) {
			return nil
		}
		return fmt.Errorf("datetime type must is time.Time")
	case TypeFloat:
		if st.Kind() == reflect.Float64 {
			return nil
		}
		return fmt.Errorf("float type must is float64")
	case TypeInt:
		if st.Kind() == reflect.Int64 {
			return nil
		}
		return fmt.Errorf("int type must is int64")
	default:
		return errors.New("invalid type:" + dt.String())
	}
}

//json 指示该字段是否用json转换
func (s *structField) json() bool {
	return !s.child && s.define.Type == TypeString &&
		(s.st.Kind() == reflect.Struct ||
			s.st.Kind() == reflect.Slice ||
			s.st.Kind() == reflect.Map)
}

//gob 指示该字段是否用gob转换
func (s *structField) gob() bool {
	if s.child {
		return false
	}
	if s.st.Kind() == reflect.Slice &&
		s.st.Elem().Kind() == reflect.Uint8 {
		return false
	}
	return s.define.Type == TypeBytea &&
		(s.st.Kind() == reflect.Struct ||
			s.st.Kind() == reflect.Slice ||
			s.st.Kind() == reflect.Map)
}

//isZero 判断一个值是否是0值
func (s *structField) isZero(val reflect.Value) bool {
	switch s.st.Kind() {
	case reflect.Slice, reflect.Map:
		if val.IsNil() {
			return true
		}
	}
	if s.child {
		return val.Len() == 0
	}
	switch s.define.Type {
	case TypeString:
		switch s.st.Kind() {
		case reflect.String, reflect.Slice, reflect.Map:
			return val.Len() == 0
		case reflect.Struct:
			return reflect.DeepEqual(reflect.Zero(s.st), val.Interface())
		}
		panic("invalid type")
	case TypeBytea:
		switch s.st.Kind() {
		case reflect.Slice, reflect.Map:
			return val.Len() == 0
		case reflect.Struct:
			return reflect.DeepEqual(reflect.Zero(s.st), val)
		}
		panic("invalid type")
	case TypeDatetime:
		return val.Interface().(time.Time).IsZero()
	case TypeFloat:
		return val.Interface().(float64) == 0
	case TypeInt:
		return val.Interface().(int64) == 0
	default:
		panic("invalid dbtype")
	}
}

//get 会自动转换
//空值且数据库可为空返回nil
//str 和 bytea类型字段，如果值是struct slice map，则返回json 和 gob
func (s *structField) get(obj reflect.Value) (interface{}, error) {
	p := s.getv(obj)
	//子表必定是[]struct
	if s.child {
		if s.isZero(p) {
			return nil, nil
		}
		clist := []map[string]interface{}{}
		for i := 0; i < p.Len(); i++ {
			cm, cd, e := struct2Row(p.Index(i))
			if e != nil {
				return nil, e
			}
			//子表不能再有子表，防止问题复杂化
			if len(cd) > 0 {
				err := fmt.Errorf("field:%s has sub child data", s.fieldName)
				return nil, err
			}
			clist = append(clist, cm)
		}
		return clist, nil
	}
	if s.define.Null && s.isZero(p) {
		return nil, nil
	}
	if s.json() {

		bys, err := json.Marshal(p.Interface())
		if err != nil {
			return nil, err
		}
		return string(bys), nil
	}
	if s.gob() {

		bys := bytes.NewBuffer(nil)
		if err := gob.NewEncoder(bys).EncodeValue(p); err != nil {
			return nil, err
		}
		return bys.Bytes(), nil
	}

	return p.Interface(), nil
}
func (s *structField) getv(obj reflect.Value) reflect.Value {
	p := obj
	for _, i := range s.valPath {
		p = p.Field(i)
	}
	for p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
	return p
}
func (s *structField) set(obj reflect.Value, val interface{}) error {
	//nil 设置为0值
	if val == nil {
		s.setv(obj, reflect.Zero(s.st))
		return nil
	}
	//子表
	if s.child {
		list := reflect.New(s.st)
		for _, row := range val.([]map[string]interface{}) {
			rv := reflect.New(s.st.Elem())
			if err := row2Struct(row, nil, rv); err != nil {
				return err
			}
			list.Elem().Set(reflect.Append(list.Elem(), rv.Elem()))
		}
		s.setv(obj, list.Elem())
		return nil
	}
	if s.json() {
		newv := reflect.New(s.st)
		if err := json.Unmarshal([]byte(val.(string)),
			newv.Interface()); err != nil {
			return err
		}
		s.setv(obj, newv.Elem())
		return nil
	}
	if s.gob() {
		newv := reflect.New(s.st)
		if err := gob.NewDecoder(bytes.NewBuffer(val.([]byte))).
			Decode(newv.Interface()); err != nil {
			return err
		}
		s.setv(obj, newv.Elem())
		return nil
	}
	s.setv(obj, reflect.ValueOf(val))
	return nil
}
func (s *structField) setv(obj reflect.Value, val reflect.Value) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(s.fieldName, s.st)
			panic(r)
		}
	}()
	p := obj
	for _, i := range s.valPath {
		p = p.Field(i)
	}
	p.Set(val)
}

//fieldsFromStruct 读取一个结构体，转换成元数据，可以接受一个*struct、
//[]struct、[]*struct,并需传入一个属性路径索引数组，方便后期赋值
//允许匿名嵌套
func fieldsFromStruct(vtype reflect.Type, parentPath []int) (rev []*structField, err error) {
	rev = []*structField{}
	for vtype.Kind() == reflect.Slice ||
		vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}
	if vtype.Kind() != reflect.Struct {
		err = errors.New("val type not is struct")
		return
	}
	for i := 0; i < vtype.NumField(); i++ {
		field := vtype.Field(i)
		newPath :=make([]int,len(parentPath))
		copy(newPath, parentPath)
		newPath = append(newPath, i)
		//嵌套结构，扁平化
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
			fieldName: field.Name,
			valPath:   newPath,
			st:        field.Type,
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

		//如果不是任意一种类型名或子表，则说明是完整的定义，即名称开始
		if !strings.HasPrefix(tags[0], TypeString.String()) &&
			!strings.HasPrefix(tags[0], TypeBytea.String()) &&
			!strings.HasPrefix(tags[0], TypeDatetime.String()) &&
			!strings.HasPrefix(tags[0], TypeFloat.String()) &&
			!strings.HasPrefix(tags[0], TypeInt.String()) &&
			!strings.HasPrefix(tags[0], "CHILD") {
			//只有两列，末列是CHILD的，说明是子表
			if len(tags) == 2 && tags[1] == "CHILD" {
				sf.child = true
				sf.childName = tags[0]
			} else {
				sf.define, err = columnDefine(tag)
				if err != nil {
					return
				}
			}
			rev = append(rev, sf)
			continue
		}
		//到这就说明，是省略名称的定义，补上大写的名称
		if tag == "CHILD" {
			sf.child = true
			sf.childName = name
		} else {
			sf.define, err = columnDefine(name + " " + tag)
			if err != nil {
				return
			}
		}
		//检查字段类型是否正确
		if err := sf.checkType(); err != nil {
			return nil, fmt.Errorf("field:%s error:%s", sf.fieldName, err)
		}
		rev = append(rev, sf)
	}
	return rev, nil
}
func struct2Table(tableName string, vtype reflect.Type, parentPath []int) ([]*Table, error) {
	list, err := fieldsFromStruct(vtype, parentPath)
	if err != nil {
		return nil, err
	}
	//剥离出其中的定义
	result := []*Table{NewTable(tableName)}
	coldefs := []*colDef{}
	for _, one := range list {

		//子表递归调用
		if one.child {
			tabs, err := struct2Table(one.childName, one.st, parentPath)
			if err != nil {
				return nil, err
			}
			result = append(result, tabs...)
		} else {
			coldefs = append(coldefs, one.define)
		}
	}

	result[0].Columns, result[0].PrimaryKeys = columnsDefine(coldefs)
	return result, nil

}

//TableFromStruct 将一个struct转换成table清单,可能有明细表
//第一个是主表，剩余是明细表
func TableFromStruct(tableName string, meta interface{}) ([]*Table, error) {
	return struct2Table(tableName, reflect.TypeOf(meta), nil)
}
func struct2Row(vval reflect.Value) (main map[string]interface{},
	child map[string][]map[string]interface{}, err error) {
	types, err := fieldsFromStruct(vval.Type(), nil)
	if err != nil {
		return
	}
	if vval.Kind() == reflect.Ptr {
		vval = vval.Elem()
	}
	main = map[string]interface{}{}
	child = map[string][]map[string]interface{}{}
	for _, v := range types {
		var clist interface{}
		if clist, err = v.get(vval); err != nil {
			return
		}
		if v.child {
			if clist != nil {
				child[v.childName] = clist.([]map[string]interface{})
			}
		} else {
			if main[v.define.Name], err = v.get(vval); err != nil {
				return
			}
		}
	}
	return
}

//Struct2Row 结构体的值转换成map
func Struct2Row(meta interface{}) (main map[string]interface{},
	detail map[string][]map[string]interface{}, err error) {
	return struct2Row(reflect.ValueOf(meta))
}

//Row2Struct map转换成结构体的值
func Row2Struct(row map[string]interface{},
	child map[string][]map[string]interface{}, vval interface{}) error {
	return row2Struct(row, child, reflect.ValueOf(vval))
}
func row2Struct(row map[string]interface{},
	child map[string][]map[string]interface{}, vval reflect.Value) error {
	types, err := fieldsFromStruct(vval.Type(), nil)
	if err != nil {
		return err
	}
	for vval.Kind() == reflect.Ptr {
		vval = vval.Elem()
	}
	/*	for vval.Kind() != reflect.Ptr {
		return fmt.Errorf("vval must is ptr")
	}*/
	for _, v := range types {
		if v.child {
			if clist, ok := child[v.childName]; ok {
				v.set(vval, clist)
			} else {
				v.set(vval, nil)
			}
		} else {
			if tv, ok := row[v.define.Name]; ok {
				v.set(vval, tv)
			} else {
				v.set(vval, nil) //没找到的属性要设置成nil，防止有旧值
			}
		}
	}
	return nil
}
