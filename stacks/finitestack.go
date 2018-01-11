package stacks

import "github.com/nathanpucheril/GoCollect/containers"

type FiniteStack struct {
	arr []interface{}
	capacity int
	index int
}

func (self *FiniteStack) Push(value interface{}) {
	if len(self.arr) < self.capacity {
		self.arr[self.index] = value
		self.index++
	}
	panic("")
}

func (self *FiniteStack) Pop() (interface{}, bool) {
	if len(self.arr) > 0 {
		ret := self.arr[self.index]
		self.index--
		return ret, true
	}
	return nil, false
}

func (self *FiniteStack) Peek() (interface{}, bool) {
	if len(self.arr) > 0 {
		return self.arr[self.index], true
	}
	return nil, false
}

func (self *FiniteStack) IsEmpty() bool {
	return len(self.arr) == 0
}

func (self *FiniteStack) Size() int {
	return len(self.arr)
}

func (self *FiniteStack) ToSlice() []interface{} {
	return self.arr
}

func (self *FiniteStack) Clear() {
	self.arr = make([]interface{}, self.capacity)
}

func (FiniteStack) Iterator() containers.Iterator {
	panic("implement me")
}

func NewFiniteStack(capacity int) FiniteStack {
	return FiniteStack{arr:make([]interface{}, capacity), capacity:capacity}
}

