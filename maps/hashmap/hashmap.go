package hashmap

type HashMap struct {
	hmap map[interface{}]interface{}
}

func New() HashMap {
	return HashMap{make(map[interface{}]interface{})}
}

func (self *HashMap) Clear() {
	self.hmap = make(map[interface{}]interface{})
}

func (self *HashMap) ContainsKey(key interface{}) bool {
	_, ok := self.hmap[key]
	return ok
}

func (self *HashMap) ContainsValue(value interface{}) interface{} {
	_, ok := self.hmap[value]
	return ok
}

func (self *HashMap) Get(key interface{}) (interface{}, bool) {
	value, ok := self.hmap[key]
	return value, ok
}

func (self *HashMap) KeySet() {

}

func (self *HashMap) IsEmpty() bool {
	return len(self.hmap) == 0
}

func (self *HashMap) Put(key, value interface{}) {
	self.hmap[key] = value
}

func (self *HashMap) PutAll() {

}

func (self *HashMap) Remove(key interface{}) bool {
	ret := self.ContainsKey(key)
	delete(self.hmap, key)
	return ret
}

func (self *HashMap) Size() int {
	return len(self.hmap)
}

func (self *HashMap) Values() {
	values := make([]interface{}, len(self.hmap))

	i := 0
	for v := range self.hmap {
		values[i] = v
		i++
	}
}
