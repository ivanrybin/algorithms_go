package helpers

import (
	"reflect"
	"sort"
	"testing"
)

func SortInts(xs []int) []int {
	cp := make([]int, len(xs))
	copy(cp, xs)
	sort.Ints(cp)
	return cp
}

func SortPairsStable[T any](xs []Pair[int, T]) []Pair[int, T] {
	cp := make([]Pair[int, T], len(xs))
	copy(cp, xs)
	sort.SliceStable(cp, func(i, j int) bool {
		return cp[i].L < cp[j].L
	})
	return cp
}

func TestArraysInt() [][]int {
	return [][]int{
		{},
		{1},
		{2, 1},
		{3, 1, 2},
		{4, 3, 2, 1},
		{5, 1, 3, 4, 2},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 0, 0, 0, -1},
		{-1, 0, 0, 0, 1},
		{-1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, -1},
		{42, 0, 1, 0, -1},
		{42, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, -1},
	}
}

func TestArraysIntNonNegative() [][]int {
	return [][]int{
		{},
		{1},
		{2, 1},
		{3, 1, 2},
		{4, 3, 2, 1},
		{5, 1, 3, 4, 2},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0},
		{42, 0, 1, 0},
		{42, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
	}
}

func TestIntSort(t *testing.T, xs []int, sort func([]int) []int) {
	want := SortInts(xs)
	got := sort(xs)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("size=%v want=%v != got=%v", len(xs), want, got)
	}
}

func TestPairSortStable[T any](t *testing.T, xs []Pair[int, T], sort func([]Pair[int, T]) []Pair[int, T]) {
	want := SortPairsStable(xs)
	got := sort(xs)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("size=%v want=%v != got=%v", len(xs), want, got)
	}
}
