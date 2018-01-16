package arraystack

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/containers/lists/arraylist"
)

type ArrayStack struct {
	*arraylist.ArrayList
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

func (self *ArrayStack) Iterator() containers.Iterator {
	panic("implement") // TODO: need reverse iterator
}

func NewArrayStack() ArrayStack {
	arr := arraylist.New()
	return ArrayStack{&arr}
}