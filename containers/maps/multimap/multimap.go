package multimap

import (
	"github.com/nathanpucheril/GoCollect/containers/maps"
	"github.com/nathanpucheril/GoCollect/containers/sets/hashset"
	"github.com/nathanpucheril/GoCollect/iterators"
)

type MultiMap struct {
	hmap map[interface{}]hashset.HashSet
	size int
}

func (self *MultiMap) PutAllMultiMap(m maps.Multimap) {
	panic("implement me")
}

func (self *MultiMap) Remove(key interface{}, values ...interface{}) bool {
	if set, ok := self.hmap[key]; ok {
		oldSetSize := set.Size()
		set.Remove(values)
		newSetSize := set.Size()
		self.size += newSetSize - oldSetSize
		if newSetSize == 0 { // special case to delete empty sets from multimap
			delete(self.hmap, key)
		}
		return newSetSize-oldSetSize < 0
	}
	return false
}

func (self *MultiMap) Put(key interface{}, values ...interface{}) bool {
	if set, ok := self.hmap[key]; ok {
		oldSetSize := set.Size()
		set.Add(values)
		newSetSize := set.Size()
		self.size += newSetSize - oldSetSize
		return newSetSize-oldSetSize > 0
	} else {
		self.hmap[key] = hashset.New()
		return self.Put(key, values)
	}
}

func (self *MultiMap) Get(key interface{}) (interface{}, bool) {
	panic("implement me")
}

func (self *MultiMap) ContainsEntry(key, value interface{}) bool {
	if set, ok := self.hmap[key]; ok {
		return set.Contains(value)
	}
	return false
}

func (self *MultiMap) ContainsKey(key interface{}) bool {
	_, ok := self.hmap[key]
	return ok
}

func (self *MultiMap) ContainsValue(value interface{}) bool {
	for _, set := range self.hmap {
		if set.Contains(value) {
			return true
		}
	}
	return false

}

func (self *MultiMap) Values() []interface{} {
	panic("implement me")
}

func (self *MultiMap) IsEmpty() bool {
	return self.size == 0
}

func (self *MultiMap) Size() int {
	return self.size
}

func (self *MultiMap) ToSlice() []interface{} {
	panic("implement me")
}

func (self *MultiMap) Clear() {
	self.size = 0
	self.hmap = make(map[interface{}]hashset.HashSet)
}

func (self *MultiMap) Iterator() iterators.Iterator {
	panic("implement me")
}
