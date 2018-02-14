package cache

type mrustore struct {
	store  map[interface{}]interface{}
	mrukey interface{}

	capacity int
}

func newMRUStore(capacity int) mrustore {
	return mrustore{make(map[interface{}]interface{}), nil, capacity}
}

func (mru *mrustore) storelookup(key interface{}) (interface{}, bool) {
	value, ok := mru.store[key]
	return value, ok
}

func (mru *mrustore) evict() (interface{}, interface{}, bool) {
	value, ok := mru.store[mru.mrukey]
	if ok {
		toDelete := mru.mrukey
		delete(mru.store, toDelete)
		mru.mrukey = nil
		return toDelete, value, ok
	}
	return nil, nil, false
}

func (mru *mrustore) storeinsert(key, value interface{}) (interface{}, interface{}, bool) {
	var retKey, retValue interface{} = nil, nil
	retOk := false
	if len(mru.store) >= mru.capacity {
		retKey, retValue, retOk = mru.evict()
	}
	mru.mrukey = key
	mru.store[key] = value
	return retKey, retValue, retOk
}
