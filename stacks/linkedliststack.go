package stacks

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists/linkedlist/doublylinkedlist"
	"github.com/pkg/errors"
)

type LinkedListStack struct {
	ll *doublylinkedlist.DoublyLinkedList
}

func NewLinkedListStack() LinkedListStack {
	dll := doublylinkedlist.NewDLL()
	return LinkedListStack{dll}
}

func (self *LinkedListStack) Push(value interface{}) {
	self.ll.Append(value)
}

func (self *LinkedListStack) Pop() (interface{}, error) {
	return self.ll.RemoveBack()
}

func (self *LinkedListStack) Peek() (interface{}, error) {
	value, ok := self.ll.GetBack()
	if ok {
		return value, nil
	}
	return nil, errors.New("NoSuchElement")
}

func (self *LinkedListStack) IsEmpty() bool {
	self.ll.IsEmpty()
}

func (self *LinkedListStack) Size() int {
	self.ll.Size()
}

func (self *LinkedListStack) ToSlice() []interface{} {
	self.ll.ToSlice()
}

func (self *LinkedListStack) Clear() {
	self.ll.Clear()
}

func (self *LinkedListStack) Iterator() containers.Iterator {
	return self.ll.Iterator()
}

