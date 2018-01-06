package sets

type Set interface {
	Add() bool
	Remove() bool
	Contains() bool
	// Collection
	//IsEmpty() bool
	//Size() int
	//ToSlice() []interface{}
	//Clear()
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
