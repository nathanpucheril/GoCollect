package hashset

import (
	"errors"
	"github.com/nathanpucheril/GoCollect/iterators"
)

type HashSetIterator struct {
	set        *HashSet
	hasNext    bool
	nextGetter <-chan interface{}
}

func NewHashSetIterator(set *HashSet) iterators.Iterator {
	var it iterators.Iterator = &HashSetIterator{set: set, hasNext: !set.IsEmpty(), nextGetter: newNextGetter(set)}
	return it
}

func (self *HashSetIterator) Next() (interface{}, error) {

	select {
	case nextItem, ok := <-self.nextGetter:
		if ok {
			return nextItem, nil
		} else {
			return nil, errors.New("StopIteration")
		}
	default:
		self.hasNext = false
		return nil, errors.New("StopIteration")
	}
}

func (self *HashSetIterator) HasNext() bool {
	return true // TODO
}

func newNextGetter(set *HashSet) <-chan interface{} {
	chn := make(chan interface{})
	go func() {
		for item, _ := range set.set {
			chn <- item
		}

		// Ensure that at the end of the loop we close the channel!
		close(chn)
	}()
	return chn
}
