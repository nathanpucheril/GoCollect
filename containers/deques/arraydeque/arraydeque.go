package arraydeque

import (
	"github.com/nathanpucheril/GoCollect/containers/lists/arraylist"
)

type ArrayDeque struct {
	*arraylist.ArrayList
}

func (ad *ArrayDeque) Peek() (interface{}, bool) {
	if ad.IsEmpty() {
		return nil, false
	}
	return ad.ArrayList.Get(0), true
}

func (ad *ArrayDeque) PeekFirst() (interface{}, bool) {
	if ad.IsEmpty() {
		return nil, false
	}
	return ad.ArrayList.Get(0), true
}

func (ad *ArrayDeque) PeekLast() (interface{}, bool) {
	if ad.IsEmpty() {
		return nil, false
	}
	return ad.ArrayList.Get(ad.Size() - 1), true
}

func (ad *ArrayDeque) Push(value interface{}) {
	panic("implement me")
}

func (ad *ArrayDeque) Pop() (interface{}, bool) {
	panic("implement me")
}

func NewArrayDeque() ArrayDeque {
	arr := arraylist.New()
	return ArrayDeque{arr}
}

func (ad *ArrayDeque) Add(value interface{}) {
	ad.ArrayList.Prepend(value)
}

func (ad *ArrayDeque) AddFirst(value interface{}) {
	ad.ArrayList.Prepend(value)
}

func (ad *ArrayDeque) AddLast(value interface{}) {
	ad.ArrayList.Append(value)
}

func (ad *ArrayDeque) Poll() (interface{}, bool) {
	if ad.IsEmpty() {
		return nil, false
	}
	return ad.ArrayList.Remove(0), true
}

func (ad *ArrayDeque) PollFirst() (interface{}, bool) {
	if ad.IsEmpty() {
		return nil, false
	}
	return ad.ArrayList.Remove(0), true
}

func (ad *ArrayDeque) PollLast() (interface{}, bool) {
	if ad.IsEmpty() {
		return nil, false
	}
	return ad.ArrayList.Remove(ad.Size() - 1), true
}
