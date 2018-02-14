package hashbimap

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/containers/maps"
	"github.com/nathanpucheril/GoCollect/containers/sets"
	"github.com/nathanpucheril/GoCollect/iterators"
)

// NOTE: HashBiMaps are by nature 1-to-1

type HashBiMap struct {
	preimage  map[interface{}]interface{}
	postimage map[interface{}]interface{}
}

func (self *HashBiMap) ContainsEntry(key, value interface{}) bool {
	actual, ok := self.preimage[key]
	return ok && actual == value
}

func New() HashBiMap {
	return HashBiMap{preimage: make(map[interface{}]interface{}), postimage: make(map[interface{}]interface{})}
}

func (self *HashBiMap) Clear() {
	self.preimage = make(map[interface{}]interface{})
	self.postimage = make(map[interface{}]interface{})
}

func (self *HashBiMap) ContainsKey(key interface{}) bool {
	_, ok := self.preimage[key]
	return ok
}

func (self *HashBiMap) ContainsValue(value interface{}) bool {
	_, ok := self.postimage[value]
	return ok
}

func (self *HashBiMap) Get(key interface{}) (interface{}, bool) {
	value, ok := self.preimage[key]
	return value, ok
}

func (self *HashBiMap) GetKey(value interface{}) (interface{}, bool) {
	key, ok := self.postimage[value]
	return key, ok
}

func (self *HashBiMap) KeySet() sets.Set {
	panic("implement me")
}

func (self *HashBiMap) IsEmpty() bool {
	return len(self.preimage) == 0
}

func (self *HashBiMap) Put(key, value interface{}) bool {
	if self.ContainsEntry(key, value) {
		return false
	}
	self.preimage[key] = value
	self.postimage[value] = key
	return true // something was modified
}

func (self *HashBiMap) PutAll(m maps.Map) {
	panic("implement me")
	//for k, v := range m.
}

func (self *HashBiMap) Remove(key interface{}) bool {
	value, ok := self.preimage[key]
	delete(self.preimage, key)
	delete(self.postimage, value)
	return ok
}

func (self *HashBiMap) RemoveValue(value interface{}) bool {
	key, ok := self.preimage[value]
	delete(self.postimage, key)
	delete(self.preimage, value)
	return ok
}

func (self *HashBiMap) Size() int {
	return len(self.preimage)
}

func (self *HashBiMap) Values() []interface{} {
	values := make([]interface{}, len(self.postimage))

	i := 0
	for v := range self.postimage {
		values[i] = v
		i++
	}
}

func (self *HashBiMap) ToSlice() []interface{} {
	panic("implement me")
}

func (self *HashBiMap) Iterator() iterators.Iterator {
	panic("implement me")
}
