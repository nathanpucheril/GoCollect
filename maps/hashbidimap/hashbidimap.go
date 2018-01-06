package hashbidimap

// NOTE: HashBidiMaps are by nature 1-to-1

type HashBidiMap struct {
	preimage  map[interface{}]interface{}
	postimage map[interface{}]interface{}
}

func New() HashBidiMap {
	return HashBidiMap{make(map[interface{}]interface{}), make(map[interface{}]interface{})}
}

func (self *HashBidiMap) Clear() {
	self.preimage = make(map[interface{}]interface{})
	self.postimage = make(map[interface{}]interface{})
}

func (self *HashBidiMap) ContainsKey(key interface{}) bool {
	_, ok := self.preimage[key]
	return ok
}

func (self *HashBidiMap) ContainsValue(value interface{}) interface{} {
	_, ok := self.postimage[value]
	return ok
}

func (self *HashBidiMap) Get(key interface{}) (interface{}, bool) {
	value, ok := self.preimage[key]
	return value, ok
}

func (self *HashBidiMap) GetKey(value interface{}) (interface{}, bool) {
	key, ok := self.postimage[value]
	return key, ok
}

func (self *HashBidiMap) KeySet() {

}

func (self *HashBidiMap) IsEmpty() bool {
	return len(self.preimage) == 0
}

func (self *HashBidiMap) Put(key, value interface{}) {
	self.preimage[key] = value
	self.postimage[value] = key
}

func (self *HashBidiMap) PutAll() {

}

func (self *HashBidiMap) Remove(key interface{}) bool {
	value, ok := self.preimage[key]
	delete(self.preimage, key)
	delete(self.postimage, value)
	return ok
}

func (self *HashBidiMap) RemoveValue(value interface{}) bool {
	key, ok := self.preimage[value]
	delete(self.postimage, key)
	delete(self.preimage, value)
	return ok
}

func (self *HashBidiMap) Size() int {
	return len(self.preimage)
}

func (self *HashBidiMap) Values() {
	values := make([]interface{}, len(self.postimage))

	i := 0
	for v := range self.postimage {
		values[i] = v
		i++
	}
}
