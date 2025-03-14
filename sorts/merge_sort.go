package sorts

// MergeSort O(n * log n) / O(n) (time / mem).
func MergeSort(xs []int) []int {
	if len(xs) <= 1 {
		return xs
	}
	m := len(xs) / 2
	l := MergeSort(xs[:m])
	r := MergeSort(xs[m:])
	return merge(l, r)
}

// MergeSortNonRecursive O(n * log n) / O(n) (time / mem) but without recursion.
func MergeSortNonRecursive(xs []int) []int {
	n := len(xs)
	for size := 1; size < n; size *= 2 {
		for l := 0; l < n; l += 2 * size {
			if m, r := min(l+size, n), min(l+2*size, n); m < r {
				copy(xs[l:r], merge(xs[l:m], xs[m:r]))
			}
		}
	}
	return xs
}

// merge O(n) / O(n) (mem / time)
func merge(l, r []int) []int {
	i, j := 0, 0
	xs := make([]int, 0, len(l)+len(r))
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			xs = append(xs, l[i])
			i++
		} else {
			xs = append(xs, r[j])
			j++
		}
	}
	xs = append(xs, l[i:]...)
	xs = append(xs, r[j:]...)
	return xs
}
