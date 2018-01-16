package sets

import "github.com/nathanpucheril/GoCollect/containers"

type Set interface {
	Add(value ...interface{}) bool
	Remove(value ...interface{}) bool
	Contains(value ...interface{}) bool

	containers.Container
}

func union(s1, s2 Set) Set {
	return nil
}

func intersection(s1, s2 Set) Set {
	return nil
}

func complement(universal, s Set) Set {
	return nil
}
