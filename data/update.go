package data

import (
	"fmt"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/render"
	"github.com/linlexing/dbx/schema"
)

//Update 是一个批量更新的类
type Update struct {
	Table            *schema.Table
	DataSQL          string
	DataUniqueFields []string
	Sets             map[string]string
	AdditionSet      string
	AdditionWhere    string
	SQLRenderArgs    interface{}
}

//Exec 执行一个更新操作，并返回影响的行数
func (u *Update) Exec(db common.DB) (icount int64, err error) {
	if len(u.Table.PrimaryKeys) == 0 {
		err = fmt.Errorf("table %s not have primary key", u.Table.FullName())
		return
	}
	if len(u.Table.PrimaryKeys) != len(u.DataUniqueFields) {
		err = fmt.Errorf("table %s primary key%v <> %v", u.Table.FullName(), u.Table.PrimaryKeys, u.DataUniqueFields)
		return
	}

	sets := []string{}
	setVals := []interface{}{}

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

			//各个数据库均有自动类型转换机制，字符串值会自动转换
			//如果有出错，则在这里增加转换
			sets = append(sets, fmt.Sprintf("%s=?", k))
			setVals = append(setVals, v)
		}
	}
	if len(u.AdditionSet) > 0 {
		sets = append(sets, u.AdditionSet)
	}
	where := []string{}
	if len(u.AdditionWhere) > 0 {
		where = append(where, u.AdditionWhere)
	}
	dataAlias := "datasql_"
	for i, field := range u.Table.PrimaryKeys {
		where = append(where, fmt.Sprintf("%s.%s = %s.%s", u.Table.FullName(), field, dataAlias, u.DataUniqueFields[i]))
	}
	strSQL := fmt.Sprintf("update %s set %s where exists(select * from (%s) %s where %s)",
		u.Table.FullName(),
		strings.Join(sets, ","),
		u.DataSQL,
		dataAlias,
		strings.Join(where, " and "),
	)

	strSQL, err = render.RenderSQL(strSQL, u.SQLRenderArgs)
	if err != nil {
		return
	}
	sr, err := db.Exec(strSQL, setVals...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, setVals...)
		return
	}
	icount, err = sr.RowsAffected()

	return
}
