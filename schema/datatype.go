package schema

import (
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"time"
)

//ErrInvalidDataType 无效的数据类型
type ErrInvalidDataType struct {
	d DataType
}

func (e *ErrInvalidDataType) Error() string {
	return fmt.Sprintf("invalid data type:%d", e.d)
}
func newErrInvalidDataType(d DataType) *ErrInvalidDataType {
	return &ErrInvalidDataType{
		d: d,
	}
}

//DataType 字段的数据类型
type DataType int

const (
	//TypeString 是字符串类型
	TypeString DataType = iota
	//TypeInt 整型
	TypeInt
	//TypeDatetime 日期
	TypeDatetime
	//TypeBytea 二进制数据
	TypeBytea
	//TypeFloat 浮点，float 和 decimal都归入此
	TypeFloat
)

//MarshalJSON 实现json的自定义的json序列化，主要是为了兼容前个直接保存字符串值的版本
func (d *DataType) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

//UnmarshalJSON 实现自定义的json反序列化，主要是为了兼容前个版本
func (d *DataType) UnmarshalJSON(v []byte) error {
	*d = ParseDataType(string(v))
	return nil
}

//ParseDataType 将一个字符串转换成类型值
func ParseDataType(d string) DataType {
	switch d {
	case "STR":
		return TypeString
	case "INT":
		return TypeInt
	case "DATE":
		return TypeDatetime
	case "FLOAT":
		return TypeFloat
	case "BYTEA":
		return TypeBytea
	default:
		panic(fmt.Errorf("invalid type:%s", d))
	}
}

//String 返回类型的字符串名称
func (d DataType) String() string {
	switch d {
	case TypeString:
		return "STR"
	case TypeInt:
		return "INT"
	case TypeDatetime:
		return "DATE"
	case TypeFloat:
		return "FLOAT"
	case TypeBytea:
		return "BYTEA"
	default:
		panic(newErrInvalidDataType(d))
	}

}

//ChineseString 返回类型的汉字名称
func (d DataType) ChineseString() string {
	switch d {
	case TypeString:
		return "字符串"
	case TypeInt:
		return "整型"
	case TypeDatetime:
		return "日期"
	case TypeFloat:
		return "浮点"
	case TypeBytea:
		return "二进制"
	default:
		panic(newErrInvalidDataType(d))
	}
}
func strToDate(s string) (tm time.Time, err error) {
	if len(s) == 8 {
		return time.Parse("20060102", s)
	}
	if len(s) == 10 {
		return time.Parse("2006-01-02", s)
	}

	tm, err = time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		tm, err = time.Parse("2006-01-02T15:04:05", s)
	}
	if err != nil {
		tm, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", s)
	}
	if err != nil {
		tm, err = time.Parse(time.RFC3339, s)
	}
	if err != nil {
		tm, err = time.Parse(time.RFC3339Nano, s)
	}
	return
}

//ParseString 将一个字符串转换成标准值
func (d DataType) ParseString(v string) interface{} {
	if len(v) == 0 {
		return nil
	}
	switch d {
	case TypeString:
		return v
	case TypeInt:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		return i

	case TypeDatetime:
		tm, err := strToDate(v)
		if err != nil {
			panic(fmt.Sprintf("%s not time value", v))
		}
		return tm
	case TypeBytea:
		str, err := base64.RawStdEncoding.DecodeString(v)
		if err != nil {
			panic("can't convert to byte array,not is base64 string," + v)
		}
		return str
	case TypeFloat:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		return f

	default:
		panic(newErrInvalidDataType(d))
	}
}

//ToString 转换成字符串形式,值必须要满足要求
func (d DataType) ToString(v interface{}) (result string) {

	//nil代表null，不需要转换，否则会出错
	if v == nil {
		return ""
	}
	switch d {
	case TypeString:
		switch tv := v.(type) {
		case []byte:
			return string(tv)
		case string:
			return tv
		default:
			panic(fmt.Sprintf("v:%#v can't to string", v))
		}
	case TypeDatetime:
		switch tv := v.(type) {
		case time.Time:
			return tv.Format("2006-01-02 15:04:05")
		case *time.Time:
			return tv.Format("2006-01-02 15:04:05")
		default:
			panic(fmt.Sprintf("v:%#v can't to string", v))
		}
	case TypeInt:
		switch tv := v.(type) {
		case int, int16, int32, int64, int8,
			uint, uint16, uint32, uint64, uint8:
			return fmt.Sprintf("%d", tv)
		default:
			panic(fmt.Sprintf("v:%#v can't to string", v))
		}
	case TypeBytea:
		switch tv := v.(type) {
		case []byte:
			return base64.RawStdEncoding.EncodeToString(tv)
		default:
			panic(fmt.Sprintf("v:%#v can't to string", v))
		}
	case TypeFloat:
		switch tv := v.(type) {
		case float32:
			return strconv.FormatFloat(float64(tv), 'f', -1, 64)
		case float64:
			return strconv.FormatFloat(tv, 'f', -1, 64)
		default:
			panic(fmt.Sprintf("v:%#v can't to string", v))
		}
	default:
		panic(newErrInvalidDataType(d))
	}

}

//ParseScan 转换数据库驱动扫描出的值，特别是time类型的，很可能是string形式
func (d DataType) ParseScan(v interface{}) interface{} {
	//nil代表null，不需要转换，否则会出错
	if v == nil {
		return nil
	}
	//空字符串当null处理
	if v == "" {
		return nil
	}
	switch d {
	case TypeString:
		switch tv := v.(type) {
		case []byte:
			return string(tv)
		default:
			panic(fmt.Sprintf("v:%#v can't to string", tv))
		}
	case TypeDatetime:
		switch tv := v.(type) {
		case time.Time:
			return tv
		case string:
			tm, err := strToDate(tv)
			if err != nil {
				panic(fmt.Sprintf("%s not is time value", tv))
			}
			return tm
		case []byte:
			tm, err := strToDate(string(tv))
			if err != nil {
				panic(fmt.Sprintf("%s not is time value", tv))
			}
			return tm
		default:
			panic(fmt.Errorf("error type,%T", v))
		}
	case TypeInt:
		switch tv := v.(type) {
		case int8:
			return int64(tv)
		case int16:
			return int64(tv)
		case int32:
			return int64(tv)
		case int:
			return int64(tv)
		case int64:
			return tv
		case uint8:
			return int64(tv)
		case uint16:
			return int64(tv)
		case uint32:
			return int64(tv)
		case uint:
			return int64(tv)
		case uint64:
			return tv
		case string:
			i, err := strconv.ParseInt(tv, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("%s not is int", tv))
			}
			return i

		case []byte:
			i, err := strconv.ParseInt(string(tv), 10, 64)
			if err != nil {
				panic(fmt.Sprintf("%s not is int", tv))
			}
			return i
		default:
			panic(fmt.Sprintf("v:%#v not is int,T:%T", tv, tv))
		}
	case TypeBytea:
		switch tv := v.(type) {
		case string:
			return []byte(tv)
		case []byte:
			return tv
		default:
			panic(fmt.Sprintf("v:%#v,T:%T not is bytea", v, v))
		}
	case TypeFloat:
		switch tv := v.(type) {
		case float32:
			return float64(tv)
		case float64:
			return tv
		case int64:
			return float64(tv)
		case int32:
			return float64(tv)
		case int16:
			return float64(tv)
		case int:
			return float64(tv)
		case int8:
			return float64(tv)
		case uint64:
			return float64(tv)
		case uint32:
			return float64(tv)
		case uint16:
			return float64(tv)
		case uint:
			return float64(tv)
		case uint8:
			return float64(tv)
		case string:
			f, err := strconv.ParseFloat(tv, 64)
			if err != nil {
				panic(fmt.Sprintf("v:%#v,T:%T not is float", tv, tv))
			}
			return f

		case []byte:
			f, err := strconv.ParseFloat(string(tv), 64)
			if err != nil {
				panic(fmt.Sprintf("v:%#v,T:%T not is float", tv, tv))
			}
			return f
		default:
			panic(fmt.Sprintf("v:%#v,T:%T not is float", tv, tv))

		}
	default:
		panic(newErrInvalidDataType(d))
	}

}

//ToJSON 转换一个字段值，方便其保存json格式，主要目地是检查数据类型和处理二进制数据和日期
func (d DataType) ToJSON(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	switch d {
	case TypeBytea: //base64
		var b []byte
		switch tv := v.(type) {
		case nil:
			b = nil
		case string:
			b = []byte(tv)
		case []byte:
			b = tv
		default:
			log.Panic(fmt.Sprintf("type:%T can't convert to bytea", v))
		}
		return base64.RawStdEncoding.EncodeToString(b), nil

	case TypeDatetime: //RFC3339
		if tv, ok := v.(time.Time); ok {
			return tv.Format(time.RFC3339), nil
		}
		return nil, fmt.Errorf("the %#v not is time", v)

	case TypeFloat:
		if tv, ok := v.(float64); ok {
			return tv, nil
		}
		return nil, fmt.Errorf("the %#v not is float64", v)

	case TypeInt:
		if tv, ok := v.(int64); ok {
			return tv, nil
		}
		return nil, fmt.Errorf("the %#v not is int64", v)

	case TypeString:

		if tv, ok := v.(string); ok {
			return tv, nil
		}
		return nil, fmt.Errorf("the %#v not is string", v)
	default:
		return nil, newErrInvalidDataType(d)
	}
}

//ParseJSON 从一个json数据中读取值，需要转换[]byte,time等类型
func (d DataType) ParseJSON(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	switch d {
	case TypeBytea: //base64
		if tv, ok := v.(string); ok {
			return base64.RawStdEncoding.DecodeString(tv)
		}
		return nil, fmt.Errorf("the %#v not is base64 string", v)
	case TypeDatetime: //RFC3339
		if tv, ok := v.(string); ok {
			return strToDate(tv)
		}
		return nil, fmt.Errorf("the %#v not is time string", v)

	case TypeFloat:
		if tv, ok := v.(float64); ok {
			return tv, nil
		}
		return nil, fmt.Errorf("the %#v not is float64", v)

	case TypeInt:
		switch tv := v.(type) {
		case string:
			return strconv.ParseInt(tv, 10, 64)
		case int:
			return int64(tv), nil
		case int64:
			return tv, nil
		case int32:
			return int64(tv), nil
		case int16:
			return int64(tv), nil
		case int8:
			return int64(tv), nil
		case uint64:
			return int64(tv), nil
		case uint32:
			return int64(tv), nil
		case uint16:
			return int64(tv), nil
		case uint8:
			return int64(tv), nil
		case float32:
			return int64(tv), nil
		case float64:
			return int64(tv), nil
		default:
			return nil, fmt.Errorf("the %#v (%T) not is int64 string", v, v)
		}
	case TypeString:
		if tv, ok := v.(string); ok {
			return tv, nil
		}
		return nil, fmt.Errorf("the %s json value not is string", v)

	default:
		return nil, newErrInvalidDataType(d)
	}
}
