package maps

import "github.com/nathanpucheril/GoCollect/containers"

type Map interface {
	Put(key, value interface{}) bool
	PutAll(m Map)

	Get(key interface{}) (interface{}, bool)

	Remove(key interface{}) bool

	ContainsEntry(key, value interface{}) bool
	ContainsKey(key interface{}) bool
	ContainsValue(key interface{}) bool

	//KeySet()
	Values() []interface{}

	containers.Container
}

// inspiration from Google's Java library, Guava
type Multimap interface {
	Put(key, value interface{}) bool
	PutAll(key interface{}, value... interface{})
	PutAllMultiMap(m Multimap)

	Get(key interface{}) (interface{}, bool)

	Remove(key, value interface{}) bool
	RemoveAll(key interface{}, value ...interface{}) bool

	ContainsEntry(key, value interface{}) bool
	ContainsKey(key interface{}) bool
	ContainsValue(key interface{}) bool

	//KeySet()
	Values() []interface{}

	containers.Container
}