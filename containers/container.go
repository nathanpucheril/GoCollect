package containers

import "github.com/nathanpucheril/GoCollect/iterators"

type Container interface {
	IsEmpty() bool
	Size() int
	ToSlice() []interface{}
	Clear()

	iterators.Iterable
}
