package stacks

import "github.com/nathanpucheril/GoCollect/containers"

type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)

	containers.Container
}
