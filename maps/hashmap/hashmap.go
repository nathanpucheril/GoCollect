package hashmap

import "github.com/nathanpucheril/GoCollect/maps"

type HashMap struct {
	hmap map[interface{}]interface{}
}

func New() HashMap {
	return HashMap{hmap:make(map[interface{}]interface{})}
}

func NewFromGoMap(hmap map[interface{}]interface{}) HashMap {
	// note that hmap is a copy of the original map passed in. This is safe
	return HashMap{hmap:hmap}
}

func (self *HashMap) Clear() {
	self.hmap = make(map[interface{}]interface{})
}

func (self *HashMap) ContainsKey(key interface{}) bool {
	_, ok := self.hmap[key]
	return ok
}

func (self *HashMap) ContainsValue(value interface{}) bool {
	_, ok := self.hmap[value]
	return ok
}

func (self *HashMap) Get(key interface{}) (interface{}, bool) {
	value, ok := self.hmap[key]
	return value, ok
}

//func (self *HashMap) KeySet() {
//
//}

func (self *HashMap) IsEmpty() bool {
	return len(self.hmap) == 0
}

func (self *HashMap) Put(key, value interface{}) bool {
	self.hmap[key] = value
}

func (self *HashMap) PutAll(m maps.Map) {

}

func (self *HashMap) Remove(key interface{}) bool {
	ret := self.ContainsKey(key)
	delete(self.hmap, key)
	return ret
}

func (self *HashMap) Size() int {
	return len(self.hmap)
}

func (self *HashMap) Values() []interface{} {
	values := make([]interface{}, len(self.hmap))

	i := 0
	for _, value := range self.hmap {
		values[i] = value
		i++
	}
	return values
}
