package sorts

// InsertionSort O(n) / O(n^2) (mem / time)
func InsertionSort(xs []int) []int {
	for i := 0; i < len(xs); i++ {
		pos := i
		for j := i + 1; j < len(xs); j++ {
			if xs[j] < xs[pos] {
				pos = j
			}
		}
		xs[i], xs[pos] = xs[pos], xs[i]
	}
	return xs
}
