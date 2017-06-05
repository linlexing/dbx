package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/linlexing/dbx/schema"
)

func getDb() (*sql.DB, error) {
	return sql.Open("postgres", "port=5432 user=test password=123456 dbname=postgres sslmode=disable")
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
			Name: "name",
			Type: schema.TypeString,
			Null: true,
		},
		&schema.Column{
			Name: "age",
			Type: schema.TypeInt,
			Null: false,
		},
	}
	schema.Find("postgres").CreateTable(db, tab)
	return
}

//创建表格测试
func TestCreateTable(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	if err = createTable(db); err != nil {
		t.Error(err)
	}
	if _, err = db.Exec("drop table test"); err != nil {
		t.Error(err)
	}
}

//表格是否存在测试
func TestTableExists(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
	}
	_, err = schema.Find("postgres").TableExists(db, "userinfo")
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
}

//复制表格测试
func TestCreateTableAs(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
	}
	strSQL := "select *from userinfo"
	pks := []string{"id", "name"}
	err = schema.Find("postgres").CreateTableAs(db, "copyuserinfo", strSQL, pks)
	defer db.Close()
	if _, err = db.Exec("drop table copyuserinfo"); err != nil {
		t.Error(err)
	}
}

//获得所以表格名称测试
func TestTableNames(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	var names []string
	names, err = schema.Find("postgres").TableNames(db)
	for key, value := range names {
		fmt.Println(key)
		fmt.Println(value)
	}
	if err != nil {
		t.Error(err)
	}

}

//创建 删除索引
func TestCreateIndexIfNotExistsAndDropIndexIfExists(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
	}
	err = schema.Find("postgres").CreateIndexIfNotExists(db, "myindex", "userinfo", "id")
	defer db.Close()
	if err != nil {
		t.Error(err)
	}
	err = schema.Find("postgres").DropIndexIfExists(db, "myindex", "userinfo")
	if err != nil {
		t.Error(err)
	}
}

//改变表格结构测试
func TestChangeTable(t *testing.T) {
	db, err := getDb()
	if err != nil {
		t.Error(err)
	}
	if err = createTable(db); err != nil {
		t.Error(err)
	}
	defer db.Close()
	tab := schema.TableSchemaChange{
		NewName:  "testChange",
		OldName:  "test",
		PKChange: true,
		PK:       []string{"username"},
	}
	tab.OriginFields = []*schema.Column{
		&schema.Column{
			Name: "id",
			Type: schema.TypeInt,
			Null: false,
		},
		&schema.Column{
			Name: "name",
			Type: schema.TypeString,
			Null: true,
		},
	}
	tab.ChangeFields = []*schema.ChangedField{
		{
			NewField: &schema.Column{
				Name: "username",
				Type: schema.TypeString,
				Null: false,
			},
			OldField: &schema.Column{
				Name: "name",
				Type: schema.TypeString,
				Null: true,
			},
		},
		{
			NewField: &schema.Column{
				Name: "age",
				Type: schema.TypeInt,
				Null: true,
			},
			OldField: &schema.Column{
				Name: "age",
				Type: schema.TypeInt,
				Null: false,
			},
		},
		{
			NewField: &schema.Column{
				Name:      "clazz",
				Type:      schema.TypeInt,
				Null:      true,
				Index:     true,
				IndexName: "myindex",
			},
		},
	}
	err = schema.Find("postgres").ChangeTable(db, &tab)
	if _, err = db.Exec("drop table testChange"); err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}
}
