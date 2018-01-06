package linkedlist

import (
	"testing"
	"fmt"
)

func TestSinglyLinkedList_Insert(t *testing.T) {
	sll := NewSLL()
	sll.InsertFront(1)
	sll.InsertFront(2)
	sll.InsertFront(3)
	fmt.Println(sll.Size())
	fmt.Println(sll.ToSlice())
}
