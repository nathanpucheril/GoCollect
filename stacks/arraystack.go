package stacks

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists/arraylist"
)

type ArrayStack struct {
	arraylist *arraylist.ArrayList
}

func (self *ArrayStack) Push(value interface{}) {
	self.arraylist.Append(value)
}

func (self *ArrayStack) Pop() (interface{}, error) {
	return self.arraylist.RemoveAtIndex(self.Size() - 1)
}

func (self *ArrayStack) Peek() (interface{}, error) {
	return self.arraylist.Get(self.Size() - 1)
}

func (self *ArrayStack) IsEmpty() bool {
	return self.arraylist.IsEmpty()
}

func (self *ArrayStack) Size() int {
	return self.arraylist.Size()
}

func (self *ArrayStack) ToSlice() []interface{} {
	return self.arraylist.ToSlice()
}

func (self *ArrayStack) Clear() {
	self.arraylist.Clear()
}

func (self *ArrayStack) Iterator() containers.Iterator {
	panic("implement me")
}

func NewArrayListStack() ArrayStack {
	arr := arraylist.New()
	return ArrayStack{arraylist:&arr}
}

