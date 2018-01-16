package bitset

import (
	"testing"
)

func TestBitset_Set(t *testing.T) {
	var numBits uint32 = 57
	var i uint32 = 0
	for ; i < numBits; i++ {
		set := New(numBits)
		set.Set(i)
		if !set.IsSet(i) {
			t.Fail()
		}
	}
}

func TestBitset_Unset(t *testing.T) {
	var numBits uint32 = 57
	var i uint32 = 0
	for ; i < numBits; i++ {
		set := New(numBits)
		set.Set(i)
		if !set.IsSet(i) {
			t.Fail()
		}
		set.Unset(i)
		if set.IsSet(i) {
			t.Fail()
		}
	}
}

func TestBitset_Set2(t *testing.T) {
	var numBits uint32 = 183
	var i uint32 = 0
	for ; i < numBits; i++ {
		set := New(numBits)
		set.Set(i)
		var j uint32 = 0
		for ; j < numBits; j++ {
			jSet := set.IsSet(j)
			if (j == i && !jSet) || (j != i && jSet) {
				t.Fail()
			}
		}
	}
}

func TestBitset_Clear(t *testing.T) {
	var numBits uint32 = 1794
	var i uint32
	set := New(numBits)
	for i = 0; i < numBits; i = i + 10 {
		set.Set(i)
	}
	set.Clear()
	for i = 0; i < numBits; i++ {
		if set.IsSet(i) {
			t.Fail()
		}
	}
}

func TestBitset_GetAllSet(t *testing.T) {
	var numBits uint32 = 1794
	var i uint32
	set := New(numBits)
	expectedSetSplice := make([]uint32, 0, numBits)
	for i = 0; i < numBits; i = i + 10 {
		set.Set(i)
		expectedSetSplice = append(expectedSetSplice, i)
	}

	//TODO
	//if  !bytes.Equal(set.GetAllSet(), expectedSetSplice) {
	//	t.Fail()
	//}
}
