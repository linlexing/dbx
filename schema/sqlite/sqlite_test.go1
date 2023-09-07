package sqlite

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/linlexing/dbx/schema"
	// _ "github.com/mattn/go-sqlite3"
)

// //
func getDb() (*sql.DB, error) {
	return sql.Open("sqlite3", "./testDB.db")

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
	schema.Find("sqlite3").CreateTable(db, tab)
	return
}

func TestCreateTable(t *testing.T) {
	db, err := getDb()
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	err = createTable(db)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("DROP TABLE test"); err != nil {
			t.Error(err)
		}
	}()
}

func TestCreateTableAs(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	if err = createTable(db); err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("DROP TABLE test"); err != nil {
			t.Error(err)
		}
	}()
	tableName := "testAS"
	strSQL := "select * from test"
	pk := []string{"ID"}
	err = schema.Find("sqlite3").CreateTableAs(db, tableName, strSQL, pk)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if _, err := db.Exec("DROP TABLE testAS"); err != nil {
			t.Error(err)
		}
	}()
}
