package trie

import "github.com/nathanpucheril/GoCollect/iterators"

type StringTrie struct {
	internalTrie Trie
}

func (self *StringTrie) IsEmpty() bool {
	return self.internalTrie.IsEmpty()
}

func (self *StringTrie) Size() int {
	return self.internalTrie.Size()
}

func (self *StringTrie) ToSlice() []interface{} {
	panic("implement me")
}

func (self *StringTrie) Clear() {
	self.internalTrie.Clear()
}

func (self *StringTrie) Iterator() iterators.Iterator {
	panic("implement me")
}

func NewStringTrie() StringTrie {
	return StringTrie{NewTrie()}
}

func (self *StringTrie) Put(str string) {
	self.internalTrie.Put(toInternalRepr(str))
}

func (self *StringTrie) Contains(str string) bool {
	return self.internalTrie.Contains(toInternalRepr(str))
}

func (self *StringTrie) Values() []string {
	strings := make([]string, 0)
	for _, iV := range self.internalTrie.Values() {
		strings = append(strings, toExternalRepr(iV))
	}
	return strings
}

func toInternalRepr(str string) []interface{} {
	internalRepr := make([]interface{}, 0, len(str))
	for _, char := range str {
		internalRepr = append(internalRepr, char)
	}
	return internalRepr
}

func toExternalRepr(internalRepr []interface{}) string {
	str := ""
	for _, char := range internalRepr {
		str += string(char.(int32))
	}
	return str
}
