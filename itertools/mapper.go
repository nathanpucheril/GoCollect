package itertools

import (
	"github.com/nathanpucheril/GoCollect/iterators"
	"github.com/pkg/errors"
)

type Mapped struct {
	it    iterators.Iterator
	mapfn func(item interface{}) interface{}
}

func (self *Mapped) Next() (interface{}, error) {
	item, err := self.it.Next()
	if err != nil {
		return nil, err
	}
	return self.mapfn(item), nil
}

func (self *Mapped) HasNext() bool {
	return self.it.HasNext()
}

func (self *Mapped) Iterator() iterators.Iterator {
	return self
}

type Filtered struct {
	it         iterators.Iterator
	curr       *interface{}
	valueReady bool
	filterfn   func(item interface{}) bool
}

func (self *Filtered) HasNext() bool {
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

func (self *Filtered) Next() (interface{}, error) {
	if self.valueReady {
		self.valueReady = false
		return *self.curr, nil
	}
	if self.HasNext() {
		return self.Next()
	}
	return nil, errors.New("StopIteration")
}

func (self *Filtered) Iterator() iterators.Iterator {
	return self
}

func mapped(it iterators.Iterable, mapfn func(item interface{}) interface{}) *Mapped {
	return &Mapped{it.Iterator(), mapfn}
}

func filtered(it iterators.Iterable, filterfn func(item interface{}) bool) *Filtered {
	f := &Filtered{it: it.Iterator(), filterfn: filterfn}
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
