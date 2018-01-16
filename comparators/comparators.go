package comparators

type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
	return a.(int) - b.(int)
}
