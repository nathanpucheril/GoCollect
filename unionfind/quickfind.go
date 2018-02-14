package unionfind

import "errors"

type QuickFind struct {
	elems map[interface{}]int
	set   []int
}

func (self *QuickFind) Add(item interface{}) (bool, error) {

	if _, ok := self.elems[item]; ok {
		return false, errors.New("item already exists")
	}
	currSize := len(self.elems)
	self.elems[item] = currSize
	self.set = append(self.set, currSize)
	return true, nil
}

func NewQuickFind(items ...interface{}) *QuickFind {
	elems := make(map[interface{}]int)
	set := make([]int, len(items), len(items))

	for idx, item := range items {
		elems[item] = idx
		set[idx] = idx // initially set all indexs to be parents of themselves
	}
	return &QuickFind{elems, set}
}

func (self *QuickFind) Union(a, b interface{}) (bool, error) {
	indexA, ok := self.elems[a]
	if !ok {
		return false, errors.New("item did not exist")
	}
	indexB, ok := self.elems[b]
	if !ok {
		return false, errors.New("item did not exist")
	}

	if indexA == indexB {
		return false, nil //  Already connected, no modifications
	}
	for idx, setNum := range self.set {
		if setNum == indexB {
			self.set[idx] = indexA // set anything that is part of set B to be part of set A
		}
	}
	return true, nil
}

func (self *QuickFind) IsConnected(a, b interface{}) (bool, error) {
	indexA, ok := self.elems[a]
	if !ok {
		return false, errors.New("item did not exist")
	}
	indexB, ok := self.elems[b]
	if !ok {
		return false, errors.New("item did not exist")
	}

	return indexA == indexB, nil
}
