package queues

import "github.com/nathanpucheril/GoCollect/containers/lists/linkedlist/singlylinkedlist"

type LinkedQueue struct {
	*singlylinkedlist.SinglyLinkedList
}

func NewLinkedQueue() LinkedQueue {
	return LinkedQueue{singlylinkedlist.NewSLL()}
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
