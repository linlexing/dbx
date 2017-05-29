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
	return f.Type == src.Type &&
		(f.MaxLength == src.MaxLength ||
			f.MaxLength <= 0 && src.MaxLength <= 0) &&
		f.Null == src.Null
}

//Clone 复制一个字段
func (f *Column) Clone() *Column {
	return &Column{f.Name, f.Type, f.MaxLength, f.Null,
		f.TrueType, f.FetchDriver, f.Index, f.IndexName, f.FormerName}
}
