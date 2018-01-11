package lists

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/pkg/errors"
)

type ListIterator struct {
	index int
	list *List
	containers.Iterator
}

func NewIndexIterator(list *List)  (ListIterator) {
	return ListIterator{list:list}
}

func (self *ListIterator) HasNext() bool {
	return self.index < (*self.list).Size()
}

func (self *ListIterator) Next() (interface{}, error) {
	if self.HasNext() {
		value,_ := (*self.list).Get(self.index)
		self.index++
		return value, nil
	}
	return nil, errors.New("StopIteration error")
}

func (self *ListIterator) NextIndex() int {
	return self.index
}

func (self *ListIterator) PrevIndex() int {
	return self.index - 1
}

func (self *ListIterator) HasPrev() bool {
	return self.index > 0
}

func (self *ListIterator) Prev() (interface{}, error) {
	if self.HasPrev() {
		self.index--
		value,_ := (*self.list).Get(self.index)
		return value, nil
	}
	return nil, errors.New("StopIteration error")
}
