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

//测试创建表
func TestCreateTable(t *testing.T) {
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
	}
	err = new(meta).CreateTable(db, tab)
	if err != nil {
		t.Error("创建表测试未通过")
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
	_, err = new(meta).TableExists(db, "test")
	if err != nil {
		t.Error("判断表是否存在测试未通过")
	}
}
