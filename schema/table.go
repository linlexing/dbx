package schema

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/linlexing/dbx/common"
)

const (
	comment = "--"
)

var (
	columnReg      *regexp.Regexp
	columnTrueType *regexp.Regexp
)

func init() {
	r, err := regexp.Compile(`(?i)([\p{Han}_a-zA-Z0-9]+)(\s+bytea|\s+date|\s+float|\s+int|\s+str\([0-9]+\)|\s+str|)(\s+null|\s+not null|)(\s+uindex|\s+index|)`)
	if err != nil {
		panic(err)
	}
	columnReg = r
	r, err = regexp.Compile(`(?i)(postgres|oci8|sqlite3|mysql)\s+(.+)`)
	if err != nil {
		panic(err)
	}
	columnTrueType = r
}

//Table 代表数据库中一个物理表
type Table struct {
	Schema      string
	Name        string
	Columns     []*Column
	FormerName  []string
	PrimaryKeys []string
}

//NewTable 返回一个新表，名称会自动依据句点拆分为shcema和name
func NewTable(name string) *Table {
	rev := new(Table)
	ns := strings.Split(name, ".")
	if len(ns) > 1 {
		rev.Schema = ns[0]
		rev.Name = ns[1]
	} else {
		rev.Name = name
	}
	return rev
}

//名称和所有字段的名称转换成大写
func (t *Table) ToUpper() {

	t.Name = strings.ToUpper(t.Name)
	for i, v := range t.PrimaryKeys {
		t.PrimaryKeys[i] = strings.ToUpper(v)
	}
	for _, v := range t.Columns {
		v.Name = strings.ToUpper(v.Name)
	}

}

//FullName 返回全名称，包括schema
func (t *Table) FullName() string {
	if len(t.Schema) > 0 {
		return t.Schema + "." + t.Name
	}
	return t.Name
}
func (t *Table) check() error {
	cm := map[string]bool{}
	for i, c := range t.Columns {
		if len(c.Name) == 0 {
			return fmt.Errorf("column %d name is empty", i)
		}
		if _, ok := cm[strings.ToUpper(c.Name)]; ok {
			return fmt.Errorf("column %d name [%s] is dup", i, c.Name)
		}
		cm[strings.ToUpper(c.Name)] = true
	}
	for _, c := range t.PrimaryKeys {
		if _, ok := cm[strings.ToUpper(c)]; !ok {
			return fmt.Errorf("primary key %s not found", c)
		}
	}
	return nil
}

//Create 创建一个新表，如果表已经存在，则失败
func (t *Table) Create(driver string, db common.DB) error {
	if err := t.check(); err != nil {
		return err
	}

	sch, err := t.extract(driver, db)
	if err != nil {
		return err
	}
	if sch.oldTable != nil {
		return fmt.Errorf("the table %s exists,can't recreate", t.Name)
	}
	list, err := sch.extract()
	if err != nil {
		return err
	}
	return common.BatchRunAndPrint(db, list)

}

//Update 更新一个表结构到数据库中
func (t *Table) Update(driver string, db common.DB) error {
	list, err := t.Extract(driver, db)
	if err != nil {
		return err
	}
	return common.BatchRunAndPrint(db, list)
}
func (t *Table) extract(driver string, db common.DB) (*tableSchema, error) {
	mt := Find(driver)
	sch := &tableSchema{
		newTable: t,
		mt:       mt,
		db:       db,
	}
	if len(t.FormerName) > 0 {
		//如果有曾用名，则需验证曾用名不能和现有名称重复
		uname := map[string]bool{
			strings.ToUpper(t.FullName()): true,
		}
		for _, v := range t.FormerName {
			v = strings.ToUpper(v)
			if _, ok := uname[v]; ok {
				return nil, fmt.Errorf("FormerName:%s dup", v)
			}
		}
		//并根据曾用名去获取之前的表结构
		for _, v := range t.FormerName {
			if b, _ := mt.TableExists(db, v); b {
				oldTable, err := mt.OpenTable(db, v)
				if err != nil {
					return nil, err
				}
				sch.oldTable = oldTable
				break
			}
		}
	}
	//如果曾用名的表找不到，则就用本来的名称，说明不需改名
	if sch.oldTable == nil {
		b, err := mt.TableExists(db, t.FullName())
		if err != nil {
			return nil, err
		}
		if b {
			sch.oldTable, err = mt.OpenTable(db, t.FullName())
			if err != nil {
				return nil, err
			}
		}
	}
	return sch, nil
}

//Extract 提取更新一个表的结构所需要的SQL语句清单
func (t *Table) Extract(driver string, db common.DB) ([]string, error) {
	if err := t.check(); err != nil {
		return nil, err
	}

	sch, err := t.extract(driver, db)
	if err != nil {
		return nil, err
	}
	return sch.extract()
}

//大小写不敏感
func (t *Table) findColumnAnyName(names ...string) *Column {
	//用map作为检索索引
	idx := map[string]bool{}
	for _, oneName := range names {
		idx[strings.ToUpper(oneName)] = true
	}
	for _, col := range t.Columns {
		if _, ok := idx[strings.ToUpper(col.Name)]; ok {
			return col
		}
	}
	return nil
}

//ColumnByName 根据一个名称返回一个字段，如果没有找到，返回nil
func (t *Table) ColumnByName(name string) *Column {
	for _, col := range t.Columns {
		if strings.ToUpper(name) == strings.ToUpper(col.Name) {
			return col
		}
	}
	return nil
}

//DefineScript 采用脚本的方式定义表，如下：
//  a str(3) not null
//  b int	--注释
//  c date not null index
//  primary key(a,c)
func (t *Table) DefineScript(src string) error {

	pks := []string{}
	columns := []*colDef{}
	for i, line := range strings.Split(strings.Replace(src, "\r\n", "\n", -1), "\n") {
		line = strings.ToUpper(strings.TrimSpace(line))
		if len(line) == 0 {
			continue
		}
		//如果是主键定义
		if strings.HasPrefix(line, "PRIMARY KEY(") {
			for _, v := range strings.Split(line[12:len(line)-1], ",") {
				pks = append(pks, strings.TrimSpace(v))
			}
		} else {
			col, err := columnDefine(line, "", "")
			if err != nil {
				return fmt.Errorf("line %d,%s", i, err)
			}
			columns = append(columns, col)
		}
	}
	t.Columns, t.PrimaryKeys = columnsDefine(columns)
	//如果单独定义了主键，则覆盖之前的主键定义
	if len(pks) > 0 {
		t.PrimaryKeys = pks
		//将所有的主键设置成NOT NULL
		for _, v := range t.PrimaryKeys {
			t.ColumnByName(v).Null = false
		}
	}
	return nil
}
