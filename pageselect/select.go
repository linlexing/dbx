package pageselect

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/data"
	"github.com/linlexing/dbx/render"
	"github.com/linlexing/dbx/scan"
	"github.com/linlexing/dbx/schema"
)

const (
	//OrderAsc 升序
	OrderAsc = "+"
	//OrderDesc 降序
	OrderDesc = "-"
	//OrderNone 无排序
	OrderNone = ""
)

var (
	opeExpre = map[string]PageSelecter{}
)

// Register 注册一个数据库接口，其实现了指定的方法
func Register(driver string, ps PageSelecter) {
	opeExpre[driver] = ps
}

// Find 根据一个驱动找到正确的Ps
func Find(driver string) PageSelecter {
	//sqlite3可能会有不同的变种
	if strings.HasPrefix(driver, "sqlite3") {
		driver = "sqlite3"
	}
	if v, ok := opeExpre[driver]; !ok {
		panic(driver + " not registe pageselectrr")
	} else {
		return v
	}

}

// PageSelect 表示一个select 类，可以附加条件和分页参数,注意，所有用到的列名会被quoted，
// 所以需要保证大小写正确
type PageSelect struct {
	DriverName string
	//是否自动把字段名加上引号
	AutoQuotedColumn bool
	SQL              string
	//非空字段
	NotNullFields []string
	ManualPage    bool
	Conditions    []*SQLCondition
	Columns       []string
	ColumnTypes   ColumnTypes //本来只用名称即可，但go 1.8中Query返回ColumnType的特性，很多驱动还不支持，需要手工传入所有可能列的类型
	Order         []string
	Divide        []string
	Limit         int
	SQLRenderArgs interface{} //sql语句在查询前，还会用template进行一次渲染，这里传入渲染的参数
}

// isNotNullField 返回一个字段是不是非空的，用于生成order by 子句
func (s *PageSelect) isNotNullField(field string) bool {
	for _, one := range s.NotNullFields {
		if one == field {
			return true
		}
	}
	return false
}
func (s *PageSelect) columnName(name string) string {
	if s.AutoQuotedColumn {
		return Find(s.DriverName).QuotedIdentifier(name)
	}
	return name
}

// BuildSQL 构造sql语句，和相应的参数值
func (s *PageSelect) BuildSQLNoRender() (strSQL string, err error) {
	if len(s.SQL) == 0 {
		return "", errors.New("sql is empty")
	}

	whereList := []string{}
	orderList := []string{}
	//全空返回sql
	if len(s.Conditions) == 0 &&
		len(s.Columns) == 0 &&
		len(s.Order) == 0 &&
		s.Limit < 0 {

		if s.ManualPage {

			strSQL, err = renderManualPageSQL(s.DriverName, s.SQL, nil, false, nil, nil, s.Limit, s.AutoQuotedColumn)
			if err != nil {
				fmt.Println("driver:", s.DriverName)
				fmt.Println("renderSQL:", s.SQL)
				fmt.Println("Limit:", s.Limit)

				return
			}
		} else {
			strSQL = s.SQL
		}
		return
	}
	//where
	if len(s.Conditions) > 0 {
		for _, v := range s.Conditions {
			if str := v.BuildWhere(s.DriverName, s.ColumnTypes, s.AutoQuotedColumn); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}
	if len(s.Order) > 0 {
		for _, v := range s.Order {
			if strings.HasPrefix(v, OrderDesc) {
				orderList = append(orderList, Find(s.DriverName).SortByDesc(s.columnName(v[1:]), s.isNotNullField(v[1:])))
			} else if strings.HasPrefix(v, OrderAsc) {
				orderList = append(orderList, Find(s.DriverName).SortByAsc(s.columnName(v[1:]), s.isNotNullField(v[1:])))
			} else {

				orderList = append(orderList, Find(s.DriverName).SortByAsc(s.columnName(v), s.isNotNullField(v)))
			}
		}
		if len(s.Divide) > 0 {
			divideCondition := &SQLCondition{
				Name:  "divide",
				Lines: buildCondition(s.Order, s.Divide),
			}
			if str := divideCondition.BuildWhere(s.DriverName, s.ColumnTypes, s.AutoQuotedColumn); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}

	if s.ManualPage {
		if strSQL, err = renderManualPageSQL(s.DriverName, s.SQL, s.Columns, false, whereList, orderList, s.Limit, s.AutoQuotedColumn); err != nil {
			return
		}
	} else {
		var where, orderby, sel string
		sel = "*"

		//select
		if len(s.Columns) > 0 {
			list := []string{}
			for _, c := range s.Columns {
				list = append(list, s.columnName(c))
			}
			sel = strings.Join(list, ",")
		}
		if len(whereList) > 0 {
			where = " where " + strings.Join(whereList, " "+AND+" ")
		}
		if len(orderList) > 0 {
			orderby = " order by " + strings.Join(orderList, ",")
		}
		if s.Limit >= 0 {
			strSQL = Find(s.DriverName).LimitSQL(sel, s.SQL, where, orderby, s.Limit)
		} else {
			//防止末尾是注释，必须换行
			strSQL = strings.TrimSpace(fmt.Sprintf("select %s from (\n%s\n) wholesql %s%s", sel, s.SQL, where, orderby))
		}
	}
	return
}

// BuildSQL 构造sql语句，和相应的参数值
func (s *PageSelect) BuildSQL() (strSQL string, err error) {
	strSQL, err = s.BuildSQLNoRender()
	if err != nil {
		return
	}
	//在最后返回前，调用render
	strSQL, err = render.RenderSQL(strSQL, s.SQLRenderArgs)

	return
}

func (s *PageSelect) QueryRows(db common.DB) ([]map[string]interface{}, []*scan.ColumnType, error) {
	return s.QueryContext(context.Background(), db)
}

// QueryRows 根据设置返回一页数据
func (s *PageSelect) QueryContext(ctx context.Context, db common.DB) (result []map[string]interface{}, cols []*scan.ColumnType, err error) {
	var strSQL string
	strSQL, err = s.BuildSQL()
	if err != nil {
		return
	}
	var rows *sql.Rows
	if rows, err = db.QueryContext(ctx, strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return
	}
	if cols, err = Find(s.DriverName).ColumnTypes(rows); err != nil {
		return
	}
	//go1.8 可以直接返回各列类型，但是oci8驱动支持有问题，number区分不了整型和浮点
	//所以，还需要从传入的字段类型中取原始类型进行修正
	for _, v := range cols {
		if fc := s.ColumnTypes.byName(v.Name); fc != nil {
			v.Type = fc.Type
		}
	}

	result = []map[string]interface{}{}
	defer rows.Close()
	for rows.Next() {
		var outList []interface{}
		outList, err = scan.TypeScan(rows, cols)
		if err != nil {
			return
		}

		oneRecord := map[string]interface{}{}

		for i, v := range outList {
			oneRecord[cols[i].Name] = v
		}
		result = append(result, oneRecord)

	}
	err = rows.Err()
	return
}

// 渲染sql
func (s *PageSelect) renderSQL() (string, error) {
	return render.RenderSQL(s.SQL, s.SQLRenderArgs)
}

// BuildTotalSQL 如果没有数值字段或者没有记录，则返回空sql
func (s *PageSelect) BuildTotalSQL(cols ...string) (strSQL string, err error) {
	totalCoumns := []string{}

	for _, col := range cols {
		// totalCoumns = append(totalCoumns, fmt.Sprintf("sum(cast(%s(%s,0) as decimal(29,6))) as %[2]s", Find(driver).IsNull(), col))
		totalCoumns = append(totalCoumns, fmt.Sprintf("%s as %s", Find(s.DriverName).Sum(
			s.columnName(col)), s.columnName(col)))
	}

	if len(totalCoumns) == 0 {
		return
	}

	if len(s.SQL) == 0 {
		err = fmt.Errorf("sql is empty")
		return
	}

	var where string
	whereList := []string{}

	//where
	if len(s.Conditions) > 0 {
		for _, v := range s.Conditions {
			if str := v.BuildWhere(s.DriverName, s.ColumnTypes, s.AutoQuotedColumn); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}
	if s.ManualPage {
		if strSQL, err = renderManualPageSQL(s.DriverName, s.SQL, totalCoumns, true, whereList, nil, -1, s.AutoQuotedColumn); err != nil {
			return
		}

	} else {
		if len(whereList) > 0 {
			where = " where " + strings.Join(whereList, " "+AND+" ")
		}
		strSQL = fmt.Sprintf("select %s from (\n%s\n) wholesql %s", strings.Join(totalCoumns, ","), s.SQL, where)
	}
	strSQL, err = render.RenderSQL(strSQL, s.SQLRenderArgs)
	return

}

// BuildRowCountSQL 构造Count的语句
func (s *PageSelect) BuildRowCountSQL() (strSQL string, err error) {

	if len(s.SQL) == 0 {
		return "", errors.New("sql is empty")
	}

	var where string
	whereList := []string{}

	//where
	if len(s.Conditions) > 0 {
		for _, v := range s.Conditions {
			if str := v.BuildWhere(s.DriverName, s.ColumnTypes, s.AutoQuotedColumn); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}
	if s.ManualPage {
		if strSQL, err = renderManualPageSQL(s.DriverName, s.SQL, []string{"COUNT(*)"}, true, whereList, nil, -1, s.AutoQuotedColumn); err != nil {
			return
		}
	} else {
		if len(whereList) > 0 {
			where = " where " + strings.Join(whereList, " "+AND+" ")
		}
		strSQL = fmt.Sprintf("select count(*) from (\n%s\n) wholesql %s", s.SQL, where)
	}
	if s, err := render.RenderSQL(strSQL, s.SQLRenderArgs); err != nil {
		log.Panic(err)
	} else {
		strSQL = s
	}

	return
}

// Total 汇总数值字段
func (s *PageSelect) Total(db common.DB, cols ...string) (result map[string]interface{}, err error) {
	var strSQL string
	if strSQL, err = s.BuildTotalSQL(cols...); err != nil {
		return
	}
	if len(strSQL) == 0 {
		return nil, errors.New("sql is emtpty")
	}
	colTypes := []*scan.ColumnType{}
	for _, v := range cols {
		colTypes = append(colTypes, &scan.ColumnType{
			Name: v,
			Type: schema.TypeFloat,
		})
	}
	var vals []interface{}
	if vals, err = scan.TypeScan(db.QueryRow(strSQL), colTypes); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return
	}
	result = map[string]interface{}{}
	for i, v := range cols {
		result[v] = vals[i]
	}

	return
}

// RowCount 根据现有设置，汇总出记录总数
func (s *PageSelect) RowCount(db common.DB, driver string) (r int64, err error) {
	r = -1
	var strSQL string
	if strSQL, err = s.BuildRowCountSQL(); err != nil {
		return
	}
	r, err = data.AsInt(db, strSQL)
	if err != nil {
		err = common.NewSQLError(err, strSQL)
	}
	return
}

// NewPageSelect 新建一个查询类
func NewPageSelect(driverName, strSQL string, colTypes ColumnTypes, manualPage bool) *PageSelect {

	return &PageSelect{
		DriverName:  driverName,
		SQL:         strSQL,
		ColumnTypes: colTypes,
		ManualPage:  manualPage,
	}

}
