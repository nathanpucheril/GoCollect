package stacks

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists/linkedlist/doublylinkedlist"
	"github.com/pkg/errors"
)

type LinkedListStack struct {
	*doublylinkedlist.DoublyLinkedList
}

func NewLinkedListStack() LinkedListStack {
	dll := doublylinkedlist.NewDLL()
	return LinkedListStack{dll}
}

func (self *LinkedListStack) Push(value interface{}) {
	self.DoublyLinkedList.Append(value)
}

func (self *LinkedListStack) Pop() (interface{}, bool) {
	return self.DoublyLinkedList.RemoveBack()
}

func (self *LinkedListStack) Peek() (interface{}, bool) {
	value, ok := self.DoublyLinkedList.GetBack()
	if ok {
		return value, true
	}
	return nil, false
}
