package pageselect

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/linlexing/dbx/common"
	"github.com/linlexing/dbx/render"
	"github.com/linlexing/dbx/scan"
	"github.com/linlexing/dbx/schema"
)

var (
	opeExpre = map[string]PageSelecter{}
)

//Register 注册一个数据库接口，其实现了指定的方法
func Register(driver string, ps PageSelecter) {
	opeExpre[driver] = ps
}

//Find 根据一个驱动找到正确的Ps
func Find(driver string) PageSelecter {
	if v, ok := opeExpre[driver]; !ok {
		panic(driver + " not registe pageselectrr")
	} else {
		return v
	}

}

//OrderColumn 排序的字段
type OrderColumn struct {
	Name  string
	Order string //DESC ASC
	Value interface{}
}

//PageSelect 表示一个select 类，可以附加条件和分页参数
type PageSelect struct {
	SQL           string
	ManualPage    bool
	Conditions    []*SQLCondition
	Columns       []string
	ColumnTypes   ColumnTypes //本来只用名称即可，但go 1.8中Query返回ColumnType的特性，很多驱动还不支持，需要手工传入所有可能列的类型
	Order         []string
	Divide        []string
	Limit         int
	SQLRenderArgs interface{} //sql语句在查询前，还会用template进行一次渲染，这里传入渲染的参数
}

//BuildSQL 构造sql语句，和相应的参数值
func (s *PageSelect) BuildSQL(driver string) (strSQL string, err error) {

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
		var renderSQL string
		renderSQL, err = s.renderSQL()
		if err != nil {
			return
		}
		if s.ManualPage {

			strSQL, err = renderManualPageSQL(driver, renderSQL, nil, nil, nil, s.Limit)
			if err != nil {
				return
			}
		} else {
			strSQL = renderSQL
		}

		return

	}
	//where
	if len(s.Conditions) > 0 {
		for _, v := range s.Conditions {
			if str := v.BuildWhere(driver, s.ColumnTypes); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}
	if len(s.Order) > 0 {
		for _, v := range s.Order {
			if strings.HasPrefix(v, "-") {
				orderList = append(orderList, Find(driver).SortByDesc(v[1:]))

			} else {
				orderList = append(orderList, Find(driver).SortByAsc(v[1:]))

			}
		}
		if len(s.Divide) > 0 {
			divideCondition := &SQLCondition{
				Name:  "divide",
				Lines: buildCondition(s.Order, s.Divide),
			}
			if str := divideCondition.BuildWhere(driver, s.ColumnTypes); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}

	if s.ManualPage {
		if strSQL, err = renderManualPageSQL(driver, s.SQL, s.Columns, whereList, orderList, s.Limit); err != nil {
			return
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
			strSQL = Find(driver).LimitSQL(sel, s.SQL, where, orderby, s.Limit)
		} else {
			strSQL = fmt.Sprintf("select %s from (%s) wholesql %s%s", sel, s.SQL, where, orderby)
		}
	}
	//在最后返回前，调用render
	strSQL, err = render.RenderSQL(strSQL, s.SQLRenderArgs)

	return
}

//QueryRows 根据设置返回一页数据
func (s *PageSelect) QueryRows(driver string, db common.DB) (result []map[string]interface{}, cols []*scan.ColumnType, err error) {
	var strSQL string
	strSQL, err = s.BuildSQL(driver)
	if err != nil {
		return
	}
	var rows *sql.Rows
	if rows, err = db.Query(strSQL); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
		return
	}
	var columns []string
	if columns, err = rows.Columns(); err != nil {
		return
	}
	//go1.8 可以直接返回各列类型，但是驱动还都不支持，以后改进
	cols = []*scan.ColumnType{}
	//先根据预置的列类型清单获取对应的字段类型
	for _, v := range columns {
		col := &scan.ColumnType{
			Name: v,
			Type: schema.TypeString,
		}
		if len(s.ColumnTypes) > 0 {
			if tCol := s.ColumnTypes.byName(col.Name); tCol != nil {
				col.Type = tCol.Type
			}
		}

		cols = append(cols, col)
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
	return
}

//渲染sql
func (s *PageSelect) renderSQL() (string, error) {
	return render.RenderSQL(s.SQL, s.SQLRenderArgs)
}

// BuildTotalSQL 如果没有数值字段或者没有记录，则返回空sql
func (s *PageSelect) BuildTotalSQL(driver string, cols ...string) (strSQL string, err error) {
	totalCoumns := []string{}
	for _, col := range cols {
		totalCoumns = append(totalCoumns, fmt.Sprintf("sum(cast(%s(%s,0) as decimal(29,6))) as %[2]s", Find(driver).IsNull(), col))
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
			if str := v.BuildWhere(driver, s.ColumnTypes); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}
	if s.ManualPage {

		if strSQL, err = renderManualPageSQL(driver, s.SQL, totalCoumns, whereList, nil, -1); err != nil {
			return
		}

	} else {
		if len(whereList) > 0 {
			where = " where " + strings.Join(whereList, " and ")
		}
		strSQL = fmt.Sprintf("select %s from (%s) wholesql %s", strings.Join(totalCoumns, ","), s.SQL, where)
	}
	strSQL, err = render.RenderSQL(strSQL, s.SQLRenderArgs)
	return

}

//BuildRowCountSQL 构造Count的语句
func (s *PageSelect) BuildRowCountSQL(driver string) (strSQL string, err error) {

	if len(s.SQL) == 0 {
		return "", errors.New("sql is empty")
	}

	var where string
	whereList := []string{}

	//where
	if len(s.Conditions) > 0 {
		for _, v := range s.Conditions {
			if str := v.BuildWhere(driver, s.ColumnTypes); len(str) > 0 {
				whereList = append(whereList, "("+str+")")
			}
		}
	}
	if s.ManualPage {
		if strSQL, err = renderManualPageSQL(driver, s.SQL, []string{"COUNT(*)"}, whereList, nil, -1); err != nil {
			return
		}
	} else {
		if len(whereList) > 0 {
			where = " where " + strings.Join(whereList, " and ")
		}
		strSQL = fmt.Sprintf("select count(*) from (%s) wholesql %s", s.SQL, where)
	}
	if s, err := render.RenderSQL(strSQL, s.SQLRenderArgs); err != nil {
		log.Panic(err)
	} else {
		strSQL = s
	}

	return
}

//Total 汇总数值字段
func (s *PageSelect) Total(db common.DB, driver string, cols ...string) (result map[string]interface{}, err error) {
	var strSQL string
	if strSQL, err = s.BuildTotalSQL(driver, cols...); err != nil {
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

//RowCount 根据现有设置，汇总出记录总数
func (s *PageSelect) RowCount(db common.DB, driver string) (r int64, err error) {
	r = -1
	var strSQL string
	if strSQL, err = s.BuildRowCountSQL(driver); err != nil {
		return
	}

	if err = db.QueryRow(strSQL).Scan(&r); err != nil {
		err = common.NewSQLError(err, strSQL)
		log.Println(err)
	}
	return

}

//NewPageSelect 新建一个查询类
func NewPageSelect(strSQL string, colTypes ColumnTypes, manualPage bool) *PageSelect {

	return &PageSelect{
		SQL:         strSQL,
		ColumnTypes: colTypes,
		ManualPage:  manualPage,
	}

}
