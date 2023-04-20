package sorts

import (
	"fmt"
	"testing"

	hs "ivanrybin/algorithms_go/helpers"
)

func TestCountingSort(t *testing.T) {
	for _, tt := range hs.TestArraysIntNonNegative() {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			hs.TestIntSort(t, tt, CountingSort)
		})
	}
}

func TestCountingSort_Random(t *testing.T) {
	maxSize, maxValue := 100, 10
	for i := 0; i < 1000; i++ {
		xs := hs.RandomInts(maxSize, maxValue)
		t.Run(fmt.Sprintf("%v", xs), func(t *testing.T) {
			hs.TestIntSort(t, xs, CountingSort)
		})
	}
}

func TestCountingSortStable(t *testing.T) {
	for _, tt := range [][]hs.Pair[int, string]{
		{},
		{{0, "A"}, {0, "B"}},
		{{1, "A"}, {0, "B"}},
		{{0, "A"}, {1, "B"}, {0, "C"}},
		{{1, "A"}, {1, "B"}, {0, "C"}},
		{{0, "A"}, {0, "B"}, {0, "C"}, {0, "D"}, {0, "E"}},
		{{0, "A"}, {1, "B"}, {0, "C"}, {1, "D"}, {0, "E"}},
		{{0, "A"}, {1, "B"}, {2, "C"}, {3, "D"}, {4, "E"}},
		{{4, "A"}, {4, "B"}, {3, "C"}, {3, "D"}, {2, "E"}, {2, "E"}},
	} {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			hs.TestPairSortStable[string](t, tt, CountingSortStable[string])
		})
	}
}

func TestCountingSortStable_Random(t *testing.T) {
	maxSize, maxValue := 100, 10
	for i := 0; i < 1000; i++ {
		xs := hs.RandomPairsIntInt(maxSize, maxValue)
		t.Run(fmt.Sprintf("%v", xs), func(t *testing.T) {
			hs.TestPairSortStable[int](t, xs, CountingSortStable[int])
		})
	}
}
