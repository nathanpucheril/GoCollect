package maps

import "github.com/nathanpucheril/GoCollect/containers"

type Map interface {
	Put(key, value interface{}) bool
	PutAll(m Map)

	Get(key interface{}) (interface{}, bool)

	Remove(key interface{}) bool

	ContainsEntry(key, value interface{}) bool
	ContainsKey(key interface{}) bool

	//KeySet()
	Values() []interface{}

	containers.Container
}

// inspiration from Guava (Google's Standard Java library)
type Multimap interface {
	Put(key interface{}, value ...interface{}) bool
	PutAllMultiMap(m Multimap)

	Get(key interface{}) (interface{}, bool)

	Remove(key interface{}, value ...interface{}) bool

	ContainsEntry(key, value interface{}) bool
	ContainsKey(key interface{}) bool

	//KeySet()
	Values() []interface{}

	containers.Container
}

type MapEntry interface {
	GetKey() interface{}
	GetValue() interface{}
}
