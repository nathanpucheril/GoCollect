package arraylist

import (
	"github.com/nathanpucheril/GoCollect/containers/lists"
	"github.com/nathanpucheril/GoCollect/iterators"
)

type ArrayList struct {
	list []interface{}
	size int
}

const INITIALCAP = 16
const RESIZEFACTOR = 2

var SHRINKFACTOR = .75

func New() *ArrayList {
	return &ArrayList{make([]interface{}, INITIALCAP, INITIALCAP), 0}
}

func NewFromList(l *lists.List) *ArrayList {
	panic("implement")
	//return &ArrayList{make([]interface{},INITIALCAP, INITIALCAP), 0}
}

func (a *ArrayList) grow() {
	oldList := a.list
	newSize := len(oldList) * RESIZEFACTOR
	a.list = make([]interface{}, newSize, newSize)
	copy(a.list, oldList)
}

func (a *ArrayList) shrink() {
	oldList := a.list
	newSize := int(float64(a.size) * SHRINKFACTOR)
	if newSize < INITIALCAP {
		newSize = INITIALCAP
	}
	a.list = make([]interface{}, newSize, newSize)
	copy(a.list, oldList)
}

func (a *ArrayList) Prepend(value interface{}) {
	if a.size >= cap(a.list) {
		a.grow()
	}
	prevItem := value
	for idx, item := range a.list {
		a.list[idx] = prevItem
		prevItem = item
	}
	a.size++
}

func (a *ArrayList) Append(value interface{}) {
	if a.size >= cap(a.list) {
		a.grow()
	}
	a.list[a.size] = value
	a.size++
}

func (a *ArrayList) Insert(index int, value interface{}) {
	if a.size >= cap(a.list) {
		a.grow()
	}
	prevItem := value
	for idx := index; idx < len(a.list); idx++ {
		a.list[idx], prevItem = prevItem, a.list[idx]
	}
	a.size++
}

func (a *ArrayList) Remove(index int) interface{} {
	a.validateIndexElsePanic(index)
	removed := a.list[index]
	for idx := index; idx < a.size-1; idx++ {
		a.list[idx] = a.list[idx+1]
	}
	if a.size <= cap(a.list)/RESIZEFACTOR {
		a.shrink()
	}
	a.size--
	return removed
}

func (a *ArrayList) Set(index int, value interface{}) {
	a.validateIndexElsePanic(index)
	a.list[index] = value
}

func (a *ArrayList) Contains(items ...interface{}) bool {
	// extremely naive
	for _, elem := range a.list {
		found := false
		for _, searchItem := range items {
			if elem == searchItem {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (a *ArrayList) IndexOf(value interface{}) int {
	for idx, elem := range a.list {
		if elem == value {
			return idx
		}
	}
	return -1
}

func (a *ArrayList) LastIndexOf(value interface{}) int {
	for idx := len(a.list) - 1; idx >= 0; idx++ {
		if a.list[idx] == value {
			return idx
		}
	}
	return -1
}

func (a *ArrayList) Sublist(from, to int) lists.List {
	a.validateIndexElsePanic(from)
	a.validateIndexElsePanic(to)
	if to > from {
		panic("invalid sublist bounds: from is greater than to")
	}
	var sublist lists.List = &arraylistview{a, from, to - from}
	return sublist
}

func (a *ArrayList) IsEmpty() bool {
	return a.size == 0
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) ToSlice() []interface{} {
	newList := make([]interface{}, a.size, a.size)
	copy(newList, a.list)
	return newList
}

func (a *ArrayList) Clear() {
	a.list = make([]interface{}, 0, INITIALCAP)
}

func (a *ArrayList) Iterator() iterators.Iterator {
	var list lists.List = a
	listIt := lists.NewListIterator(&list)
	var it iterators.Iterator = &listIt
	return it
}

func (a *ArrayList) ReverseIterator() iterators.Iterator {
	panic("implement")
	//var list lists.List = *self
	//return lists.NewIndexIterator(&list)
}

func (a *ArrayList) Get(index int) interface{} {
	a.validateIndexElsePanic(index)
	return a.list[index]
}

func (a *ArrayList) validateIndexElsePanic(index int) {
	if index < 0 || index >= a.size {
		panic("index out of bounds")
	}
}
