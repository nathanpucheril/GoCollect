package multiset

import "github.com/nathanpucheril/GoCollect/iterators"

type MultiSet struct {
	set map[interface{}]int
	size int
}


func New() MultiSet {
	return MultiSet{set:make(map[interface{}]int)}
}

func (self *MultiSet) Clear() {
	self.set = make(map[interface{}]int)
}

func (self *MultiSet) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := self.set[item]; !ok {
			return false
		}
	}
	return true
}

func (self *MultiSet) IsEmpty() bool {
	return len(self.set) == 0
}

func (self *MultiSet) Add(items ...interface{}) bool {
	for _, item := range items {
		self.set[item]++
		self.size++
	}
	return true // set will always change because duplicates are allowed
}

func (self *MultiSet) Remove(items ...interface{}) bool {
	changed := false
	for _, item := range items {
		_, ok := self.set[item]
		if ok { // if the set contained the item, we successfully will remove it
			changed = true
			self.set[item]--
			self.size--
			if self.set[item] == 0 {
				delete(self.set, item)
			}
		}
	}
	return changed
}

func (self *MultiSet) Size() int {
	return self.size
}

func (self *MultiSet) ToSlice() []interface{} {
	keys := make([]interface{}, self.size)

	i := 0
	for key, count := range self.set {
		for j := 0; j <  count; j ++ {
			keys[i] = key
			i++
		}
	}
	return keys
}

func (self *MultiSet) Iterator() iterators.Iterator {
	panic("implement me")
}
