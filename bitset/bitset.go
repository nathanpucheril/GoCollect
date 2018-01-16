package bitset

import (
	"unsafe"
)

type Bitset struct {
	words []uint32
	size  uint32
}

func New(numBits uint32) Bitset {
	wordIndex, _ := getOffsets(numBits)
	return Bitset{make([]uint32, wordIndex+1, wordIndex+1), numBits}
}

func (self *Bitset) Set(i uint32) {
	self.validateIndex(i)
	wordIndex, bitOffset := getOffsets(i)
	self.words[wordIndex] |= 1 << (31 - bitOffset)
}

func (self *Bitset) Unset(i uint32) {
	self.validateIndex(i)
	wordIndex, bitOffset := getOffsets(i)
	self.words[wordIndex] &= ^(1 << (31 - bitOffset))
}

func (self *Bitset) Toggle(i uint32) {
	self.validateIndex(i)
	wordIndex, bitOffset := getOffsets(i)
	self.words[wordIndex] ^= 1 << (31 - bitOffset)
}

func (self *Bitset) IsSet(i uint32) bool {
	self.validateIndex(i)
	wordIndex, bitOffset := getOffsets(i)
	return self.words[wordIndex]&(1<<(31-bitOffset)) != 0
}

func (self *Bitset) Clear() {
	wordIndex, _ := getOffsets(self.size)
	self.words = make([]uint32, wordIndex+1, wordIndex+1)
}

func (self *Bitset) Xor(bitset *Bitset) []uint32 {
	panic("implement")
}
func (self *Bitset) Or(bitset *Bitset) []uint32 {
	panic("implement")
}

func (self *Bitset) And(bitset *Bitset) []uint32 {
	panic("implement")
}

func (self *Bitset) Cardinality() uint32 {
	var count uint32 = 0
	for _, word := range self.words {
		for word > 0 {
			word &= word - 1
			count++
		}
	}
	return count
}

func (self *Bitset) GetAllSet() []uint32 {
	set := make([]uint32, 0, self.size)
	var i uint32 = 0
	for ; i < self.size; i++ {
		if self.IsSet(i) {
			set = append(set, i)
		}
	}
	return set
}

func (self *Bitset) validateIndex(i uint32) {
	if i < 0 || i >= self.size {
		panic("OutOfBounds: Bit set out of bounds")
	}
}

func getOffsets(i uint32) (uint32, uint32) {
	return i >> 5, i % 32 // divide by 32
}

func (self *Bitset) toByteArray() {
	byteArray := make([]byte, 0, self.size>>3)
	for _, word := range self.words {
		byteArrPtr := (*[4]byte)(unsafe.Pointer(&word))[:]
		byteArray = append(byteArray, byteArrPtr...)
	}
}
