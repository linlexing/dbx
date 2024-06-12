package data

import (
	"database/sql"
	"fmt"
	"strings"
	"text/template"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/render"
	"github.com/linlexing/dbx/schema"
)

// UpdateSet 是update set的一个子句
type UpdateSet struct {
	Column string
	Value  string
}

// Update 是一个批量更新的类
type Update struct {
	Table            *schema.Table
	DataSQL          string
	DataUniqueFields []string
	Sets             []UpdateSet
	AdditionSet      string
	AdditionWhere    string
	BeforeSQL        string
	SQLRenderArgs    interface{}
	SQLRenderFunc    template.FuncMap
}
type NotifyAndKill interface {
	AddProgress(string, interface{})
	AddProgressTx(common.DB, string, interface{})
	IsKill() bool
}

func UpdateLogNewColumnName(col string) string {
	return col + "_2"
}
func UpdateLogOldColumnName(col string) string {
	return col + "_1"
}

// Exec 执行一个更新操作，并返回影响的行数,记录日志到临时表中
// 临时表结构为：主键，更新字段1，更新字段1_1,...,其中_1记录旧值,_2记录新值，
// 不用字母是因为有大小写问题
func (u *Update) ExecLog(nt NotifyAndKill, db common.DB, tmpTableName string, param ...interface{}) (icount int64, err error) {
	if err = u.checkSets(); err != nil {
		return
	}
	tmpTable := schema.NewTable(tmpTableName)
	//先增加主键
	// mapPk := map[string]struct{}{}
	for _, one := range u.Table.PrimaryKeys {
		tmpTable.Columns = append(tmpTable.Columns, u.Table.ColumnByName(one).Clone())
		// mapPk[one] = struct{}{}
	}
	//再增加每个更新的字段
	for _, one := range u.Sets {
		col := u.Table.ColumnByName(one.Column).Clone()
		col.Name = UpdateLogNewColumnName(col.Name)
		// //如果是主键，则不要增加重复了
		// if _, ok := mapPk[one.Column]; !ok {
		tmpTable.Columns = append(tmpTable.Columns, col)
		// }
		colOld := u.Table.ColumnByName(one.Column).Clone()
		colOld.Name = UpdateLogOldColumnName(colOld.Name)
		tmpTable.Columns = append(tmpTable.Columns, colOld)
	}
	tmpTable.PrimaryKeys = u.Table.PrimaryKeys
	nt.AddProgress(fmt.Sprintf("创建临时表[%s]...", tmpTableName), nil)
	//创建临时表
	if err = tmpTable.Update(db.DriverName(), db); err != nil {
		return
	}
	insertColNames := make([]string, len(u.Table.PrimaryKeys))
	selectColNames := make([]string, len(u.Table.PrimaryKeys))
	for i, one := range u.Table.PrimaryKeys {
		insertColNames[i] = one
		selectColNames[i] = one
	}
	dataAlias := "datasrc_"
	sets := []ColMap{}
	for _, one := range u.Sets {
		v := one.Value
		//取出表达式
		if v[0] == '`' && v[len(v)-1] == '`' {
			v = v[1 : len(v)-1]
			if v == "" {
				v = "NULL"
			}
		} else {
			//各种数据库都有类型转换，如果出错，改用类型判断后转换
			v = "'" + strings.ReplaceAll(v, "'", "''") + "'"
		}
		insertColNames = append(insertColNames, UpdateLogNewColumnName(one.Column))
		selectColNames = append(selectColNames, v)
		insertColNames = append(insertColNames, UpdateLogOldColumnName(one.Column))
		selectColNames = append(selectColNames, one.Column)
		sets = append(sets, ColMap{Dest: one.Column, Src: UpdateLogNewColumnName(one.Column)})
	}
	where := []string{}
	if len(u.AdditionWhere) > 0 {
		where = append(where, "("+u.AdditionWhere+")")
	}
	for i, field := range u.Table.PrimaryKeys {
		where = append(where, fmt.Sprintf("%s.%s = %s.%s", u.Table.FullName(), field, dataAlias, u.DataUniqueFields[i]))
	}
	//数据插入临时表
	strSQL := fmt.Sprintf("insert into %s(%s)select %s from %s where exists(select 1 from (%s) %s where %s)",
		tmpTableName, strings.Join(insertColNames, ","), strings.Join(selectColNames, ","), u.Table.FullName(),
		u.DataSQL, dataAlias, strings.Join(where, " and "))

	sr, err := db.Exec(strSQL)
	if err != nil {
		return
	}
	ic, err := sr.RowsAffected()
	if err != nil {
		return
	}
	nt.AddProgress(fmt.Sprintf("%d 条数据插入临时表[%s]", ic, tmpTableName), nil)
	pksList := make([]ColMap, len(u.Table.PrimaryKeys))
	for i, v := range u.Table.PrimaryKeys {
		pksList[i] = ColMap{Dest: v, Src: v}
	}
	strSQL = Find(db.DriverName()).UpdateFrom(u.Table.FullName(),
		"select * from "+tmpTableName, u.AdditionSet, pksList, sets)

	strSQL, err = render.RenderSQL(strSQL, u.SQLRenderArgs, u.SQLRenderFunc)
	if err != nil {
		return
	}
	if len(u.BeforeSQL) > 0 {
		var strBefore string
		funcmap := template.FuncMap{}
		for k, v := range u.SQLRenderFunc {
			funcmap[k] = v
		}
		funcmap["DataSQL"] = func() string {
			return u.DataSQL
		}
		funcmap["tmpTableName"] = func() string {
			return tmpTableName
		}
		strBefore, err = render.RenderSQL(u.BeforeSQL, u.SQLRenderArgs, funcmap)
		if err != nil {
			return
		}
		nt.AddProgress("执行Before脚本...", nil)
		var br sql.Result
		br, err = db.Exec(strBefore, param...)
		if err != nil {
			return
		}
		var brCount int64
		brCount, err = br.RowsAffected()
		if err != nil {
			return
		}
		nt.AddProgress(fmt.Sprintf("%d 条记录受执行Before脚本影响", brCount), nil)
	}
	nt.AddProgress(fmt.Sprintf("从临时表[%s]更新到[%s]中...", tmpTableName, u.Table.FullName()), nil)
	sr, err = db.Exec(strSQL, param...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, param...)
		return
	}
	icount, err = sr.RowsAffected()
	if err != nil {
		return
	}
	nt.AddProgress(fmt.Sprintf("更新完成，共涉及 %d 条记录", icount), nil)
	return
}
func (u *Update) checkSets() (err error) {

	for _, set := range u.Sets {
		k := set.Column
		v := set.Value
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
	}
	return nil
}

// Exec 执行一个更新操作，并返回影响的行数
func (u *Update) Exec(db common.DB, param ...interface{}) (icount int64, err error) {
	if len(u.Table.PrimaryKeys) == 0 {
		err = fmt.Errorf("table %s not have primary key", u.Table.FullName())
		return
	}
	if len(u.Table.PrimaryKeys) != len(u.DataUniqueFields) {
		err = fmt.Errorf("table %s primary key%v <> %v", u.Table.FullName(), u.Table.PrimaryKeys, u.DataUniqueFields)
		return
	}
	if err = u.checkSets(); err != nil {
		return
	}
	sets := []string{}
	setVals := []interface{}{}

	for _, set := range u.Sets {
		k := set.Column
		v := set.Value
		//取出表达式
		if v[0] == '`' && v[len(v)-1] == '`' {
			v = v[1 : len(v)-1]
			if v == "" {
				v = "NULL"
			}
			sets = append(sets, fmt.Sprintf("%s=%s", k, v))
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
	strSQL := fmt.Sprintf("update %s set %s where exists(select 1 from (%s) %s where %s)",
		u.Table.FullName(),
		strings.Join(sets, ","),
		u.DataSQL,
		dataAlias,
		strings.Join(where, " and "),
	)

	strSQL, err = render.RenderSQL(strSQL, u.SQLRenderArgs, u.SQLRenderFunc)
	if err != nil {
		return
	}
	if len(u.BeforeSQL) > 0 {
		var strBefore string
		funcmap := template.FuncMap{}
		for k, v := range u.SQLRenderFunc {
			funcmap[k] = v
		}
		funcmap["DataSQL"] = func() string {
			return u.DataSQL
		}
		strBefore, err = render.RenderSQL(u.BeforeSQL, u.SQLRenderArgs, funcmap)
		if err != nil {
			return
		}
		if _, err = db.Exec(strBefore, param...); err != nil {
			return
		}
	}
	sr, err := db.Exec(strSQL, append(setVals, param...)...)
	if err != nil {
		err = common.NewSQLError(err, strSQL, append(setVals, param...)...)
		return
	}
	icount, err = sr.RowsAffected()

	return
}
func init() {
	render.AddFunc(template.FuncMap{
		"concat_sql": func(driver string, val ...string) string {
			return Find(driver).Concat(val...)
		}})
}
