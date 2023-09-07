package schema

import "testing"

func TestDefine(t *testing.T) {
	col, err := columnDefine("A str(1) index")
	if err != nil {
		t.Fatal(err)
	}
	if col.Name != "A" ||
		!col.Null ||
		col.Type != TypeString ||
		!col.Index {
		t.Fatal("error")
	}
}
