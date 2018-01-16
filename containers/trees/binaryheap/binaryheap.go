package binaryheap

import (
	"fmt"
	"github.com/nathanpucheril/GoCollect"
)

type BinaryHeap struct {
	heap       []interface{}
	comparator GoCollect.Comparator
}

func NewIntBinaryHeap() BinaryHeap {
	return BinaryHeap{make([]interface{}, 0), GoCollect.IntComparator}
}

func (self *BinaryHeap) Push(value interface{}) {
	self.heap = append(self.heap, value)
	self.bubbleUp()
}

func (self *BinaryHeap) Pop() interface{} {
	popped := self.heap[0]

	if len(self.heap) > 1 {
		self.swap(0, len(self.heap)-1)
	}
	self.heap = self.heap[:len(self.heap)-1]
	self.sinkDown()
	return popped
}
func (self *BinaryHeap) swap(i, i2 int) {
	tmp := self.heap[i]
	self.heap[i] = self.heap[i2]
	self.heap[i2] = tmp
}
func (self *BinaryHeap) sinkDown() {
	size := len(self.heap)
	for currIndex := 0; currIndex < size/2-1; {
		currValue := self.heap[currIndex]

		swapIndex := 2*currIndex + 1
		swapValue := self.heap[swapIndex]
		if swapIndex < size && self.comparator(currValue, swapValue) > 0 {
			swapIndex++
			swapValue = self.heap[swapIndex]
		}
		if self.comparator(currValue, swapValue) >= 0 {
			return
		}

		self.swap(currIndex, swapIndex)
		currIndex = swapIndex
	}
}

func (self *BinaryHeap) bubbleUp() {
	for i := len(self.heap) - 1; i > 0 && self.comparator(self.heap[i], self.heap[i/2]) < 0; i = i / 2 {
		self.swap(i, i/2)
	}
}

// For Testing Purposes
func (self *BinaryHeap) isHeap(i, n int) bool {
	// If a leaf node
	if i > (n-2)/2 {

		return true
	}

	// If an internal node and is greater than its children, and
	// same is recursively true for the children
	if self.comparator(self.heap[i], self.heap[2*i+1]) >= 0 && self.comparator(self.heap[i], self.heap[2*(i+1)]) >= 0 {
		return self.isHeap(2*i+1, n) && self.isHeap(2*i+2, n)
	}
	fmt.Print("here")
	return false
}
