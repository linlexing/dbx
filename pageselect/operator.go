package pageselect

import "errors"
import "encoding/json"

//Operator 表示条件中的运算符
type Operator int

const (
	//OperatorEqu =
	OperatorEqu Operator = iota
	//OperatorNotEqu <>
	OperatorNotEqu
	//OperatorGreaterThan >
	OperatorGreaterThan
	//OperatorGreaterThanOrEqu >=
	OperatorGreaterThanOrEqu
	//OperatorLessThan <
	OperatorLessThan
	//OperatorLessThanOrEqu <=
	OperatorLessThanOrEqu
	//OperatorLike 包含
	OperatorLike
	//OperatorNotLike 不包含
	OperatorNotLike
	//OperatorPrefix 前缀
	OperatorPrefix
	//OperatorNotPrefix 非前缀
	OperatorNotPrefix
	//OperatorSuffix 后缀
	OperatorSuffix
	//OperatorNotSuffix 非后缀
	OperatorNotSuffix
	//OperatorIn 在列表
	OperatorIn
	//OperatorNotIn 不在列表
	OperatorNotIn
	//OperatorRegexp 正则
	OperatorRegexp
	//OperatorNotRegexp 非正则
	OperatorNotRegexp
	//OperatorIsNull 为空
	OperatorIsNull
	//OperatorIsNotNull is not null
	OperatorIsNotNull
	//OperatorLengthEqu 长度等于
	OperatorLengthEqu
	//OperatorLengthNotEqu 长度不等于
	OperatorLengthNotEqu
	//OperatorLengthGreaterThan 长度大于
	OperatorLengthGreaterThan
	//OperatorLengthGreaterThanOrEqu 长度 >=
	OperatorLengthGreaterThanOrEqu
	//OperatorLengthLessThan 长度 <
	OperatorLengthLessThan
	//OperatorLengthLessThanOrEqu 长度<=
	OperatorLengthLessThanOrEqu
)

var (
	//ErrInvalidOperator 表明一个不存在的运算符
	ErrInvalidOperator = errors.New("invalid operate")
)

//MarshalJSON 实现json的自定义的json序列化，主要是为了兼容前个直接保存字符串值的版本
func (o *Operator) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}

//UnmarshalJSON 实现自定义的json反序列化，主要是为了兼容前个版本
func (o *Operator) UnmarshalJSON(v []byte) error {
	outstr := ""
	if err := json.Unmarshal(v, &outstr); err != nil {
		return err
	}
	opt, err := ParseString(outstr)
	if err != nil {
		return err
	}
	*o = opt
	return err
}

//ParseString 将一个字符串转换成Operator值
func ParseString(str string) (Operator, error) {
	switch str {
	case "=":
		return OperatorEqu, nil
	case "!=":
		return OperatorNotEqu, nil
	case ">":
		return OperatorGreaterThan, nil
	case ">=":
		return OperatorGreaterThanOrEqu, nil
	case "<":
		return OperatorLessThan, nil
	case "<=":
		return OperatorLessThanOrEqu, nil
	case "?":
		return OperatorLike, nil
	case "!?":
		return OperatorNotLike, nil
	case "?>":
		return OperatorPrefix, nil
	case "!?>":
		return OperatorNotPrefix, nil
	case "<?":
		return OperatorSuffix, nil
	case "!<?":
		return OperatorNotSuffix, nil
	case "in":
		return OperatorIn, nil
	case "!in":
		return OperatorNotIn, nil
	case "~":
		return OperatorRegexp, nil
	case "!~":
		return OperatorNotRegexp, nil
	case "e":
		return OperatorIsNull, nil
	case "!e":
		return OperatorIsNotNull, nil
	case "_":
		return OperatorLengthEqu, nil
	case "!_":
		return OperatorLengthNotEqu, nil
	case "_>":
		return OperatorLengthGreaterThan, nil
	case "_>=":
		return OperatorLengthGreaterThanOrEqu, nil
	case "_<":
		return OperatorLengthLessThan, nil
	case "_<=":
		return OperatorLengthLessThanOrEqu, nil
	default:
		return 0, ErrInvalidOperator
	}

}

//String 返回字符串形式
func (o Operator) String() string {
	switch o {
	case OperatorEqu:
		return "="
	case OperatorNotEqu:
		return "!="
	case OperatorGreaterThan:
		return ">"
	case OperatorGreaterThanOrEqu:
		return ">="
	case OperatorLessThan:
		return "<"
	case OperatorLessThanOrEqu:
		return "<="
	case OperatorLike:
		return "?"
	case OperatorNotLike:
		return "!?"
	case OperatorPrefix:
		return "?>"
	case OperatorNotPrefix:
		return "!?>"
	case OperatorSuffix:
		return "<?"
	case OperatorNotSuffix:
		return "!<?"
	case OperatorIn:
		return "in"
	case OperatorNotIn:
		return "!in"
	case OperatorRegexp:
		return "~"
	case OperatorNotRegexp:
		return "!~"
	case OperatorIsNull:
		return "e"
	case OperatorIsNotNull:
		return "!e"
	case OperatorLengthEqu:
		return "_"
	case OperatorLengthNotEqu:
		return "!_"
	case OperatorLengthGreaterThan:
		return "_>"
	case OperatorLengthGreaterThanOrEqu:
		return "_>="
	case OperatorLengthLessThan:
		return "_<"
	case OperatorLengthLessThanOrEqu:
		return "_<="
	default:
		panic(ErrInvalidOperator)
	}
}

//ChineseString 返回中文名称
func (o Operator) ChineseString() string {
	switch o {
	case OperatorEqu:
		return "等于"
	case OperatorNotEqu:
		return "不等于"
	case OperatorGreaterThan:
		return "大于"
	case OperatorGreaterThanOrEqu:
		return "大于等于"
	case OperatorLessThan:
		return "小于"
	case OperatorLessThanOrEqu:
		return "小于等于"
	case OperatorLike:
		return "包含"
	case OperatorNotLike:
		return "不包含"
	case OperatorPrefix:
		return "前缀"
	case OperatorNotPrefix:
		return "非前缀"
	case OperatorSuffix:
		return "后缀"
	case OperatorNotSuffix:
		return "非后缀"
	case OperatorIn:
		return "在列表"
	case OperatorNotIn:
		return "不在列表"
	case OperatorRegexp:
		return "正则"
	case OperatorNotRegexp:
		return "非正则"
	case OperatorIsNull:
		return "为空"
	case OperatorIsNotNull:
		return "非空"
	case OperatorLengthEqu:
		return "长度等于"
	case OperatorLengthNotEqu:
		return "长度不等于"
	case OperatorLengthGreaterThan:
		return "长度大于"
	case OperatorLengthGreaterThanOrEqu:
		return "长度大于等于"
	case OperatorLengthLessThan:
		return "长度小于"
	case OperatorLengthLessThanOrEqu:
		return "长度小于等于"
	default:
		panic(ErrInvalidOperator)
	}
}
