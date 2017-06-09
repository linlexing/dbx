package schema

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type T用户 struct {
	C代码 string  `dbx:"代码 STR(50) PRIMARY KEY"`
	C名称 string  `dbx:"名称 STR(50) PRIMARY KEY"`
	C描述 string  `dbx:"描述 STR"`
	C数量 float64 `dbx:"数量 FLOAT NOT NULL INDEX"`
}

func TestTableFromStruct(t *testing.T) {
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
			tab := NewTable("TEST")
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
			tab, err := TableFromStruct("TEST", new(test))
			if err != nil {
				t.Error(err)
			}
			return tab
		}()},
		struct {
			name string
			tab1 *Table
			tab2 *Table
		}{
			"汉字",
			func() *Table {
				tab := NewTable("用户")
				if err := tab.DefineScript(`
			代码  STR(50) PRIMARY KEY
			名称  STR(50) PRIMARY KEY
			描述  STR
			数量  FLOAT NOT NULL INDEX
		`); err != nil {
					t.Error(err)
				}
				return tab
			}(), func() *Table {

				tab, err := TableFromStruct("用户", new(T用户))
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
func TestStruct2Row(t *testing.T) {
	tt := []struct {
		name string
		row1 map[string]interface{}
		row2 map[string]interface{}
	}{struct {
		name string
		row1 map[string]interface{}
		row2 map[string]interface{}
	}{
		"基本测试",
		func() map[string]interface{} {
			user := &T用户{
				C代码: "001",
				C名称: "name",
				C数量: 11.2,
			}
			row, err := Struct2Row(user)
			if err != nil {
				t.Error(err)
			}
			return row
		}(),
		map[string]interface{}{
			"代码": "001",
			"名称": "name",
			"描述": nil,
			"数量": float64(11.2),
		},
	},
	}
	for _, one := range tt {
		if err := mapEqu(one.row1, one.row2); err != nil {
			t.Error(err)
		}
	}
}

func TestRow2Struct(t *testing.T) {
	tt := []struct {
		name string
		row1 *T用户
		row2 *T用户
	}{struct {
		name string
		row1 *T用户
		row2 *T用户
	}{
		"基本测试",
		func() *T用户 {
			user := &T用户{}
			if err := Row2Struct(
				map[string]interface{}{
					"代码": "001",
					"名称": "name",
					"描述": nil,
					"数量": float64(11.2),
				}, user); err != nil {
				t.Error(err)
			}
			return user
		}(),
		&T用户{
			C代码: "001",
			C名称: "name",
			C数量: 11.2,
			C描述: "",
		},
	},
	}
	for _, one := range tt {
		if !reflect.DeepEqual(one.row1, one.row2) {
			t.Error(fmt.Errorf("row1:%v <> %v", one.row1, one.row2))
		}
	}
}
func mapEqu(row1, row2 map[string]interface{}) error {
	if len(row1) != len(row2) {
		return fmt.Errorf("len %d <> %d", len(row1), len(row2))
	}
	for k, v := range row1 {
		ov, ok := row2[k]
		if !ok {
			return fmt.Errorf("key:%s not found at row2", k)
		}
		if !reflect.DeepEqual(v, ov) {
			return fmt.Errorf("key:%s,val: %v <> %v", k, v, ov)
		}
	}
	for k := range row2 {
		_, ok := row1[k]
		if !ok {
			return fmt.Errorf("key:%s not found at row1", k)
		}
	}
	return nil
}
