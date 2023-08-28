package data

import "testing"

func TestBind(t *testing.T) {
	list := [][]string{
		{"select 1 where a=?", "select 1 where a=$1"},
		{"select 1 where a='?'", "select 1 where a='?'"},
		{"select 1 where a='???'", "select 1 where a='???'"},
		{"select 1 where a='???''?'?", "select 1 where a='???''?'$1"},
		{"select 1 where a='??' and b=?", "select 1 where a='??' and b=$1"},
	}
	for _, one := range list {

		if Rebind(DOLLAR, one[0]) != one[1] {
			t.Error("error", one[0], Rebind(DOLLAR, one[0]))
		}
	}
}
