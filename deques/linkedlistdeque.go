package deques

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists/linkedlist"
	"github.com/nathanpucheril/GoCollect/lists/linkedlist/doublylinkedlist"
)

type LinkedListDeque struct {
	 *doublylinkedlist.DoublyLinkedList
}


func (self *LinkedListDeque) Peek() (interface{}, bool) {
	return self.DoublyLinkedList.GetFront()
}

func (self *LinkedListDeque) PeekFirst() (interface{}, bool) {
	return self.DoublyLinkedList.GetFront()
}

func (self *LinkedListDeque) PeekLast() (interface{}, bool) {
	return self.DoublyLinkedList.GetBack()
}


func (self *LinkedListDeque) Push(value interface{}) {
	panic("implement me")
}

func (self *LinkedListDeque) Pop() (interface{}, error) {
	panic("implement me")
}

func NewLLDeque() LinkedListDeque {
	return LinkedListDeque{doublylinkedlist.NewDLL()}
}


func (self *LinkedListDeque) Add(value interface{}) {
	self.AddFirst(value)
}

func (self *LinkedListDeque) AddFirst(value interface{}) {
	self.DoublyLinkedList.Prepend(value)
}

func (self *LinkedListDeque) AddLast(value interface{}) {
	self.DoublyLinkedList.Append(value)
}


func (self *LinkedListDeque) Remove() (interface{}, bool) {
	return self.DoublyLinkedList.RemoveFront()
}

func (self *LinkedListDeque) RemoveFirst() (interface{}, bool){
	return self.DoublyLinkedList.RemoveFront()
}

func (self *LinkedListDeque) RemoveLast() (interface{}, bool) {
	return self.DoublyLinkedList.RemoveBack()
}


