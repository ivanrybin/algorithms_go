package sorts

import (
	"fmt"
	"testing"

	hs "github.com/ivanrybin/algorithms_go/helpers"
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
		{{L: 0, R: "A"}, {L: 0, R: "B"}},
		{{L: 1, R: "A"}, {L: 0, R: "B"}},
		{{L: 0, R: "A"}, {L: 1, R: "B"}, {L: 0, R: "C"}},
		{{L: 1, R: "A"}, {L: 1, R: "B"}, {L: 0, R: "C"}},
		{{L: 0, R: "A"}, {L: 0, R: "B"}, {L: 0, R: "C"}, {L: 0, R: "D"}, {L: 0, R: "E"}},
		{{L: 0, R: "A"}, {L: 1, R: "B"}, {L: 0, R: "C"}, {L: 1, R: "D"}, {L: 0, R: "E"}},
		{{L: 0, R: "A"}, {L: 1, R: "B"}, {L: 2, R: "C"}, {L: 3, R: "D"}, {L: 4, R: "E"}},
		{{L: 4, R: "A"}, {L: 4, R: "B"}, {L: 3, R: "C"}, {L: 3, R: "D"}, {L: 2, R: "E"}, {L: 2, R: "E"}},
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
