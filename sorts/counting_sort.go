package sorts

import hs "github.com/ivanrybin/algorithms_go/helpers"

// CountingSort O(n + m) / O(n + m) (mem / time) where m - max element in xs.
//
// WARN: ONLY FOR NON-NEGATIVE NUMBERS.
func CountingSort[T hs.IntegerType](xs []T) []T {
	if len(xs) <= 1 {
		return xs
	}
	m := xs[0]
	for _, x := range xs {
		if x > m {
			m = x
		}
	}
	count := make([]int, m+1)
	for _, x := range xs {
		count[x]++
	}
	sorted := make([]T, 0, len(xs))
	for x, c := range count {
		for ; c > 0; c-- {
			sorted = append(sorted, T(x))
		}
	}
	return sorted
}

// CountingSortStable O(n + m) / O(n + m) (mem / time) where m - max element in xs.
//
// WARN: ONLY FOR NON-NEGATIVE NUMBERS.
func CountingSortStable[L hs.IntegerType, R any](xs []hs.Pair[L, R]) []hs.Pair[L, R] {
	if len(xs) <= 1 {
		return xs
	}
	m := xs[0].L
	for _, p := range xs {
		if p.L > m {
			m = p.L
		}
	}
	count := make([]int, m+1)
	for _, p := range xs {
		count[p.L]++
	}
	// partial sum to restore order of equal elements in the initial array
	partialSum := make([]int, m+1)
	for i := 1; i < len(count); i++ {
		partialSum[i] = partialSum[i-1] + count[i-1]
	}
	sorted := make([]hs.Pair[L, R], len(xs))
	for _, x := range xs {
		sorted[partialSum[x.L]] = x
		partialSum[x.L]++
	}
	return sorted
}
