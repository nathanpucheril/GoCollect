package deques

import (
	"github.com/nathanpucheril/GoCollect/lists/linkedlist"
	"github.com/nathanpucheril/GoCollect/lists/linkedlist/doublylinkedlist"
)

type LinkedListDeque struct {
	ll *doublylinkedlist.DoublyLinkedList
}

func NewLLDeque() LinkedListDeque {
	return LinkedListDeque{doublylinkedlist.NewDLL()}
}


func (self *LinkedListDeque) Add(value interface{}) {
	self.AddFirst(value)
}

func (self *LinkedListDeque) AddFirst(value interface{}) {
	self.ll.Prepend(value)
}

func (self *LinkedListDeque) AddLast(value interface{}) {
	self.ll.Append(value)
}


func (self *LinkedListDeque) Contains(value interface{}) bool {
	//return self.ll.Contains(value)
}


func (self *LinkedListDeque) GetFirst(value interface{}) (interface{}, bool) {
	value, ok := self.ll.GetFront()
	return value, ok
}

func (self *LinkedListDeque) GetLast(value interface{})(interface{}, bool)  {
	value, ok := self.ll.GetLast()
	return value, ok
}


func (self *LinkedListDeque) RemoveFirst(value interface{}) {
	self.ll.RemoveFront(value)
}

func (self *LinkedListDeque) RemoveLast(value interface{}) {
	self.ll.RemoveBack(value)
}


func (self *LinkedListDeque) Size() int {
	return self.ll.Size()
}