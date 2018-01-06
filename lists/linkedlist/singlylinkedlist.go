package linkedlist

type SinglyLinkedList struct {
	sentinelFront *slnode
	tailPtr       *slnode
	size          int
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

func (self *SinglyLinkedList) Remove(index int, value interface{}) {

}

func (self *SinglyLinkedList) RemoveFront(value interface{}) {
	self.Remove(0, value)
}

func (self *SinglyLinkedList) RemoveBack(value interface{}) {
	self.Remove(self.Size()-1, value)
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
func (self *SinglyLinkedList) IsEmpty() bool {
	return self.size == 0

}

type slnode struct {
	value interface{}
	next  *slnode
}
