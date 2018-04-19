package postgres

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/linlexing/dbx/schema"
)

func getdb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=test password=123456 dbname=postgres sslmode=disable")
	return db, err
}
func tableTest() *schema.Table {
	tab := schema.NewTable("test")
	tab.PrimaryKeys = []string{"ID"}
	tab.Columns = []*schema.Column{
		&schema.Column{
			Name: "ID",
			Type: schema.TypeInt,
			Null: false,
		},
		&schema.Column{
			Name: "name",
			Type: schema.TypeString,
			Null: true,
		},
		&schema.Column{
			Name: "birthday",
			Type: schema.TypeDatetime,
			Null: true,
		},
		&schema.Column{
			Name: "salary",
			Type: schema.TypeFloat,
			Null: true,
		},
		&schema.Column{
			Name: "phone",
			Type: schema.TypeBytea,
			Null: true,
		},
	}
	return tab
}

//测试创建表
func TestCreateTable(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	err = new(meta).CreateTable(db, tableTest())
	if err != nil {
		t.Error(err)
	}
	if _, err := db.Exec("drop table test"); err != nil {
		t.Error(err)
	}
}

//测试表是否存在
func TestTableExists(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	b, err := new(meta).TableExists(db, "test")
	if err != nil {
		t.Error(err)
	}
	if b {
		t.Error("test 表不应该存在")
	}

	if err = new(meta).CreateTable(db, tableTest()); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("drop table test"); err != nil {
			t.Error(err)
		}
	}()
	b, err = new(meta).TableExists(db, "test")
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("test 应该存在")
	}

}

//测试新增单字段索引
func TestCreateColumnIndex(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	err = createColumnIndex(db, "test01", "name")
	if err != nil {
		t.Error(err)
	}
}

//测试删除多列
func TestRemoveColumns(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err = db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	cols := []string{"birthday", "phone"}
	err = removeColumns(db, "test01", cols)
	if err != nil {
		t.Error("删除多列测试未通过")
	}
}

//测试表重命名
func TestTableRename(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err = db.Exec("drop table newtest01"); err != nil {
			t.Error(err)
		}
	}()
	err = tableRename(db, "test01", "newtest01")
	if err != nil {
		t.Error("表重命名测试未通过")
	}
}

//测试删除列索引
func TestDropColumIndex(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err = db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	err = createColumnIndex(db, "test01", "name")
	if err != nil {
		t.Error(err)
	}
	err = dropColumnIndex(db, "newtest01", "test01_name_idx")
	if err != nil {
		t.Error("删除列索引测试未通过")
	}

}

//测试删除表主键
func TestDropTablePrimaryKey(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err = db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	err = dropTablePrimaryKey(db, "test01")
	if err != nil {
		t.Error("删除表主键测试未通过")
	}
}

//测试新增主键
func TestAddTablePrimaryKey(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err = db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	err = dropTablePrimaryKey(db, "test01")
	if err != nil {
		t.Error(err)
	}
	pks := []string{"ID", "name"}
	err = addTablePrimaryKey(db, "test01", pks)
	if err != nil {
		t.Error("新增主键测试未通过")
	}
}

//测试获取主键字段
func TestGetPk(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err = db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	_, err = getPk(db, "test01")
	if err != nil {
		t.Error("测试获取逐渐字段未通过")
	}
}

//测试获取表的列
func TestGetTableColumns(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err = db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	_, err = getTableColumns(db, "test", "test")
	if err != nil {
		t.Error(err)
		t.Error("获取表的列测试未通过")
	}
}

//测试获取表索引
// func TestGetTableIndexes(t *testing.T) {
// 	db, err := getdb()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	defer db.Close()
// 	tab := tableTest()
// 	tab.Name = "test01"
// 	if err = new(meta).CreateTable(db, tab); err != nil {
// 		t.Error(err)
// 	}
// 	defer func() {
// 		if _, err = db.Exec("drop table test01"); err != nil {
// 			t.Error(err)
// 		}
// 	}()
// 	err = createColumnIndex(db, "test01", "name")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	_, err = getTableIndexes(db, "test", "test01")
// 	if err != nil {
// 		t.Error("获取表索引测试未通过")
// 	}
// }

// //获取所有列
// func TestGetColumns(t *testing.T) {
// 	db, err := getdb()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	defer db.Close()
// 	tab := tableTest()
// 	tab.Name = "test01"
// 	if err = new(meta).CreateTable(db, tab); err != nil {
// 		t.Error(err)
// 	}
// 	defer func() {
// 		if _, err = db.Exec("drop table test01"); err != nil {
// 			t.Error(err)
// 		}
// 	}()
// 	_, err = getColumns(db, "test", "test01")
// 	if err != nil {
// 		t.Error(err)
// 		t.Error("测试未通过")
// 	}
// }

//测试 OpenTable
func TestOpenTable(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := tableTest()
	tab.Name = "test01"
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err = db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	err = createColumnIndex(db, "test01", "name")
	if err != nil {
		t.Error(err)
	}
	tab, err = new(meta).OpenTable(db, "test01")
	if err != nil {
		t.Error(err)
		t.Error("测试未通过")
	}
	if !tab.ColumnByName("name").Index {
		t.Error("没有索引")
	}
}
func TestArrayColumn(t *testing.T) {
	db, err := getdb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := schema.NewTable("test")
	tab.PrimaryKeys = []string{"ID"}
	tab.Columns = []*schema.Column{
		&schema.Column{
			Name: "ID",
			Type: schema.TypeInt,
			Null: false,
		},
		&schema.Column{
			Name:        "strarr",
			Type:        schema.TypeString,
			Null:        true,
			FetchDriver: "postgres",
			TrueType:    "character varying[]",
		},
	}
	if err = new(meta).CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err = db.Exec("drop table test"); err != nil {
			t.Error(err)
		}
	}()
}
