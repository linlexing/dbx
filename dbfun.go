package dbx

import (
	"bytes"
	"database/sql"
	"dbweb/lib/safe"
	"dbweb/lib/tempext"
	"encoding/binary"
	"fmt"

	"math/rand"
	"sort"
	"strings"
	"text/template"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/linlexing/mapfun"

	"github.com/jmoiron/sqlx"
)

//SQLError 表示一个sql语句执行出错
type SQLError struct {
	SQL    string
	Params interface{}
	Err    error
}

//NewSQLError 构造
func NewSQLError(sql string, params interface{}, err error) SQLError {
	return SQLError{
		SQL:    sql,
		Params: params,
		Err:    err,
	}
}
func (e SQLError) Error() string {
	l := 0
	content := fmt.Sprintf("%#v", e.Params)
	switch tv := e.Params.(type) {
	case []interface{}:
		l = len(tv)
	case map[string]interface{}:
		l = len(tv)
		list := []string{}
		for k, v := range tv {
			switch subtv := v.(type) {
			case time.Time:
				list = append(list, fmt.Sprintf("%s(time):%s", k, subtv.Format(time.RFC3339)))
			default:
				list = append(list, fmt.Sprintf("%s(%T):%#v", k, v, v))
			}

		}
		content = strings.Join(list, "\n")
	}
	return fmt.Sprintf("%s\n%s\nparams len is %d,content is:\n%s", e.Err, e.SQL, l, content)
}

//DB 接口表示一个数据库操作，主要用来统一DBConnection 和 DBTrans的方法
//使得函数可以接收两种参数传入，依赖于sqlx包
type DB interface {
	Select(dest interface{}, query string, args ...interface{}) error
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	DriverName() string
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(string) string
	BindNamed(string, interface{}) (string, []interface{}, error)
	Exec(string, ...interface{}) (sql.Result, error)
	MustExec(string, ...interface{}) sql.Result
	Get(dest interface{}, query string, args ...interface{}) error
}

//Columns 返回一个SQL语句执行后应该返回的列名清单
func Columns(db DB, strSQL string, p map[string]interface{}) ([]string, error) {
	str, pam := BindSQL(db, strSQL, p)
	rows, err := db.Queryx(str, pam...)

	if err != nil {
		return nil, NewSQLError(strSQL, pam, err)
	}
	defer rows.Close()
	r, err := rows.Columns()
	if err == nil {
		for i, v := range r {
			r[i] = strings.ToUpper(v)
		}
	}

	return r, err
}

//TableNames 返回一个数据库所有的表名
func TableNames(db DB) (names []string, err error) {
	return Meta(db).TableNames(db)
}

//NameGet 类似Get，不过可以用命名参数
func NameGet(db DB, d interface{}, strSQL string, p map[string]interface{}) error {
	str, pam := BindSQL(db, strSQL, p)
	if err := db.Get(d, str, pam...); err != nil {
		return NewSQLError(strSQL, p, err)
	}
	return nil
}

//NameSelect 类似Select
func NameSelect(db DB, d interface{}, strSQL string, p map[string]interface{}) error {
	str, pam := BindSQL(db, strSQL, p)
	if err := db.Select(d, str, pam...); err != nil {
		return NewSQLError(strSQL, p, err)
	}
	return nil
}

//QueryRecord 返回一个结果集，并返回字段名称列表（转换为大写）
func QueryRecord(db DB, strSQL string, p map[string]interface{}) (result []map[string]interface{},
	cols []string, err error) {
	var rows *sqlx.Rows
	str, pam := BindSQL(db, strSQL, p)
	if rows, err = db.Queryx(str, pam...); err != nil {
		log.Println(str)
		err = NewSQLError(strSQL, p, err)

		return
	}
	result = []map[string]interface{}{}
	defer rows.Close()
	cols, err = rows.Columns()
	if err == nil {
		for i, v := range cols {
			cols[i] = strings.ToUpper(v)
		}
	}
	for rows.Next() {
		oneRecord := map[string]interface{}{}
		if err = rows.MapScan(oneRecord); err != nil {
			err = NewSQLError(strSQL, p, err)
			return
		}

		result = append(result, mapfun.UpperKeys(oneRecord))
	}
	return
}

//IsNull 返回数据库isnull函数的名称，因为不同的数据库使用不同的名称，而这个函数又非常常用
func IsNull(db DB) string {
	return Meta(db).IsNull()
}

//Exists 返回sql语句有没有返回值，性能比较快
func Exists(db DB, strSQL string, p map[string]interface{}) (result bool, err error) {
	str, pam := BindSQL(db, strSQL, p)

	var rows *sqlx.Rows
	if rows, err = db.Queryx(str, pam...); err != nil {
		err = NewSQLError(strSQL, p, err)
		return
	}
	defer rows.Close()
	result = rows.Next()
	return
}

//CreateTableAs 执行create table as select语句
func CreateTableAs(db DB, tableName, strSQL string, pks []string) error {
	return Meta(db).CreateTableAs(db, tableName, strSQL, pks)
}

//RemoveColumns 删除表字段
func RemoveColumns(db DB, tabName string, cols []string) error {
	return Meta(db).RemoveColumns(db, tabName, cols)
}

//GetSlice 返回一个字符串数组
func GetSlice(db DB, strSQL string, params map[string]interface{}) ([]string, error) {
	s, p := BindSQL(db, strSQL, params)

	rev := []sql.NullString{}
	if err := db.Select(&rev, s, p...); err != nil {
		println(err)
		return nil, NewSQLError(strSQL, p, err)
	}
	strList := []string{}
	for _, one := range rev {
		if one.Valid {
			strList = append(strList, one.String)
		} else {
			strList = append(strList, "")
		}
	}
	return strList, nil
}

//MustGetSlice 返回一个字符串数组，如果有错误则发生异常
func MustGetSlice(db DB, strSQL string, params map[string]interface{}) []string {
	rev, err := GetSlice(db, strSQL, params)
	if err != nil {
		log.Panic(err)
	}
	return rev
}

//MustGetSliceAndSort 返回一个字符串数组，其值自动排序,如果有错误则发生异常
func MustGetSliceAndSort(db DB, strSQL string, params map[string]interface{}) []string {
	rev, err := GetSlice(db, strSQL, params)
	if err != nil {
		log.Panic(err)
	}
	sort.Strings(rev)
	return rev
}

//TableRename 表更名
func TableRename(db DB, oldName, newName string) error {
	return Meta(db).TableRename(db, oldName, newName)
}

//TableExists 返回一个表是否存在
func TableExists(db DB, tableName string) (bool, error) {
	return Meta(db).TableExists(db, tableName)
}

//GetSQLFun 返回第一列第一行的值，如果没有结果，返回nil，不出错
func GetSQLFun(db DB, strSQL string, p map[string]interface{}) (result interface{}, err error) {
	str, pam := BindSQL(db, strSQL, p)

	var rows *sqlx.Rows
	if rows, err = db.Queryx(str, pam...); err != nil {
		err = NewSQLError(strSQL, p, err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&result); err != nil {
			err = NewSQLError(strSQL, p, err)
			return
		}
	}
	return
}

//MustGetSQLFun 类似GetSQLFun，不过出错会抛出异常
func MustGetSQLFun(db DB, strSQL string, p map[string]interface{}) (result interface{}) {
	var err error
	if result, err = GetSQLFun(db, strSQL, p); err != nil {
		log.Printf("sql err:%s\n%s\n", err, strSQL)
		log.Panic(err)
	}
	return
}

//MustQueryRecord 类似QueryRecord，出错换成异常
func MustQueryRecord(db DB, strSQL string, p map[string]interface{}) (result []map[string]interface{}, cols []string) {
	var err error
	if result, cols, err = QueryRecord(db, strSQL, p); err != nil {
		log.Panic(err)
	}
	return
}

//MustRow 类似Row，出错换成异常
func MustRow(db DB, strSQL string, p map[string]interface{}) (map[string]interface{}, []string) {
	var err error
	result, cols, err := QueryRecord(db, strSQL, p)
	if err != nil {
		log.Panic(err)
	}
	if len(result) == 0 {
		log.Panic(NewSQLError(strSQL, p, sql.ErrNoRows))
	}
	return result[0], cols
}

//GetTempTableName 获取一个临时表名
func GetTempTableName(db DB, prev string) (string, error) {
	if len(prev) == 0 {
		return "", fmt.Errorf("prev can't empty")
	}
	//确定名称
	tableName := ""
	rand.Seed(time.Now().UnixNano())
	bys := make([]byte, 4)
	icount := 0
	for {
		binary.BigEndian.PutUint32(bys, rand.Uint32())
		tableName = fmt.Sprintf("%s%X", prev, bys)
		if exists, err := TableExists(db, tableName); err != nil {
			return "", err
		} else if !exists {
			break
		}
		icount++
		if icount > 100 {
			return "", fmt.Errorf("find table name too much")
		}
	}
	return tableName, nil
}

//BatchExec 批量执行，分号换行的会被分开执行
func BatchExec(db DB, strSQL string, params map[string]interface{}) error {
	for _, v := range strings.Split(strSQL, ";\n") {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		_, err := db.NamedExec(v, params)
		if err != nil {
			return err
		}
	}
	return nil
}

//RenderSQLError 表示一个SQL语句渲染错误
type RenderSQLError struct {
	Template   string
	SQLParam   map[string]interface{}
	RenderArgs interface{}
	Err        error
}

func (r *RenderSQLError) Error() string {
	return fmt.Sprintf("Template:\n%s\nSqlParam:\n%#v\nRenderArgs:\n%#v\nError:\n%s", r.Template, r.SQLParam, r.RenderArgs, r.Err)
}

//RenderSQL 修改{{P}}的语法，因为后期的交叉汇总等需要sql传递的功能，生成参数就无法实现了，改成内嵌的字符串
//×渲染一个sql，可以用{{P val}}的语法加入一个参数，就不用考虑字符串转义了
//后期如果速度慢，可以加入一个模板缓存
func RenderSQL(strSQL string, renderArgs interface{}) (string, error) {

	if len(strSQL) == 0 {
		return strSQL, nil
	}
	var err error
	var t *template.Template
	if t, err = template.New("sql").Funcs(tempext.GetFuncMap()).Parse(strSQL); err != nil {
		return "", &RenderSQLError{strSQL, nil, renderArgs, err}
	}

	out := bytes.NewBuffer(nil)
	if err = t.Execute(out, renderArgs); err != nil {
		return "", &RenderSQLError{strSQL, nil, renderArgs, err}
	}
	strSQL = out.String()
	return strSQL, nil
}

//Count 统计一个SQL语句的返回行数，采用外套select count(*) from() 的方式
func Count(db DB, strSQL string, params map[string]interface{}) (result int64, err error) {
	v, err := GetSQLFun(db, fmt.Sprintf("SELECT COUNT(*) FROM (%s) count_sql", strSQL), params)
	if err != nil {
		return
	}
	result = safe.Int(v)
	return
}

//BindSQLWithError 类似BindSQL，返回错误，不抛出异常
func BindSQLWithError(db DB, strSQL string, params map[string]interface{}) (result string, paramsValues []interface{}, err error) {
	//转换in的条件
	sql, pam, err := sqlx.Named(strSQL, params)
	if err != nil {
		err = NewSQLError(strSQL, params, err)
		return
	}
	sql, pam, err = sqlx.In(sql, pam...)
	if err != nil {
		err = NewSQLError(strSQL, params, err)
		return
	}

	result = db.Rebind(sql)
	paramsValues = pam
	return
}

//BindSQL 转换绑定命名语法至常规，出错返回异常
func BindSQL(db DB, strSQL string, params map[string]interface{}) (result string, paramsValues []interface{}) {
	//转换in的条件
	sql, pam, err := sqlx.Named(strSQL, params)
	if err != nil {
		log.Panic(NewSQLError(strSQL, params, err))
	}
	sql, pam, err = sqlx.In(sql, pam...)
	if err != nil {
		log.Panic(NewSQLError(strSQL, params, err))
	}

	result = db.Rebind(sql)
	paramsValues = pam
	return
}

//CreateColumnIndex 新增单字段索引
func CreateColumnIndex(db DB, tableName, colName string) error {
	return Meta(db).CreateColumnIndex(db, tableName, colName)
}

//DropColumnIndex 删除单字段索引
func DropColumnIndex(db DB, tableName, indexName string) error {
	return Meta(db).DropColumnIndex(db, tableName, indexName)
}

//新增主键
func AddTablePrimaryKey(db DB, tableName string, pks []string) error {
	var strSql string
	ns := strings.Split(tableName, ".")
	var clearTableName string
	if len(ns) > 1 {
		clearTableName = ns[1]
	} else {
		clearTableName = tableName
	}
	switch db.DriverName() {
	case "postgres", "mysql":
		strSql = fmt.Sprintf("alter table %s add primary key(%s)", tableName, strings.Join(pks, ","))
	case "oci8":
		strSql = fmt.Sprintf("alter table %s add constraint %s_pk primary key(%s)", tableName, clearTableName, strings.Join(pks, ","))
	default:
		log.Panic("not impl," + db.DriverName())
	}
	if _, err := db.Exec(strSql); err != nil {
		return NewSQLError(strSql, nil, err)
	}
	return nil
}

//删除主键
func DropTablePrimaryKey(db DB, tableName string) error {
	log.WithFields(log.Fields{
		"table": tableName,
	}).Debug("dropkey")
	switch db.DriverName() {
	case "postgres":
		//先获取主键索引的名称，然后删除索引
		strSql := fmt.Sprintf(
			"select b.relname from  pg_index a inner join pg_class b on a.indexrelid =b.oid where indisprimary and indrelid='%s'::regclass",
			tableName)
		pkCons := ""
		if err := db.Get(&pkCons, strSql); err != nil {
			return NewSQLError(strSql, nil, err)
		}
		strSql = fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, pkCons)
		if _, err := db.Exec(strSql); err != nil {
			return NewSQLError(strSql, nil, err)
		}
	case "oci8":
		ns := strings.Split(tableName, ".")
		var strSql string
		if len(ns) > 1 {
			strSql = fmt.Sprintf(
				"select constraint_name from ALL_CONSTRAINTS where owner = '%s' and table_name ='%s' and constraint_type='P'",
				strings.ToUpper(ns[0]),
				strings.ToUpper(ns[1]))
		} else {
			strSql = fmt.Sprintf(
				"select constraint_name from user_CONSTRAINTS where table_name ='%s' and constraint_type='P'",
				strings.ToUpper(tableName))
		}
		pkCons := ""
		if rows, _, err := QueryRecord(db, strSql, nil); err != nil {
			return NewSQLError(strSql, nil, err)
		} else {
			if len(rows) > 0 {
				pkCons = rows[0]["CONSTRAINT_NAME"].(string)
			} else {
				return nil
			}
		}
		strSql = fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, pkCons)
		if _, err := db.Exec(strSql); err != nil {
			return NewSQLError(strSql, nil, err)
		}
	case "mysql":
		strSql := fmt.Sprintf("ALTER TABLE %s DROP PRIMARY KEY", tableName)
		if _, err := db.Exec(strSql); err != nil {
			return NewSQLError(strSql, nil, err)
		}
	default:
		log.Panic("not impl," + db.DriverName())
	}
	return nil
}

//返回一个字段值的字符串表达式
func ValueExpress(db DB, dataType int, value string) string {
	switch dataType {
	case TypeFloat, TypeInt:
		return value
	case TypeString:
		return safe.SignString(value)
	case TypeDatetime:
		switch db.DriverName() {
		case "oci8":
			if len(value) == 10 {
				return fmt.Sprintf("to_date(%s,'yyyy-mm-dd')", safe.SignString(value))
			} else if len(value) == 19 {
				return fmt.Sprintf("to_date(%s,'yyyy-mm-dd hh24:mi:ss')", safe.SignString(value))
			} else {
				log.Panic(fmt.Errorf("invalid datetime:%s", value))
				return ""
			}
		default:
			log.Panic(fmt.Errorf("not impl datetime,dbtype:%s", db.DriverName()))
			return ""
		}
	default:
		log.Panic(fmt.Errorf("not impl ValueExpress,type:%d", dataType))
		return ""
	}
}
func MustExec(db DB, strSql string, params ...interface{}) {
	if _, err := db.Exec(strSql, params...); err != nil {
		log.Panic(NewSQLError(strSql, params, err))
	}
	return
}

//返回差集的sql
func Minus(db DB, table1, where1, table2, where2 string, primaryKeys, cols []string) string {
	strSql := ""
	if len(where1) > 0 {
		where1 = "where " + where1
	}
	if len(where2) > 0 {
		where2 = "where " + where2
	}

	switch db.DriverName() {
	case "oci8":
		strSql = fmt.Sprintf(
			"select %s from %s %s minus select %s from %s %s",
			strings.Join(cols, ","),
			table1,
			where1,
			strings.Join(cols, ","),
			table2,
			where2)
	case "postgres":
		strSql = fmt.Sprintf(
			"select %s from %s %s EXCEPT select %s from %s %s",
			strings.Join(cols, ","),
			table1,
			where1,
			strings.Join(cols, ","),
			table2,
			where2)
	case "mysql":
		keyMap := map[string]bool{}
		for _, v := range primaryKeys {
			keyMap[v] = true
		}
		join := []string{}
		cols_l := []string{}
		for _, str := range cols {
			cols_l = append(cols_l, fmt.Sprintf("l_a.%s", str))
			//如果是主键，则不需要检查null
			if _, ok := keyMap[str]; ok {
				join = append(join, fmt.Sprintf("l_a.%[1]s=l_b.%[1]s", str))
			} else {
				join = append(join, fmt.Sprintf("(l_a.%s is null and l_b.%[1]s is null or l_a.%[1]s=l_b.%[1]s)", str))
			}
		}
		var from1 string
		var from2 string
		if len(where1) > 0 {
			from1 = fmt.Sprintf("(select * from %s %s)", table1, where1)
		} else {
			from1 = table1
		}
		if len(where2) > 0 {
			from2 = fmt.Sprintf("(select * from %s %s)", table2, where2)
		} else {
			from2 = table2
		}
		strSql = fmt.Sprintf(
			"select %s from %s l_a left join %s l_b on %s where %s",
			strings.Join(cols_l, ",\n"),
			from1,
			from2,
			strings.Join(join, " and\n"),
			fmt.Sprintf("l_b.%s is null", primaryKeys[0]))
	default:
		log.Panic("not impl")
	}
	return strSql
}
func DropIndexIfExists(db DB, indexName string) error {
	var strSQL string
	switch db.DriverName() {
	case "oci8":
		strSQL = fmt.Sprintf(`
		DECLARE
		  COUNT_INDEXES INTEGER;
		BEGIN
		  SELECT COUNT(*) INTO COUNT_INDEXES
		    FROM USER_INDEXES
		    WHERE INDEX_NAME = '%s';

		  IF COUNT_INDEXES = 1 THEN
		    EXECUTE IMMEDIATE %s;
		  END IF;
		END;`, indexName,
			safe.SignString(fmt.Sprintf("drop index %s", indexName)))
	case "sqlite3", "postgres":
		strSQL = fmt.Sprintf("drop index if exists %s", indexName)
	default:
		return fmt.Errorf("invalid driver")
	}
	if _, err := db.Exec(strSQL); err != nil {
		return NewSQLError(strSQL, nil, err)
	}
	return nil
}

func CreateIndexIfNotExists(db DB, indexName, tableName, express string) error {
	var strSQL string
	switch db.DriverName() {
	case "oci8":
		strSQL = fmt.Sprintf(`
		DECLARE
		  COUNT_INDEXES INTEGER;
		BEGIN
		  SELECT COUNT(*) INTO COUNT_INDEXES
		    FROM USER_INDEXES
		    WHERE INDEX_NAME = '%s';

		  IF COUNT_INDEXES = 0 THEN
		    EXECUTE IMMEDIATE %s;
		  END IF;
		END;`, indexName,
			safe.SignString(fmt.Sprintf("create index %s on %s(%s)", indexName, tableName, express)))
	case "sqlite3", "postgres":
		strSQL = fmt.Sprintf("create index if not exists %s on %s(%s)", indexName, tableName, express)
	default:
		return fmt.Errorf("invalid driver")
	}
	if _, err := db.Exec(strSQL); err != nil {
		return NewSQLError(strSQL, nil, err)
	}
	return nil
}

//在一个事务中运行，自动处理commit 和rollback
func RunAtTx(db *sqlx.DB, callback func(DB) error) (err error) {
	var tx *sqlx.Tx
	if tx, err = db.Beginx(); err != nil {
		return err
	}
	finish := false
	defer func() {
		//如果没有设置，说明是中途跳出，发生了异常
		//这里不捕获异常是要保留现场
		if !finish {
			tx.Rollback()
		}
	}()
	if err = callback(tx); err != nil {
		tx.Rollback()
	} else {
		err = tx.Commit()
	}
	finish = true
	return
}
