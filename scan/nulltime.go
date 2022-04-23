package scan

import (
	"database/sql/driver"
	"time"

	"github.com/linlexing/dbx/schema"
)

// NullTime represents a time.Time that may be null. NullTime implements the
// sql.Scanner interface so it can be used as a scan destination, similar to
// sql.NullString.
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	//字符串的，尝试各种解析
	if !nt.Valid && value != nil {
		str, ok := value.(string)
		if ok {
			tm, err := schema.StrToDate(str)
			if err != nil {
				return err
			}
			nt.Time = tm
			nt.Valid = true
			return nil
		}
	}
	//UTC时区的，需要转换成当前时区
	// println(nt.Time.Location().String())
	if nt.Time.Location().String() == "UTC" {
		nt.Time = time.Date(nt.Time.Year(), nt.Time.Month(), nt.Time.Day(),
			nt.Time.Hour(), nt.Time.Minute(), nt.Time.Second(),
			nt.Time.Nanosecond(), time.Local)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
