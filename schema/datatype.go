package schema

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// ErrInvalidDataType 无效的数据类型
type ErrInvalidDataType struct {
	d DataType
}

// 字符串转换成时间，按照各种可能的格式
func StrToDate(s string) (tm time.Time, err error) {
	tm, err = time.ParseInLocation("2006-1-2", s, time.Local)
	if err == nil {
		return
	}
	tm, err = time.ParseInLocation("2006-1-2 15:4:5", s, time.Local)
	if err == nil {
		return
	}
	tm, err = time.ParseInLocation("2006-1-2T15:4:5", s, time.Local)
	if err == nil {
		return
	}
	tm, err = time.ParseInLocation("2006/1/2", s, time.Local)
	if err == nil {
		return
	}
	tm, err = time.ParseInLocation("2006/1/2 15:4:5", s, time.Local)
	if err == nil {
		return
	}
	tm, err = time.ParseInLocation("2006-1-2 15:4:5-07:00", s, time.Local)
	if err == nil {
		return
	}
	tm, err = time.ParseInLocation("20060102", s, time.Local)
	if err == nil {
		return
	}
	tm, err = time.ParseInLocation("2006-1-2 15:4:5.999999999 -0700 MST", s, time.Local)
	if err == nil {
		return
	}
	tm, err = time.ParseInLocation(time.RFC3339, s, time.Local)
	if err == nil {
		return
	}
	tm, err = time.ParseInLocation(time.RFC3339Nano, s, time.Local)
	return
}
func (e *ErrInvalidDataType) Error() string {
	return fmt.Sprintf("invalid data type:%d", e.d)
}
func newErrInvalidDataType(d DataType) *ErrInvalidDataType {
	return &ErrInvalidDataType{
		d: d,
	}
}

// DataType 字段的数据类型
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
const (
	//TypeStringString 字符串
	TypeStringString = "STR"
	//TypeIntString 整型
	TypeIntString = "INT"
	//TypeDateString 日期
	TypeDateString = "DATE"
	//TypeFloatString 浮点
	TypeFloatString = "FLOAT"
	//TypeByteaString 二进制
	TypeByteaString = "BYTEA"
)

// MarshalJSON 实现json的自定义的json序列化，主要是为了兼容前个直接保存字符串值的版本
func (d DataType) MarshalJSON() ([]byte, error) {
	str, err := d.String()
	if err != nil {
		return nil, err
	}
	return json.Marshal(str)
}

// MarshalYAML 是支持yaml序列化
func (d DataType) MarshalYAML() (interface{}, error) {
	return d.String()
}

// UnmarshalYAML 支持yaml反序列化
func (d *DataType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var outstr string
	if err := unmarshal(&outstr); err != nil {
		return err
	}
	val, err := ParseDataType(outstr)
	if err != nil {
		return nil
	}
	*d = val
	return nil
}

// UnmarshalJSON 实现自定义的json反序列化，主要是为了兼容前个版本
func (d *DataType) UnmarshalJSON(v []byte) error {
	var str string
	//v 是一个字符串，带有两端双引号，需要转换
	if err := json.Unmarshal(v, &str); err != nil {
		return err
	}
	val, err := ParseDataType(str)
	if err != nil {
		return err
	}
	*d = val
	return nil
}

// ParseDataType 将一个字符串转换成类型值
func ParseDataType(d string) (DataType, error) {
	switch d {
	case TypeStringString:
		return TypeString, nil
	case TypeIntString:
		return TypeInt, nil
	case TypeDateString:
		return TypeDatetime, nil
	case TypeFloatString:
		return TypeFloat, nil
	case TypeByteaString:
		return TypeBytea, nil
	default:
		return TypeString, fmt.Errorf("invalid type:%#v", d)
	}
}

// String 返回类型的字符串名称
func (d DataType) String() (string, error) {
	switch d {
	case TypeString:
		return TypeStringString, nil
	case TypeInt:
		return TypeIntString, nil
	case TypeDatetime:
		return TypeDateString, nil
	case TypeFloat:
		return TypeFloatString, nil
	case TypeBytea:
		return TypeByteaString, nil
	default:
		return "", newErrInvalidDataType(d)
	}

}

// ChineseString 返回类型的汉字名称
func (d DataType) ChineseString() (string, error) {
	switch d {
	case TypeString:
		return "字符串", nil
	case TypeInt:
		return "整型", nil
	case TypeDatetime:
		return "日期", nil
	case TypeFloat:
		return "浮点", nil
	case TypeBytea:
		return "二进制", nil
	default:
		return "", newErrInvalidDataType(d)
	}
}

// ParseString 将一个字符串转换成标准值
func (d DataType) ParseString(v string) (interface{}, error) {
	if len(v) == 0 {
		return nil, nil
	}
	switch d {
	case TypeString:
		return v, nil
	case TypeInt:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			//补充识别 11.0000的格式，自动去掉后缀0
			arr := strings.Split(v, ".")
			if len(arr) == 2 {
				if len(arr[1]) == 0 || arr[1] == strings.Repeat("0", len(arr[1])) {
					i1, er := strconv.ParseInt(arr[0], 10, 64)
					if er == nil {
						return i1, nil
					}
				}
			}
			return nil, err
		}
		return i, nil

	case TypeDatetime:
		tm, err := StrToDate(v)
		if err != nil {
			return nil, fmt.Errorf("[%s] not time value", v)
		}
		return tm, nil
	case TypeBytea:
		str, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return nil, errors.New("can't convert to byte array,not is base64 string," + v)
		}
		return str, nil
	case TypeFloat:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		return f, nil

	default:
		return nil, newErrInvalidDataType(d)
	}
}

// ToString 转换成字符串形式,值必须要满足要求
func (d DataType) ToString(v interface{}) (result string, err error) {
	//nil代表null，不需要转换，否则会出错
	if v == nil {
		result = ""
		return
	}
	switch d {
	case TypeString:
		switch tv := v.(type) {
		case []byte:
			result = string(tv)
		case string:
			result = tv
		default:
			err = fmt.Errorf("v:%#v(%T) can't to string", v, v)
		}
	case TypeDatetime:
		switch tv := v.(type) {
		case time.Time:
			result = tv.Format("2006-01-02 15:04:05")
		case *time.Time:
			result = tv.Format("2006-01-02 15:04:05")
		default:
			err = fmt.Errorf("v:%#v(%T) can't to date", v, v)
		}
	case TypeInt:
		switch tv := v.(type) {
		case int, int16, int32, int64, int8,
			uint, uint16, uint32, uint64, uint8:
			result = fmt.Sprintf("%d", tv)
		default:
			err = fmt.Errorf("v:%#v(%T) can't to int", v, v)
		}
	case TypeBytea:
		switch tv := v.(type) {
		case []byte:
			result = base64.StdEncoding.EncodeToString(tv)
		default:
			err = fmt.Errorf("v:%#v(%T) can't to bytea", v, v)
		}
	case TypeFloat:
		switch tv := v.(type) {
		case float32:
			result = strconv.FormatFloat(float64(tv), 'f', -1, 64)
		case float64:
			result = strconv.FormatFloat(tv, 'f', -1, 64)
		case int, int16, int32, int64, int8,
			uint, uint16, uint32, uint64, uint8:
			result = fmt.Sprintf("%d", tv)
		default:
			err = fmt.Errorf("v:%#v(%T) can't to float", v, v)
		}
	default:
		err = newErrInvalidDataType(d)
	}
	return
}

// ParseScan 转换数据库驱动扫描出的值，特别是time类型的，很可能是string形式
func (d DataType) ParseScan(v interface{}) (result interface{}, err error) {
	//nil代表null，不需要转换，否则会出错
	if v == nil {
		return nil, nil
	}
	//空字符串当null处理
	if v == "" {
		return nil, nil
	}
	switch d {
	case TypeString:
		switch tv := v.(type) {
		case []byte:
			result = string(tv)
		case string:
			result = tv
		default:
			err = fmt.Errorf("v:%#v can't to string", tv)
		}
	case TypeDatetime:
		switch tv := v.(type) {
		case time.Time, *time.Time:
			result = tv
		case string:
			result, err = StrToDate(tv)
		case []byte:
			result, err = StrToDate(string(tv))
		default:
			err = fmt.Errorf("error type,%T", v)
		}
	case TypeInt:
		switch tv := v.(type) {
		case int8:
			result = int64(tv)
		case int16:
			result = int64(tv)
		case int32:
			result = int64(tv)
		case int:
			result = int64(tv)
		case int64:
			result = tv
		case uint8:
			result = int64(tv)
		case uint16:
			result = int64(tv)
		case uint32:
			result = int64(tv)
		case uint:
			result = int64(tv)
		case uint64:
			result = tv
		case string:
			result, err = strconv.ParseInt(tv, 10, 64)

		case []byte:
			result, err = strconv.ParseInt(string(tv), 10, 64)
		default:
			err = fmt.Errorf("v:%#v not is int,T:%T", tv, tv)
		}
	case TypeBytea:
		switch tv := v.(type) {
		case string:
			result = []byte(tv)
		case []byte:
			result = tv
		default:
			err = fmt.Errorf("v:%#v,T:%T not is bytea", v, v)
		}
	case TypeFloat:
		switch tv := v.(type) {
		case float32:
			result = float64(tv)
		case float64:
			result = tv
		case int64:
			result = float64(tv)
		case int32:
			result = float64(tv)
		case int16:
			result = float64(tv)
		case int:
			result = float64(tv)
		case int8:
			result = float64(tv)
		case uint64:
			result = float64(tv)
		case uint32:
			result = float64(tv)
		case uint16:
			result = float64(tv)
		case uint:
			result = float64(tv)
		case uint8:
			result = float64(tv)
		case string:
			result, err = strconv.ParseFloat(tv, 64)

		case []byte:
			result, err = strconv.ParseFloat(string(tv), 64)
		default:
			err = fmt.Errorf("v:%#v,T:%T not is float", tv, tv)
		}
	default:
		err = newErrInvalidDataType(d)
	}
	return

}

// ToJSON 转换一个字段值，方便其保存json格式，主要目地是检查数据类型和处理二进制数据和日期
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
		return base64.StdEncoding.EncodeToString(b), nil

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

// ParseJSON 从一个json数据中读取值，需要转换[]byte,time等类型
func (d DataType) ParseJSON(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	switch d {
	case TypeBytea: //base64
		if tv, ok := v.(string); ok {
			return base64.StdEncoding.DecodeString(tv)
		}
		return nil, fmt.Errorf("the %#v not is base64 string", v)
	case TypeDatetime: //RFC3339
		if tv, ok := v.(string); ok {
			return StrToDate(tv)
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
