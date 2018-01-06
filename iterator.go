package GoCollect

type Iterator interface {
	Next() interface{}
	HasNext() bool
	//Remove() interface{}
}
