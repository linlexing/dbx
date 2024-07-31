package schema

import "strings"

const (
	//NoIndex 没有索引
	NoIndex = ""
	//Index 正常的索引
	Index = "i"
	//UniqueIndex 唯一索引
	UniqueIndex = "ui"
)

//Column 字段定义
type Column struct {
	Name         string
	Type         DataType
	MaxLength    int            `json:",omitempty"`
	Null         bool           `json:",omitempty"`
	TrueType     string         `json:",omitempty"`
	FetchDriver  string         `json:",omitempty"` //上次获取字段信息时，数据库驱动的名称
	Index        string         `json:",omitempty"`
	IndexName    string         `json:",omitempty"` //如果该字段有索引，存放数据库中索引的名称
	FormerName   []string       `json:",omitempty"` //曾用名，可以放多个，一个字段改名后，旧名称应当永久不使用，方便从任意版本更新到最新版本
	PropertyName string         `json:",omitempty"` //属性名称，只有从struct转换来时才有，一般是原样大小写的名称
	Extended     map[string]any `json:",omitempty"` //扩展属性，内部没有使用，自由引用
}

//Eque 判定两个字段定义是否相等
func (f *Column) Eque(src *Column) bool {
	if strings.EqualFold(f.Name, src.Name) {
		return false
	}
	// FetchDriver为空，则说明是内存定义，不是从数据库取回，truetype是手工赋予，可以适用更新数据库
	if (len(f.FetchDriver) == 0 || len(src.FetchDriver) == 0 ||
		strings.EqualFold(f.FetchDriver, src.FetchDriver)) &&
		len(f.TrueType) > 0 && len(src.TrueType) > 0 {
		return f.TrueType == src.TrueType &&
			f.Index == src.Index
	}
	if f.Type != src.Type {
		return false
	}
	switch f.Type {
	//日期、数值、整型不需要判断长度
	case TypeDatetime, TypeFloat, TypeInt:
		return f.Null == src.Null &&
			f.Index == src.Index
	default:
		//历史原因，MaxLength <=0 只有一个含义，无限的长度
		//历史代码中有时用了-1,有时是0，所以都是<=0的视为相等
		return (f.MaxLength == src.MaxLength ||
			f.MaxLength <= 0 && src.MaxLength <= 0) &&
			f.Null == src.Null &&
			f.Index == src.Index
	}
}

//EqueNoIndexAndName 不判断索引和名称
func (f *Column) EqueNoIndexAndName(src *Column) bool {
	if !f.EqueType(src) {
		return false
	}
	return f.Null == src.Null
}
func (f *Column) EqueType(src *Column) bool {
	if (len(f.FetchDriver) == 0 || len(src.FetchDriver) == 0 ||
		f.FetchDriver == src.FetchDriver) &&
		len(f.TrueType) > 0 && len(src.TrueType) > 0 {
		return f.TrueType == src.TrueType
	}

	if f.Type != src.Type {
		return false
	}
	switch f.Type {
	//日期、数值、整型不需要判断长度
	case TypeDatetime, TypeFloat, TypeInt:
		return true
	default:
		//历史原因，MaxLength <=0 只有一个含义，无限的长度
		//历史代码中有时用了-1,有时是0，所以都是<=0的视为相等
		return (f.MaxLength == src.MaxLength ||
			f.MaxLength <= 0 && src.MaxLength <= 0)
	}
}

//Clone 复制一个字段
func (f *Column) Clone() *Column {
	var fns []string
	if len(f.FormerName) > 0 {
		fns = make([]string, len(f.FormerName))
		copy(fns, f.FormerName)
	}
	cpmap := map[string]any{}
	for k, v := range f.Extended {
		cpmap[k] = v
	}

	return &Column{f.Name, f.Type, f.MaxLength, f.Null,
		f.TrueType, f.FetchDriver, f.Index, f.IndexName, fns, f.PropertyName, cpmap}
}
