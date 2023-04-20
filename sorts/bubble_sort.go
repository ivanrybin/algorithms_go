package sorts

// BubbleSort O(n) / O(n^2) (mem / time).
func BubbleSort(xs []int) []int {
	for i := 0; i+1 < len(xs); i++ {
		for j := 0; j < len(xs)-i-1; j++ {
			if xs[j] > xs[j+1] {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
		}
	}
	return xs
}
