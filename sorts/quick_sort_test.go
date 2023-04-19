package sorts

import (
	"fmt"
	"testing"

	"ivanrybin/algorithms_go/helpers"
)

func TestQuickSort(t *testing.T) {
	for _, tt := range helpers.TestArraysInt() {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			helpers.TestSort(t, tt, QuickSort)
		})
	}
}

func TestQuickSort_Random(t *testing.T) {
	maxSize, maxValue := 100, 10
	for i := 0; i < 1000; i++ {
		xs := helpers.RandomInts(maxSize, maxValue)
		t.Run(fmt.Sprintf("%v", xs), func(t *testing.T) {
			helpers.TestSort(t, xs, QuickSort)
		})
	}
}
