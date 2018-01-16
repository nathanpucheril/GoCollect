package deques

import (
	"github.com/nathanpucheril/GoCollect/containers"
)

type Deque interface {
	Add(value interface{})
	AddFirst(value interface{})
	AddLast(value interface{})

	Remove() (interface{}, bool)
	RemoveFirst() (interface{}, bool)
	RemoveLast() (interface{}, bool)

	Peek() (interface{}, bool)
	PeekFirst() (interface{}, bool)
	PeekLast() (interface{}, bool)

	Contains(items ...interface{}) bool

	Push(value interface{})
	Pop() (interface{}, bool)

	containers.Container
}
