package containers

type Container interface {
	IsEmpty() bool
	Size() int
	ToSlice() []interface{}
	Clear()

	Iterable
}

