readme
=============
The database enhancement operations library.
Support mysql postgresql oracle dmdb hive sqlite database type.
Provides data structure definitions, data manipulation, and paging classes.

-------------
```go
package main

import (
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/linlexing/dbx/ddb"
	_ "github.com/linlexing/dbx/postgres"
	"github.com/linlexing/dbx/schema"
)

func main() {
	db, err := ddb.Openx("pgx", "postgres://common:123456@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	//table struct define
	//create table
	tabTest := schema.NewTable("TestTab")
	tabTest.Columns = []*schema.Column{
		{
			Name:      "ID",
			Type:      schema.TypeString,
			MaxLength: 24,
		},
		{
			Name:      "Name",
			Type:      schema.TypeString,
			MaxLength: 50,
		},
		{
			Name: "Remark",
			Type: schema.TypeString,
		},
		{
			Name: "Age",
			Type: schema.TypeInt,
		},
		{
			Name: "Birthday",
			Type: schema.TypeDatetime,
		},
	}
	tabTest.PrimaryKeys = []string{"ID"}
	// list, err := tabTest.Extract(db.DriverName(), db)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, l := range list {
	// 	println(l)
	// }
	//db change
	if err := tabTest.Update(db.DriverName(), db); err != nil {
		panic(err)
	}
	defer func() {
		if _, err := db.Exec("drop table TestTab"); err != nil {
			panic(err)
		}
	}()
	//change struct
	tabTest.ColumnByName("Name").MaxLength = 100
	tabTest.ColumnByName("Age").Type = schema.TypeFloat
	//remove Remark
	tabTest.Columns = append(tabTest.Columns[:2], tabTest.Columns[3:]...)
	tabTest.Columns = append(tabTest.Columns, &schema.Column{
		Name:      "Detail",
		Type:      schema.TypeString,
		MaxLength: 100,
	})
	// list, err = tabTest.Extract(db.DriverName(), db)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, l := range list {
	// 	println(l)
	// }
	//db change
	if err := tabTest.Update(db.DriverName(), db); err != nil {
		panic(err)
	}
}

```