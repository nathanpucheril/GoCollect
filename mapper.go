package GoCollect

//
//import (
//	"github.com/nathanpucheril/GoCollect/containers"
//	"github.com/nathanpucheril/GoCollect/lists"
//	"github.com/nathanpucheril/GoCollect/lists/arraylist"
//	"github.com/nathanpucheril/GoCollect/sets"
//	"github.com/nathanpucheril/GoCollect/sets/hashset"
//)
//
//func mapToList(container containers.Container, mapfn func(item interface{}) interface{}) *lists.List {
//	var mappedContainer lists.List = arraylist.New()
//	for _, item := range container.ToSlice() {
//		mappedContainer.Append(mapfn(item))
//	}
//	return mappedContainer
//}
//
//func mapToSet(container containers.Container, mapfn func(item interface{}) interface{}) *sets.Set {
//	mappedContainer := hashset.New()
//	for _, item := range container.ToSlice() {
//		mappedContainer.Add(mapfn(item))
//	}
//	return mappedContainer
//}
//
//func Filter(container containers.Container, filterfn func(item interface{}) bool) *lists.List {
//	var mappedContainer lists.List = arraylist.New()
//	for _, item := range container.ToSlice() {
//		if filterfn(item) {
//			mappedContainer.Append(item)
//		}
//	}
//	return mappedContainer
//}
//
//func Foreach(container containers.Container, fn func(item interface{})) {
//	for _, item := range container.ToSlice() {
//		fn(item)
//	}
//}
//
//func Any(container containers.Container, fn func(item interface{}) bool) bool {
//	for _, item := range container.ToSlice() {
//		if fn(item) {
//			return true
//		}
//	}
//	return false
//}
//
//func All(container containers.Container, fn func(item interface{}) bool) bool {
//	for _, item := range container.ToSlice() {
//		if !fn(item) {
//			return false
//		}
//	}
//	return true
//}
//
//
