package schema

//Column 字段定义
type Column struct {
	Name        string
	Type        DataType
	MaxLength   int
	Null        bool
	TrueType    string
	FetchDriver string //上次获取字段信息时，数据库驱动的名称

	Index      bool
	IndexName  string   //如果该字段有索引，存放数据库中索引的名称
	FormerName []string //曾用名，可以放多个，一个字段改名后，旧名称应当永久不使用，方便从任意版本更新到最新版本
}

//Eque 判定两个字段定义是否相等
func (f *Column) Eque(src *Column) bool {
	if f.Name != src.Name {
		return false
	}
	if f.FetchDriver == src.FetchDriver &&
		len(f.TrueType) > 0 && len(src.TrueType) > 0 {
		return f.TrueType == src.TrueType
	}
	//历史原因，MaxLength <=0 只有一个含义，无限的长度
	//历史代码中有时用了-1,有时是0，所以都是<=0的视为相等
	return f.Type == src.Type &&
		(f.MaxLength == src.MaxLength ||
			f.MaxLength <= 0 && src.MaxLength <= 0) &&
		f.Null == src.Null &&
		f.Index == src.Index
}

//EqueNotIndex 不判断索引
func (f *Column) EqueNoIndex(src *Column) bool {
	if f.Name != src.Name {
		return false
	}
	if f.FetchDriver == src.FetchDriver &&
		len(f.TrueType) > 0 && len(src.TrueType) > 0 {
		return f.TrueType == src.TrueType
	}
	//历史原因，MaxLength <=0 只有一个含义，无限的长度
	//历史代码中有时用了-1,有时是0，所以都是<=0的视为相等
	return f.Type == src.Type &&
		(f.MaxLength == src.MaxLength ||
			f.MaxLength <= 0 && src.MaxLength <= 0) &&
		f.Null == src.Null
}

//Clone 复制一个字段
func (f *Column) Clone() *Column {
	var fns []string
	if f.FormerName != nil {
		copy(fns, f.FormerName)
	}
	return &Column{f.Name, f.Type, f.MaxLength, f.Null,
		f.TrueType, f.FetchDriver, f.Index, f.IndexName, fns}
}
