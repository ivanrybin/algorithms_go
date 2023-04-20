package sorts

// InsertionSort O(n) / O(n^2) (mem / time) - big amount of elements copying.
func InsertionSort(xs []int) []int {
	for i := 0; i < len(xs); i++ {
		for j := i; j > 0; j-- {
			if xs[j-1] > xs[j] {
				xs[j-1], xs[j] = xs[j], xs[j-1]
			}
		}
	}
	return xs
}
