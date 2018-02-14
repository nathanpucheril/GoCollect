package fuzzymap

import (
	"github.com/nathanpucheril/GoCollect/containers/maps"
	"github.com/nathanpucheril/GoCollect/containers/maps/hashmap"
)

type FuzzyMap struct {
	writeFn, readFn func(interface{}) interface{}
	baseMap         maps.Map
}

func NewFuzzyHashMap(writeFn, readFn func(interface{}) interface{}) FuzzyMap {
	hmap := hashmap.New()
	var baseMap maps.Map = &hmap
	return FuzzyMap{writeFn, readFn, baseMap}
}

func FuzzifyMap(baseMap maps.Map, writeFn, readFn func(interface{}) interface{}) FuzzyMap {
	return FuzzyMap{writeFn, readFn, baseMap}
}

func (f *FuzzyMap) Put(key, value interface{}) bool {
	return f.baseMap.Put(f.writeFn(key), value)
}

func (f *FuzzyMap) PutAll(m maps.Map) {
	panic("implement me")
}

func (f *FuzzyMap) Get(key interface{}) (interface{}, bool) {
	return f.baseMap.Get(f.readFn(key))
}

func (f *FuzzyMap) Remove(key interface{}) bool {
	return f.baseMap.Remove(f.writeFn(key))
}

func (f *FuzzyMap) ContainsEntry(key, value interface{}) bool {
	return f.baseMap.ContainsEntry(f.readFn(key), value)
}

func (f *FuzzyMap) ContainsKey(key interface{}) bool {
	return f.baseMap.ContainsKey(f.readFn(key))
}
