package queues

import "github.com/nathanpucheril/GoCollect/containers"

type Queue interface {
	Add(value interface{})
	Poll() (interface{}, bool)
	Peek() (interface{}, bool)

	containers.Container
}
