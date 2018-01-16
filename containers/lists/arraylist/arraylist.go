package arraylist

import (
	"github.com/nathanpucheril/GoCollect/containers/lists"
	"github.com/nathanpucheril/GoCollect/iterators"
)

type ArrayList struct{}

func (*ArrayList) Prepend(value interface{}) {
	panic("implement me")
}

func (*ArrayList) Append(value interface{}) {
	panic("implement me")
}

func (*ArrayList) Insert(index int, value interface{}) {
	panic("implement me")
}

func (*ArrayList) Remove(index int) (interface{}, bool) {
	panic("implement me")
}

func (*ArrayList) Set(index int, value interface{}) {
	panic("implement me")
}

func (*ArrayList) Contains(items ...interface{}) bool {
	panic("implement me")
}

func (*ArrayList) IndexOf(value interface{}) int {
	panic("implement me")
}

func (*ArrayList) LastIndexOf(value interface{}) int {
	panic("implement me")
}

func (*ArrayList) Sublist(to, from int) lists.List {
	panic("implement me")
}

func (*ArrayList) IsEmpty() bool {
	panic("implement me")
}

func (*ArrayList) Size() int {
	panic("implement me")
}

func (*ArrayList) ToSlice() []interface{} {
	panic("implement me")
}

func (*ArrayList) Clear() {
	panic("implement me")
}

func (self *ArrayList) Iterator() iterators.Iterator {
	var list lists.List = self
	return lists.NewListIterator(&list)
}

func (self *ArrayList) ReverseIterator() iterators.Iterator {
	panic("implement")
	//var list lists.List = *self
	//return lists.NewIndexIterator(&list)
}

func New() ArrayList {
	return ArrayList{}
}

func (*ArrayList) Get(index int) (interface{}, bool) {
	return nil, false
}
