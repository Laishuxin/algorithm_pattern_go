package ArrayList

import (
	"fmt"
	"testing"
)

func TestArrayList_Set(t *testing.T) {
	// case1: out of range
	name1 := "out of range"
	list := NewArrayList()
	wantErr := list.Set(1, 100)
	if wantErr != nil {
		t.Log("success, list.String() = ", list.String())
	} else {
		t.Fatal("fail ", name1)
	}

	// case2: normal
	name2 := "normal"
	list2 := NewArrayList()
	list2.Append("str1")
	list2.Append("str2")
	list2.Append("str3")
	want2 := list2.Set(1, "str22222")
	if want2 == nil {
		t.Log("success, list.String() = ", list2.String())
	} else {
		t.Fatal("fail, ", name2)
	}
}

func TestArrayList_Iterator(t *testing.T) {
	list := NewArrayList()
	list.Append("str1")
	list.Append("str2")
	list.Append("str3")
	list.Append("str4")
	list.Append("str5")

	it := list.Iterator()
	for it.HasNext() {
		result, _ := it.Next()
		if result == "str3" {
			it.Remove()
		}
	}

	for it := list.Iterator(); it.HasNext(); {
		result, _ := it.Next()
		fmt.Println(result)
	}
}
