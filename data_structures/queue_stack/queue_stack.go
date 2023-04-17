package queue_list

import "ivanrybin/algorithms_go/data_structures"

type Queue[T any, S data_structures.Stack[T]] struct {
	S1 S
	S2 S
}

// Push - O(1)
func (q *Queue[T, S]) Push(v T) {
	q.S1.Push(v)
}

// Pop - amortized O(1)
func (q *Queue[T, S]) Pop() T {
	if q.S2.Size() == 0 {
		for q.S1.Size() != 0 {
			q.S2.Push(q.S1.Pop())
		}
	}
	return q.S2.Pop()
}

// Size - O(1)
func (q *Queue[T, S]) Size() int {
	return q.S1.Size() + q.S2.Size()
}
