package mysql

import (
	"database/sql"
	"testing"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/linlexing/dbx/schema"
	_ "github.com/linlexing/dbx/schema/mysql"
)

func getDb() (*sql.DB, error) {
	return sql.Open("mysql", "test:123456@tcp(localhost:3306)/test?charset=utf8")
}
func createTable(db *sql.DB) (err error) {
	tab := schema.NewTable("test")
	tab.PrimaryKeys = []string{"ID"}
	tab.Columns = []*schema.Column{
		&schema.Column{
			Name:      "ID",
			Type:      schema.TypeInt,
			Null:      false,
			Index:     true,
			IndexName: "oldindex",
		},
		&schema.Column{
			Name: "NAME",
			Type: schema.TypeString,
			Null: true,
		},
		&schema.Column{
			Name: "AGE",
			Type: schema.TypeInt,
			Null: false,
		},
	}
	schema.Find("mysql").CreateTable(db, tab)
	return
}

//打开表格测试
func TestOpenTable(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if err = createTable(db); err != nil {
		t.Error(err)
	}
	var tab *schema.Table
	tab, err = schema.Find("mysql").OpenTable(db, "test")
	if err != nil {
		t.Error(err)
	}
	for _, value := range tab.PrimaryKeys {
		fmt.Println(value)
	}
	for key, value := range tab.Columns {
		fmt.Println(key)
		fmt.Println(value)
	}
	if _, err = db.Exec("DROP TABLE test"); err != nil {
		t.Error(err)
	}

}
