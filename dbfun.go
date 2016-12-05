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

var (
	PKValuesNumberError = fmt.Errorf("the pk values number error")
)

type SqlError struct {
	Sql    string
	Params interface{}
	Err    error
}

func (e SqlError) Error() string {
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
	return fmt.Sprintf("%s\n%s\nparams len is %d,content is:\n%s", e.Err, e.Sql, l, content)
}

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

func Columns(db DB, strSql string, pam map[string]interface{}) ([]string, error) {
	rows, err := db.NamedQuery(strSql, pam)
	if err != nil {
		return nil, SqlError{strSql, pam, err}
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
func TableNames(db DB) (names []string) {
	var strSql string
	switch db.DriverName() {
	case "postgres":
		strSql = "SELECT table_name FROM information_schema.tables WHERE table_schema = current_schema()"
	case "oci8":
		strSql = "SELECT table_name FROM user_tables"
	case "mysql":
		strSql = "SELECT table_name FROM information_schema.tables WHERE table_schema = schema()"
	default:
		log.Panic("not impl," + db.DriverName())
	}
	names = []string{}
	if err := db.Select(&names, strSql); err != nil {
		log.Panic(err)
	}
	for i, v := range names {
		names[i] = strings.ToUpper(v)
	}
	sort.Strings(names)
	return
}
func NameGet(db DB, d interface{}, strSql string, p map[string]interface{}) error {
	str, pam := BindSql(db, strSql, p)
	if err := db.Get(d, str, pam...); err != nil {
		return SqlError{strSql, p, err}
	}
	return nil
}
func NameSelect(db DB, d interface{}, strSql string, p map[string]interface{}) error {
	str, pam := BindSql(db, strSql, p)
	if err := db.Select(d, str, pam...); err != nil {
		return SqlError{strSql, p, err}
	}
	return nil
}
func QueryRecord(db DB, strSql string, p map[string]interface{}) (result []map[string]interface{}, err error) {
	var rows *sqlx.Rows
	str, pam := BindSql(db, strSql, p)
	if rows, err = db.Queryx(str, pam...); err != nil {
		err = SqlError{strSql, p, err}
		return
	}
	result = []map[string]interface{}{}
	defer rows.Close()
	for rows.Next() {
		oneRecord := map[string]interface{}{}
		if err = rows.MapScan(oneRecord); err != nil {
			err = SqlError{strSql, p, err}
			return
		}

		result = append(result, mapfun.UpperKeys(oneRecord))
	}
	return
}
func IsNull(db DB) string {
	switch db.DriverName() {
	case "oci8":
		return "nvl"
	case "postgres":
		return "COALESCE"
	case "mysql", "sqlite3":
		return "ifnull"
	default:
		log.Panic("not impl")
		return ""
	}

}
func Exists(db DB, strSql string, p map[string]interface{}) (result bool, err error) {
	str, pam := BindSql(db, strSql, p)

	var rows *sqlx.Rows
	if rows, err = db.Queryx(str, pam...); err != nil {
		err = SqlError{strSql, p, err}
		return
	}
	defer rows.Close()
	result = rows.Next()
	return
}

//执行create table as select语句
func CreateTableAs(db DB, tableName, strSql string, pks []string) error {
	switch db.DriverName() {
	case "postgres", "mysql", "oci8":
		s := fmt.Sprintf("CREATE TABLE %s as %s", tableName, strSql)
		if _, err := db.Exec(s); err != nil {
			return SqlError{s, nil, err}
		}
		s = fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY(%s)", tableName, strings.Join(pks, ","))
		if _, err := db.Exec(s); err != nil {
			return SqlError{s, nil, err}
		}
	default:
		log.Panic("not impl create table as")
	}
	return nil
}

//删除表字段
func TableRemoveColumns(db DB, tabName string, cols []string) error {
	var strSql string
	switch db.DriverName() {
	case "postgres", "mysql":
		strList := []string{}
		for _, v := range cols {
			strList = append(strList, "DROP COLUMN "+v)
		}
		strSql = fmt.Sprintf("ALTER table %s %s", tabName, strings.Join(strList, ","))
	case "oci8":
		strSql = fmt.Sprintf("ALTER table %s drop(%s)", tabName, strings.Join(cols, ","))
	default:
		return fmt.Errorf("not impl," + db.DriverName())
	}
	log.Println(strSql)
	if _, err := db.Exec(strSql); err != nil {
		return SqlError{strSql, nil, err}
	}
	return nil

}

//表更名
func TableRename(db DB, oldName, newName string) error {
	var strSql string
	switch db.DriverName() {
	case "postgres", "sqlite3":
		strSql = fmt.Sprintf("ALTER table %s RENAME TO %s", oldName, newName)
	case "oci8", "mysql":
		strSql = fmt.Sprintf("rename table %s TO %s", oldName, newName)
	default:
		return fmt.Errorf("not impl," + db.DriverName())
	}
	log.Println(strSql)
	if _, err := db.Exec(strSql); err != nil {
		return SqlError{strSql, nil, err}
	}
	return nil
}
func TableExists(db DB, tableName string) (bool, error) {
	schema := ""
	ns := strings.Split(tableName, ".")
	tname := ""
	if len(ns) > 1 {
		schema = ns[0]
		tname = ns[1]
	} else {

		tname = tableName
	}
	var strSql string
	switch db.DriverName() {
	case "postgres":
		if len(schema) == 0 {
			schema = safe.String(MustGetSqlFun(db, "select current_schema()", nil))
		}

		strSql = fmt.Sprintf(
			"SELECT count(*) FROM information_schema.tables WHERE table_schema ilike '%s' and table_name ilike :tname", schema)
	case "oci8":
		if len(schema) == 0 {
			schema = safe.String(MustGetSqlFun(db, "select user from dual", nil))
		}
		strSql = fmt.Sprintf("SELECT count(*) FROM all_tables where owner='%s' and table_name=:tname", schema)
	case "mysql":
		if len(schema) == 0 {
			schema = safe.String(MustGetSqlFun(db, "select schema()", nil))
		}
		strSql = fmt.Sprintf(
			"SELECT count(*) FROM information_schema.tables WHERE table_schema = '%s' and UPPER(table_name)=:tname", schema)
	default:
		return false, fmt.Errorf("not impl," + db.DriverName())
	}
	var iCount int64
	p := map[string]interface{}{"tname": strings.ToUpper(tname)}
	if err := NameGet(db, &iCount, strSql, p); err != nil {
		return false, SqlError{strSql, p, err}
	}
	return iCount > 0, nil
}
func GetSqlFun(db DB, strSql string, p map[string]interface{}) (result interface{}, err error) {
	str, pam := BindSql(db, strSql, p)

	var rows *sqlx.Rows
	if rows, err = db.Queryx(str, pam...); err != nil {
		err = SqlError{strSql, p, err}
		return
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&result); err != nil {
			err = SqlError{strSql, p, err}
			return
		}
	}
	return
}

func MustGetSqlFun(db DB, strSql string, p map[string]interface{}) (result interface{}) {
	var err error
	if result, err = GetSqlFun(db, strSql, p); err != nil {
		log.Printf("sql err:%s\n%s\n", err, strSql)
		log.Panic(err)
	}
	return
}
func MustQueryRecord(db DB, strSql string, p map[string]interface{}) (result []map[string]interface{}) {
	var err error
	if result, err = QueryRecord(db, strSql, p); err != nil {
		log.Panic(err)
	}
	return
}
func MustRow(db DB, strSql string, p map[string]interface{}) map[string]interface{} {
	var err error
	result, err := QueryRecord(db, strSql, p)
	if err != nil {
		log.Panic(err)
	}
	if len(result) == 0 {
		log.Panic(SqlError{strSql, p, sql.ErrNoRows})
	}
	return result[0]
}

//获取一个临时表名
func GetTempTableName(db DB, prev string) (string, error) {
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

//批量执行，分号换行的会被分开执行
func BatchExec(db DB, strSql string, params map[string]interface{}) error {
	for _, v := range strings.Split(strSql, ";\n") {
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

type RenderSqlError struct {
	Template   string
	SqlParam   map[string]interface{}
	RenderArgs interface{}
	Err        error
}

func (r *RenderSqlError) Error() string {
	return fmt.Sprintf("Template:\n%s\nSqlParam:\n%#v\nRenderArgs:\n%#v\nError:\n%s", r.Template, r.SqlParam, r.RenderArgs, r.Err)
}

//修改{{P}}的语法，因为后期的交叉汇总等需要sql传递的功能，生成参数就无法实现了，改成内嵌的字符串
//×渲染一个sql，可以用{{P val}}的语法加入一个参数，就不用考虑字符串转义了
//后期如果速度慢，可以加入一个模板缓存
func RenderSql(strSql string, renderArgs interface{}) (string, error) {

	if len(strSql) == 0 {
		return strSql, nil
	}
	var err error
	var t *template.Template
	if t, err = template.New("sql").Funcs(tempext.GetFuncMap()).Parse(strSql); err != nil {
		return "", &RenderSqlError{strSql, nil, renderArgs, err}
	}

	out := bytes.NewBuffer(nil)
	if err = t.Execute(out, renderArgs); err != nil {
		return "", &RenderSqlError{strSql, nil, renderArgs, err}
	}
	strSql = out.String()
	return strSql, nil
}
func Count(db DB, strSql string, params map[string]interface{}) (result int64, err error) {
	v, err := GetSqlFun(db, fmt.Sprintf("SELECT COUNT(*) FROM (%s) count_sql", strSql), params)
	if err != nil {
		return
	}
	result = safe.Int(v)
	return
}
func BindSqlWithError(db DB, strSql string, params map[string]interface{}) (result string, paramsValues []interface{}, err error) {
	//转换in的条件
	sql, pam, err := sqlx.Named(strSql, params)
	if err != nil {
		err = &SqlError{strSql, params, err}
		return
	}
	sql, pam, err = sqlx.In(sql, pam...)
	if err != nil {
		err = &SqlError{strSql, params, err}
		return
	}

	result = db.Rebind(sql)
	paramsValues = pam
	return
}
func BindSql(db DB, strSql string, params map[string]interface{}) (result string, paramsValues []interface{}) {
	//转换in的条件
	sql, pam, err := sqlx.Named(strSql, params)
	if err != nil {
		log.Panic(&SqlError{strSql, params, err})
	}
	sql, pam, err = sqlx.In(sql, pam...)
	if err != nil {
		log.Panic(&SqlError{strSql, params, err})
	}

	result = db.Rebind(sql)
	paramsValues = pam
	return
}

//新增单字段索引
func CreateColumnIndex(db DB, tableName, colName string) error {
	ns := strings.Split(tableName, ".")
	schema := ""
	tname := ""
	if len(ns) > 1 {
		schema = ns[0] + "."
		tname = ns[1]
	} else {
		tname = tableName
	}
	var strSql string
	switch db.DriverName() {
	case "postgres":
		strSql = fmt.Sprintf("create index on %s(%s)", tableName, colName)
	case "oci8", "mysql", "sqlite3":
		//这里会有问题，如果表名和字段名比较长就会出错
		strSql = fmt.Sprintf("create index %si%s%s on %s(%s)", schema, tname, colName, tableName, colName)
	default:
		log.Panic("not impl " + db.DriverName())
	}
	if _, err := db.Exec(strSql); err != nil {
		return SqlError{strSql, nil, err}
	}
	log.Println(strSql)
	return nil
}

//删除单字段索引
func DropColumnIndex(db DB, tableName, indexName string) error {
	var strSql string

	switch db.DriverName() {
	case "postgres", "oci8", "sqlite3":
		strSql = fmt.Sprintf("drop index %s", indexName)
	case "mysql":
		strSql = fmt.Sprintf("drop index %s on %s", indexName, tableName)
	default:
		log.Panic("not impl," + db.DriverName())
	}
	if _, err := db.Exec(strSql); err != nil {
		return SqlError{strSql, nil, err}
	}
	return nil
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
		return SqlError{strSql, nil, err}
	}
	return nil
}

//删除主键
func DropTablePrimaryKey(db DB, tableName string) error {
	switch db.DriverName() {
	case "postgres":
		//先获取主键索引的名称，然后删除索引
		strSql := fmt.Sprintf(
			"select b.relname from  pg_index a inner join pg_class b on a.indexrelid =b.oid where indisprimary and indrelid='%s'::regclass",
			tableName)
		pkCons := ""
		if err := db.Get(&pkCons, strSql); err != nil {
			return SqlError{strSql, nil, err}
		}
		strSql = fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, pkCons)
		if _, err := db.Exec(strSql); err != nil {
			return SqlError{strSql, nil, err}
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
		if rows, err := QueryRecord(db, strSql, nil); err != nil {
			return SqlError{strSql, nil, err}
		} else {
			if len(rows) > 0 {
				pkCons = rows[0]["CONSTRAINT_NAME"].(string)
			} else {
				return nil
			}
		}
		strSql = fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, pkCons)
		if _, err := db.Exec(strSql); err != nil {
			return SqlError{strSql, nil, err}
		}
	case "mysql":
		strSql := fmt.Sprintf("ALTER TABLE %s DROP PRIMARY KEY", tableName)
		if _, err := db.Exec(strSql); err != nil {
			return SqlError{strSql, nil, err}
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
		log.Panic(SqlError{strSql, params, err})
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
func CreateIndexIfNotExists(db DB, indexName, tableName, express string) error {
	var strSql string
	switch db.DriverName() {
	case "oci8":
		strSql = fmt.Sprintf(`
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
	default:
		return fmt.Errorf("invalid driver")
	}
	if _, err := db.Exec(strSql); err != nil {
		return SqlError{strSql, nil, err}
	}
	return nil
}

//在一个事务中运行，自动处理commit 和rollback
func RunAtTx(db *sqlx.DB, callback func(DB) error) (err error) {
	var tx *sqlx.Tx
	if tx, err = db.Beginx(); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		if r := recover(); r != nil {
			tx.Rollback()
			log.Panic(r)
		}
		err = tx.Commit()
	}()
	err = callback(tx)
	return
}
