package coltype

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/linlexing/dbx/schema"
)

// 尝试识别出数据类型，如果实在无法识别的，返回false
func RecognizeColumnType(col *sql.ColumnType) (schema.DataType, bool) {
	atype := col.ScanType()
	if atype == nil {
		return schema.TypeBytea, false
	}
	switch atype.Kind() {
	case reflect.Int, reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return schema.TypeInt, true
	case reflect.Float32,
		reflect.Float64:
		return schema.TypeFloat, true
	case reflect.String:
		return schema.TypeString, true
	default:
		//判断日期型
		if col.ScanType() == reflect.TypeOf(time.Now()) {
			return schema.TypeDatetime, true
		}
	}
	return schema.TypeBytea, false
}
