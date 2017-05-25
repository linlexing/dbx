package dbx

import (
	"bytes"
	"dbweb/lib/safe"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/jmoiron/sqlx"
)

//条件一行
type ConditionLine struct {
	LeftBrackets  string
	ColumnName    string
	Operators     string
	Value         string
	RightBrackets string
	Logic         string
}

//模板条件
type SqlCondition struct {
	Name      string
	Lines     []*ConditionLine
	PlainText string
}

//排序的字段
type OrderColumn struct {
	Name  string
	Order string //DESC ASC
	Value interface{}
}
type SqlSelect struct {
	sql           string
	ManualPage    bool
	Table         *DBTable //对应数据库中的表，用于探查字段的数据类型，如果为空或者字段在表中不存在，则是STR
	Conditions    []*SqlCondition
	Columns       []string
	Order         []string
	Divide        []string
	Limit         int64
	SqlRenderArgs interface{} //sql语句在查询前，还会用template进行一次渲染，这里传入渲染的参数
}

func FieldValueToString(val interface{}) string {
	return safe.String(val)
}
func StringToFieldValue(str string, dataType int) (val interface{}) {
	var err error

	switch dataType {
	case TypeBytea:
		log.Panic("the bytea not impl")
	case TypeDatetime:
		if val, err = time.Parse("2006-01-02 15:04:05", str); err != nil {
			log.Panic(err)
		}
	case TypeFloat:
		if val, err = strconv.ParseFloat(str, 64); err != nil {
			log.Panic(err)
		}

	case TypeInt:
		if val, err = strconv.ParseInt(str, 10, 64); err != nil {
			log.Panic(err)
		}

	case TypeString:
		val = str
	default:
		log.Panic("not impl StringToFieldValue")
	}
	return
}

func (c *ConditionLine) GetExpress(db DB, dataType int) (strSql string) {
	//需要考虑到null的情况
	switch c.Operators {
	case "=": //等于
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is null", c.ColumnName)
		} else {
			strSql = fmt.Sprintf("%s = %s", c.ColumnName, ValueExpress(db, dataType, c.Value))
		}
	case "!=": //不等于
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is not null", c.ColumnName)
		} else {
			strSql = fmt.Sprintf("(%s <> %s or %[1]s is null)", c.ColumnName, ValueExpress(db, dataType, c.Value))
		}
	case ">": //大于
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is not null", c.ColumnName)
		} else {
			strSql = fmt.Sprintf("%s > %s", c.ColumnName, ValueExpress(db, dataType, c.Value))
		}
	case ">=": //大于等于
		if c.Value == "" {
			strSql = "1=1"
		} else {
			strSql = fmt.Sprintf("%s >= %s", c.ColumnName, ValueExpress(db, dataType, c.Value))
		}
	case "<": //小于
		if c.Value == "" {
			strSql = "1=2"
		} else {
			strSql = fmt.Sprintf("(%s < %s or %[1]s is null)", c.ColumnName, ValueExpress(db, dataType, c.Value))
		}
	case "<=": //小于等于
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is null", c.ColumnName)

		} else {
			strSql = fmt.Sprintf("(%s <= %s or %[1]s is null)", c.ColumnName, ValueExpress(db, dataType, c.Value))
		}
	case "?": //包含
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is null", c.ColumnName)
		} else {
			strSql = fmt.Sprintf("%s like %s", c.ColumnName, ValueExpress(db, dataType, "%"+c.Value+"%"))
		}
	case "!?": //不包含
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is not null", c.ColumnName)
		} else {
			strSql = fmt.Sprintf("%s not like %s", c.ColumnName, ValueExpress(db, dataType, "%"+c.Value+"%"))
		}
	case "?>": //前缀
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is null", c.ColumnName)
		} else {
			strSql = fmt.Sprintf("%s like %s", c.ColumnName, ValueExpress(db, dataType, c.Value+"%"))
		}
	case "!?>": //非前缀
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is not null", c.ColumnName)
		} else {
			strSql = fmt.Sprintf("%s not like %s", c.ColumnName, ValueExpress(db, dataType, c.Value+"%"))
		}
	case "<?": //后缀
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is null", c.ColumnName)
		} else {
			strSql = fmt.Sprintf("%s like %s", c.ColumnName, ValueExpress(db, dataType, "%"+c.Value))
		}
	case "!<?": //非后缀
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is not null", c.ColumnName)
		} else {
			strSql = fmt.Sprintf("%s not like %s", c.ColumnName, ValueExpress(db, dataType, "%"+c.Value))
		}
	case "in": //在列表
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is null", c.ColumnName)
		} else {
			//在列表简化起见，不再类型化
			if array, err := csv.NewReader(strings.NewReader(c.Value)).Read(); err != nil {
				log.Panic(err)
			} else {
				list := []string{}
				for _, v := range array {
					list = append(list, ValueExpress(db, dataType, v))
				}
				strSql = fmt.Sprintf("%s in (%s)", c.ColumnName, strings.Join(list, ",\n"))
			}

		}
	case "!in": //不在列表
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is not null", c.ColumnName)
		} else {
			//在列表简化起见，不再类型化
			if array, err := csv.NewReader(strings.NewReader(c.Value)).Read(); err != nil {
				log.Panic(err)
			} else {
				list := []string{}
				for _, v := range array {
					list = append(list, ValueExpress(db, dataType, v))
				}
				strSql = fmt.Sprintf("%s not in (%s)", c.ColumnName, strings.Join(list, ",\n"))
			}
		}
	case "~": //正则
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is null", c.ColumnName)
		} else {
			switch db.DriverName() {
			case "oci8":
				strSql = fmt.Sprintf("regexp_like(%s,%s)", c.ColumnName, ValueExpress(db, dataType, c.Value))
			case "postgres":
				strSql = fmt.Sprintf("%s ~ %s", c.ColumnName, ValueExpress(db, dataType, c.Value))
			case "mysql":
				strSql = fmt.Sprintf("%s REGEXP %s", c.ColumnName, ValueExpress(db, dataType, c.Value))
			default:
				log.Panic("not impl GetExpress")
			}
		}
	case "!~": //非正则
		if c.Value == "" {
			strSql = fmt.Sprintf("%s is not null", c.ColumnName)
		} else {
			switch db.DriverName() {
			case "oci8":
				strSql = fmt.Sprintf("not regexp_like(%s,%s)", c.ColumnName, ValueExpress(db, dataType, c.Value))
			case "postgres":
				strSql = fmt.Sprintf("%s !~ %s", c.ColumnName, ValueExpress(db, dataType, c.Value))
			case "mysql":
				strSql = fmt.Sprintf("%s not REGEXP %s", c.ColumnName, ValueExpress(db, dataType, c.Value))
			default:
				log.Panic("not impl GetExpress")
			}
		}
	case "e": //为空
		strSql = fmt.Sprintf("%s is null", c.ColumnName)
	case "!e": //不为空
		strSql = fmt.Sprintf("%s is not null", c.ColumnName)
	case "_": //长度等于
		switch db.DriverName() {
		case "oci8", "postgres":
			strSql = fmt.Sprintf("length(%s) = %s", c.ColumnName, c.Value)
		case "mysql":
			strSql = fmt.Sprintf("char_length(%s) = %s", c.ColumnName, c.Value)
		default:
			log.Panic("not impl GetExpress")
		}
	case "!_": //长度不等于
		switch db.DriverName() {
		case "oci8", "postgres":
			strSql = fmt.Sprintf("length(%s) <> %s", c.ColumnName, c.Value)
		case "mysql":
			strSql = fmt.Sprintf("char_length(%s) <> %s", c.ColumnName, c.Value)
		default:
			log.Panic("not impl GetExpress")
		}
	case "_>": //长度大于
		switch db.DriverName() {
		case "oci8", "postgres":
			strSql = fmt.Sprintf("length(%s) > %s", c.ColumnName, c.Value)
		case "mysql":
			strSql = fmt.Sprintf("char_length(%s) > %s", c.ColumnName, c.Value)
		default:
			log.Panic("not impl GetExpress")
		}

	case "_<": //长度小于
		switch db.DriverName() {
		case "oci8", "postgres":
			strSql = fmt.Sprintf("length(%s) < %s", c.ColumnName, c.Value)
		case "mysql":
			strSql = fmt.Sprintf("char_length(%s) < %s", c.ColumnName, c.Value)
		default:
			log.Panic("not impl GetExpress")
		}
	default:
		log.Panic(fmt.Errorf("the opt:%s not impl", c.Operators))
	}
	//加上括号
	strSql = fmt.Sprintf("%s%s%s", c.LeftBrackets, strSql, c.RightBrackets)
	return
}
func (c *SqlCondition) BuildWhere(db DB, table *DBTable) string {
	strLines := []string{}

	if len(c.Lines) > 0 {
		//最后一行的逻辑设置为空
		c.Lines[len(c.Lines)-1].Logic = ""
		for i, v := range c.Lines {
			dataType := TypeString
			if table != nil {
				if field := table.Field(v.ColumnName); field != nil {
					dataType = field.GoType()
				}
			}
			exp := v.GetExpress(db, dataType)
			//最后一行不需要加逻辑
			if i < len(c.Lines)-1 {
				strLines = append(strLines, exp+" "+v.Logic)
			} else {
				strLines = append(strLines, exp)
			}

		}
	}
	if len(c.PlainText) > 0 {
		if len(strLines) > 0 {
			return fmt.Sprintf("(\n%s\n) and (\n%s\n)", strings.Join(strLines, "\n"), c.PlainText)
		} else {
			return c.PlainText
		}
	} else {
		if len(strLines) > 0 {
			return strings.Join(strLines, "\n")
		} else {
			return ""
		}
	}
}
func buildCondition(order, divide []string) []*ConditionLine {
	result := []*ConditionLine{}
	//a=:a and b=:b and c>:c or
	//a=:a and b>:b or
	//a>:a
	for i := len(divide) - 1; i >= 0; i-- {
		lines := []*ConditionLine{}
		for j := i; j >= 0; j-- {
			colName := order[j]
			isDesc := strings.HasPrefix(colName, "-")
			if isDesc {
				colName = colName[1:]
			}
			//尾部指标是用大于或者小于（倒序）
			if j == i {
				opt := ">"
				if isDesc {
					opt = "<"
				}
				lines = append(lines, &ConditionLine{
					ColumnName: colName,
					Operators:  opt,
					Value:      divide[j],
					Logic:      "AND",
				})
			} else {
				lines = append(lines, &ConditionLine{
					ColumnName: colName,
					Operators:  "=",
					Value:      divide[j],
					Logic:      "AND",
				})
			}
		}
		if len(lines) > 1 {
			lines[0].LeftBrackets = "("
			lines[len(lines)-1].RightBrackets = ")"
		}
		lines[len(lines)-1].Logic = "OR"
		result = append(result, lines...)
	}
	return result
}

func renderManualPageSql(db DB, strSql string, columnList, whereList, orderbyList []string, limit int64) (string, error) {
	tmpl, err := template.New("ManualPage").Delims("<<", ">>").Parse(strSql)
	if err != nil {
		return "", err
	}
	var where string
	var columns string
	var orderby string
	bys := bytes.NewBuffer(nil)
	if len(whereList) > 0 {
		where = "(" + strings.Join(whereList, " and ") + ")"
	}
	if len(columnList) > 0 {
		columns = strings.Join(columnList, ",")
	}
	if len(orderbyList) > 0 {
		orderby = strings.Join(orderbyList, ",")
	}
	if err = tmpl.Execute(bys, map[string]interface{}{
		"Driver":  db.DriverName(),
		"Columns": columns,
		"Where":   where,
		"OrderBy": orderby,
		"Limit":   limit,
	}); err != nil {
		return "", err
	}
	return bys.String(), nil
}

//构造sql语句，和相应的参数值
func (s *SqlSelect) BuildSql(db DB) (strSql string) {

	renderSql, err := s.renderSql()
	if err != nil {
		log.Panic(err)
	}
	if len(renderSql) == 0 {
		log.Panic("sql is empty")
	}

	whereList := []string{}
	orderList := []string{}
	//全空返回sql
	if len(s.Conditions) == 0 &&
		len(s.Columns) == 0 &&
		len(s.Order) == 0 &&
		s.Limit < 0 {
		if s.ManualPage {
			var err error
			strSql, err = renderManualPageSql(db, renderSql, nil, nil, nil, s.Limit)
			if err != nil {
				log.Panic(err)
			}
		} else {
			strSql = renderSql
		}
		return

	}
	//where
	if len(s.Conditions) > 0 {
		for _, v := range s.Conditions {
			if str := v.BuildWhere(db, s.Table); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}
	if len(s.Order) > 0 {
		for _, v := range s.Order {
			if strings.HasPrefix(v, "-") {
				if db.DriverName() == "mysql" ||
					db.DriverName() == "sqlite3" {
					orderList = append(orderList, v[1:]+" DESC")
				} else {
					orderList = append(orderList, v[1:]+" DESC NULLS LAST")
				}
			} else {
				if db.DriverName() == "mysql" ||
					db.DriverName() == "sqlite3" {
					orderList = append(orderList, v)
				} else {
					orderList = append(orderList, v+" NULLS FIRST")
				}
			}
		}
		if len(s.Divide) > 0 {
			divideCondition := &SqlCondition{
				Name:  "divide",
				Lines: buildCondition(s.Order, s.Divide),
			}
			if str := divideCondition.BuildWhere(db, s.Table); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}

	if s.ManualPage {
		if str, err := renderManualPageSql(db, renderSql, s.Columns, whereList, orderList, s.Limit); err != nil {
			log.Panic(err)
		} else {
			strSql = str
		}
	} else {
		var where, orderby, sel string
		sel = "*"

		//select
		if len(s.Columns) > 0 {
			sel = strings.Join(s.Columns, ",")
		}
		if len(whereList) > 0 {
			where = " where " + strings.Join(whereList, " and ")
		}
		if len(orderList) > 0 {
			orderby = " order by " + strings.Join(orderList, ",")
		}
		if s.Limit >= 0 {
			switch db.DriverName() {
			case "oci8":
				strSql = fmt.Sprintf(
					"select * from (select %s from (%s) wholesql %s%s) where rownum<=%d",
					sel, renderSql, where, orderby, s.Limit)
			case "mysql", "postgres", "sqlite3":
				strSql = fmt.Sprintf("select %s from (%s) wholesql %s%s limit %d",
					sel, renderSql, where, orderby, s.Limit)
			default:
				log.Panic("not impl BuildSql")
			}

		} else {
			strSql = fmt.Sprintf("select %s from (%s) wholesql %s%s", sel, renderSql, where, orderby)
		}
	}
	//在最后返回前，还需要一次render，防止条件中引用了模板
	if s, err := RenderSQL(strSql, s.SqlRenderArgs); err != nil {
		log.Panic(err)
	} else {
		strSql = s
	}

	return
}
func (s *SqlSelect) convertRow(row map[string]interface{}) map[string]interface{} {
	if s.Table != nil {
		return s.Table.ConvertToTrueType(row)
	} else {
		transRecord := map[string]interface{}{}
		for k, v := range row {
			k = strings.ToUpper(k)
			switch tv := v.(type) {
			case []byte:
				transRecord[k] = string(tv)
			default:
				transRecord[k] = tv
			}
		}
		return transRecord
	}
}
func (s *SqlSelect) QueryRows(db DB) (result []map[string]interface{}, cols []*ColumnType, err error) {
	strSQL := s.BuildSql(db)
	var rows *sqlx.Rows
	if rows, err = db.Queryx(strSQL); err != nil {
		err = NewSQLError(strSQL, nil, err)
		return
	}
	var columns []string
	if columns, err = rows.Columns(); err != nil {
		return
	}
	cols = []*ColumnType{}
	//先根据预置的表获取对应的字段类型
	for _, v := range columns {
		col := &ColumnType{
			Name: strings.ToUpper(v),
		}
		if s.Table != nil {
			if tCol := s.Table.Field(col.Name); tCol != nil {
				col.Type = tCol.Type
			}
		}
		if len(col.Type) == 0 {
			col.Type = "STR"
		}

		cols = append(cols, col)
	}

	result = []map[string]interface{}{}
	defer rows.Close()
	for rows.Next() {
		oneRecord := map[string]interface{}{}
		if err = rows.MapScan(oneRecord); err != nil {
			err = NewSQLError(strSQL, nil, err)
			return
		}
		result = append(result, s.convertRow(oneRecord))
		//再检查所有没有类型的字段的值，根据值来设置类型，nil值的确实没有办法了
		for _, v := range cols {
			if len(v.Type) == 0 {
				//发现Oracle的数值、整型返回的的是字符串，得想其他办法弥补
				switch oneRecord[v.Name].(type) {
				//由于字符串返回[]byte，所以bytea就没了
				//case []byte:
				//	v.Type= "BYTEA"
				case time.Time, *time.Time:
					v.Type = "DATE"
				case float32, float64:
					v.Type = "FLOAT"
				case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
					v.Type = "INT"
				case string, []byte, nil: //nil作为str处理
					v.Type = "STR"
				default:
					log.Panic("not impl QueryRows")
				}
			}
		}
	}
	return
}

//渲染sql
func (s *SqlSelect) renderSql() (strSql string, err error) {
	return RenderSQL(s.sql, s.SqlRenderArgs)
}

//如果没有数值字段或者没有记录，则返回空sql
func (s *SqlSelect) BuildTotalSql(db DB, cols ...string) (strSql string, err error) {
	totalCoumns := []string{}
	for _, col := range cols {
		totalCoumns = append(totalCoumns, fmt.Sprintf("sum(cast(%s(%s,0) as decimal(29,6))) as %[2]s", IsNull(db), col))
	}
	if len(totalCoumns) == 0 {
		return
	}

	renderSql, err := s.renderSql()
	if err != nil {
		return
	}
	if len(renderSql) == 0 {
		err = fmt.Errorf("sql is empty")
		return
	}

	var where string
	whereList := []string{}

	//where
	if len(s.Conditions) > 0 {
		for _, v := range s.Conditions {
			if str := v.BuildWhere(db, s.Table); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}
	if s.ManualPage {
		var str string
		if str, err = renderManualPageSql(db, renderSql, totalCoumns, whereList, nil, -1); err != nil {
			return
		} else {
			strSql = str
		}
	} else {
		if len(whereList) > 0 {
			where = " where " + strings.Join(whereList, " and ")
		}
		strSql = fmt.Sprintf("select %s from (%s) wholesql %s", strings.Join(totalCoumns, ","), renderSql, where)
	}
	strSql, err = RenderSQL(strSql, s.SqlRenderArgs)
	return

}
func (s *SqlSelect) BuildRowCountSql(db DB) (strSQL string) {
	renderSql, err := s.renderSql()
	if err != nil {
		log.Panic(err)
	}
	if len(renderSql) == 0 {
		log.Panic("sql is empty")
	}

	var where string
	whereList := []string{}

	//where
	if len(s.Conditions) > 0 {
		for _, v := range s.Conditions {
			if str := v.BuildWhere(db, s.Table); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}
	if s.ManualPage {
		if str, err := renderManualPageSql(db, renderSql, []string{"COUNT(*)"}, whereList, nil, -1); err != nil {
			log.Panic(err)
		} else {
			strSQL = str
		}
	} else {
		if len(whereList) > 0 {
			where = " where " + strings.Join(whereList, " and ")
		}
		strSQL = fmt.Sprintf("select count(*) from (%s) wholesql %s", renderSql, where)
	}
	if s, err := RenderSQL(strSQL, s.SqlRenderArgs); err != nil {
		log.Panic(err)
	} else {
		strSQL = s
	}

	return
}

//汇总数值字段
func (s *SqlSelect) Total(db DB, cols ...string) (result map[string]interface{}, err error) {
	strSql, err := s.BuildTotalSql(db, cols...)
	if err != nil || len(strSql) == 0 {
		return
	}
	rows, _, err := QueryRecord(db, strSql, nil)
	if err != nil {
		return
	}
	if len(rows) == 0 {
		err = fmt.Errorf("sql:%s\npam:%v\ntotal is nil", strSql, nil)
		return
	}
	result = rows[0]
	return
}
func (s *SqlSelect) RowCount(db DB) (r int64, err error) {
	r = -1

	strSql := s.BuildRowCountSql(db)
	if err = db.Get(&r, strSql); err != nil {
		err = NewSQLError(strSql, nil, err)
	}
	return

}

func NewSqlSelect(strSql string, table *DBTable, manualPage bool) *SqlSelect {
	//如果没有sql语句而有表名，则生成一个sql
	if len(strSql) > 0 {
		return &SqlSelect{
			sql:        strSql,
			Table:      table,
			ManualPage: manualPage,
		}
	}
	if table == nil {
		log.Panic("no table")
	}
	strSql = fmt.Sprintf(`<<if eq "oci8" .Driver>>
	<<if ge .Limit 0>>
	SELECT * FROM(
	<<end>>
	  SELECT <<if .Columns>><<.Columns>><<else>>*<<end>>
	  FROM %s wholesql
	  <<if .Where>>WHERE <<.Where>><<end>>
	  <<if .OrderBy>>ORDER BY <<.OrderBy>><<end>>
	<<if ge .Limit 0>>
	)WHERE ROWNUM<=<<.Limit>>
	<<end>>
<<else>>
	SELECT <<if .Columns>><<.Columns>><<else>>*<<end>>
	FROM %[1]s wholesql
	<<if .Where>>WHERE <<.Where>><<end>>
	<<if .OrderBy>>ORDER BY <<.OrderBy>><<end>>
	<<if ge .Limit 0>>LIMIT <<.Limit>><<end>>
<<end>>`, table.Name())
	return &SqlSelect{
		sql:        strSql,
		Table:      table,
		ManualPage: true,
	}

}
