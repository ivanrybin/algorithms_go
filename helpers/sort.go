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

func TestSort(t *testing.T, xs []int, sort func([]int) []int) {
	want := SortInts(xs)
	got := sort(xs)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("size=%v want=%v != got=%v", len(xs), want, got)
	}
}
