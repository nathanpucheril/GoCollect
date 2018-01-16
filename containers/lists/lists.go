package lists

import (
	"github.com/nathanpucheril/GoCollect/containers"
)

type List interface {
	Prepend(value interface{})
	Append(value interface{})
	Insert(index int, value interface{})

	Remove(index int) (interface{}, bool)
	Set(index int, value interface{})
	Get(index int) (interface{}, bool)
	Contains(items ...interface{}) bool

	IndexOf(value interface{}) int
	LastIndexOf(value interface{}) int

	Sublist(to, from int) List

	containers.Container
}
