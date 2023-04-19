package sorts

import "ivanrybin/algorithms_go/data_structures/heap"

// HeapSort O(n) / O(n * log n) (mem / time).
func HeapSort(xs []int) []int {
	h := heap.NewBinaryHeap[int](append([]int{}, xs...), heap.MaxIntComparator)
	for i := len(xs) - 1; i >= 0; i-- {
		xs[i] = h.ExtractTop()
	}
	return xs
}
