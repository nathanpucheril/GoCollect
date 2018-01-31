package streams

import (
	"fmt"
	"github.com/nathanpucheril/GoCollect/comparators"
	"github.com/nathanpucheril/GoCollect/iterators"
)

type Stream interface {
	Limit(int) Stream
	Map(func(interface{}) interface{}) Stream
	Filter(func(interface{}) bool) Stream
	Sorted(c comparators.Comparator) Stream
	Distinct() Stream

	Count() int

	Max(c comparators.Comparator) interface{}
	Min(c comparators.Comparator) interface{}

	ToSlice() []interface{}

	ForEach(func(interface{}))

	source
}

type Streamable interface {
	Stream() Stream
}

type source interface {
	next() (interface{}, bool)
}

type simplestream struct {
	out source
}

func (self *simplestream) Limit(lim int) Stream {
	var s Stream = &simplestream{&limitfunnel{src: self.out, lim: lim}}
	return s
}

func (self *simplestream) Map(fn func(interface{}) interface{}) Stream {
	var s Stream = &simplestream{&mapfunnel{self.out, fn}}
	return s
}

func (self *simplestream) Filter(fn func(interface{}) bool) Stream {
	var s Stream = &simplestream{&filterfunnel{self.out, fn}}
	return s
}

func (self *simplestream) Sorted(c comparators.Comparator) Stream {
	panic("implement me")
}

func (self *simplestream) Distinct() Stream {
	var s Stream = &simplestream{&distinctfunnel{self.out, make(map[interface{}]struct{})}}
	return s
}

func (self *simplestream) Max(c comparators.Comparator) interface{} {
	maxer := minmaxterminal{self.out, c, true}
	return maxer.terminate()
}

func (self *simplestream) Min(c comparators.Comparator) interface{} {
	minner := minmaxterminal{self.out, c, false}
	return minner.terminate()
}

func (self *simplestream) ToSlice() []interface{} {
	slice := make([]interface{}, 0, 10)
	for true {
		item, ok := self.out.next()
		if ok {
			slice = append(slice, item)
		} else {
			break
		}
	}
	return slice
}

func (self *simplestream) ForEach(fn func(interface{})) {
	for true {
		item, ok := self.out.next()
		if ok {
			fn(item)
		} else {
			break
		}
	}
}

func (self *simplestream) next() (interface{}, bool) {
	return self.out.next()
}

func (self *simplestream) Count() int {
	i := 0
	for true {
		_, ok := self.out.next()
		if ok {
			i++
		} else {
			break
		}
	}
	return i
}

// Sourcers

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

type splicesourcer struct {
	splice []interface{}
	idx    int
}

func (self *splicesourcer) next() (interface{}, bool) {
	if self.idx < len(self.splice) {
		ret := self.splice[self.idx]
		self.idx++
		return ret, true
	}
	return nil, false
}

// Funnels

type allpassfunnel struct {
	src source
}

func (self *allpassfunnel) next() (interface{}, bool) {
	fmt.Println("allpassfunnel next")
	return self.src.next()
}

type filterfunnel struct {
	src source
	fn  func(interface{}) bool
}

func (self *filterfunnel) next() (interface{}, bool) {
	for true {
		item, ok := self.src.next()
		fmt.Println("filter allpassfunnel next")
		if !ok {
			break
		}
		if self.fn(item) {
			return item, true
		}
	}
	return nil, false
}

type mapfunnel struct {
	src source
	fn  func(interface{}) interface{}
}

func (self *mapfunnel) next() (interface{}, bool) {
	for true {
		item, ok := self.src.next()
		fmt.Println("filter allpassfunnel next")
		if !ok {
			break
		}
		return self.fn(item), true
	}
	return nil, false
}

type limitfunnel struct {
	src source
	lim int
	cnt int
}

func (self *limitfunnel) next() (interface{}, bool) {
	for self.cnt < self.lim {
		item, ok := self.src.next()
		fmt.Println("limit allpassfunnel next")
		if !ok {
			break
		}
		self.cnt++
		return item, true
	}
	return nil, false
}

type distinctfunnel struct {
	src source
	set map[interface{}]struct{}
}

func (self *distinctfunnel) next() (interface{}, bool) {
	for true {
		item, ok := self.src.next()
		fmt.Println("funnel distinct next")
		if !ok {
			break
		}
		if _, inSet := self.set[item]; inSet {
			return item, true
		}
	}
	return nil, false
}

// Terminals

type minmaxterminal struct {
	src   source
	c     comparators.Comparator
	isMax bool
}

func (self *minmaxterminal) terminate() interface{} {
	var currMinMax interface{}

	if item, ok := self.src.next(); ok {
		currMinMax = item
	} else {
		return nil
	}

	for true {
		item, ok := self.src.next()
		if !ok {
			break
		}
		comparison := self.c(item, currMinMax)
		if (self.isMax && comparison > 0) || (!self.isMax && comparison < 0) {
			currMinMax = item
		}
	}
	return currMinMax
}
