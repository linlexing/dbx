package data

import (
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

// UpdateLogNewColumnName 生成记录新值的临时列名
func UpdateLogNewColumnName(col string) string {
	return col + "_2"
}

// UpdateLogOldColumnName 生成记录旧值的临时列名
func UpdateLogOldColumnName(col string) string {
	return col + "_1"
}

// ====================== 内部辅助方法 ======================

// validatePrimaryKeyMatch 校验主键字段与 DataUniqueFields 是否匹配
func (u *Update) validatePrimaryKeyMatch() error {
	if len(u.Table.PrimaryKeys) == 0 {
		return fmt.Errorf("table %s has no primary key", u.Table.FullName())
	}
	if len(u.Table.PrimaryKeys) != len(u.DataUniqueFields) {
		return fmt.Errorf("primary key count mismatch: table has %d, DataUniqueFields has %d",
			len(u.Table.PrimaryKeys), len(u.DataUniqueFields))
	}
	return nil
}

// checkSets 校验 UpdateSet 的 Value 格式是否合法（保持原有逻辑）
func (u *Update) checkSets() error {
	for _, set := range u.Sets {
		k := set.Column
		v := set.Value
		if len(v) == 0 {
			return fmt.Errorf("%s the value is empty", k)
		}
		if v == "`" {
			return fmt.Errorf("%s the value is %s", k, v)
		}
		if v[0] == '`' && v[len(v)-1] != '`' {
			return fmt.Errorf("%s the value is %s", k, v)
		}
	}
	return nil
}

// parseSetExpression 解析单个 UpdateSet 的值，返回：
// expr: 表达式的 SQL 片段（如果是反引号包裹，则返回内部表达式；否则返回 "?"）
// val:  如果是普通值，返回该值用于参数化；否则返回 nil
// isParam: 是否需要使用参数占位符
func (u *Update) parseSetExpression(v string) (expr string, val interface{}, isParam bool) {
	if len(v) >= 2 && v[0] == '`' && v[len(v)-1] == '`' {
		expr = v[1 : len(v)-1]
		if expr == "" {
			expr = "NULL"
		}
		return expr, nil, false
	}
	return "?", v, true
}

// buildJoinCondition 构建主表与数据源的 JOIN 条件（用于 exists 子查询）
func (u *Update) buildJoinCondition(dataAlias string) []string {
	conds := make([]string, len(u.Table.PrimaryKeys))
	for i, field := range u.Table.PrimaryKeys {
		conds[i] = fmt.Sprintf("%s.%s = %s.%s",
			u.Table.FullName(), field, dataAlias, u.DataUniqueFields[i])
	}
	return conds
}

// renderTemplate 渲染 SQL 模板，支持额外的自定义函数
func (u *Update) renderTemplate(tmpl string, extraFuncs template.FuncMap) (string, error) {
	if tmpl == "" {
		return "", nil
	}
	funcMap := template.FuncMap{}
	for k, v := range u.SQLRenderFunc {
		funcMap[k] = v
	}
	for k, v := range extraFuncs {
		funcMap[k] = v
	}
	return render.RenderSQL(tmpl, u.SQLRenderArgs, funcMap)
}

// execBeforeSQL 执行 BeforeSQL 语句并返回影响行数（通用版本）
func (u *Update) execBeforeSQL(db common.DB, extraFuncs template.FuncMap, param ...interface{}) (int64, error) {
	if u.BeforeSQL == "" {
		return 0, nil
	}
	sqlStr, err := u.renderTemplate(u.BeforeSQL, extraFuncs)
	if err != nil {
		return 0, err
	}
	res, err := db.Exec(sqlStr, param...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// ====================== 主要导出方法 ======================

// ExecLog 执行一个更新操作，并返回影响的行数，记录日志到临时表中
// 临时表结构为：主键，更新字段1，更新字段1_1,...,其中_1记录旧值,_2记录新值
func (u *Update) ExecLog(nt NotifyAndKill, db common.DB, tmpTableName string, param ...interface{}) (icount int64, err error) {
	// 基础校验
	if err = u.validatePrimaryKeyMatch(); err != nil {
		return
	}
	if err = u.checkSets(); err != nil {
		return
	}

	tmpTable := schema.NewTable(tmpTableName)
	// 先增加主键列
	for _, pk := range u.Table.PrimaryKeys {
		tmpTable.Columns = append(tmpTable.Columns, u.Table.ColumnByName(pk).Clone())
	}
	// 再增加每个更新字段的新旧值列
	for _, set := range u.Sets {
		col := u.Table.ColumnByName(set.Column).Clone()
		col.Name = UpdateLogNewColumnName(col.Name)
		tmpTable.Columns = append(tmpTable.Columns, col)

		colOld := u.Table.ColumnByName(set.Column).Clone()
		colOld.Name = UpdateLogOldColumnName(colOld.Name)
		tmpTable.Columns = append(tmpTable.Columns, colOld)
	}
	tmpTable.PrimaryKeys = u.Table.PrimaryKeys

	nt.AddProgress(fmt.Sprintf("创建临时表[%s]...", tmpTableName), nil)
	if err = tmpTable.Update(db.DriverName(), db); err != nil {
		return
	}

	// 构建插入临时表的列与数据
	insertCols := make([]string, len(u.Table.PrimaryKeys))
	selectCols := make([]string, len(u.Table.PrimaryKeys))
	for i, pk := range u.Table.PrimaryKeys {
		insertCols[i] = pk
		selectCols[i] = pk
	}

	dataAlias := "datasrc_"
	setsMap := []ColMap{} // 用于后续更新语句的字段映射

	for _, set := range u.Sets {
		expr, val, isParam := u.parseSetExpression(set.Value)
		var sqlValue string
		if isParam {
			// 注意：此处构造的 SQL 是直接拼接进 SELECT 列表的，不支持参数化占位符，
			// 保持原有逻辑：将字符串值转义后直接拼接
			sqlValue = "'" + strings.ReplaceAll(val.(string), "'", "''") + "'"
		} else {
			sqlValue = expr
		}

		newColName := UpdateLogNewColumnName(set.Column)
		oldColName := UpdateLogOldColumnName(set.Column)

		insertCols = append(insertCols, newColName)
		selectCols = append(selectCols, sqlValue)

		insertCols = append(insertCols, oldColName)
		selectCols = append(selectCols, set.Column)

		setsMap = append(setsMap, ColMap{Dest: set.Column, Src: newColName})
	}

	// 构建 WHERE 条件（主表与数据源关联）
	whereConds := []string{}
	if u.AdditionWhere != "" {
		whereConds = append(whereConds, "("+u.AdditionWhere+")")
	}
	whereConds = append(whereConds, u.buildJoinCondition(dataAlias)...)

	// 将数据插入临时表
	strSQL := fmt.Sprintf("INSERT INTO %s(%s) SELECT %s FROM %s WHERE EXISTS(SELECT 1 FROM (%s) %s WHERE %s)",
		tmpTableName,
		strings.Join(insertCols, ","),
		strings.Join(selectCols, ","),
		u.Table.FullName(),
		u.DataSQL,
		dataAlias,
		strings.Join(whereConds, " AND "))

	sr, err := db.Exec(strSQL)
	if err != nil {
		return
	}
	ic, err := sr.RowsAffected()
	if err != nil {
		return
	}
	nt.AddProgress(fmt.Sprintf("%d 条数据插入临时表[%s]", ic, tmpTableName), nil)

	// --- 新增：删除不需要更新的记录（新旧值完全相同）---
	if len(u.Sets) > 0 {
		var delConds []string
		for _, set := range u.Sets {
			newCol := UpdateLogNewColumnName(set.Column)
			oldCol := UpdateLogOldColumnName(set.Column)
			delConds = append(delConds, fmt.Sprintf("(%s = %s OR (%s IS NULL AND %s IS NULL))", newCol, oldCol, newCol, oldCol))
		}
		delSQL := fmt.Sprintf("DELETE FROM %s WHERE %s", tmpTableName, strings.Join(delConds, " AND "))
		delRes, err := db.Exec(delSQL)
		if err != nil {
			return 0, err
		}
		delCount, _ := delRes.RowsAffected()
		if delCount > 0 {
			nt.AddProgress(fmt.Sprintf("跳过 %d 条无需更新的记录，已从临时表中删除", delCount), nil)
		}

		// 检查临时表是否还有剩余记录
		var remainCount int64
		row := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tmpTableName))
		if err := row.Scan(&remainCount); err != nil {
			return 0, err
		}
		if remainCount == 0 {
			nt.AddProgress("没有需要实际更新的记录，操作结束", nil)
			return 0, nil
		}
	}
	// --- 新增结束 ---

	// 构建主键映射（用于更新语句的 JOIN 条件）
	pksList := make([]ColMap, len(u.Table.PrimaryKeys))
	for i, pk := range u.Table.PrimaryKeys {
		pksList[i] = ColMap{Dest: pk, Src: pk}
	}

	// 生成更新语句
	strSQL = Find(db.DriverName()).UpdateFrom(
		u.Table.FullName(),
		"SELECT * FROM "+tmpTableName,
		u.AdditionSet,
		pksList,
		setsMap,
	)

	// 渲染更新 SQL 模板
	strSQL, err = u.renderTemplate(strSQL, nil)
	if err != nil {
		return
	}

	// 执行 BeforeSQL
	beforeCount, err := u.execBeforeSQL(db, template.FuncMap{
		"DataSQL":      func() string { return u.DataSQL },
		"tmpTableName": func() string { return tmpTableName },
	}, param...)
	if err != nil {
		return
	}
	if beforeCount > 0 {
		nt.AddProgress(fmt.Sprintf("%d 条记录受执行Before脚本影响", beforeCount), nil)
	}

	// 执行实际更新
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

// Exec 执行一个更新操作，并返回影响的行数
func (u *Update) Exec(db common.DB, param ...interface{}) (icount int64, err error) {
	// 基础校验
	if err = u.validatePrimaryKeyMatch(); err != nil {
		return
	}
	if err = u.checkSets(); err != nil {
		return
	}

	// 构建 SET 子句
	sets := []string{}
	setVals := []interface{}{}
	for _, set := range u.Sets {
		k := set.Column
		expr, val, isParam := u.parseSetExpression(set.Value)
		if isParam {
			sets = append(sets, fmt.Sprintf("%s=?", k))
			setVals = append(setVals, val)
		} else {
			sets = append(sets, fmt.Sprintf("%s=%s", k, expr))
		}
	}
	if u.AdditionSet != "" {
		sets = append(sets, u.AdditionSet)
	}

	// 构建 WHERE 条件
	dataAlias := "datasql_"
	whereConds := []string{}
	if u.AdditionWhere != "" {
		whereConds = append(whereConds, u.AdditionWhere)
	}
	whereConds = append(whereConds, u.buildJoinCondition(dataAlias)...)

	// 生成 SQL
	strSQL := fmt.Sprintf("UPDATE %s SET %s WHERE EXISTS(SELECT 1 FROM (%s) %s WHERE %s)",
		u.Table.FullName(),
		strings.Join(sets, ","),
		u.DataSQL,
		dataAlias,
		strings.Join(whereConds, " AND "),
	)

	// 渲染 SQL 模板
	strSQL, err = u.renderTemplate(strSQL, nil)
	if err != nil {
		return
	}

	// 执行 BeforeSQL
	if _, err = u.execBeforeSQL(db, template.FuncMap{
		"DataSQL": func() string { return u.DataSQL },
	}, param...); err != nil {
		return
	}

	// 执行更新
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
		},
	})
}
