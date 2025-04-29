package bitset

import (
	"fmt"
	"math"

	hs "github.com/ivanrybin/algorithms_go/helpers"
)

var bits = []uint8{
	1 << 7,
	1 << 6,
	1 << 5,
	1 << 4,
	1 << 3,
	1 << 2,
	1 << 1,
	1 << 0,
}

type BitSet struct {
	l  int
	c  int
	bs []uint8
}

func NewBitSet(n uint) *BitSet {
	capacity := int(math.Pow(2.0, math.Ceil(math.Log2(float64(n)))))
	return &BitSet{
		l:  int(n),
		c:  capacity,
		bs: make([]uint8, hs.MaxInt(1, capacity/8)),
	}
}

func (s *BitSet) Len() int {
	return s.l
}

func (s *BitSet) Cap() int {
	return s.c
}

func (s *BitSet) Get(i int) uint8 {
	if i > s.Len() {
		panic(fmt.Sprintf("%v > len=%v", i, s.Len()))
	}
	return s.bs[i/8] & bits[i%8]
}

func (s *BitSet) Set(i int) {
	if i > s.Len() {
		panic(fmt.Sprintf("%v > len=%v", i, s.Len()))
	}
	s.bs[i/8] |= bits[i%8]
}
