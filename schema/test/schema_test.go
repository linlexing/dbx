package postgres

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/linlexing/dbx/schema"
	_ "github.com/linlexing/dbx/schema/postgres"
)

//测试表格创建
func TestCreateTable(t *testing.T) {
	db, err := sql.Open("postgres", "port=5432 user=postgres password=123456 dbname=postgres sslmode=disable")
	tab := schema.NewTable("test")
	tab.PrimaryKeys = []string{"ID"}
	tab.Columns = []*schema.Column{
		&schema.Column{
			Name: "ID",
			Type: schema.TypeInt,
			Null: false,
		},
		&schema.Column{
			Name:      "testName",
			Type:      schema.TypeString,
			MaxLength: 100,
		},
	}
	err = schema.Find("postgres").CreateTable(db, tab)
	if err != nil {
		t.Error("测试未通过")
	} else {
		t.Log("测试通过")
	}
	db.Exec("drop table test")
	db.Close()

}
