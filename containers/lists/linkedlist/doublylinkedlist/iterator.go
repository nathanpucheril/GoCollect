package doublylinkedlist

import (
	"github.com/nathanpucheril/GoCollect/iterators"
	"github.com/pkg/errors"
)

type DLLIterator struct {
	node *dlnode
	list *DoublyLinkedList
	iterators.Iterator
}

func NewDLLIterator(list *DoublyLinkedList) DLLIterator {
	return DLLIterator{node: list.sentinelFront, list: list}
}

func (self *DLLIterator) HasNext() bool {
	return self.node.next != self.list.sentinelBack
}

func (self *DLLIterator) Next() (interface{}, error) {
	if self.HasNext() {
		value := self.node.next.value
		self.node = self.node.next
		return value, nil
	}
	return nil, errors.New("StopIteration error")
}

type DLLReverseIterator struct {
	node *dlnode
	list *DoublyLinkedList
	iterators.Iterator
}

func NewDLLReverseIterator(list *DoublyLinkedList) DLLIterator {
	return DLLIterator{node: list.sentinelBack, list: list}
}

func (self *DLLReverseIterator) HasNext() bool {
	return self.node.prev != self.list.sentinelFront
}

func (self *DLLReverseIterator) Next() (interface{}, error) {
	if self.HasNext() {
		value := self.node.prev.value
		self.node = self.node.prev
		return value, nil
	}
	return nil, errors.New("StopIteration error")
}
