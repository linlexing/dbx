package schema

import "reflect"

//MetaField 一个struct field和dbfield的结合体
type MetaField struct {
	Name     string
	DBName   string
	IsChild  bool
	Children []*MetaField
}

//Struct2MetaFields 将一个struct转换成元数据field
func Struct2MetaFields(meta interface{}) ([]*MetaField, error) {
	vtype := reflect.TypeOf(meta)
	if vtype.Kind() == reflect.Ptr {
		vtype = vtype.Elem()
	}
	conv, ok := meta.(converFieldName)
	if !ok {
		conv = nil
	}
	return struct2MetaFields(vtype, conv, "", nil, true)
}
func struct2MetaFields(vtype reflect.Type, conv converFieldName,
	parentName string, parentPath []int, root bool) ([]*MetaField, error) {
	list, err := fieldsFromStruct(vtype, conv, parentName, parentPath, root)
	if err != nil {
		return nil, err
	}
	rev := []*MetaField{}
	for _, one := range list {
		if one.child {

			slist, err := struct2MetaFields(one.st, conv, one.fieldName, parentPath, false)
			if err != nil {
				return nil, err
			}
			rev = append(rev, &MetaField{
				Name:     one.fieldName,
				DBName:   one.childName,
				IsChild:  true,
				Children: slist,
			})
		} else {
			rev = append(rev, &MetaField{
				Name:   one.fieldName,
				DBName: one.define.Name,
			})

		}
	}
	return rev, nil
}
