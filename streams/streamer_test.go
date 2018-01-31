package streams

import (
	"fmt"
	"github.com/nathanpucheril/GoCollect/comparators"
	"github.com/nathanpucheril/GoCollect/containers/lists/linkedlist/singlylinkedlist"
	"testing"
)

func TestMapped(t *testing.T) {
	dll := singlylinkedlist.NewSLL()
	dll.Append(1)
	dll.Append(2)
	dll.Append(2)
	dll.Append(-3)
	dll.Append(4)
	dll.Append(-3)
	dll.Append(2)
	dll.Append(3)
	dll.Append(99)
	dll.Append(3)
	dll.Append(4)
	dll.Append(-4)
	dll.Append(20)
	dll.Append(3)
	dll.Append(4)
	dll.Append(-14)

	its := iteratorsourcer{dll.Iterator()}
	var ss source = &its
	sstream := simplestream{ &allpassfunnel{ss}}
	//pos := func (i interface{}) bool {
	//	return i.(int) > 0
	//}
	fmt.Println(sstream.Max(comparators.IntComparator))
	//if sstream.Max() != 100 {
	//	t.Fail()
	//}

}
