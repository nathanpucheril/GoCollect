package GoCollect

import (
	"github.com/nathanpucheril/GoCollect/containers"
	"github.com/nathanpucheril/GoCollect/lists"
	"github.com/nathanpucheril/GoCollect/lists/arraylist"
	"github.com/nathanpucheril/GoCollect/sets"
	"github.com/nathanpucheril/GoCollect/sets/hashset"
)

func mapToList(container containers.Container, mapfn func(item interface{}) interface{}) lists.List {
	mappedContainer := arraylist.New()
	for _, item := range container.ToSlice() {
		mappedContainer.Append(mapfn(item))
	}
	return mappedContainer
}

func mapToSet(container containers.Container, mapfn func(item interface{}) interface{}) sets.Set {
	mappedContainer := hashset.New()
	for _, item := range container.ToSlice() {
		mappedContainer.Add(mapfn(item))
	}
	return mappedContainer
}
