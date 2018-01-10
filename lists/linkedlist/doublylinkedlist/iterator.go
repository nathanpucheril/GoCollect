package doublylinkedlist

import (
	"github.com/nathanpucheril/GoCollect"
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/pkg/errors"
)

type DLLIterator struct {
	node *dlnode
	list *DoublyLinkedList
	containers.Iterator
}

func NewDLLIterator(list *DoublyLinkedList)  (DLLIterator) {
	return DLLIterator{node:list.sentinelFront, list:list}
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
