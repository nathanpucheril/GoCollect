package containers

import "reflect"

type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
	if reflect.TypeOf(a).Kind() ==reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
		return b.(int) - a.(int)
	}
	panic("IntComparator can not compare interfaces that are not Ints")
}