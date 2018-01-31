package trie


type StringTrie struct {
	internalTrie Trie
}

func  NewStringTrie() StringTrie {
	return StringTrie{NewTrie()}
}

func (self *StringTrie) Put(str string) {
	self.internalTrie.Put(toInternalRepr(str))
}

func (self *StringTrie) Contains(str string) bool {
	return self.internalTrie.Contains(toInternalRepr(str))
}

func (self *StringTrie) Values() []string {
	internalValues := self.internalTrie.Values()
	strings := make([]string, 0, len(internalValues))
	for _, iV := range internalValues {
		strings = append(strings, toExternalRepr(iV))
	}
	return strings
}

func toInternalRepr(str string) []interface{} {
	internalRepr := make([]interface{}, 0 , len(str))
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
