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
	sstream := simplestream{ss}
	//pos := func (i interface{}) bool {
	//	return i.(int) > 0
	//}
	fmt.Println(sstream.Sorted(comparators.IntComparator).next())
	//if sstream.Max() != 100 {
	//	t.Fail()
	//}

}

func TestMapped_SpliceSourcer(t *testing.T) {

	interfaceArr := make([]interface{}, 0)
	intarr := []int{1235, 1523, 15, 52, 125, 13}
	for _, s := range intarr {
		fmt.Println(s)
		interfaceArr = append(interfaceArr, s)
	}
	splicesource := splicesourcer{splice: interfaceArr}
	var ss source = &splicesource
	sstream := simplestream{ss}
	//pos := func (i interface{}) bool {
	//	return i.(int) > 0
	//}
	fmt.Println(sstream.Limit(4).ToSlice())
	//if sstream.Max() != 100 {
	//	t.Fail()
	//}

}
