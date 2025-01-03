package search

import "fmt"

type BinarySearchStrategy = int

const (
	BinarySearchFirst BinarySearchStrategy = iota
	BinarySearchLeftMost
	BinarySearchRightMost
)

// BinarySearch O(log n) worst case.
func BinarySearch(x int, xs []int, strategy BinarySearchStrategy) int {
	switch strategy {
	case BinarySearchFirst:
		return binarySearchFirst(x, xs, 0, len(xs))
	case BinarySearchLeftMost:
		return binarySearchLeftMost(x, xs, 0, len(xs))
	case BinarySearchRightMost:
		return binarySearchRightMost(x, xs, 0, len(xs))
	default:
		panic(fmt.Sprintf("unknown strategy: %v", strategy))
	}
}

// binarySearchFirst segment [l, r).
func binarySearchFirst(x int, xs []int, l, r int) int {
	for l < r {
		m := l + (r-l)/2
		if xs[m] == x {
			return m
		}
		if xs[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	return -1
}

// binarySearchLeftMost segment [l, r).
func binarySearchLeftMost(x int, xs []int, l, r int) int {
	id := -1
	for l < r {
		m := l + (r-l)/2
		if xs[m] >= x {
			r = m
			id = m
		} else {
			l = m + 1
		}
	}
	return id
}

// binarySearchRightMost segment [l, r).
func binarySearchRightMost(x int, xs []int, l, r int) int {
	id := -1
	for l < r {
		m := l + (r-l)/2
		if xs[m] <= x {
			l = m + 1
			id = m
		} else {
			r = m
		}
	}
	return id
}
