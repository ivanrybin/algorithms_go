package sorts

import (
	"fmt"
	"testing"

	hs "ivanrybin/algorithms_go/helpers"
)

func TestInsertionSort(t *testing.T) {
	for _, tt := range hs.TestArraysInt() {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			hs.TestIntSort(t, tt, InsertionSort)
		})
	}
}

func TestInsertionSort_Random(t *testing.T) {
	maxSize, maxValue := 100, 10
	for i := 0; i < 1000; i++ {
		xs := hs.RandomInts(maxSize, maxValue)
		t.Run(fmt.Sprintf("%v", xs), func(t *testing.T) {
			hs.TestIntSort(t, xs, InsertionSort)
		})
	}
}
