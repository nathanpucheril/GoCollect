package hashset

type HashSet struct {
	set map[interface{}]struct{}
}

type empty struct{}

func New() HashSet {
	return HashSet{make(map[interface{}]struct{})}
}

func (self *HashSet) Clear() {
	self.set = make(map[interface{}]struct{})
}

func (self *HashSet) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := self.set[item]; !ok {
			return false
		}
	}
	return true
}

func (self *HashSet) IsEmpty() bool {
	return len(self.set) == 0
}

func (self *HashSet) Add(items ...interface{}) {
	for _, item := range items {
		self.set[item] = empty{}
	}
}

func (self *HashSet) Remove(item interface{}) bool {
	ret := self.Contains(item)
	delete(self.set, item)
	return ret
}

func (self *HashSet) Size() int {
	return len(self.set)
}

func (self *HashSet) ToSlice() []interface{} {
	keys := make([]interface{}, len(self.set))

	i := 0
	for key := range self.set {
		keys[i] = key
		i++
	}
	return keys
}
