package deques

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists/linkedlist"
	"github.com/nathanpucheril/GoCollect/lists/linkedlist/doublylinkedlist"
	"github.com/nathanpucheril/GoCollect/stacks"
)

type Deque interface {
	Add(value interface{})
	AddFirst(value interface{})
	AddLast(value interface{})

	Remove() (interface{}, bool)
	RemoveFirst() (interface{}, bool)
	RemoveLast() (interface{}, bool)

	Peek()(interface{}, bool)
	PeekFirst() (interface{}, bool)
	PeekLast() (interface{}, bool)

	Contains(items ...interface{}) bool

	Push(value interface{})
	Pop()(interface{}, bool)

	containers.Container
}
