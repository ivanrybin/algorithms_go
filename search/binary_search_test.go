package search

import (
	"fmt"
	"testing"

	hs "github.com/ivanrybin/algorithms_go/helpers"
)

func TestBinarySearch_First_Uniq(t *testing.T) {
	for size := 0; size < 100; size++ {
		xs := hs.UniqSequentialInts(uint(size))
		for want, x := range xs {
			got := BinarySearch(x, xs, BinarySearchFirst)
			if got != want {
				t.Errorf("got=%v != want=%v for xs[%v]=%v", got, want, want, x)
			}
		}
	}
}

func TestBinarySearchLeftMost_Uniq(t *testing.T) {
	for size := 0; size < 100; size++ {
		xs := hs.UniqSequentialInts(uint(size))
		for want, x := range xs {
			got := BinarySearch(x, xs, BinarySearchLeftMost)
			if got != want {
				t.Errorf("got=%v != want=%v for xs[%v]=%v", got, want, want, x)
			}
		}
	}
}

func TestBinarySearchRightMost_Uniq(t *testing.T) {
	for size := 0; size < 100; size++ {
		xs := hs.UniqSequentialInts(uint(size))
		for want, x := range xs {
			got := BinarySearch(x, xs, BinarySearchRightMost)
			if got != want {
				t.Errorf("got=%v != want=%v for xs[%v]=%v", got, want, want, x)
			}
		}
	}
}

func TestBinarySearch_LeftMost_RightMost(t *testing.T) {
	for _, tt := range []struct {
		xs            []int
		x             int
		wantLeftMost  int
		wantRightMost int
	}{
		{
			xs:            []int{},
			x:             42,
			wantLeftMost:  -1,
			wantRightMost: -1,
		},
		{
			xs:            []int{0},
			x:             0,
			wantLeftMost:  0,
			wantRightMost: 0,
		},
		{
			xs:            []int{0, 0},
			x:             0,
			wantLeftMost:  0,
			wantRightMost: 1,
		},
		{
			xs:            []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			x:             0,
			wantLeftMost:  0,
			wantRightMost: 5,
		},
		{
			xs:            []int{-1, -1, -1, -1, 0, 0, 0, 0, 1, 1, 1, 1},
			x:             0,
			wantLeftMost:  4,
			wantRightMost: 7,
		},
		{
			xs:            []int{-1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			x:             0,
			wantLeftMost:  1,
			wantRightMost: 2,
		},
		{
			xs:            []int{-1, 0, 0},
			x:             0,
			wantLeftMost:  1,
			wantRightMost: 2,
		},
		{
			xs:            []int{-1, -1, 0, 0},
			x:             0,
			wantLeftMost:  2,
			wantRightMost: 3,
		},
		{
			xs:            []int{-1, -1, -1, 0, 0},
			x:             0,
			wantLeftMost:  3,
			wantRightMost: 4,
		},
		{
			xs:            []int{-1, -1, -1, -1, 0, 0, 0, 0},
			x:             0,
			wantLeftMost:  4,
			wantRightMost: 7,
		},
		{
			xs:            []int{-1, -1, -1, -1, -1, 0, 0, 0, 0, 0, 0},
			x:             0,
			wantLeftMost:  5,
			wantRightMost: 10,
		},
		{
			xs:            []int{-1, -1, -1, -1, -1, -1, 0, 0, 0, 0, 0, 0},
			x:             0,
			wantLeftMost:  6,
			wantRightMost: 11,
		},
		{
			xs:            []int{-1, -1, -1, -1, -1, -1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
			x:             0,
			wantLeftMost:  7,
			wantRightMost: 21,
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.xs), func(t *testing.T) {
			if got := BinarySearch(tt.x, tt.xs, BinarySearchLeftMost); got != tt.wantLeftMost {
				t.Errorf("got=%v != wantLeftMost=%v", got, tt.wantLeftMost)
			}
			if got := BinarySearch(tt.x, tt.xs, BinarySearchRightMost); got != tt.wantRightMost {
				t.Errorf("got=%v != wantRightMost=%v", got, tt.wantRightMost)
			}
		})
	}
}
