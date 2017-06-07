package schema

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/linlexing/dbx/common"
)

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

//FullName 返回全名称，包括schema
func (t *Table) FullName() string {
	if len(t.Schema) > 0 {
		return t.Schema + "." + t.Name
	}
	return t.Name
}

//Update 更新一个表的结构至数据库中，会自动处理表改名、字段改名以及字段修改、索引修改等操作,
//先自动去数据库取出旧表结构
func (t *Table) Update(driver string, db common.DB) error {
	mt := Find(driver)
	sch := &tableSchema{
		newTable: t,
		mt:       mt,
		db:       db,
	}
	if len(t.FormerName) > 0 {
		//如果有曾用名，则需验证曾用名不能和现有名称重复
		uname := map[string]bool{
			t.FullName(): true,
		}
		for _, v := range t.FormerName {
			if _, ok := uname[v]; ok {
				return fmt.Errorf("FormerName:%s dup", v)
			}
		}
		//并根据曾用名去获取之前的表结构
		for _, v := range t.FormerName {
			if b, _ := mt.TableExists(db, v); b {
				oldTable, err := mt.OpenTable(db, v)
				if err != nil {
					return nil
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
			return err
		}
		if b {
			sch.oldTable, err = mt.OpenTable(db, t.FullName())
			if err != nil {
				return err
			}
		}
	}
	return sch.update()
}
func (t *Table) findColumnAnyName(names ...string) *Column {
	//用map作为检索索引
	idx := map[string]bool{}
	for _, oneName := range names {
		idx[oneName] = true
	}
	for _, col := range t.Columns {
		if _, ok := idx[col.Name]; ok {
			return col
		}
	}
	return nil
}

//ColumnByName 根据一个名称返回一个字段，如果没有找到，返回nil
func (t *Table) ColumnByName(name string) *Column {
	for _, col := range t.Columns {
		if name == col.Name {
			return col
		}
	}
	return nil
}

//DefineScript 采用脚本的方式定义表，如下：
//  a str(3) not null
//  b int
//  c date not null index
//  primary key(a,c)
func (t *Table) DefineScript(src string) error {
	lineReg, err := regexp.Compile(`(?i)([\p{Han}_a-zA-Z0-9]+)(\s+bytea|\s+date|\s+float|\s+int|\s+str\([0-9]+\)|\s+str|)(\s+null|\s+not null|)(\s+index|)`)
	if err != nil {
		return err
	}
	pks := []string{}
	columns := []*Column{}
	var prevColumn *Column
	for i, line := range strings.Split(strings.Replace(src, "\r\n", "\n", -1), "\n") {

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//如果是主键定义
		if strings.HasPrefix(line, "primary key(") {
			for _, v := range strings.Split(line[12:len(line)-1], ",") {
				pks = append(pks, strings.TrimSpace(v))
			}
		} else {
			lineList := lineReg.FindStringSubmatch(line)
			if len(lineList) == 0 {
				return fmt.Errorf("line %d:%s error", i, line)
			}
			//第一个是整行，需要去除
			lineList = lineList[1:]
			if len(lineList) == 0 {
				return fmt.Errorf("line %d:%s error", i, line)
			}
			colName := lineList[0]
			if len(strings.TrimSpace(lineList[1])) == 0 {
				//如果只有列名，则自动从上一个字段取出数据类型等定义
				if prevColumn == nil {
					return fmt.Errorf("line %d:%s not data type", i, line)
				}
				col := prevColumn.Clone()
				col.Name = colName
				columns = append(columns, col)
				prevColumn = col
				continue

			}
			dataType := strings.ToUpper(strings.TrimSpace(lineList[1]))
			notNull := false
			index := false
			var maxLength int64 = -1
			if len(lineList) > 2 {
				switch str := strings.ToLower(strings.TrimSpace(lineList[2])); str {
				case "not null":
					notNull = true
				case "null":
					notNull = false
				case "":
				default:
					return fmt.Errorf("line %d:%s ,error define %s", i, line, str)
				}
			}
			if len(lineList) > 3 {
				switch str := strings.ToLower(strings.TrimSpace(lineList[3])); str {
				case "index":
					index = true
				case "":
				default:
					return fmt.Errorf("line %d:%s ,error define %s", i, line, str)
				}
			}
			if strings.HasPrefix(dataType, "str(") {
				maxLength, err = strconv.ParseInt(dataType[4:len(dataType)-1], 10, 64)
				if err != nil {
					return err
				}
				dataType = "STR"
			} else {
				dataType = strings.ToUpper(dataType)
			}
			prevColumn = &Column{
				Name:      colName,
				Type:      ParseDataType(dataType),
				MaxLength: int(maxLength),
				Null:      !notNull,
				Index:     index,
			}
			columns = append(columns, prevColumn)
		}
	}
	t.Columns = columns
	t.PrimaryKeys = pks
	return nil
}
