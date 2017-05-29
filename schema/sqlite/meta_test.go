package sqlite

import (
	"fmt"
	"testing"

	"os"

	"github.com/jmoiron/sqlx"
	"github.com/linlexing/dbx"
	_ "github.com/mattn/go-sqlite3"
)

type testDB struct {
	fileName string

	DB *sqlx.DB
}

func createTestDB() *testDB {
	rev := testDB{fileName: "e:\\temp\\test.sq3"}

	if _, err := os.Stat(rev.fileName); err == nil {
		os.Remove(rev.fileName)
	}
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
	if _, err = db.NamedExec("insert into aaa(a,b)values(:a,:b)", &v); err != nil {
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
func Test_RemoveColumns(t *testing.T) {
	testDB := createTestDB()
	db := testDB.DB
	defer testDB.Close()
	var err error
	if _, err = db.Exec("create table aaa(a varchar(200) primary key,b integer,C text)"); err != nil {
		t.Fatal(err)
	}
	if err = new(meta).RemoveColumns(db, "aaa", []string{"b"}); err != nil {
		t.Error(err)
	}
	strSQL := fmt.Sprintf("PRAGMA table_info(%s)", "aaa")
	tabCols := []tableColumn{}
	if err = db.Select(&tabCols, strSQL); err != nil {
		t.Error(dbx.NewSQLError(strSQL, nil, err))
	}
	if len(tabCols) != 2 {
		t.Error("not 2")
	}
	if tabCols[0].Name != "a" || tabCols[1].Name != "C" {
		t.Error("error")
	}
}
