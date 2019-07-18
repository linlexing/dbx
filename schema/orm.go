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

type converFieldName interface {
	ConvertFieldName(childName string, str string) string
}
type structField struct {
	fieldName  string          //结构属性名称
	define     *colDef         //字段定义
	st         reflect.Type    //Go类型
	child      bool            //是否明细表
	childName  string          //如是明细表，这里是表名称
	conv       converFieldName //用于明细表
	valPath    []int           //属性路径，即每层属性序号
	parentName string          //如果是子表属性，这里是子表属性名称，用于指标名称转换
}

//checkType 检查数据库类型和实际类型是否相容
//数据库类型 Go类型
//String string struct slice map bool
//Bytea  []byte struct slice map
//Datetime time.Time
//Float  float64
//Int    int64,bool
//child []struct
func (s *structField) checkType(root bool) error {
	st := s.st
	if s.child {
		if !root {
			return errors.New("child table must at root level")
		}
		if st.Kind() == reflect.Slice &&
			st.Elem().Kind() == reflect.Struct {
			return nil
		}
		return errors.New("child table must is slice of struct")
	}
	dt := s.define.Type
	switch dt {
	case TypeString:
		switch st.Kind() {
		case
			reflect.Bool,
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
		switch st.Kind() {
		case reflect.Int64, reflect.Bool:
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
		case reflect.Bool:
			return val.Bool() == false
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
		switch s.st.Kind() {
		case reflect.Bool:
			return val.Bool() == false
		case reflect.Int64:
			return val.Interface().(int64) == 0
		}
		panic("invalid dbtype")
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
			cm, e := childStruct2Row(p.Index(i), s.conv, s.fieldName)
			if e != nil {
				return nil, e
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
	//数据库没有Bool类型，需要转换
	if s.st.Kind() == reflect.Bool {
		switch s.define.Type {
		case TypeString:
			if p.Bool() {
				return "1", nil
			}
			return "0", nil

		case TypeInt:
			if p.Bool() {
				return int64(1), nil
			}
			return int64(0), nil

		}
	}
	return p.Interface(), nil
}
func (s *structField) getv(obj reflect.Value) reflect.Value {
	p := obj.FieldByIndex(s.valPath)
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
			if err := childRow2Struct(row, rv, s.conv, s.fieldName); err != nil {
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
	//数据库没有Bool类型，需要转换
	if s.st.Kind() == reflect.Bool {
		switch s.define.Type {
		case TypeString:
			switch val.(string) {
			case "1":
				s.setv(obj, reflect.ValueOf(true))
			case "0":
				s.setv(obj, reflect.ValueOf(false))
			default:
				return fmt.Errorf("string %s can't convert to bool,valid is [0,1]", val)
			}
			return nil
		case TypeInt:
			switch val.(int64) {
			case 1:
				s.setv(obj, reflect.ValueOf(true))
			case 0:
				s.setv(obj, reflect.ValueOf(false))
			default:
				return fmt.Errorf("int64 %s can't convert to bool,valid is [0,1]", val)
			}
			return nil
		}
	}
	s.setv(obj, reflect.ValueOf(val))
	return nil
}
func (s *structField) setv(obj reflect.Value, val reflect.Value) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(s.fieldName, s.define.Type, s.st, val.Interface())
			panic(r)
		}
	}()
	p := obj.FieldByIndex(s.valPath)
	p.Set(val)
}

//fieldsFromStruct 读取一个结构体，转换成元数据，可以接受一个*struct、
//[]struct、[]*struct,并需传入一个属性路径索引数组，方便后期赋值
//允许匿名嵌套,所有未导出的字段被忽略
func fieldsFromStruct(vtype reflect.Type, conv converFieldName, parentName string,
	parentPath []int, root bool) (rev []*structField, err error) {
	rev = []*structField{}
	for vtype.Kind() == reflect.Slice ||
		vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}
	if vtype.Kind() != reflect.Struct {
		err = fmt.Errorf("val type %s kind %s not is struct,parent:%s,path:%v",
			vtype.Name(), vtype.Kind(), parentName, parentPath)
		return
	}
	var preDefineTag, preTrueTypeTag string
	for i := 0; i < vtype.NumField(); i++ {
		field := vtype.Field(i)
		newPath := append(parentPath, i)
		//嵌套结构，扁平化
		if field.Anonymous {
			list, e := fieldsFromStruct(field.Type, conv, parentName, newPath, root)
			if e != nil {
				err = e
				return
			}
			rev = append(rev, list...)
			continue
		}
		//unexported
		if len(field.PkgPath) > 0 {
			continue
		}

		sf := &structField{
			fieldName:  field.Name,
			valPath:    newPath,
			st:         field.Type,
			conv:       conv,
			parentName: parentName,
		}
		name := strings.ToUpper(field.Name)
		if conv != nil {
			name = strings.ToUpper(conv.ConvertFieldName(sf.parentName, field.Name))
		}
		defineTag, defineOk := field.Tag.Lookup("dbx")
		trueTypeTag, typeTypeOk := field.Tag.Lookup("dbx_t")
		formerNameTag, _ := field.Tag.Lookup("dbx_fname")
		//只有定义和实际类型都没定义时，才从上个字段复制定义
		if !defineOk && !typeTypeOk {
			defineTag = preDefineTag
			trueTypeTag = preTrueTypeTag
		} else {
			preDefineTag = defineTag
			preTrueTypeTag = trueTypeTag
		}
		if len(defineTag) == 0 {
			//没有定义，则只有名称
			defineTag = name
		}
		defineTag = strings.ToUpper(defineTag)
		//如果有定义，则必定是类型名称或者字段名+类型
		tags := strings.Fields(defineTag)

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
				sf.define, err = columnDefine(defineTag, trueTypeTag, formerNameTag)
				if err != nil {
					return
				}
			}
			rev = append(rev, sf)
			continue
		}
		//到这就说明，是省略名称的定义，补上大写的名称
		if defineTag == "CHILD" {
			sf.child = true
			sf.childName = name
		} else {
			sf.define, err = columnDefine(name+" "+defineTag, trueTypeTag, formerNameTag)
			if err != nil {
				return
			}
		}
		//检查字段类型是否正确
		if err := sf.checkType(root); err != nil {
			return nil, fmt.Errorf("%s field:%s error:%s", vtype.Name(), sf.fieldName, err)
		}
		rev = append(rev, sf)
	}
	return rev, nil
}

func struct2Table(tableName string, vtype reflect.Type, conv converFieldName,
	parentName string, parentPath []int, root bool) ([]*Table, error) {
	list, err := fieldsFromStruct(vtype, conv, parentName, parentPath, root)
	if err != nil {
		return nil, err
	}
	//剥离出其中的定义
	result := []*Table{NewTable(tableName)}
	coldefs := []*colDef{}
	for _, one := range list {

		//子表递归调用
		if one.child {
			tabs, err := struct2Table(one.childName, one.st, conv,
				one.fieldName, parentPath, false)
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
func TableFromStruct(meta interface{}, tabNames ...string) ([]*Table, error) {
	vtype := reflect.TypeOf(meta)
	if vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}
	if len(tabNames) > 1 {
		return nil, fmt.Errorf("table names %v len > 1", tabNames)
	}
	var tabName string
	if len(tabNames) > 0 {
		tabName = tabNames[0]
	} else {
		tabName = strings.ToUpper(vtype.Name())
	}
	if len(tabNames) > 1 {
		panic("tableNames must is one or none")
	}
	conv, ok := meta.(converFieldName)
	if !ok {
		conv = nil
	}
	return struct2Table(tabName, reflect.TypeOf(meta), conv, "", nil, true)
}
func mainStruct2Row(vval reflect.Value, conv converFieldName) (main map[string]interface{},
	child map[string][]map[string]interface{}, err error) {
	types, err := fieldsFromStruct(vval.Type(), conv, "", nil, true)
	if err != nil {
		return
	}
	if vval.Kind() == reflect.Ptr {
		vval = vval.Elem()
	}
	main = map[string]interface{}{}
	child = map[string][]map[string]interface{}{}
	for _, v := range types {
		var tmpVal interface{}
		if tmpVal, err = v.get(vval); err != nil {
			return
		}
		if v.child {
			if tmpVal != nil {
				child[v.childName] = tmpVal.([]map[string]interface{})
			}
		} else {
			main[v.define.Name] = tmpVal
		}
	}
	return
}
func childStruct2Row(vval reflect.Value, conv converFieldName,
	parentName string) (main map[string]interface{}, err error) {
	types, err := fieldsFromStruct(vval.Type(), conv, parentName, nil, false)
	if err != nil {
		return
	}
	if vval.Kind() == reflect.Ptr {
		vval = vval.Elem()
	}
	main = map[string]interface{}{}
	for _, v := range types {
		if main[v.define.Name], err = v.get(vval); err != nil {
			return
		}
	}
	return
}

//Struct2Row 结构体的值转换成map
func Struct2Row(meta interface{}) (main map[string]interface{},
	detail map[string][]map[string]interface{}, err error) {
	conv, ok := meta.(converFieldName)
	if !ok {
		conv = nil
	}
	return mainStruct2Row(reflect.ValueOf(meta), conv)
}

//Row2Struct map转换成结构体的值
func Row2Struct(row map[string]interface{},
	child map[string][]map[string]interface{}, vval interface{}) error {
	conv, ok := vval.(converFieldName)
	if !ok {
		conv = nil
	}
	return mainRow2Struct(row, child, reflect.ValueOf(vval), conv)
}
func mainRow2Struct(row map[string]interface{},
	child map[string][]map[string]interface{}, vval reflect.Value, conv converFieldName) error {
	types, err := fieldsFromStruct(vval.Type(), conv, "", nil, true)
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

//childRow2Struct 将一个子表记录转换成struct
func childRow2Struct(row map[string]interface{}, vval reflect.Value,
	conv converFieldName, parentName string) error {
	types, err := fieldsFromStruct(vval.Type(), conv, parentName, nil, false)
	if err != nil {
		return err
	}
	for vval.Kind() == reflect.Ptr {
		vval = vval.Elem()
	}
	for _, v := range types {
		if tv, ok := row[v.define.Name]; ok {
			v.set(vval, tv)
		} else {
			v.set(vval, nil) //没找到的属性要设置成nil，防止有旧值
		}
	}
	return nil
}
