package queue_list

import "ivanrybin/algorithms_go/data_structures"

type Queue[T any, L data_structures.List[T]] struct {
	L L
}

// Push
// - O(1) singly/doubly linked list
// - amortized O(1) array list
func (q *Queue[T, L]) Push(v T) {
	q.L.PushBack(v)
}

// Pop
// - O(1) singly/doubly linked list
// - O(1) array list
func (q *Queue[T, L]) Pop() T {
	return q.L.PopFront()
}

// Size - O(1)
func (q *Queue[T, L]) Size() int {
	return q.L.Size()
}
