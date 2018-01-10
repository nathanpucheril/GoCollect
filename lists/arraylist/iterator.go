package arraylist

import (
	"github.com/nathanpucheril/GoCollect"
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/pkg/errors"
)

type IndexIterator struct {
	index int
	list *ArrayList

	containers.Iterator
}

func NewIndexIterator(list *ArrayList)  (IndexIterator) {
	return IndexIterator{index:0, list:list}
}

func (self *IndexIterator) HasNext() bool {
	return self.index < self.list.Size()
}

func (self *IndexIterator) Next() (interface{}, error) {
	if self.HasNext() {
		value,_ := self.list.Get(self.index)
		self.index++
		return value, nil
	}
	return nil, errors.New("StopIteration error")
}
