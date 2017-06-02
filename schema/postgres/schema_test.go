package postgres

import (
	"database/sql"
	"testing"

	_ "github.com/bmizerany/pq"
	"github.com/linlexing/dbx/schema"
	_ "github.com/linlexing/dbx/schema"
)

func getdb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=postgres sslmode=disable")
	return db, err
}

//测试创建表
func TestCreateTable(t *testing.T) {
	db, err := getdb()
	tab := schema.NewTable("test")
	tab.PrimaryKeys = []string{"ID"}
	tab.Columns = []*schema.Column{
		&schema.Column{
			Name: "ID",
			Type: schema.TypeInt,
			Null: false,
		},
	}
	err = schema.Find("postgres").CreateTable(db, tab)
	if err != nil {
		t.Error("创建表测试未通过")
	} else {
		t.Log("创建表测试通过")
	}
	db.Exec("drop table test")
	db.Close()
}

//测试表是否存在
func TestTableExists(t *testing.T) {
	db, err := getdb()
	_, err = schema.Find("postgres").TableExists(db, "test")
	if err != nil {
		t.Error("判断表是否存在测试未通过")
	} else {
		t.Log("判断表是否存在测试通过")
	}
}
