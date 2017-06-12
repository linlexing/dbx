package schema

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type otherType struct {
	Code string  `dbx:"STR(50) PRIMARY KEY"`
	Name string  `dbx:"STR(50) PRIMARY KEY"`
	Num  float64 `dbx:"FLOAT"`
}
type innerT struct {
	C嵌入 string `dbx:"嵌入 STR(50)"`
}
type T用户 struct {
	C代码  string       `dbx:"代码 STR(50) PRIMARY KEY"`
	C名称  string       `dbx:"名称 STR(50) PRIMARY KEY"`
	C描述  string       `dbx:"描述 STR"`
	C时间  time.Time    `dbx:"时间 DATE"`
	C数量  float64      `dbx:"数量 FLOAT NOT NULL INDEX"`
	C其他  []*otherType `dbx:"其他 STR"`
	C其他1 []otherType  `dbx:"其他1 BYTEA"`
	C明细  []otherType  `dbx:"明细 CHILD"`
	innerT
}

func init() {
	gob.Register(T用户{})
}
func TestInnerStruct(t *testing.T) {
	type tabT struct {
		C代码 string `dbx:"代码 STR(50) PRIMARY KEY"`
		C名称 string `dbx:"名称 STR(50) PRIMARY KEY"`
		innerT
	}
	fields, err := fieldsFromStruct(reflect.TypeOf(tabT{}), nil, true)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fields[2].valPath, []int{2, 0}) {
		t.Fatal("error", fields[2])
	}
}
func TestTableFromStruct(t *testing.T) {
	tt := []struct {
		name string
		tab1 []*Table
		tab2 []*Table
	}{struct {
		name string
		tab1 []*Table
		tab2 []*Table
	}{
		"数据类型测试",
		func() []*Table {
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
			return []*Table{tab}
		}(), func() []*Table {
			type test struct {
				Code  string  `dbx:"STR(50) PRIMARY KEY"`
				Name  string  `dbx:"STR(50) PRIMARY KEY"`
				Desc  string  `dbx:"STR"`
				Num   float64 `dbx:"FLOAT NOT NULL INDEX"`
				Num1  int64   `dbx:"INT"`
				Num2  int64
				Date1 time.Time `dbx:"DATE NULL INDEX"`
				Date2 time.Time
				Byte1 []byte `dbx:"bytea"`
			}
			tabs, err := TableFromStruct("TEST", new(test))
			if err != nil {
				t.Error(err)
			}
			return tabs
		}()},
		struct {
			name string
			tab1 []*Table
			tab2 []*Table
		}{
			"汉字",
			func() []*Table {
				tab := NewTable("用户")
				if err := tab.DefineScript(`
					代码  STR(50) PRIMARY KEY
					名称  STR(50) PRIMARY KEY
					描述  STR
					时间 DATE
					数量  FLOAT NOT NULL INDEX
					其他  STR
					其他1 BYTEA
					嵌入 STR(50)
				`); err != nil {
					t.Error(err)
				}
				tab1 := NewTable("明细")
				if err := tab1.DefineScript(`
					CODE  STR(50) PRIMARY KEY
					name  STR(50) PRIMARY KEY
					num  FLOAT
				`); err != nil {
					t.Error(err)
				}
				return []*Table{tab, tab1}
			}(), func() []*Table {

				tabs, err := TableFromStruct("用户", new(T用户))
				if err != nil {
					t.Error(err)
				}
				return tabs
			}()},
	}
	for _, v := range tt {
		if len(v.tab1) != len(v.tab2) {
			t.Error(fmt.Errorf("%s len %d <> %d", v.name,
				len(v.tab1), len(v.tab2)))
		}
		for i := range v.tab1 {
			if err := tableEqu(v.tab1[i], v.tab2[i]); err != nil {
				t.Error(fmt.Errorf("%s,tab:%s <> %s,%s", v.name,
					v.tab1[i].Name, v.tab2[i].Name, err))
			}

		}
	}
}

func TestStruct2Row(t *testing.T) {
	var row1 map[string]interface{}
	var child1 map[string][]map[string]interface{}
	var row2 map[string]interface{}
	var child2 map[string][]map[string]interface{}
	user := &T用户{
		C代码: "001",
		C名称: "name",
		C数量: 11.2,
		C描述: "",
		C时间: time.Time{},
		C其他: []*otherType{
			&otherType{
				Name: "n1",
				Code: "c1",
				Num:  111.22,
			}, &otherType{
				Name: "n2",
				Code: "c2",
				Num:  111.11,
			},
		},
		C其他1: []otherType{
			otherType{
				Name: "n1",
				Code: "c1",
				Num:  111.22,
			}, otherType{
				Name: "n2",
				Code: "c2",
				Num:  111.11,
			},
		},
		C明细: []otherType{
			otherType{
				Name: "n1",
				Code: "c1",
				Num:  111.22,
			}, otherType{
				Name: "n2",
				Code: "c2",
				Num:  111.11,
			},
		},
	}

	row1, child1, err := Struct2Row(user)
	if err != nil {
		t.Error(err)
	}
	bys := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(bys).Encode([]otherType{
		otherType{
			Name: "n1",
			Code: "c1",
			Num:  111.22,
		}, otherType{
			Name: "n2",
			Code: "c2",
			Num:  111.11,
		},
	}); err != nil {
		t.Error(err)
	}
	row2 = map[string]interface{}{
		"代码":  "001",
		"名称":  "name",
		"描述":  nil,
		"时间":  nil,
		"数量":  float64(11.2),
		"其他":  `[{"Code":"c1","Name":"n1","Num":111.22},{"Code":"c2","Name":"n2","Num":111.11}]`,
		"其他1": bys.Bytes(),
		"嵌入":  nil,
	}
	child2 = map[string][]map[string]interface{}{
		"明细": []map[string]interface{}{
			map[string]interface{}{
				"CODE": "c1",
				"NAME": "n1",
				"NUM":  float64(111.22),
			},
			map[string]interface{}{
				"CODE": "c2",
				"NAME": "n2",
				"NUM":  float64(111.11),
			},
		},
	}
	if err := mapEqu(row1, row2); err != nil {
		t.Error(err)
	}
	if len(child1) != len(child2) {
		t.Error("child not equ")
	}
	for k := range child1 {
		c1 := child1[k]
		c2 := child2[k]
		if len(c1) != len(c2) {
			t.Error("child ", k, " rows not equ")
		}
		for i := range c1 {
			if err := mapEqu(c1[i], c2[i]); err != nil {
				t.Error("child", k, "row", i, err)
			}
		}
	}
	//测试时间
	dt := time.Date(2016, 1, 1, 12, 1, 1, 0, time.Local)
	user.C时间 = dt
	row1, child1, err = Struct2Row(user)
	if err != nil {
		t.Error(err)
	}
	row2["时间"] = time.Date(2016, 1, 1, 12, 1, 1, 0, time.Local)
	if err := mapEqu(row1, row2); err != nil {
		t.Error(err)
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
			bys := bytes.NewBuffer(nil)
			if err := gob.NewEncoder(bys).Encode([]otherType{
				otherType{
					Name: "n1",
					Code: "c1",
					Num:  111.22,
				}, otherType{
					Name: "n2",
					Code: "c2",
					Num:  111.11,
				},
			}); err != nil {
				t.Error(err)
			}
			user := &T用户{}
			if err := Row2Struct(
				map[string]interface{}{
					"代码":  "001",
					"名称":  "name",
					"描述":  nil,
					"数量":  float64(11.2),
					"其他":  `[{"name":"n1","code":"c1","num":111.22},{"name":"n2","code":"c2","num":111.11}]`,
					"其他1": bys.Bytes(),
					"嵌入":  "aaa",
				}, map[string][]map[string]interface{}{
					"明细": []map[string]interface{}{
						map[string]interface{}{
							"CODE": "c1",
							"NAME": "n1",
							"NUM":  float64(111.22),
						},
						map[string]interface{}{
							"CODE": "c2",
							"NAME": "n2",
							"NUM":  float64(111.11),
						},
					},
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
			C其他: []*otherType{
				&otherType{
					Name: "n1",
					Code: "c1",
					Num:  111.22,
				}, &otherType{
					Name: "n2",
					Code: "c2",
					Num:  111.11,
				},
			},
			C其他1: []otherType{
				otherType{
					Name: "n1",
					Code: "c1",
					Num:  111.22,
				}, otherType{
					Name: "n2",
					Code: "c2",
					Num:  111.11,
				},
			},
			C明细: []otherType{
				otherType{
					Name: "n1",
					Code: "c1",
					Num:  111.22,
				}, otherType{
					Name: "n2",
					Code: "c2",
					Num:  111.11,
				},
			},
			innerT: innerT{
				C嵌入: "aaa",
			},
		},
	},
	}
	for _, one := range tt {
		if !reflect.DeepEqual(one.row1, one.row2) {
			t.Error(fmt.Errorf("row not equ:\n%#v\n%#v", one.row1, one.row2))
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
			return fmt.Errorf("key:%s,val not equ:\n%v\n%v", k, v, ov)
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
