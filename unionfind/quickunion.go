package unionfind

import "github.com/pkg/errors"

type QuickUnion struct {
	elems    map[interface{}]int
	set      []int
	setSizes []int
}

func (self *QuickUnion) Add(item interface{}) (bool, error) {

	if _, ok := self.elems[item]; ok {
		return false, errors.New("item already exists")
	}
	currSize := len(self.elems)
	self.elems[item] = currSize
	self.set = append(self.set, currSize)
	return true, nil
}

func NewQuickUnion(items ...interface{}) *QuickUnion {
	elems := make(map[interface{}]int)
	set := make([]int, len(items), len(items))
	setSizes := make([]int, len(items), len(items))

	for idx, item := range items {
		elems[item] = idx
		set[idx] = idx // initially set all indexs to be parents of themselves
		setSizes[idx] = 1
	}
	return &QuickUnion{elems, set, setSizes}
}

func (self *QuickUnion) Union(a, b interface{}) (bool, error) {
	indexA, ok := self.elems[a]
	if !ok {
		return false, errors.New("item did not exist")
	}
	indexB, ok := self.elems[b]
	if !ok {
		return false, errors.New("item did not exist")
	}

	immParentA, immParentB := self.set[indexA], self.set[indexB]
	if immParentA == immParentB {
		return false, nil //  Already connected, no modifications
	}

	if setSizeA, setSizeB := self.setSizes[immParentA], self.setSizes[immParentB]; setSizeA < setSizeB {
		self.set[immParentB] = immParentA
		self.setSizes[immParentA] += self.setSizes[immParentB]
	} else {
		self.set[immParentA] = immParentB
		self.setSizes[immParentB] += self.setSizes[immParentA]
	}
	return true, nil
}

func (self *QuickUnion) IsConnected(a, b interface{}) (bool, error) {
	indexA, ok := self.elems[a]
	if !ok {
		return false, errors.New("item did not exist")
	}
	indexB, ok := self.elems[b]
	if !ok {
		return false, errors.New("item did not exist")
	}

	return self.parent(indexA) == self.parent(indexB), nil
}

// implements path compression
func (self *QuickUnion) parent(index int) int {
	if parent := self.set[index]; index == parent {
		return index
	} else {
		self.set[index] = self.parent(parent)
		return self.set[index]
	}
}
