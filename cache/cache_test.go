package cache

import (
	"fmt"
	"testing"
)

func TestCache_SetEvictionPolicy(t *testing.T) {
	c := cache{}
	c.SetCacheCapacity(5)
	c.SetEvictionPolicy(0)
	c.SetUpdateFunction(func(k, v interface{}) bool {
		return true
	})
	c.SetLookupFunction(func(k interface{}) (interface{}, bool) {
		fmt.Println("Sending update: ", k.(int), "->", k.(int))
		return (k.(int) - 72) % 10, true
	})

	fmt.Println(c.updatefn)
	fmt.Println(c.Lookup(1))
	fmt.Println(c.Lookup(1))
	fmt.Println(c.Lookup(2))
	fmt.Println(c.Lookup(3))
	fmt.Println(c.Lookup(4))
	fmt.Println(c.Lookup(5))
	fmt.Println(c.Lookup(5))
	fmt.Println(c.Lookup(6))
	fmt.Println(c.Lookup(1))
	c.Update(1, 1)

}

func TestCache_SetEvictionPolicyMRU(t *testing.T) {
	c := cache{}
	c.SetCacheCapacity(5)
	c.SetEvictionPolicy(1)
	c.SetUpdateFunction(func(k, v interface{}) bool {
		return true
	})
	c.SetLookupFunction(func(k interface{}) (interface{}, bool) {
		return (k.(int) - 72) % 10, true
	})

	// Misses
	fmt.Println(c.Lookup(1))
	fmt.Println(c.Lookup(2))
	fmt.Println(c.Lookup(3))
	fmt.Println(c.Lookup(4))
	fmt.Println(c.Lookup(5))

	// Hits
	fmt.Println(c.Lookup(1))
	fmt.Println(c.Lookup(2))
	fmt.Println(c.Lookup(3))
	fmt.Println(c.Lookup(4))
	fmt.Println(c.Lookup(5))

	// miss and evict 5
	fmt.Println(c.Lookup(6))

	// hits
	fmt.Println(c.Lookup(1))
	fmt.Println(c.Lookup(2))
	fmt.Println(c.Lookup(3))
	fmt.Println(c.Lookup(4))
	fmt.Println(c.Lookup(6))

	// miss
	fmt.Println(c.Lookup(5))
	// hit and evict 6
	fmt.Println(c.Lookup(5))

}
