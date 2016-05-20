package dbx

import (
	"fmt"
	"strings"
)

type Update struct {
	Table            *DBTable
	DataSql          string
	DataUniqueFields []string
	Sets             map[string]string
	AdditionSet      string
	SqlRenderArgs    interface{}
}

func (u *Update) Exec() (icount int64, err error) {
	if len(u.Table.PrimaryKeys()) == 0 {
		err = fmt.Errorf("table %s not primary key", u.Table.Name)
		return
	}
	if len(u.Table.PrimaryKeys()) != len(u.DataUniqueFields) {
		err = fmt.Errorf("table %s primary key%v <> %v", u.Table.Name, u.Table.PrimaryKeys(), u.DataUniqueFields)
		return
	}
	params := map[string]interface{}{}

	sets := []string{}
	idx := 0
	for k, v := range u.Sets {
		val := v
		if len(v) == 0 {
			err = fmt.Errorf("%s the value is empty", k)
			return
		}
		if v == "`" {
			err = fmt.Errorf("%s the value is %s", k, v)
			return
		}
		if v[0] == '`' && v[len(v)-1] != '`' {
			err = fmt.Errorf("%s the value is %s", k, v)
			return
		}
		//取出表达式
		if v[0] == '`' && v[len(v)-1] == '`' {
			val = v[1 : len(v)-1]
			if val == "" {
				val = "NULL"
			}
			sets = append(sets, fmt.Sprintf("%s=%s", k, val))
		} else {
			idx++
			//各个数据库均有自动类型转换机制，字符串值会自动转换
			//如果有出错，则在这里增加转换
			pname := fmt.Sprintf("up_%d", idx)
			params[pname] = v
			sets = append(sets, fmt.Sprintf("%s=:%s", k, pname))
		}
	}
	if len(u.AdditionSet) > 0 {
		sets = append(sets, u.AdditionSet)
	}
	where := []string{}
	dataAlias := "datasql_"
	for i, field := range u.Table.PrimaryKeys() {
		where = append(where, fmt.Sprintf("%s.%s = %s.%s", u.Table.Name, field, dataAlias, u.DataUniqueFields[i]))
	}
	strSql := fmt.Sprintf("update %s set %s where exists(select * from (%s) %s where %s)",
		u.Table.Name,
		strings.Join(sets, ","),
		u.DataSql,
		dataAlias,
		strings.Join(where, " and "),
	)
	str := RenderSql(strSql, u.SqlRenderArgs)

	if sr, e := u.Table.Db.NamedExec(str, params); e != nil {
		err = SqlError{str, params, e}
		return
	} else {
		icount, err = sr.RowsAffected()
	}
	return
}
