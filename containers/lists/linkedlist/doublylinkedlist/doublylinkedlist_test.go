package doublylinkedlist

import (
	"fmt"
	"github.com/nathanpucheril/GoCollect/containers"
	"testing"
)

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

	for item := range containers.ToChanIter(dll.Iterator()) {
		fmt.Println(item)
	}
}
