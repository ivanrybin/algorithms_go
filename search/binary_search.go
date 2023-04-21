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

func binarySearchFirst(x int, xs []int, l, r int) int {
	if l >= r {
		return -1
	}
	m := l + (r-l)/2
	if xs[m] == x {
		return m
	}
	if xs[m] < x {
		return binarySearchFirst(x, xs, m+1, r)
	}
	return binarySearchFirst(x, xs, l, m)
}

func binarySearchLeftMost(x int, xs []int, l, r int) int {
	if l >= r {
		return -1
	}
	m := l + (r-l)/2
	if xs[m] == x {
		for leftMost := binarySearchLeftMost(x, xs, l, m); leftMost != -1; leftMost = binarySearchLeftMost(x, xs, l, m) {
			m = leftMost
		}
		return m
	}
	if xs[m] < x {
		return binarySearchLeftMost(x, xs, m+1, r)
	}
	return binarySearchLeftMost(x, xs, l, m)
}

func binarySearchRightMost(x int, xs []int, l, r int) int {
	if l >= r {
		return -1
	}
	m := l + (r-l)/2
	if xs[m] == x {
		for rightMost := binarySearchRightMost(x, xs, m+1, r); rightMost != -1; rightMost = binarySearchRightMost(x, xs, m+1, r) {
			m = rightMost
		}
		return m
	}
	if xs[m] < x {
		return binarySearchRightMost(x, xs, m+1, r)
	}
	return binarySearchRightMost(x, xs, l, m)
}
