package queues

type Queue interface {
	Add(value interface{})
	Get() (interface{}, bool)
	Peek() (interface{}, bool)
}