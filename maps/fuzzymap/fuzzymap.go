package fuzzymap

import "github.com/nathanpucheril/GoCollect/maps"

type FuzzyMap struct {
	writeFn, readFn func(interface{}) interface{}
	baseMap         maps.Map
}

func (self *FuzzyMap) Put(key, value interface{}) bool {
	return self.baseMap.Put(self.writeFn(key), value)
}

func (self *FuzzyMap) PutAll(m maps.Map) {
	panic("implement me")
}

func (self *FuzzyMap) Get(key interface{}) (interface{}, bool) {
	return self.baseMap.Get(self.readFn(key))
}

func (self *FuzzyMap) Remove(key interface{}) bool {
	return self.baseMap.Remove(self.writeFn(key))
}

func (self *FuzzyMap) ContainsEntry(key, value interface{}) bool {
	return self.baseMap.ContainsEntry(self.readFn(key), value)
}

func (self *FuzzyMap) ContainsKey(key interface{}) bool {
	return self.baseMap.ContainsKey(self.readFn(key))
}
