package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/linlexing/dbx/schema"
	_ "github.com/linlexing/dbx/schema/postgres"
)

func getDb() (*sql.DB, error) {
	return sql.Open("postgres", "user=test password=123456 dbname=postgres sslmode=disable")
}

func TestCreateTable(t *testing.T) {
	db, err := getDb()
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
			Name: "name",
			Type: schema.TypeString,
			Null: true,
		},
	}
	if err = schema.Find("postgres").CreateTable(db, tab); err != nil {
		t.Error(err)
	}
	if _, err = db.Exec("drop table test"); err != nil {
		t.Error(err)
	}
}

func TestTableExists(t *testing.T) {
	db, err := getDb()
	var ok bool
	ok, err = schema.Find("postgres").TableExists(db, "userinfo")
	db.Close()
	if err == nil && ok {
		t.Log("测试通过")
	} else {
		t.Error("测试未通过")
	}
}

func TestCreateTableAs(t *testing.T) {
	db, err := getDb()
	strSQL := "select *from userinfo"
	pks := []string{"id", "name"}
	err = schema.Find("postgres").CreateTableAs(db, "copyuserinfo", strSQL, pks)
	db.Exec("drop table copyuserinfo")
	db.Close()
	fmt.Println(err)
	if err == nil {
		t.Log("测试通过")
	} else {
		t.Error("测试未通过")
	}
}

func TestTableNames(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
	}

	var names []string
	names, err = schema.Find("postgres").TableNames(db)
	fmt.Println(err)
	fmt.Println(len(names))
	for key, value := range names {
		fmt.Println(key)
		fmt.Println(value)
	}
	if err == nil {
		t.Log("测试通过")
	} else {
		t.Error("测试未通过")
	}

}
