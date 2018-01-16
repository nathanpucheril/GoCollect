package queues

import "github.com/nathanpucheril/GoCollect/containers"

type Queue interface {
	Add(value interface{})
	Get() (interface{}, bool)
	Peek() (interface{}, bool)

	containers.Container
}
