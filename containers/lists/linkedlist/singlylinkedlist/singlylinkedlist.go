package singlylinkedlist

import (
	"github.com/nathanpucheril/GoCollect/containers/lists"
	"github.com/nathanpucheril/GoCollect/iterators"
)

var _ lists.List = (*SinglyLinkedList)(nil)

type SinglyLinkedList struct {
	sentinelFront *slnode
	tailPtr       *slnode
	size          int
	lists.List
}

func (self *SinglyLinkedList) Set(index int, value interface{}) {
	panic("implement me")
}

func (self *SinglyLinkedList) Get(index int) interface{} {
	// TODO: check index validity
	curr := self.sentinelFront.next
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.value
}

func (self *SinglyLinkedList) IndexOf(value interface{}) int {
	panic("implement me")
}

func (self *SinglyLinkedList) LastIndexOf(value interface{}) int {
	panic("implement me")
}

func (self *SinglyLinkedList) Sublist(to, from int) lists.List {
	panic("implement me")
}

func (self *SinglyLinkedList) IsEmpty() bool {
	panic("implement me")
}

func (self *SinglyLinkedList) Clear() {
	panic("implement me")
}

func NewSLL() *SinglyLinkedList {
	sentinel := &slnode{}
	return &SinglyLinkedList{sentinelFront: sentinel, tailPtr: sentinel}
}

func (self *SinglyLinkedList) Insert(index int, value interface{}) {
	if index == self.size-1 {
		self.Append(value)
	}
	curr, next := self.sentinelFront, self.sentinelFront.next
	for i := 0; i < index; i++ {
		curr, next = next, next.next
	}

	curr.next = &slnode{value, next}
	if self.size == 0 { // need to set tailptr,
		self.tailPtr = curr.next
	}
	self.size++
}

func (self *SinglyLinkedList) Prepend(value interface{}) {
	self.Insert(0, value)
}

func (self *SinglyLinkedList) Append(value interface{}) {
	self.tailPtr.next = &slnode{value: value}
	self.tailPtr = self.tailPtr.next
	self.size++
}

// TODO with iterator
func (self *SinglyLinkedList) Contains(values ...interface{}) bool {
	//for _, value := range values {
	//	for _, listItem := range
	//}
	return false
}

func (self *SinglyLinkedList) Remove(index int) interface{} {
	panic("implement")
}

func (self *SinglyLinkedList) RemoveFront() interface{} {
	return self.Remove(0)
}

func (self *SinglyLinkedList) RemoveBack() interface{} {
	return self.Remove(self.Size() - 1)
}

func (self *SinglyLinkedList) Size() int {
	return self.size
}

func (self *SinglyLinkedList) ToSlice() []interface{} {
	slice := make([]interface{}, self.size)
	curr := self.sentinelFront.next
	for i := 0; i < self.size; i++ {
		slice[i] = curr.value
		curr = curr.next
	}
	return slice
}

//func (self *SinglyLinkedList) IsEmpty() bool {
//	return self.size == 0
//
//}

func (self *SinglyLinkedList) Iterator() iterators.Iterator {
	sllIt := NewSLLIterator(self)
	return &sllIt
}

type slnode struct {
	value interface{}
	next  *slnode
}
