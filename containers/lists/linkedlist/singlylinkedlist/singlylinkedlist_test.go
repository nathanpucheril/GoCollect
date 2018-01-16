package singlylinkedlist

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedList_Insert(t *testing.T) {
	sll := NewSLL()
	sll.Prepend(1)
	sll.Prepend(2)
	sll.Prepend(3)
	fmt.Println(sll.Size())
	fmt.Println(sll.ToSlice())

	it := sll.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
	fmt.Println(it.Next())
	fmt.Println(it.Next())
	fmt.Println(it.Next())
}
