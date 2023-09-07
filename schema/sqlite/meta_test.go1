package sqlite

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/linlexing/dbx/schema"
	// _ "github.com/mattn/go-sqlite3"
)

type testDB struct {
	fileName string

	DB *sqlx.DB
}

var sqdb = "e:\\temp\\dump932726563"

//var sqdb = "E:\\SQLite\\test.db"

// create
func createTestDB() *testDB {
	rev := testDB{fileName: sqdb}

	/*	if _, err := os.Stat(rev.fileName); err == nil {
		os.Remove(rev.fileName)
	}*/
	db, err := sqlx.Open("sqlite3", rev.fileName)
	if err != nil {
		panic(err)
	}
	rev.DB = db
	return &rev
}
func (t *testDB) Close() {
	t.DB.Close()
	os.Remove(t.fileName)
}

func Test_CreateTable(t *testing.T) {
	testDB := createTestDB()
	db := testDB.DB
	defer testDB.Close()
	tab := schema.NewTable("DEPT")
	tab.Columns = []*schema.Column{
		&schema.Column{Name: "CODE", Type: schema.TypeString, MaxLength: 50, Null: false},
		&schema.Column{Name: "NAME", Type: schema.TypeString, MaxLength: 50, Null: false},
		&schema.Column{Name: "DLEVEL", Type: schema.TypeInt, Null: false},
	}
	tab.PrimaryKeys = []string{"CODE"}
	if err := tab.Update("sqlite3", db); err != nil {
		t.Error(err)
	}
}
func Test_exesql(t *testing.T) {
	testDB := createTestDB()
	db := testDB.DB
	defer testDB.Close()
	if _, err := db.Exec(
		`CREATE TABLE dept(
CODE TEXT(50) NOT NULL,
NAME TEXT(50) NOT NULL,
DLEVEL INTEGER NOT NULL,
CONSTRAINT dept_pkey PRIMARY KEY(CODE))`); err != nil {
		t.Fatal(err)
	}

}

func Test_TableNames(t *testing.T) {
	testDB := createTestDB()
	db := testDB.DB
	defer testDB.Close()
	if _, err := db.Exec("create table aaa(a varchar(200) primary key,b integer)"); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if _, err := db.Exec("drop table aaa"); err != nil {
			t.Error(err)
		}
	}()
	list, err := new(meta).TableNames(db)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) == 0 {
		t.Fatal("no list")
	}
	if list[0] != "aaa" {
		t.Fatal("not aaa", list[0])
	}
}

func Test_CreateTableAs(t *testing.T) {
	testDB := createTestDB()
	db := testDB.DB
	defer testDB.Close()
	type testStr struct {
		A string
		B int64
	}
	v := testStr{
		A: "1234",
		B: 111,
	}
	vNew := testStr{}
	var err error
	if _, err = db.Exec("create table aaa(a varchar(200) primary key,b integer)"); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if _, err := db.Exec("drop table aaa"); err != nil {
			t.Error(err)
		}
	}()
	if _, err = db.NamedExec("insert into aaa(a,b)values(:a,:b)", &v); err != nil {

		t.Fatal(err)
	}
	if err = new(meta).CreateTableAs(db, "bbb", "select * from aaa", []string{"a"}); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if _, err := db.Exec("drop table bbb"); err != nil {
			t.Error(err)
		}
	}()
	if err = db.Get(&vNew, "select * from bbb"); err != nil {
		t.Fatal(err)
	}
	if v != vNew {
		t.Error("not equ")
	}

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
	tab01.PrimaryKeys = []string{"name"}
	tab01.Columns = []*schema.Column{
		&schema.Column{
			Name:      "ID",
			Type:      schema.TypeInt,
			MaxLength: 20,
			Index:     true,
			Null:      false,
		},
		&schema.Column{
			Name:      "name",
			Type:      schema.TypeString,
			MaxLength: 200,
			Index:     false,
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
	}
	return tab01
}
func tableTest02() *schema.Table {
	tab02 := schema.NewTable("test02")
	tab02.PrimaryKeys = []string{"ID"}
	tab02.Columns = []*schema.Column{
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
		&schema.Column{
			Name:      "sex",
			Type:      schema.TypeBytea,
			MaxLength: 200,
			Null:      true,
		},
	}
	return tab02
}

// tablechange
// 修改已经存在的字段
func TestChangeTable(t *testing.T) {
	testDB := createTestDB()
	db := testDB.DB
	defer testDB.Close()
	tab := tableTest()
	err := new(meta).CreateTable(db, tab)
	if err != nil {
		t.Error(err)
	}

	tab01 := tableTest01()
	defer func() {
		if _, err := db.Exec("drop table test01"); err != nil {
			t.Error(err)
		}
	}()
	tab01.FormerName = []string{"test", "test03"}
	err = tab01.Update("sqlite3", db)
	if err != nil {
		t.Error(err)
	}
}

// 增加新的字段
func TestChangeTable01(t *testing.T) {
	testDB := createTestDB()
	db := testDB.DB
	defer testDB.Close()
	tab := tableTest()
	err := new(meta).CreateTable(db, tab)
	if err != nil {
		t.Error(err)
	}

	tab02 := tableTest02()
	defer func() {
		if _, err := db.Exec("drop table test02"); err != nil {
			t.Error(err)
		}
	}()
	tab02.FormerName = []string{"test", "test03"}
	err = tab02.Update("sqlite3", db)
	if err != nil {
		t.Error(err)
	}

}
