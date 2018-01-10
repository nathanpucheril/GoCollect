package lists

import (
	"github.com/nathanpucheril/GoCollect/containers"
)

type List interface {
	Prepend(value interface{})
	Append(value interface{})
	Insert(index int, value interface{})

	Remove(index int) (interface{}, bool)
	Set(index int, value interface{})
	Get(index int) (interface{}, bool)
	Contains(items ...interface{}) bool

	IndexOf(value interface{}) int
	LastIndexOf(value interface{}) int

	Sublist(to, from int) List

	containers.Container
}
//
//type list struct {
//	List
//}
//
//func (self *list) Prepend(value interface{}) {
//	self.Insert(0, value)
//}
//
//func (self *list) Append(value interface{}) {
//	self.Insert(self.Size(), value)
//}


//func validIndex(list List, index int) bool {
//	return 0 <= index && index < list.Size()
//}
