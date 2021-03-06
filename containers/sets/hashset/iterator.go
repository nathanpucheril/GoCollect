package hashset

import (
	"errors"
	"github.com/nathanpucheril/GoCollect/iterators"
)

type HashSetIterator struct {
	set          *HashSet
	next         interface{}
	nextPrepared bool
	nextGetter   <-chan interface{}
}

func newHashSetIterator(set *HashSet) iterators.Iterator {
	hashsetIterator := &HashSetIterator{set: set, nextGetter: newNextGetter(set)}
	hashsetIterator.prepareNext()
	return hashsetIterator
}

func (self *HashSetIterator) Next() (interface{}, error) {
	if self.nextPrepared {
		ret := self.next
		self.prepareNext()
		return ret, nil
	} else if self.prepareNext() {
		return self.Next()
	} else {
		return nil, errors.New("StopIteration: no more elements")
	}
}

func (self *HashSetIterator) HasNext() bool {
	return self.prepareNext()
}

func (self *HashSetIterator) prepareNext() bool {
	if self.nextPrepared {
		return self.nextPrepared
	}

	self.next = nil
	self.nextPrepared = false
	select {
	case nextItem, ok := <-self.nextGetter:
		if ok {
			self.next = nextItem
			self.nextPrepared = true
		}
	}
	return self.nextPrepared
}

func newNextGetter(set *HashSet) <-chan interface{} {
	chn := make(chan interface{})
	go func() {
		for item := range set.set {
			chn <- item
		}

		// Ensure that at the end of the loop we close the channel!
		close(chn)
	}()
	return chn
}
