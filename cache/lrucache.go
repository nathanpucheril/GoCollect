package cache

import (
	"fmt"
	"github.com/nathanpucheril/GoCollect/containers/lists/linkedlist/doublylinkedlist"
)

type lrustore struct {
	store map[interface{}]*storeentry
	queue *doublylinkedlist.DoublyLinkedList

	capacity int
}

func newLRUStore(capacity int) lrustore {
	dll := doublylinkedlist.NewDLL()
	return lrustore{make(map[interface{}]*storeentry), dll, capacity}
}

func (lru *lrustore) storelookup(key interface{}) (interface{}, bool) {
	item, ok := lru.store[key]
	if ok {
		return item.value, ok
	}
	return nil, ok
}

func (lru *lrustore) evict() (interface{}, interface{}, bool) {
	if item, ok := lru.queue.RemoveBack(); ok {
		citem := item.(*storeentry)
		delete(lru.store, citem.key)
		return citem.key, citem.value, ok
	} else {
		return nil, nil, ok
	}
}

func (lru *lrustore) storeinsert(key, value interface{}) (interface{}, interface{}, bool) {
	fmt.Println(lru.capacity, len(lru.store))
	if len(lru.store) >= lru.capacity {
		return lru.evict()
	}
	citem := &storeentry{key, value}
	lru.queue.Prepend(citem)
	lru.store[key] = citem
	return nil, nil, false
}

type storeentry struct {
	key, value interface{}
}
