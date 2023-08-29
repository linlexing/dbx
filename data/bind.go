package data

import (
	"strconv"
	"strings"
)

const (
	UNKNOWN = iota
	QUESTION
	DOLLAR
	NAMED
	AT
)

// 是否是postgresql数据库
func IsPostgres(driver string) bool {
	switch driver {
	case "postgres", "pgx-opengauss", "opengauss", "pgx", "pq-timeouts", "cloudsqlpostgres":
		return true
	}
	return false
}

// BindType returns the bindtype for a given database given a drivername.
func BindType(driverName string) int {
	if IsPostgres(driverName) {
		return DOLLAR
	}
	switch driverName {
	case "mysql":
		return QUESTION
	case "sqlite3":
		return QUESTION
	case "oci8", "ora", "goracle", "oracle":
		return NAMED
	case "sqlserver":
		return AT
	}
	return UNKNOWN
}

// FIXME: this should be able to be tolerant of escaped ?'s in queries without
// losing much speed, and should be to avoid confusion.

// Rebind a query from the default bindtype (QUESTION) to the target bindtype.
func Rebind(bindType int, query string) string {
	switch bindType {
	case QUESTION, UNKNOWN:
		return query
	}

	// Add space enough for 10 params before we have to allocate
	rqb := make([]byte, 0, len(query)+10)

	var i, j int
	var leftSign = 0
	for i = strings.Index(query, "?"); i != -1; i = strings.Index(query, "?") {
		//检查单引号个数，只有偶数个才满足条件,奇数个说明是在字符串中，需要跳过
		if leftSign += strings.Count(query[:i], "'"); leftSign%2 == 1 {
			//问号要添加到rqb
			rqb = append(rqb, query[:i+1]...)
			query = query[i+1:]
			continue
		}
		rqb = append(rqb, query[:i]...)

		switch bindType {
		case DOLLAR:
			rqb = append(rqb, '$')
		case NAMED:
			rqb = append(rqb, ':', 'a', 'r', 'g')
		case AT:
			rqb = append(rqb, '@', 'p')
		}

		j++
		rqb = strconv.AppendInt(rqb, int64(j), 10)
		query = query[i+1:]

	}

	return string(append(rqb, query...))
}
