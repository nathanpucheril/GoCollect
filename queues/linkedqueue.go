package queues

import (
	"github.com/nathanpucheril/GoCollect/lists/linkedlist/singlylinkedlist"
)

type LinkedQueue struct {
	 *singlylinkedlist.SinglyLinkedList
	 Queue
}

func New() LinkedQueue {
	return LinkedQueue{singlylinkedlist.NewSLL(), nil}
}

func (self *LinkedQueue) Add(value interface{}) {
	self.SinglyLinkedList.Append(value)
}

func (self *LinkedQueue) Get() (interface{}, bool) {
	return self.SinglyLinkedList.RemoveFront()
}

func (self *LinkedQueue) Peek() (interface{}, bool) {
	return self.SinglyLinkedList.Get(0)
}