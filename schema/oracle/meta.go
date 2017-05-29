package oracle

import (
	"dbweb/lib/safe"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/schema"
)

const driverName = "oci8"

type meta struct {
}

func (m *meta) IsNull() string {
	return "nvl"
}

//CreateTableAs 执行create table as select语句
func (m *meta) CreateTableAs(db common.DB, tableName, strSQL string, pks []string) error {
	s := fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSQL)
	if _, err := db.Exec(s); err != nil {
		log.Println(s)
		return common.NewSQLError(err, s)
	}
	s = fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ","))
	if _, err := db.Exec(s); err != nil {
		return common.NewSQLError(err, s)
	}
	return nil
}

func (m *meta) TableNames(db common.DB) (names []string, err error) {
	strSQL := "SELECT table_name FROM user_tables"
	names = []string{}
	rows, err := db.Query(strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			log.Println(err)
			return nil, err
		}
	}

	sort.Strings(names)
	return
}
func (m *meta) TableExists(db common.DB, tabName string) (bool, error) {

	return tableExists(db, tabName)
}

func init() {
	schema.Register(driverName, new(meta))
}

func (m *meta) ValueExpress(db common.DB, dataType schema.DataType, value string) string {
	switch dataType {
	case schema.TypeFloat, schema.TypeInt:
		return value
	case schema.TypeString:
		return safe.SignString(value)
	case schema.TypeDatetime:
		if len(value) == 10 {
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd')", value)
		} else if len(value) == 19 {
			return fmt.Sprintf("TO_DATE('%s','yyyy-mm-dd hh24:mi:ss')", value)
		} else {
			panic(fmt.Errorf("invalid datetime:%s", value))
		}
	default:
		panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
	}
}
func truncateTimeZone(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), 0, time.UTC)
}

//Prepare拦截保存前，oracle，需要去除时间中的时区，以免触发ORA-01878错误
func (m *meta) Prepare(dataType schema.DataType, v interface{}) interface{} {
	if dataType == schema.TypeDatetime {
		switch tv := v.(type) {
		case time.Time:
			return truncateTimeZone(tv)
		case *time.Time:
			return truncateTimeZone(*tv)
		default:
			log.Panic("not is time")
		}
	}
	return v
}
func (m *meta) CreateTable(db common.DB, tab *schema.Table) error {
	cols := []string{}
	for _, v := range tab.Columns {
		cols = append(cols, dbDefine(v))
	}
	var strSQL string
	if len(tab.PrimaryKeys) > 0 {
		strSQL = fmt.Sprintf(
			"CREATE TABLE %s(\n%s,\nCONSTRAINT %s_pkey PRIMARY KEY(%s)\n)",
			tab.FullName(), strings.Join(cols, ",\n"), tab.Name, strings.Join(tab.PrimaryKeys, ","))
	} else {
		strSQL = fmt.Sprintf(
			"CREATE TABLE %s(\n%s\n)",
			tab.FullName(), strings.Join(cols, ",\n"))
	}
	if _, err := db.Exec(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return err
	}
	log.Println(strSQL)
	//最后处理索引
	for _, col := range tab.Columns {
		if col.Index {
			if err := createColumnIndex(db, tab.FullName(), col.Name); err != nil {
				return err
			}
		}
	}
	return nil
}
