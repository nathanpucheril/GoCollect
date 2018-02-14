package trie

import (
	"github.com/nathanpucheril/GoCollect/containers/sets/hashset"
	"github.com/nathanpucheril/GoCollect/iterators"
)

type Trie struct {
	root node
	size int
}

func (self *Trie) IsEmpty() bool {
	return self.size == 0
}

func (self *Trie) Size() int {
	return self.size
}

func (self *Trie) ToSlice() []interface{} {
	panic("implement me")
}

func (self *Trie) Clear() {
	self.root = node{value: nil, children: make(map[interface{}]*node)}
}

func (self *Trie) Iterator() iterators.Iterator {
	panic("implement me")
}

type node struct {
	value    interface{}
	children map[interface{}]*node
	isEnd    bool
}

func NewTrie() Trie {
	return Trie{root: node{value: nil, children: make(map[interface{}]*node)}}
}

func (self *Trie) Put(item []interface{}) {
	if self.root.insert(item) {
		self.size++
	}
}

func (self *Trie) Contains(item []interface{}) bool {
	return self.root.contains(item)
}

func (self *Trie) Values() [][]interface{} {
	values := make([][]interface{}, 0)

	stack, prefix := make([]*node, 0), make([]interface{}, 0)
	for _, v := range self.root.children {
		stack = append(stack, v)
	}

	visited := hashset.New()
	for len(stack) > 0 {
		if peek := stack[len(stack)-1]; visited.Contains(peek) {
			prefix = prefix[:len(prefix)-1]
			stack = stack[:len(stack)-1]
		} else {
			// visit
			visited.Add(peek)

			prefix = append(prefix, peek.value)
			if peek.isEnd {
				value := make([]interface{}, len(prefix))
				copy(value, prefix)
				values = append(values, value)
			}

			// add children to stack
			for _, v := range peek.children {
				stack = append(stack, v)
			}

		}
	}
	return values
}

func (self *node) insert(item []interface{}) bool {
	if len(item) == 0 {
		wasEnd := self.isEnd
		self.isEnd = true
		return !wasEnd
	} else {
		toInsert := item[0]
		childNode, ok := self.children[toInsert]
		if !ok {
			self.children[toInsert] = &node{value: toInsert, children: make(map[interface{}]*node)}
			childNode = self.children[toInsert]
		}
		return childNode.insert(item[1:])
	}
}

func (self *node) contains(suffix []interface{}) bool {
	if len(suffix) == 0 {
		return self.isEnd == true
	} else if childNode, ok := self.children[suffix[0]]; ok {
		return childNode.contains(suffix[1:])
	}
	return false
}
