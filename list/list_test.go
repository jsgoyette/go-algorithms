package list

import "testing"

func Test_New(t *testing.T) {
	list := New()

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(10)
	list.Insert(103)
	list.Insert(56)

	if list.Has(103) != true {
		t.Error("[Error] Has(103, list) doesn't work as expected'")
	}

	if list.Length() != 6 {
		t.Error("[Error] Length(list) doensn't work as expected'")
	}

	list.Remove(10)

	if list.Has(10) != false {
		t.Error("[Error] Has(10, list) doesn't work as expected after removing 10 from list'")
	}

	if list.Length() != 5 {
		t.Error("[Error] Length(list) doensn't work as expected after removing data from list'")
	}
}
