package containers

type Iterator interface {
	Next() (interface{}, error)
	HasNext() bool
	//Remove() interface{}
}

type Iterable interface {
	Iterator() Iterator
}