package singlylinkedlist

import (
	"github.com/nathanpucheril/GoCollect/iterators"
	"github.com/pkg/errors"
)

type SLLIterator struct {
	node *slnode
	list *SinglyLinkedList
	iterators.Iterator
}

func NewSLLIterator(list *SinglyLinkedList) SLLIterator {
	return SLLIterator{node: list.sentinelFront, list: list}
}

func (self *SLLIterator) HasNext() bool {
	return *self.node != *self.list.tailPtr
}

func (self *SLLIterator) Next() (interface{}, error) {
	if self.HasNext() {
		value := self.node.next.value
		self.node = self.node.next
		return value, nil
	}
	return nil, errors.New("StopIteration error")
}
