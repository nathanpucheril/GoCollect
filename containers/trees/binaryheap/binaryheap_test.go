package binaryheap

import (
	//"fmt"
	"fmt"
	"testing"
)

func TestBinaryHeap_Push(t *testing.T) {
	heap := NewIntBinaryHeap()
	for i := 0; i < 400; i++ {
		heap.Push(i)
	}
	fmt.Println(heap.heap)
	fmt.Println(heap.isHeap(0, len(heap.heap)))
	heap.Push(0)
	for i := 0; i < 16; i++ {
		fmt.Println(heap.Pop())
	}
}
