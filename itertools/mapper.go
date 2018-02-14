package itertools

import (
	"github.com/nathanpucheril/GoCollect/iterators"
	"github.com/pkg/errors"
)

type mapped struct {
	it    iterators.Iterator
	mapfn func(item interface{}) interface{}
}

func (self *mapped) Next() (interface{}, error) {
	item, err := self.it.Next()
	if err != nil {
		return nil, err
	}
	return self.mapfn(item), nil
}

func (self *mapped) HasNext() bool {
	return self.it.HasNext()
}

func (self *mapped) Iterator() iterators.Iterator {
	return self
}

type filtered struct {
	it         iterators.Iterator
	curr       *interface{}
	valueReady bool
	filterfn   func(item interface{}) bool
}

func (self *filtered) HasNext() bool {
	if self.valueReady {
		return true
	}
	for self.it.HasNext() {
		item, err := self.it.Next()
		if err != nil {
			return false
		} else if self.filterfn(item) {
			self.valueReady = true
			self.curr = &item
			return true
		}
	}
	return false
}

func (self *filtered) Next() (interface{}, error) {
	if self.valueReady {
		self.valueReady = false
		return *self.curr, nil
	}
	if self.HasNext() {
		return self.Next()
	}
	return nil, errors.New("StopIteration")
}

func (self *filtered) Iterator() iterators.Iterator {
	return self
}

func Mapped(it iterators.Iterable, mapfn func(item interface{}) interface{}) *mapped {
	return &mapped{it.Iterator(), mapfn}
}

func Filtered(it iterators.Iterable, filterfn func(item interface{}) bool) *filtered {
	f := &filtered{it: it.Iterator(), filterfn: filterfn}
	f.HasNext()
	return f
}

func Foreach(it iterators.Iterable, fn func(item interface{})) {
	for item := range iterators.ToChanIter(it.Iterator()) {
		fn(item)
	}
}

func Any(it iterators.Iterable, fn func(item interface{}) bool) bool {
	for item := range iterators.ToChanIter(it.Iterator()) {
		if fn(item) {
			return true
		}
	}
	return false
}

func All(it iterators.Iterable, fn func(item interface{}) bool) bool {
	for item := range iterators.ToChanIter(it.Iterator()) {
		if !fn(item) {
			return false
		}
	}
	return true
}
