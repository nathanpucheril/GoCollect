package linkedlist

import (
	"testing"
	"fmt"
)

func TestDoublyLinkedList_Insert(t *testing.T) {
	dll := NewDLL()
	dll.InsertFront(1)
	dll.InsertFront(2)
	dll.InsertFront(3)
	dll.InsertFront(4)
	dll.InsertFront(5)
	dll.InsertFront(6)

	//dll.Insert(2, 0)
	fmt.Println(dll.Size())
	fmt.Println(dll.ToSlice())
}
