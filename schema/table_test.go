package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func tableEqu(tab1, tab2 *Table) error {
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
func TestTableDefineScript(t *testing.T) {
	tt := []struct {
		name string
		tab1 *Table
		tab2 *Table
	}{struct {
		name string
		tab1 *Table
		tab2 *Table
	}{
		"数据类型测试",
		func() *Table {
			tab := NewTable("test")
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
		}(), func() *Table {
			tab := NewTable("test")
			tab.Columns = []*Column{
				&Column{Name: "CODE", Type: TypeString, Null: true, MaxLength: 50},
				&Column{Name: "NAME", Type: TypeString, Null: true, MaxLength: 50},
				&Column{Name: "DESC", Type: TypeString, Null: true, MaxLength: 0},
				&Column{Name: "NUM", Type: TypeFloat, Index: true},
				&Column{Name: "NUM1", Type: TypeInt, Null: true},
				&Column{Name: "NUM2", Type: TypeInt, Null: true},
				&Column{Name: "DATE1", Type: TypeDatetime, Null: true, Index: true},
				&Column{Name: "DATE2", Type: TypeDatetime, Null: true, Index: true},
				&Column{Name: "BYTE1", Type: TypeBytea, Null: true},
			}
			tab.PrimaryKeys = []string{"CODE", "NAME"}
			return tab
		}()},
		struct {
			name string
			tab1 *Table
			tab2 *Table
		}{
			"主键另一种写法测试",
			func() *Table {
				tab := NewTable("test")
				if err := tab.DefineScript(`
			CODE  STR(50)
			NAME  STR(50)
			DESC  STR
			NUM	  FLOAT NOT NULL INDEX
			NUM1  INT
			NUM2
			DATE1 DATE NULL INDEX
			DATE2
			PRIMARY key(code,name)
		`); err != nil {
					t.Error(err)
				}
				return tab
			}(), func() *Table {
				tab := NewTable("test")
				tab.Columns = []*Column{
					&Column{Name: "CODE", Type: TypeString, Null: true, MaxLength: 50},
					&Column{Name: "NAME", Type: TypeString, Null: true, MaxLength: 50},
					&Column{Name: "DESC", Type: TypeString, Null: true, MaxLength: 0},
					&Column{Name: "NUM", Type: TypeFloat, Index: true},
					&Column{Name: "NUM1", Type: TypeInt, Null: true},
					&Column{Name: "NUM2", Type: TypeInt, Null: true},
					&Column{Name: "DATE1", Type: TypeDatetime, Null: true, Index: true},
					&Column{Name: "DATE2", Type: TypeDatetime, Null: true, Index: true},
				}
				tab.PrimaryKeys = []string{"CODE", "NAME"}
				return tab
			}()},
	}
	for _, v := range tt {
		if err := tableEqu(v.tab1, v.tab2); err != nil {
			t.Error(v.name, err)
		}
	}
}
