package streams

import (
	"github.com/nathanpucheril/GoCollect/comparators"
	"github.com/nathanpucheril/GoCollect/iterators"
	"math"
)

type Stream interface {
	Limit() Stream
	Map() Stream
	Filter() Stream
	Sorted() Stream
	Distinct() Stream

	Count() int

	Max() interface{}
	Min() interface{}

	ToSlice() []interface{}

	ForEach(func(interface{}))

	streamsource
}

type Streamable interface {
	Stream() Stream
}

type streamsource interface {
	next() (interface{}, bool)
}

type simplestream struct {
	src streamsource
}

func (self *simplestream) Limit() Stream {
	panic("implement me")
}

func (self *simplestream) Map() Stream {
	panic("implement me")
}

func (self *simplestream) Filter() Stream {
	panic("implement me")
}

func (self *simplestream) Sorted() Stream {
	panic("implement me")
}

func (self *simplestream) Distinct() Stream {
	panic("implement me")
}

func (self *simplestream) Max() interface{} {
	mmst := minmaxstreamterminal{self, comparators.IntComparator, 0, math.MinInt32}
	return mmst.terminate()
}

func (self *simplestream) Min() interface{} {
	mmst := minmaxstreamterminal{self, comparators.IntComparator, 1, math.MaxInt32}
	return mmst.terminate()
}

func (self *simplestream) ToSlice() []interface{} {
	panic("implement me")
}

func (self *simplestream) ForEach(fn func(interface{})) {
	fet := foreachstreamterminal{self.src, fn}
	fet.terminate()
}

func (self *simplestream) next() (interface{}, bool) {
	return self.src.next()
}

func (self *simplestream) Count() int {
	cst := countstreamterminal{src: self.src}
	return cst.terminate()
}

type iteratorsourcer struct {
	it iterators.Iterator
}

func (self *iteratorsourcer) next() (interface{}, bool) {
	if self.it.HasNext() {
		item, err := self.it.Next()
		if err != nil {
			return nil, false
		}
		return item, true
	}
	return nil, false
}

// Terminals

type foreachstreamterminal struct {
	src streamsource
	fn  func(interface{})
}

func (self *foreachstreamterminal) terminate() {
	for true {
		item, ok := self.src.next()
		if ok {
			self.fn(item)
		} else {
			break
		}
	}
}

type countstreamterminal struct {
	src streamsource
	cnt int
}

func (self *countstreamterminal) terminate() int {
	for true {
		if _, ok := self.src.next(); ok {
			self.cnt++
		} else {
			break
		}
	}
	return self.cnt
}

type minmaxstreamterminal struct {
	src  streamsource
	cmp  comparators.Comparator
	mode int // max : 1; min : 0
	val  interface{}
}

func (self *minmaxstreamterminal) terminate() interface{} {
	for true {
		if item, ok := self.src.next(); ok {
			if self.cmp(self.val, item) > 0 && self.mode == 1 {
				self.val = item
			} else if self.cmp(self.val, item) < 0 && self.mode == 0 {
				self.val = item
			}
		} else {
			break
		}
	}
	return self.val
}
