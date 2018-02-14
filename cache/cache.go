package cache

import (
	"time"
)

const (
	LRU = iota
	MRU
	PRIORITY
)

type PolicyBasedStore interface {
	storelookup(key interface{}) (interface{}, bool)
	storeinsert(key, value interface{}) (interface{}, interface{}, bool)
}

type Cache interface {
	Lookup(key interface{}) (interface{}, bool, bool)
	Update(key, value interface{}) bool
	Contains(key interface{}) bool
	Keys() []interface{}
	Clear()
}

type simplecache struct {
}

type cache struct {
	stats      CacheStats
	localstore PolicyBasedStore

	capacity int

	writeback bool

	lookupfn func(key interface{}) (interface{}, bool)
	updatefn func(key, value interface{}) bool
}

func (c *cache) Lookup(key interface{}) (interface{}, bool, bool) {
	before := time.Now()

	var retLookup interface{} = nil
	retOk := false
	inCache := false

	if item, ok := c.localstore.storelookup(key); ok {
		retLookup, retOk = item, ok
		inCache = true
	} else if retrievedItem, ok := c.lookupfn(key); ok {
		keyEvicted, valueEvicted, didEvict := c.localstore.storeinsert(key, retrievedItem)
		if didEvict && c.writeback {
			c.updatefn(keyEvicted, valueEvicted)
		}
		retLookup, retOk = retrievedItem, ok
	} else {
		retLookup, retOk = nil, false
	}

	timeToLookup := time.Since(before).Nanoseconds()
	c.stats.recordHitTimeSample(timeToLookup)
	return retLookup, inCache, retOk //
}

func (c *cache) Update(key, value interface{}) bool {
	needsUpdate := false
	if item, ok := c.localstore.storelookup(key); ok {
		needsUpdate = item != value
	} else if retrievedItem, ok := c.lookupfn(key); ok {
		needsUpdate = retrievedItem != value
	} else {
		needsUpdate = true
	}

	if needsUpdate {
		keyEvicted, valueEvicted, didEvict := c.localstore.storeinsert(key, value)
		if didEvict && c.writeback && keyEvicted != key { //  TODO CHECK THIS
			c.updatefn(keyEvicted, valueEvicted)
		} else if keyEvicted != key && valueEvicted != value && didEvict {
			c.updatefn(keyEvicted, valueEvicted)
		}
	}

	return needsUpdate
}

func (c *cache) SetEvictionPolicy(mode int) *cache {
	switch mode {
	case LRU:
		lru := newLRUStore(c.capacity)
		c.localstore = &lru
	case MRU:
		mru := newMRUStore(c.capacity)
		c.localstore = &mru
	case PRIORITY:
		mru := newMRUStore(c.capacity)
		c.localstore = &mru
	default:
		lru := newLRUStore(c.capacity)
		c.localstore = &lru
	}
	return c
}

func (c *cache) SetOnEvictionFunction(func(key, value interface{})) *cache {
	panic("implement me")
	return c
}

func (c *cache) SetCacheMissFunction(func()) *cache {
	panic("implement me")
	return c
}

func (c *cache) SetCacheHitFunction(func()) *cache {
	panic("implement me")
	return c
}

func (c *cache) SetLookupFunction(lookupfn func(interface{}) (interface{}, bool)) *cache {
	c.lookupfn = lookupfn
	return c
}

func (c *cache) SetUpdateFunction(updatefn func(key, value interface{}) bool) *cache {
	c.updatefn = updatefn
	return c
}

func (c *cache) GetCacheStats(func()) *CacheStats {
	panic("implement me")
	return nil
}

func (c *cache) SetCacheCapacity(capacity int) *cache {
	c.capacity = capacity
	return c
}

func (c *cache) Clear(size int) *cache {
	panic("implement me")
	return c
}
