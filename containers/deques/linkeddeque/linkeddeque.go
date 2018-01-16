package linkeddeque

import (
	"github.com/nathanpucheril/GoCollect/containers/lists/linkedlist/doublylinkedlist"
)

type LinkedDeque struct {
	*doublylinkedlist.DoublyLinkedList
}

func (self *LinkedDeque) Peek() (interface{}, bool) {
	return self.DoublyLinkedList.GetFront()
}

func (self *LinkedDeque) PeekFirst() (interface{}, bool) {
	return self.DoublyLinkedList.GetFront()
}

func (self *LinkedDeque) PeekLast() (interface{}, bool) {
	return self.DoublyLinkedList.GetBack()
}

func (self *LinkedDeque) Push(value interface{}) {
	panic("implement me")
}

func (self *LinkedDeque) Pop() (interface{}, error) {
	panic("implement me")
}

func NewLLDeque() LinkedDeque {
	return LinkedDeque{doublylinkedlist.NewDLL()}
}

func (self *LinkedDeque) Add(value interface{}) {
	self.AddFirst(value)
}

func (self *LinkedDeque) AddFirst(value interface{}) {
	self.DoublyLinkedList.Prepend(value)
}

func (self *LinkedDeque) AddLast(value interface{}) {
	self.DoublyLinkedList.Append(value)
}

func (self *LinkedDeque) Remove() (interface{}, bool) {
	return self.DoublyLinkedList.RemoveFront()
}

func (self *LinkedDeque) RemoveFirst() (interface{}, bool) {
	return self.DoublyLinkedList.RemoveFront()
}

func (self *LinkedDeque) RemoveLast() (interface{}, bool) {
	return self.DoublyLinkedList.RemoveBack()
}
