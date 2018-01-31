package doublylinkedlist

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedList_Clear(t *testing.T) {
	dll := NewDLL()
	dll.Prepend(1)
	dll.Prepend(2)
	dll.Prepend(3)
	dll.Append(4)
	dll.Append(5)
	dll.Append(6)
	dll.Append(7)

	if dll.Size() != 7 {
		t.Fail()
	}

	// Clear everything
	dll.Clear()

	// Check that the list *thinks* its cleared
	if dll.Size() != 0 {
		t.Fail()
	}

	// Check that it is actually cleared in the forward direction
	for i, ptr := 0, dll.sentinelFront.next; ptr != dll.sentinelBack; i++ {
		t.Fail()
		ptr = ptr.next
	}
	// Check that it is actually cleared in the backward direction
	for i, ptr := 0, dll.sentinelBack.prev; ptr != dll.sentinelFront; i++ {
		t.Fail()
		ptr = ptr.prev
	}

	dll.Prepend(1)
	dll.Prepend(2)
	dll.Prepend(3)
	dll.Append(4)
	dll.Append(5)
	dll.Append(6)
	dll.Append(7)

	if dll.Size() != 7 {
		t.Fail()
	}

	// Check we successfully added items back in
	i := 0
	for ptr := dll.sentinelFront.next; ptr != dll.sentinelBack; i++ {
		ptr = ptr.next
	}

	// Check that items were added back in
	if dll.Size() != i {
		t.Fail()
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	dll := NewDLL()
	dll.Prepend(1)
	dll.Prepend(2)
	dll.Prepend(3)
	dll.Prepend(4)
	dll.Prepend(5)
	dll.Prepend(6)

	//dll.Insert(2, 0)
	fmt.Println(dll.Size())
	fmt.Println(dll.ToSlice())
	it := dll.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
	it = dll.ReverseIterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
	fmt.Println(it.Next())
	fmt.Println(it.Next())
	fmt.Println(it.Next())

	//for item := range containers.ToChanIter(dll.Iterator()) {
	//	fmt.Println(item)
	//}
}
