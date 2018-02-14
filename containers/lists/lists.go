package lists

import (
	"github.com/nathanpucheril/GoCollect/comparators"
	"github.com/nathanpucheril/GoCollect/containers"
)

type List interface {
	Prepend(value interface{})
	Append(value interface{})
	Insert(index int, value interface{})

	Remove(index int) interface{}
	Set(index int, value interface{})
	Get(index int) interface{}
	Contains(items ...interface{}) bool

	IndexOf(value interface{}) int
	LastIndexOf(value interface{}) int

	Sublist(to, from int) List

	containers.Container
}

type SortableList struct {
	list       List
	comparator comparators.Comparator
}

func (sl *SortableList) Len() int {
	return sl.list.Size()
}

func (sl *SortableList) Less(i, j int) bool {
	itemI, ok := sl.list.Get(i)
	if !ok {
		panic("shouldn't happen")
	}
	itemJ, ok := sl.list.Get(j)
	if !ok {
		panic("shouldn't happen")
	}
	return sl.comparator(itemI, itemJ) <= 0
}

func (sl *SortableList) Swap(i, j int) {
	itemI, ok := sl.list.Get(i)
	if !ok {
		panic("shouldn't happen")
	}
	itemJ, ok := sl.list.Get(j)
	if !ok {
		panic("shouldn't happen")
	}

	sl.list.Set(i, itemJ)
	sl.list.Set(j, itemI)
}
