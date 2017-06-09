package orm

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/linlexing/dbx/schema"
)

type T用户 struct {
	C代码 string  `dbx:"代码 STR(50) PRIMARY KEY"`
	C名称 string  `dbx:"名称 STR(50) PRIMARY KEY"`
	C描述 string  `dbx:"描述 STR"`
	C数量 float64 `dbx:"数量 FLOAT NOT NULL INDEX"`
}

func (t *T用户) TableName() string {
	return "用户"
}

func TestTableFromStruct(t *testing.T) {
	tt := []struct {
		name string
		tab1 *schema.Table
		tab2 *schema.Table
	}{struct {
		name string
		tab1 *schema.Table
		tab2 *schema.Table
	}{
		"数据类型测试",
		func() *schema.Table {
			tab := schema.NewTable("TEST")
			if err := tab.DefineScript(`
			CODE  STR(50) PRIMARY KEY
			NAME  STR(50) PRIMARY KEY
			DESC  STR
			NUM	  FLOAT NOT NULL INDEX
			NUM1  INT
			NUM2
			DATE1 DATE NULL INDEX
			DATE2
			BYTE1 BYTEA
		`); err != nil {
				t.Error(err)
			}
			return tab
		}(), func() *schema.Table {
			type test struct {
				Code  string  `dbx:"STR(50) PRIMARY KEY"`
				Name  string  `dbx:"STR(50) PRIMARY KEY"`
				Desc  string  `dbx:"STR"`
				Num   float64 `dbx:"FLOAT NOT NULL INDEX"`
				Num1  int     `dbx:"INT"`
				Num2  int
				Date1 time.Time `dbx:"DATE NULL INDEX"`
				Date2 time.Time
				Byte1 []byte `dbx:"bytea"`
			}
			tab, err := TableFromStruct(new(test))
			if err != nil {
				t.Error(err)
			}
			return tab
		}()},
		struct {
			name string
			tab1 *schema.Table
			tab2 *schema.Table
		}{
			"汉字",
			func() *schema.Table {
				tab := schema.NewTable("用户")
				if err := tab.DefineScript(`
			代码  STR(50) PRIMARY KEY
			名称  STR(50) PRIMARY KEY
			描述  STR
			数量  FLOAT NOT NULL INDEX
		`); err != nil {
					t.Error(err)
				}
				return tab
			}(), func() *schema.Table {

				tab, err := TableFromStruct(new(T用户))
				if err != nil {
					t.Error(err)
				}
				return tab
			}()},
	}
	for _, v := range tt {
		if err := tableEqu(v.tab1, v.tab2); err != nil {
			t.Error(v.name, err)
		}
	}
}

func tableEqu(tab1, tab2 *schema.Table) error {
	if tab1.FullName() != tab2.FullName() {
		return fmt.Errorf("table name %s <> %s", tab1.FullName(), tab2.FullName())
	}
	if !reflect.DeepEqual(tab1.PrimaryKeys, tab2.PrimaryKeys) {
		return fmt.Errorf("table pk %v <> %v", tab1.PrimaryKeys, tab2.PrimaryKeys)
	}
	if len(tab1.Columns) != len(tab2.Columns) {
		return fmt.Errorf("col num %d <> %d", len(tab1.Columns), len(tab2.Columns))
	}
	for i := range tab1.Columns {
		if !tab1.Columns[i].Eque(tab2.Columns[i]) {
			return fmt.Errorf("col not equ :\n%v\n%v", tab1.Columns[i], tab2.Columns[i])
		}
	}
	return nil
}
