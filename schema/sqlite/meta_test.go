package sqlite

import (
	"dbweb/lib/ddb"
	"os"
	"testing"

	"github.com/linlexing/dbx/schema"
	_ "github.com/mattn/go-sqlite3"
)

type testDB struct {
	fileName string

	DB ddb.TxDB
}

var sqdb = "e:\\temp\\dump932726563"

//var sqdb = "E:\\SQLite\\test.db"

//create
func createTestDB() *testDB {
	rev := testDB{fileName: sqdb}

	/*	if _, err := os.Stat(rev.fileName); err == nil {
		os.Remove(rev.fileName)
	}*/
	db, err := ddb.Openx("sqlite3", rev.fileName)
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

/*
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
	if _, err = db.Exec("insert into aaa(a,b)values(?,?)", "1324", 111); err != nil {
		t.Fatal(err)
	}
	if err = new(meta).CreateTableAs(db, "bbb", "select * from aaa", []string{"a"}); err != nil {
		t.Fatal(err)
	}

		if err = db.Get(&vNew, "select * from bbb"); err != nil {
			t.Fatal(err)
		}
		if v != vNew {
			t.Error("not equ")
		}

}
*/
