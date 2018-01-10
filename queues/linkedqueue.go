package queues

import "github.com/nathanpucheril/GoCollect/lists/linkedlist/singlylinkedlist"

type LinkedQueue struct {
	sll *singlylinkedlist.SinglyLinkedList
}

func New() LinkedQueue {
	return LinkedQueue{singlylinkedlist.NewSLL()}
}

func (self *LinkedQueue) Add(value interface{}) {
	self.sll.Append(value)
}

func (self *LinkedQueue) Get() (interface{}, bool) {
	return self.sll.RemoveFront()
}

func (self *LinkedQueue) Peek() (interface{}, bool) {
	return self.sll.Get(0)
}