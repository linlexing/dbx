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

type DriverType int

const (
	Oracle = iota
	Postgres
	Mysql
	Sqlite
	SqlServer
	DBWeb
	Hive
	Unknown
)

// 废弃：用ParseDriverType代替
func IsPostgres(driver string) bool {
	return ParseDriverType(driver) == Postgres
}
func ParseDriverType(driver string) DriverType {
	switch driver {
	case "postgres", "pgx-opengauss", "opengauss", "pgx", "pq-timeouts", "cloudsqlpostgres", "pgx-un":
		return Postgres
	case "oci8", "dmdb", "oracle", "godror":
		return Oracle
	case "mysql":
		return Mysql
	case "sqlserver":
		return SqlServer
	case "dbweb":
		return DBWeb
	case "sqlite":
		return Sqlite
	case "hive":
		return Hive
	default:
		if strings.HasPrefix(driver, "sqlite3") {
			return Sqlite
		} else {
			return Unknown
		}
	}
}

// BindType returns the bindtype for a given database given a drivername.
func BindType(driverName string) int {
	switch ParseDriverType(driverName) {
	case Postgres:
		return DOLLAR
	case Mysql, Sqlite:
		return QUESTION
	case Oracle:
		return NAMED
	case SqlServer:
		return AT
	}
	return UNKNOWN
}

// FIXME: this should be able to be tolerant of escaped ?'s in queries without
// losing much speed, and should be to avoid confusion.

// Rebind a query from the default bindtype (QUESTION) to the target bindtype.
// func Rebind(bindType int, query string) string {
// 	switch bindType {
// 	case QUESTION, UNKNOWN:
// 		return query
// 	}

// 	// Add space enough for 10 params before we have to allocate
// 	rqb := make([]byte, 0, len(query)+10)

// 	var i, j int
// 	var leftSign = 0
// 	for i = strings.Index(query, "?"); i != -1; i = strings.Index(query, "?") {
// 		//检查单引号个数，只有偶数个才满足条件,奇数个说明是在字符串中，需要跳过
// 		if leftSign += strings.Count(query[:i], "'"); leftSign%2 == 1 {
// 			//问号要添加到rqb
// 			rqb = append(rqb, query[:i+1]...)
// 			query = query[i+1:]
// 			continue
// 		}
// 		rqb = append(rqb, query[:i]...)

// 		switch bindType {
// 		case DOLLAR:
// 			rqb = append(rqb, '$')
// 		case NAMED:
// 			rqb = append(rqb, ':', 'a', 'r', 'g')
// 		case AT:
// 			rqb = append(rqb, '@', 'p')
// 		}

// 		j++
// 		rqb = strconv.AppendInt(rqb, int64(j), 10)
// 		query = query[i+1:]

// 	}

//		return string(append(rqb, query...))
//	}
func Rebind(bindType int, query string) string {
	switch bindType {
	case QUESTION, UNKNOWN:
		return query
	}
	sb := new(strings.Builder)
	// Add space enough for 10 params before we have to allocate
	// rqb := make([]byte, 0, len(query)+10)

	var i, j int
	var leftSign = 0
	for i = strings.Index(query, "?"); i != -1; i = strings.Index(query, "?") {
		//检查单引号个数，只有偶数个才满足条件,奇数个说明是在字符串中，需要跳过
		if leftSign += strings.Count(query[:i], "'"); leftSign%2 == 1 {
			//问号要添加到rqb
			if _, err := sb.WriteString(query[:i+1]); err != nil {
				panic(err)
			}
			// rqb = append(rqb, query[:i+1]...)
			query = query[i+1:]
			continue
		}
		if _, err := sb.WriteString(query[:i]); err != nil {
			panic(err)
		}
		// rqb = append(rqb, query[:i]...)

		switch bindType {
		case DOLLAR:
			if _, err := sb.WriteRune('$'); err != nil {
				panic(err)
			}
			// rqb = append(rqb, '$')
		case NAMED:
			if _, err := sb.WriteString(":arg"); err != nil {
				panic(err)
			}

			// rqb = append(rqb, ':', 'a', 'r', 'g')
		case AT:
			if _, err := sb.WriteString(":@p"); err != nil {
				panic(err)
			}
			// rqb = append(rqb, '@', 'p')
		}

		j++
		// rqb = strconv.AppendInt(rqb, int64(j), 10)
		if _, err := sb.WriteString(strconv.Itoa(j)); err != nil {
			panic(err)
		}

		query = query[i+1:]

	}
	if _, err := sb.WriteString(query); err != nil {
		panic(err)
	}
	return sb.String()
}
