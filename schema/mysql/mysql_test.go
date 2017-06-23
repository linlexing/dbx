package mysql

import (
	"database/sql"
	"testing"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/linlexing/dbx/schema"
)

func getmysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	return db, err
}
func tableTest() *schema.Table {
	tab := schema.NewTable("test")
	tab.PrimaryKeys = []string{"ID"}
	tab.Columns = []*schema.Column{
		&schema.Column{
			Name:      "ID",
			Type:      schema.TypeInt,
			MaxLength: 20,
			Null:      false,
		},
		&schema.Column{
			Name:      "name",
			Type:      schema.TypeString,
			MaxLength: 200,
			Index:     true,
			Null:      true,
		},
		&schema.Column{
			Name:      "birthday",
			Type:      schema.TypeDatetime,
			MaxLength: 200,
			Null:      true,
		},
		&schema.Column{
			Name:      "salary",
			Type:      schema.TypeFloat,
			MaxLength: 200,
			Null:      true,
		},
		&schema.Column{
			Name:      "phone",
			Type:      schema.TypeBytea,
			MaxLength: 200,
			Null:      true,
		},
	}
	return tab
}
func tableTest01() *schema.Table {
	tab01 := schema.NewTable("test01")
	tab01.Columns = []*schema.Column{
		&schema.Column{
			Name:      "ID",
			Type:      schema.TypeInt,
			MaxLength: 20,
			Null:      false,
		},
		&schema.Column{
			Name:      "name",
			Type:      schema.TypeString,
			MaxLength: 200,
			Null:      true,
		},
	}
	return tab01
}

//创建表
func TestCreateTable(t *testing.T) {
	db, err := getmysql()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	err = new(meta).CreateTable(db, tableTest())
	if err != nil {
		t.Error(err)
		t.Error("测试不通过")
	}
	err = new(meta).CreateTable(db, tableTest01())
	if err != nil {
		t.Error(err)
		t.Error("测试不通过")
	}
	defer func() {
		if _, err := db.Exec("drop table test"); err != nil {
			t.Error(err)
		}
		if _, err := db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()

}

//测试表是否存在
func TestTabExists(t *testing.T) {
	db, err := getmysql()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	var ok bool
	ok, err = new(meta).TableExists(db, "test01")
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Error("表不应该存在")
	}
	tab := tableTest()
	tab.Name = "test01"
	err = new(meta).CreateTable(db, tab)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	ok, err = new(meta).TableExists(db, "test01")
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("表应该存在")
	}
}

//表名
func TestTableNames(t *testing.T) {
	db, err := getmysql()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	err = new(meta).CreateTable(db, tab)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	tab.Name = "test02"
	err = new(meta).CreateTable(db, tab)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("drop table test02"); err != nil {
			t.Error(err)
		}
	}()
	var names []string
	names, err = new(meta).TableNames(db)
	if err != nil {
		t.Error(err)
	}
	if names[0] != "test01" || names[1] != "test02" {
		t.Error("测试未通过")
	}
}

//复制表
func TestCreateTableAs(t *testing.T) {
	db, err := getmysql()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	err = new(meta).CreateTable(db, tab)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	strSQL := "select * from test01"
	var pks []string
	pks = []string{"ID", "name"}
	err = new(meta).CreateTableAs(db, "test02", strSQL, pks)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("drop table test02"); err != nil {
			t.Error(err)
		}
	}()
}

//ChangeTable
//这是一个执行失败的测试用例
func TestChangeTable(t *testing.T) {
	db, err := getmysql()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "oldtest"
	err = new(meta).CreateTable(db, tab)
	if err != nil {
		t.Error(err)
	}
	// defer func() {
	// 	if _, err := db.Exec("drop table newtest"); err != nil {
	// 		t.Error(err)
	// 	}
	// }()
	// tab, err = new(meta).OpenTable(db, "oldtest")
	// if err != nil {
	// 	t.Error(err)
	// }
	change := schema.TableSchemaChange{
		OldName:      "oldtest",
		NewName:      "newtest",
		PKChange:     true,
		PK:           []string{"name"},
		RemoveFields: []string{"birthday"},
	}
	// change.OriginFields = tab.Columns
	change.OriginFields = []*schema.Column{
		&schema.Column{
			Name:      "ID",
			Type:      schema.TypeInt,
			MaxLength: 20,
			Null:      false,
		},
		&schema.Column{
			Name:      "name",
			Type:      schema.TypeString,
			MaxLength: 200,
			Null:      true,
		},
	}
	change.ChangeFields = []*schema.ChangedField{
		{
			OldField: &schema.Column{
				Name:      "ID",
				Type:      schema.TypeInt,
				MaxLength: 20,
				Index:     false,
				Null:      false,
			},
			NewField: &schema.Column{
				Name:      "ID",
				Type:      schema.TypeInt,
				MaxLength: 30,
				Index:     true,
				Null:      false,
			},
		},
		{
			OldField: &schema.Column{
				Name:      "name",
				Type:      schema.TypeString,
				MaxLength: 200,
				Index:     true,
				IndexName: "ioldtestname",
				Null:      true,
			},
			NewField: &schema.Column{
				Name:      "newname",
				Type:      schema.TypeString,
				MaxLength: 300,
				Index:     false,
				Null:      false,
			},
		},
	}
	if err = new(meta).ChangeTable(db, &change); err != nil {
		t.Error(err)
	}
}
