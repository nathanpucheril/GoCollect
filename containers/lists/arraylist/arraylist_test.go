package arraylist

import (
	"fmt"
	"testing"
)

func TestArrayList_Append(t *testing.T) {
	al := New()
	al.Append(1)
	al.Append(2)
	al.Prepend(0)
	al.Prepend(-1)
	al.Prepend(-2)
	al.Insert(5, 5)
	al.Insert(6, 6)
	al.Insert(7, 7)
	al.Insert(5, 5)
	al.Insert(6, 6)
	al.Insert(7, 7)
	al.Insert(5, 5)
	al.Insert(6, 6)
	al.Insert(7, 7)
	al.Remove(0)
	al.Remove(0)
	al.Remove(0)
	al.Remove(0)
	al.Remove(0)
	al.Remove(0)
	al.Remove(0)
	al.Remove(0)
	fmt.Println(len(al.list))
	fmt.Println(al.ToSlice())
	sublist := al.Sublist(2, 3)
	fmt.Println(sublist.ToSlice())
	fmt.Println(al.ToSlice())

	sublist.Set(0, 15)

	fmt.Println(sublist.ToSlice())
	fmt.Println(al.ToSlice())

	//fmt.Println(sublist.Sublist(1,2).ToSlice())
}
