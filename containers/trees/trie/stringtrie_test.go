package trie

import (
	"fmt"
	"testing"
)

func TestNewStringTrie(t *testing.T) {
	trie := NewStringTrie()

	trie.Put("hello")

	if !trie.Contains("hello") {
		t.Fail()
	}

	if trie.Contains("olleh") {
		t.Fail()
	}

	if trie.Contains("hel") {
		t.Fail()
	}

	trie.Put("hel")
	if !trie.Contains("hel") {
		t.Fail()
	}

	trie.Put("hel")
	trie.Put("h23el")
	trie.Put("he513l")
	trie.Put("153hel")

	fmt.Println(trie.Values())
}
