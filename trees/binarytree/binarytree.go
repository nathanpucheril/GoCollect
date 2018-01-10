package binarytree

import "github.com/nathanpucheril/GoCollect/containers"

type BinarySearchTree struct {
	root *node
	comparator containers.Comparator
}



func NewBinarySearchTree() BinarySearchTree{
	return BinarySearchTree{}
}

func (self *BinarySearchTree) Insert(value interface{}) {
	if self.root == nil {
		self.root = &node{value:value}
	}

}



type node struct {
	value interface{}
	lchild *node
	rchild *node
}

func (self *node) IsLeaf() bool {
	return self.lchild == nil && self.rchild == nil
}