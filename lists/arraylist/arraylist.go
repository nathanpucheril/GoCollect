package arraylist

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists"
)

type ArrayList struct {
	lists.List
}

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

func (*ArrayList) Iterator() containers.Iterator {
	panic("implement me")
}

func New() ArrayList {
	return ArrayList{}
}

func (*ArrayList) Get(index int) (interface{}, bool) {
	return nil, false
}