package itertools

import (
	"fmt"
	"github.com/nathanpucheril/GoCollect/containers/lists/linkedlist/singlylinkedlist"
	"github.com/nathanpucheril/GoCollect/iterators"
	"testing"
)

func TestMapped(t *testing.T) {
	dll := singlylinkedlist.NewSLL()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Append(4)

	double := func(item interface{}) interface{} {
		return item.(int) * 2
	}
	greaterThan3 := func(item interface{}) bool {
		return item.(int) > 3
	}
	square := func(item interface{}) interface{} {
		val := item.(int)
		return val * val
	}

	squared := mapped(filtered(mapped(dll, double), greaterThan3), square)

	for item := range iterators.ToChanIter(squared) {
		fmt.Println(item)
	}

}
