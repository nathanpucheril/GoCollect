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

	Remove() (interface{}, error)
	RemoveFirst() (interface{}, error)
	RemoveLast() (interface{}, error)

	Peek()(interface{}, error)
	PeekFirst() (interface{}, error)
	PeekLast() (interface{}, error)

	Poll() (interface{}, bool)
	PollFirst() (interface{}, bool)
	PollLast() (interface{}, bool)

	Contains(items ...interface{}) bool

	Push(value interface{})
	Pop()(interface{}, error)

	containers.Container
}
