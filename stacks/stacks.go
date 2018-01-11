package stacks

import "github.com/nathanpucheril/GoCollect/containers"

type Stack interface {
	Push(value interface{})
	Pop() (interface{}, bool)
	Peek() (interface{}, bool)

	containers.Container
}
