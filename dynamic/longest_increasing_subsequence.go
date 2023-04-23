package dynamic

// https://cp-algorithms.com/sequences/longest_increasing_subsequence.html

import (
	"math"
	"sort"
)

// LongestIncreasingSubsequence O(n) / O(n^2) (mem / time).
func LongestIncreasingSubsequence(xs []int) []int {
	// d[i] - the length of the longest increasing subsequence that ends on ith element
	prev, d := make([]int, len(xs)), make([]int, len(xs))
	for i := 0; i < len(prev); i++ {
		prev[i], d[i] = -1, 1
	}
	// subsequence calculation
	for i := 1; i < len(xs); i++ {
		for j := 0; j < i; j++ {
			if xs[j] < xs[i] && d[j]+1 > d[i] {
				d[i] = d[j] + 1
				prev[i] = j
			}
		}
	}
	// the end of the LIS
	maxPos := 0
	for i := 0; i < len(d); i++ {
		if d[i] > d[maxPos] {
			maxPos = i
		}
	}
	// indices restoring
	lis := make([]int, 0, d[maxPos])
	lis = append(lis, maxPos)
	for p := prev[maxPos]; p != -1; p = prev[p] {
		lis = append(lis, p)
	}
	// reversing indices
	for i := 0; i < len(lis)/2; i++ {
		lis[i], lis[len(lis)-i-1] = lis[len(lis)-i-1], lis[i]
	}
	return lis
}

// LongestIncreasingSubsequenceModified O(n) / O(n^2) (mem / time).
func LongestIncreasingSubsequenceModified(xs []int) []int {
	// d[l] - the smallest element at which an increasing subsequence of length l ends
	prev, d := make([]int, len(xs)+1), make([]int, len(xs)+1)
	prev[0], d[0] = -1, math.MinInt
	for i := 1; i < len(d); i++ {
		d[i] = math.MaxInt
		prev[i] = -1
	}
	// subsequence calculation
	for i := 0; i < len(xs); i++ {
		for l := 1; l <= len(xs); l++ {
			if d[l-1] < xs[i] && xs[i] < d[l] {
				d[l] = xs[i]
				prev[l] = i
			}
		}
	}
	// indices restoring
	lis := make([]int, 0, len(xs))
	for _, p := range prev {
		if p != -1 {
			lis = append(lis, p)
		}
	}
	return lis
}

// LongestIncreasingSubsequenceModifiedFast O(n) / O(n * log n) (mem / time).
func LongestIncreasingSubsequenceModifiedFast(xs []int) []int {
	// d[l] - the smallest element at which an increasing subsequence of length l ends
	prev, d := make([]int, len(xs)+1), make([]int, len(xs)+1)
	prev[0], d[0] = -1, math.MinInt
	for i := 1; i < len(d); i++ {
		d[i] = math.MaxInt
		prev[i] = -1
	}
	// subsequence calculation
	for i := 0; i < len(xs); i++ {
		if l := sort.SearchInts(d, xs[i]); d[l-1] < xs[i] && xs[i] < d[l] {
			d[l] = xs[i]
			prev[l] = i
		}
	}
	// indices restoring
	lis := make([]int, 0, len(xs))
	for _, p := range prev {
		if p != -1 {
			lis = append(lis, p)
		}
	}
	return lis
}
