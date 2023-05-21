package helpers

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue_Push_Pop(t *testing.T) {
	n := 100
	pq := NewPriorityQueue[int](func(l, r int) bool {
		return l < r
	})
	for i := 100; i > 0; i-- {
		heap.Push(pq, &Elem[int]{Value: i})
	}
	for i := 1; i <= n; i++ {
		x := heap.Pop(pq).(*Elem[int])
		if x.Value != i {
			t.Errorf("want=%v got=%v", i, x)
		}
		if x.Index() != -1 {
			t.Errorf("x.Index()=%v != %v", x.Index(), -1)
		}
		if pq.Len() != n-i {
			t.Errorf("want=%v got=%v", n-i, pq.Len())
		}
	}
}
