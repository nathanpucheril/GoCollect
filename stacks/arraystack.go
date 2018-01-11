package stacks

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists/arraylist"
)

type ArrayStack struct {
	*arraylist.ArrayList
	Stack
}

func (self *ArrayStack) Push(value interface{}) {
	self.ArrayList.Append(value)
}

func (self *ArrayStack) Pop() (interface{}, bool) {
	return self.ArrayList.Remove(self.Size() - 1)
}

func (self *ArrayStack) Peek() (interface{}, bool) {
	return self.ArrayList.Get(self.Size() - 1)
}

func (self *LinkedListStack) Iterator() containers.Iterator {
	return self.DoublyLinkedList.Iterator()// TODO need to do in reverse
}

func NewArrayListStack() ArrayStack {
	arr := arraylist.New()
	return ArrayStack{&arr, nil}
}

