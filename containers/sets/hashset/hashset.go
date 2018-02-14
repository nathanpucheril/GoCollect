package hashset

import (
	"github.com/nathanpucheril/GoCollect/iterators"
)

type HashSet struct {
	set map[interface{}]struct{}
}

type empty struct{}

func New() HashSet {
	return HashSet{set: make(map[interface{}]struct{})}
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

func (self *HashSet) Add(items ...interface{}) bool {
	changed := false
	for _, item := range items {
		if _, ok := self.set[item]; !ok { // if the set did not contain the item, we successfully will add it
			changed = true
		}
		self.set[item] = empty{}
	}
	return changed
}

func (self *HashSet) Remove(items ...interface{}) bool {
	changed := false
	for _, item := range items {
		if _, ok := self.set[item]; ok { // if the set contained the item, we successfully will remove it
			changed = true
		}
		delete(self.set, item)
	}
	return changed
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

func (self *HashSet) Iterator() iterators.Iterator {
	return newHashSetIterator(self)
}
