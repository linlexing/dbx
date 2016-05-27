package dbx

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

//该结构完成表结构的调整，自动处理
//1.表改名
//2.主键更改
//3.字段改名
//4.字段调整
//5.单字段的索引调整
type TableSchema struct {
	OldTable *DBTable
	NewTable *DBTable
}

//检查新表的字段定义是否合法：
//1.字段名（含曾用名）不能重复
//2.必须有主键
func (t *TableSchema) CheckTableColumns(tab *DBTable) error {
	if len(tab.PrimaryKeys()) == 0 {
		return fmt.Errorf("primary key is null")
	}
	uname := map[string]bool{}
	for _, c := range tab.AllField() {
		if _, ok := uname[c.Name]; ok {
			return fmt.Errorf("column:%s dup", c.Name)
		}
		uname[c.Name] = true
		if len(c.FormerName) > 0 {
			for _, fc := range c.FormerName {
				if _, ok := uname[fc]; ok {
					return fmt.Errorf("column:%s former name:%s dup", c.Name, fc)
				}
				uname[fc] = true
			}
		}
	}
	return nil
}
func (t *TableSchema) Update() error {
	//如果没有旧表，则是新增表
	if t.OldTable == nil {
		cols := []string{}
		for _, v := range t.NewTable.AllField() {
			cols = append(cols, v.DBDefine(t.NewTable.Db.DriverName()))
		}
		var strSql string
		if len(t.NewTable.PrimaryKeys()) > 0 {
			strSql = fmt.Sprintf(
				"CREATE TABLE %s(\n%s,\nCONSTRAINT %s_pkey PRIMARY KEY(%s)\n)",
				t.NewTable.Name(), strings.Join(cols, ",\n"), t.NewTable.TableName, strings.Join(t.NewTable.PrimaryKeys(), ","))
		} else {
			strSql = fmt.Sprintf(
				"CREATE TABLE %s(\n%s\n)",
				t.NewTable.Name(), strings.Join(cols, ",\n"))
		}
		if _, err := t.NewTable.Db.Exec(strSql); err != nil {
			return SqlError{strSql, nil, err}
		}
		log.Println(strSql)
		//最后处理索引
		for _, col := range t.NewTable.AllField() {
			if col.Index {
				if err := CreateColumnIndex(t.NewTable.Db, t.NewTable.Name(), col.Name); err != nil {
					return err
				}
			}
		}
		return nil
	} else {
		if err := t.CheckTableColumns(t.NewTable); err != nil {
			return err
		}
		//处理表更名,处理过后，所有后续操作都在新表名上进行
		if t.OldTable.Name() != t.NewTable.Name() {
			if err := TableRename(t.NewTable.Db, t.OldTable.Name(), t.NewTable.Name()); err != nil {
				return err
			}
		}
		pkChanged := false
		//如果主键变更，则需要先除去主键
		if !reflect.DeepEqual(t.OldTable.PrimaryKeys(), t.NewTable.PrimaryKeys()) {
			if err := DropTablePrimaryKey(t.NewTable.Db, t.NewTable.Name()); err != nil {
				return err
			}
			pkChanged = true
		}
		//逐个处理字段，每处理一个字段，旧表字段就标上标记，最后删除没有标记的字段
		oldColumnProcesses := map[string]bool{}
		for _, v := range t.OldTable.Columns() {
			oldColumnProcesses[v] = false
		}
		for _, col := range t.NewTable.AllField() {
			var oldCol *DBTableColumn
			//如果有曾用名，则用曾用名去旧表中获取旧字段
			if len(col.FormerName) > 0 {
				for _, v := range col.FormerName {
					if o := t.OldTable.Field(v); o != nil {
						oldCol = o
						oldColumnProcesses[v] = true
						break
					}
				}
			}
			//如果没有找到曾用名的旧字段，则用当前名称去找旧字段
			if oldCol == nil {
				if o := t.OldTable.Field(col.Name); o != nil {
					oldCol = o
					oldColumnProcesses[col.Name] = true
				}
			}
			if err := t.processColumn(oldCol, col); err != nil {
				return err
			}
		}
		//最后删除没有处理过的旧字段
		deleteCols := []string{}
		for k, prc := range oldColumnProcesses {
			if !prc {
				deleteCols = append(deleteCols, k)
			}
		}
		if len(deleteCols) > 0 {
			if err := TableRemoveColumns(t.NewTable.Db, t.NewTable.Name(), deleteCols); err != nil {
				return err
			}
		}
		//如果主键变过，则新增主键
		if pkChanged {
			if err := AddTablePrimaryKey(t.NewTable.Db, t.NewTable.Name(), t.NewTable.PrimaryKeys()); err != nil {
				return err
			}
		}
	}
	return nil
}
func (t *TableSchema) processColumn(oldCol, newCol *DBTableColumn) error {
	var strSql string
	//如果是新增字段
	if oldCol == nil {
		switch t.NewTable.Db.DriverName() {
		case "postgres", "oci8", "mysql", "sqlite3":
			strSql = fmt.Sprintf("alter table %s add %s", t.NewTable.Name(), newCol.DBDefine(t.NewTable.Db.DriverName()))
		default:
			panic("not impl " + t.NewTable.Db.DriverName())
		}
		if _, err := t.NewTable.Db.Exec(strSql); err != nil {
			return SqlError{strSql, nil, err}
		}
		log.Println(strSql)
		//处理索引
		if newCol.Index {
			if err := CreateColumnIndex(t.NewTable.Db, t.NewTable.Name(), newCol.Name); err != nil {
				return err
			}
			log.Printf("table %s add column index %s\n", t.NewTable.Name(), newCol.Name)
		}
		return nil
	}
	//如果是更名，需要先处理
	if oldCol.Name != newCol.Name {
		switch t.NewTable.Db.DriverName() {
		case "postgres":
			strSql = fmt.Sprintf("alter table %s rename %s to %s", t.NewTable.Name(), oldCol.Name, newCol.Name)
			if _, err := t.NewTable.Db.Exec(strSql); err != nil {
				return SqlError{strSql, nil, err}
			}
			log.Println(strSql)
		case "oci8":
			strSql = fmt.Sprintf("alter table %s rename column %s to %s", t.NewTable.Name(), oldCol.Name, newCol.Name)
			if _, err := t.NewTable.Db.Exec(strSql); err != nil {
				return SqlError{strSql, nil, err}
			}
			log.Println(strSql)
		case "mysql":
			strSql = fmt.Sprintf("alter table %s CHANGE column %s %s", t.NewTable.Name(), oldCol.Name, newCol.DBDefine(t.NewTable.Db.DriverName()))
			if _, err := t.NewTable.Db.Exec(strSql); err != nil {
				return SqlError{strSql, nil, err}
			}
			log.Println(strSql)
		default:
			panic("not impl " + t.NewTable.Db.DriverName())
		}
	}
	//如果字段定义不相等且不是mysql则需要再次修改字段定义
	if !oldCol.Eque(newCol) && t.NewTable.Db.DriverName() != "mysql" {
		switch t.NewTable.Db.DriverName() {
		case "postgres":
			//先改类型,如果都有truetype，则直接判断truetype
			if (oldCol.FetchDriver == newCol.FetchDriver &&
				len(oldCol.TrueType) > 0 && len(newCol.TrueType) > 0 &&
				oldCol.TrueType != newCol.TrueType) ||
				(oldCol.Type != newCol.Type ||
					(oldCol.Type == "STR" &&
						oldCol.MaxLength != newCol.MaxLength)) {
				//去掉最后的notnull
				strSql = fmt.Sprintf(
					"alter table %s ALTER COLUMN %s type %s",
					t.NewTable.Name(), newCol.Name, newCol.DBType(t.NewTable.Db.DriverName()))
				//去掉定义中的字段名，因为中间多了个type字样
				if _, err := t.NewTable.Db.Exec(strSql); err != nil {
					return SqlError{strSql, nil, err}
				}
				log.Println(strSql)
			}
			//再改not null
			if oldCol.Null && !newCol.Null {
				strSql = fmt.Sprintf(
					"alter table %s alter column %s set not null",
					t.NewTable.Name(), newCol.Name)
				if _, err := t.NewTable.Db.Exec(strSql); err != nil {
					return SqlError{strSql, nil, err}
				}
				log.Println(strSql)
			}
			if !oldCol.Null && newCol.Null {
				strSql = fmt.Sprintf(
					"alter table %s alter column %s drop not null",
					t.NewTable.Name(), newCol.Name)
				if _, err := t.NewTable.Db.Exec(strSql); err != nil {
					return SqlError{strSql, nil, err}
				}
				log.Println(strSql)
			}

		case "mysql":
			strSql = fmt.Sprintf("alter table %s MODIFY %s", t.NewTable.Name(), newCol.DBDefine(t.NewTable.Db.DriverName()))
			if _, err := t.NewTable.Db.Exec(strSql); err != nil {
				return SqlError{strSql, nil, err}
			}
			log.Println(strSql)
		case "oci8":
			if oldCol.Null != newCol.Null {
				strSql = fmt.Sprintf("alter table %s MODIFY %s", t.NewTable.Name(), newCol.DBDefineNull(t.NewTable.Db.DriverName()))

			} else {
				strSql = fmt.Sprintf("alter table %s MODIFY %s %s", t.NewTable.Name(), newCol.Name, newCol.DBType(t.NewTable.Db.DriverName()))

			}
			if _, err := t.NewTable.Db.Exec(strSql); err != nil {
				return SqlError{strSql, nil, err}
			}
			log.Println(strSql)

		default:
			panic("not impl " + t.NewTable.Db.DriverName())
		}
	}
	//处理索引,字段更名的操作，oracle、postgres、mysql都是安全的，所以不需处理
	//ref:http://stackoverflow.com/questions/6732896/does-rename-column-take-care-of-indexes
	if oldCol.Index && !newCol.Index {
		//删除索引
		if err := DropColumnIndex(t.NewTable.Db, t.NewTable.Name(), oldCol.IndexName); err != nil {
			return err
		}
		log.Printf("drop table %s column %s index %s\n", t.NewTable.Name(), newCol.Name, oldCol.IndexName)
	} else if !oldCol.Index && newCol.Index {
		//新增索引
		if err := CreateColumnIndex(t.NewTable.Db, t.NewTable.Name(), oldCol.Name); err != nil {
			return err
		}
		log.Printf("create table %s column %s index\n", t.NewTable.Name(), newCol.Name)
	}
	return nil
}
