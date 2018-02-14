package linkedstack

import (
	"github.com/nathanpucheril/GoCollect/containers/lists/linkedlist/doublylinkedlist"
	"github.com/nathanpucheril/GoCollect/iterators"
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

func (self *LinkedListStack) Iterator() iterators.Iterator {
	return self.DoublyLinkedList.Iterator() // TODO need to do in reverse
}
