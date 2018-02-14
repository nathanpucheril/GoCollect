package comparators

type Comparator func(a, b interface{}) int

type Comparable interface {
	Comparator() Comparator
}

func IntComparator(a, b interface{}) int {
	return a.(int) - b.(int)
}

func Reverse(comparator Comparator) func(a, b interface{}) int {
	return func(a, b interface{}) int {
		return -comparator(a, b)
	}
}
