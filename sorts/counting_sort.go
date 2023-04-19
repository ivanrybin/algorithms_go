package sorts

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
	c := make([]int, max+1)
	for _, x := range xs {
		c[x]++
	}
	sorted := make([]int, 0, len(xs))
	for x, count := range c {
		for ; count > 0; count-- {
			sorted = append(sorted, x)
		}
	}
	return sorted
}
