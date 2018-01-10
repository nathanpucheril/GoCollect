package singlylinkedlist

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists"
)

var _ lists.List = (*SinglyLinkedList)(nil)

type SinglyLinkedList struct {
	sentinelFront *slnode
	tailPtr       *slnode
	size          int
	lists.List
}

func NewSLL() *SinglyLinkedList {
	return &SinglyLinkedList{sentinelFront: &slnode{}}
}

func (self *SinglyLinkedList) Insert(index int, value interface{}) {
	if index == self.size-1 {
		self.InsertBack(value)
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

func (self *SinglyLinkedList) InsertFront(value interface{}) {
	self.Insert(0, value)
}

func (self *SinglyLinkedList) InsertBack(value interface{}) {
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


func (self *SinglyLinkedList) Remove(index int) (interface{}, bool) {

}

func (self *SinglyLinkedList) RemoveFront()  (interface{}, bool) {
	self.Remove(0)
}

func (self *SinglyLinkedList) RemoveBack()  (interface{}, bool) {
	self.Remove(self.Size()-1)
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

func (self *SinglyLinkedList) Iterator() containers.Iterator {
	return nil
}

type slnode struct {
	value interface{}
	next  *slnode
}