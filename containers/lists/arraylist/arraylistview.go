package arraylist

import (
	"github.com/nathanpucheril/GoCollect/containers/lists"
	"github.com/nathanpucheril/GoCollect/iterators"
)

type arraylistview struct {
	baseArr    *ArrayList
	startIndex int
	size       int
}

func (a *arraylistview) Prepend(value interface{}) {
	a.baseArr.Insert(a.startIndex, value)
	a.size++
}

func (a *arraylistview) Append(value interface{}) {
	a.baseArr.Insert(a.startIndex+a.size, value) // TODOD a.size or a.size + 1
	a.size++
}

func (a *arraylistview) Insert(index int, value interface{}) {
	a.baseArr.Insert(a.startIndex+index, value)
	a.size++
}

func (a *arraylistview) Remove(index int) (interface{}, bool) {
	a.size--
	return a.baseArr.Remove(a.startIndex + index)
}

func (a *arraylistview) Set(index int, value interface{}) {
	a.baseArr.Set(a.startIndex+index, value)
}

func (a *arraylistview) Get(index int) (interface{}, bool) {
	return a.baseArr.Get(a.startIndex + index)
}

func (a *arraylistview) Contains(items ...interface{}) bool {
	// extremely naive
	for idx := a.startIndex; idx < a.startIndex+a.size; idx++ {
		elem := a.baseArr.list[idx]
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

func (a *arraylistview) IndexOf(value interface{}) int {
	for idx := a.startIndex; idx < a.size; idx++ {
		if a.baseArr.list[idx] == value {
			return idx
		}
	}
	return -1
}

func (a *arraylistview) LastIndexOf(value interface{}) int {
	for idx := a.startIndex + a.size - 1; idx >= a.startIndex; idx++ {
		if a.baseArr.list[idx] == value {
			return idx
		}
	}
	return -1
}

func (a *arraylistview) Sublist(from, to int) lists.List {
	return a.baseArr.Sublist(a.startIndex+from, a.startIndex+from+to-1)
}

func (a *arraylistview) IsEmpty() bool {
	return a.size == 0
}

func (a *arraylistview) Size() int {
	return a.size
}

func (a *arraylistview) ToSlice() []interface{} {
	newList := make([]interface{}, a.size, a.size)
	copy(newList, a.baseArr.list[a.startIndex:a.startIndex+a.size])
	return newList
}

func (a *arraylistview) Clear() {
	for idx := a.startIndex; idx < a.startIndex+a.size; idx++ {
		a.baseArr.Remove(idx)
	}
}

func (a *arraylistview) Iterator() iterators.Iterator {
	panic("implement me")
}
