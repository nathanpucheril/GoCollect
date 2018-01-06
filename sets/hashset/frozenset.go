package hashset

type ImmutableSet struct {
	set map[interface{}]struct{}
	HashSet
}

func of(items ... interface{}) ImmutableSet {
	set := make(map[interface{}]struct{})
	for _, item := range items {
		set[item] = empty{}
	}
	return ImmutableSet{set}
}

func (self *ImmutableSet) Clear() {
	panic("frozen set can not be mutated")
}

func (self *ImmutableSet) Add(items ...interface{}) {
	panic("frozen set can not be mutated")
}

func (self *ImmutableSet) Remove(item interface{}) bool {
	panic("frozen set can not be mutated")
}
