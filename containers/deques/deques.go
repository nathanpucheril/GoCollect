package deques

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/containers/queues"
)

type Deque interface {
	//Add(value interface{})
	AddFirst(value interface{})
	AddLast(value interface{})

	//Poll() (interface{}, bool)
	PollFirst() (interface{}, bool)
	PollLast() (interface{}, bool)

	//Peek() (interface{}, bool)
	PeekFirst() (interface{}, bool)
	PeekLast() (interface{}, bool)

	Contains(items ...interface{}) bool

	Push(value interface{})
	Pop() (interface{}, bool)

	//containers.Container
	queues.Queue
}
