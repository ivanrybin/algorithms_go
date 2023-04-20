package sorts

import hs "ivanrybin/algorithms_go/helpers"

// CountingSort O(n + m) / O(n + m) (mem / time) where m - max element in xs.
//
// WARN: ONLY FOR NON-NEGATIVE NUMBERS.
func CountingSort(xs []int) []int {
	if len(xs) <= 1 {
		return xs
	}
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	count := make([]int, max+1)
	for _, x := range xs {
		count[x]++
	}
	sorted := make([]int, 0, len(xs))
	for x, c := range count {
		for ; c > 0; c-- {
			sorted = append(sorted, x)
		}
	}
	return sorted
}

// CountingSortStable O(n + m) / O(n + m) (mem / time) where m - max element in xs.
//
// WARN: ONLY FOR NON-NEGATIVE NUMBERS.
func CountingSortStable[R any](xs []hs.Pair[int, R]) []hs.Pair[int, R] {
	if len(xs) <= 1 {
		return xs
	}
	max := xs[0].L
	for _, p := range xs {
		if p.L > max {
			max = p.L
		}
	}
	count := make([]int, max+1)
	for _, p := range xs {
		count[p.L]++
	}
	// partial sum to restore order of equal elements in the initial array
	partialSum := make([]int, max+1)
	for i := 1; i < len(count); i++ {
		partialSum[i] = partialSum[i-1] + count[i-1]
	}
	sorted := make([]hs.Pair[int, R], len(xs))
	for _, x := range xs {
		sorted[partialSum[x.L]] = x
		partialSum[x.L]++
	}
	return sorted
}
