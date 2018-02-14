package multiset

import (
	"errors"
	"github.com/nathanpucheril/GoCollect/iterators"
)

type MultiSetIterator struct {
	set          *MultiSet
	next         interface{}
	nextPrepared bool
	nextGetter   <-chan interface{}
}

func newMultiSetIterator(set *MultiSet) iterators.Iterator {
	multisetIterator := &MultiSetIterator{set: set, nextGetter: newNextGetter(set)}
	multisetIterator.prepareNext()
	return multisetIterator
}

func (self *MultiSetIterator) Next() (interface{}, error) {
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

func (self *MultiSetIterator) HasNext() bool {
	return self.prepareNext()
}

func (self *MultiSetIterator) prepareNext() bool {
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

func newNextGetter(set *MultiSet) <-chan interface{} {
	chn := make(chan interface{})
	go func() {
		for item, count := range set.set {
			for i := 0; i < count; i++ {
				chn <- item
			}
		}

		// Ensure that at the end of the loop we close the channel!
		close(chn)
	}()
	return chn
}
