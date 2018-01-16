package arraydeque

import (
	"github.com/nathanpucheril/GoCollect/containers/lists/arraylist"
)

type ArrayDeque struct {
	*arraylist.ArrayList
}

func (self *ArrayDeque) Peek() (interface{}, bool) {
	return self.ArrayList.Get(0)
}

func (self *ArrayDeque) PeekFirst() (interface{}, bool) {
	return self.ArrayList.Get(0)
}

func (self *ArrayDeque) PeekLast() (interface{}, bool) {
	return self.ArrayList.Get(self.Size() - 1)
}

func (self *ArrayDeque) Push(value interface{}) {
	panic("implement me")
}

func (self *ArrayDeque) Pop() (interface{}, bool) {
	panic("implement me")
}

func NewArrayDeque() ArrayDeque {
	arr := arraylist.New()
	return ArrayDeque{&arr}
}

func (self *ArrayDeque) Add(value interface{}) {
	self.ArrayList.Prepend(value)
}

func (self *ArrayDeque) AddFirst(value interface{}) {
	self.ArrayList.Prepend(value)
}

func (self *ArrayDeque) AddLast(value interface{}) {
	self.ArrayList.Append(value)
}

func (self *ArrayDeque) Remove() (interface{}, bool) {
	return self.ArrayList.Remove(0)
}

func (self *ArrayDeque) RemoveFirst() (interface{}, bool) {
	return self.ArrayList.Remove(0)
}

func (self *ArrayDeque) RemoveLast() (interface{}, bool) {
	return self.ArrayList.Remove(self.Size() - 1)
}
