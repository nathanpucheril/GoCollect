package linkedlist

type DoublyLinkedList struct {
	sentinelFront *dlnode
	sentinelBack  *dlnode
	size          int
}

func NewDLL() *DoublyLinkedList {
	sentinelFront, sentinelBack := &dlnode{}, &dlnode{}
	sentinelFront.next = sentinelBack
	sentinelBack.prev = sentinelFront
	return &DoublyLinkedList{sentinelFront, sentinelBack, 0}
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

func (self *DoublyLinkedList) InsertFront(value interface{}) {
	self.Insert(0, value)
}

func (self *DoublyLinkedList) InsertBack(value interface{}) {
	self.Insert(self.Size()-1, value)
}

func (self *DoublyLinkedList) Remove(index int, value interface{}) {

}

func (self *DoublyLinkedList) RemoveFront(value interface{}) {
	self.Remove(0, value)
}

func (self *DoublyLinkedList) RemoveBack(value interface{}) {
	self.Remove(self.Size()-1, value)
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
