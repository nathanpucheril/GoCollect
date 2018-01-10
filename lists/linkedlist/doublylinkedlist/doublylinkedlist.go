package doublylinkedlist

import (
	"container/list"
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists"
)

type DoublyLinkedList struct {
	sentinelFront *dlnode
	sentinelBack  *dlnode
	size          int
}



func (self *DoublyLinkedList) Set(index int, value interface{}) {
	var curr *dlnode
	if index <= self.size/2 {
		for i := 0; i < index; i++ {
			curr = curr.next
		}
	} else {
		for i := 0; i < self.size-index-1; i++ {
			curr = curr.prev
		}

	}
	curr.value = value
}

func (self *DoublyLinkedList) IndexOf(value interface{}) int {
	panic("implement me")
}

func (self *DoublyLinkedList) LastIndexOf(value interface{}) int {
	panic("implement me")
}

func (self *DoublyLinkedList) Sublist(to, from int) lists.List {
	panic("implement me")
}

func (self *DoublyLinkedList) IsEmpty() bool {
	return self.size == 0
}

func (self *DoublyLinkedList) Clear() {
	sentinelFront, sentinelBack := &dlnode{}, &dlnode{}
	sentinelFront.next = sentinelBack
	sentinelBack.prev = sentinelFront
	// reset
	self.size = 0
	self.sentinelFront = sentinelFront
	self.sentinelBack = sentinelBack
}

func (self *DoublyLinkedList) Iterator() containers.Iterator {
	panic("implement me")
}

func NewDLL() *DoublyLinkedList {
	sentinelFront, sentinelBack := &dlnode{}, &dlnode{}
	sentinelFront.next = sentinelBack
	sentinelBack.prev = sentinelFront
	return &DoublyLinkedList{sentinelFront:sentinelFront, sentinelBack:sentinelBack}
}

func (self *DoublyLinkedList) Insert(index int, value interface{}) {
	var prev, next *dlnode
	if index <= self.size/2 {
		prev, next = self.sentinelFront, self.sentinelFront.next
		for i := 0; i < index; i++ {
			prev, next = next, next.next
		}
	} else {
		next, prev = self.sentinelBack, self.sentinelBack.prev
		for i := 0; i < self.size-index-1; i++ {
			prev, next = prev.prev, prev
		}

	}

	nodeToInsert := &dlnode{value, prev, next}
	prev.next = nodeToInsert
	next.prev = nodeToInsert

	self.size++
}

func (self *DoublyLinkedList) Prepend(value interface{}) {
	self.Insert(0, value)
}

func (self *DoublyLinkedList) Append(value interface{}) {
	self.Insert(self.Size()-1, value)
}


// TODO with iterator
func (self *DoublyLinkedList) Contains(values ...interface{}) bool {
	for _, value := range values {
		it := NewDLLIterator(self)
		found := false
		for it.HasNext()  {
			listItem, err :=it.Next()
			if err != nil {
				panic(err)
			}
			if value == listItem {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}



func (self *DoublyLinkedList) Get(index int) (interface{}, bool) {
	if self.size == 0 {
		return nil, false
	}
	return self.sentinelFront.next.value, true
}

func (self *DoublyLinkedList) GetFront() (interface{}, bool) {
	return self.Get(0)
}

func (self *DoublyLinkedList) GetBack() (interface{}, bool) {
	return self.Get(self.size - 1)
}



func (self *DoublyLinkedList) RemoveAtIndex(index int) (interface{}, error){
	return nil, nil
}

func (self *DoublyLinkedList) RemoveValue(value interface{})(interface{}, error) {
 return nil, nil
}

func (self *DoublyLinkedList) RemoveFront() (interface{}, error){
	return self.RemoveAtIndex(0)
}

func (self *DoublyLinkedList) RemoveBack() (interface{}, error){
	return self.RemoveAtIndex(self.Size()-1)
}



func (self *DoublyLinkedList) Size() int {
	return self.size
}

func (self *DoublyLinkedList) ToSlice() []interface{} {
	slice := make([]interface{}, self.size)
	curr := self.sentinelFront.next
	for i := 0; i < self.size; i++ {
		slice[i] = curr.value
		curr = curr.next
	}
	return slice
}

type dlnode struct {
	value interface{}
	prev  *dlnode
	next  *dlnode
}
