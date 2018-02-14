package trie

import (
	"fmt"
	"testing"
)

func TestTrie_Put(t *testing.T) {
	trie := NewTrie()

	arr := make([]interface{}, 0, 10)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	trie.Put(arr)

	arr2 := make([]interface{}, 0, 10)
	arr2 = append(arr2, 5)
	arr2 = append(arr2, 4)
	arr2 = append(arr2, 3)

	arr3 := make([]interface{}, 0, 10)
	arr3 = append(arr3, 1)
	arr3 = append(arr3, 2)
	if !trie.Contains(arr) {
		t.Fail()
	}

	if trie.Contains(arr2) {
		t.Fail()
	}

	if trie.Contains(arr3) {
		t.Fail()
	}

	trie.Put(arr3)
	trie.Put(arr3)
	trie.Put(arr3)
	trie.Put(arr3)
	if !trie.Contains(arr3) {
		t.Fail()
	}

	trie.Put(arr2)
	if !trie.Contains(arr2) {
		t.Fail()
	}

	for _, item := range trie.Values() {
		fmt.Println(item)
	}

	fmt.Println(trie.size)
}
