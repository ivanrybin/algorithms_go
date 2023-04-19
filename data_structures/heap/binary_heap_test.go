package heap

import (
	"fmt"
	"reflect"
	"testing"

	"ivanrybin/algorithms_go/helpers"
)

func TestBinaryHeap_NewBinaryHeap(t *testing.T) {
	for _, tt := range []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{},
			want: []int{},
		},
		{
			in:   []int{0},
			want: []int{0},
		},
		{
			in:   []int{0, 1},
			want: []int{0, 1},
		},
		{
			in:   []int{1, 0},
			want: []int{0, 1},
		},
		{
			in:   []int{0, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			in:   []int{2, 1, 0},
			want: []int{0, 1, 2},
		},
		{
			in:   []int{0, 1, 2, 3},
			want: []int{0, 1, 2, 3},
		},
		{
			in:   []int{3, 2, 1, 0},
			want: []int{0, 2, 1, 3},
		},
		{
			in:   []int{0, 1, 2, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			heap := NewBinaryHeap[int](tt.in, MinIntComparator)
			if !reflect.DeepEqual(heap.xs, tt.want) {
				t.Errorf("heap=%v != want=%v", heap.xs, tt.want)
			}
		})
	}
}

func TestBinaryHeap_Insert(t *testing.T) {
	for _, tt := range helpers.TestArraysInt() {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			// empty heap
			heap := NewBinaryHeap[int]([]int{}, MinIntComparator)
			for _, x := range tt {
				heap.Insert(x)
			}
			validateHeap(t, heap, tt)
		})
	}
}

func TestBinaryHeap_Insert_Random(t *testing.T) {
	maxSize, maxValue := 100, 10
	for i := 0; i < 1000; i++ {
		xs := helpers.RandomInts(maxSize, maxValue)
		t.Run(fmt.Sprintf("%v", xs), func(t *testing.T) {
			// empty heap
			heap := NewBinaryHeap[int]([]int{}, MinIntComparator)
			for _, x := range xs {
				heap.Insert(x)
			}
			validateHeap(t, heap, xs)
		})
	}
}

func TestBinaryHeap_ExtractTop(t *testing.T) {
	for _, tt := range helpers.TestArraysInt() {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			// initialized heap
			heap := NewBinaryHeap[int](append([]int{}, tt...), MinIntComparator)
			validateHeap(t, heap, tt)
		})
	}
}

func TestBinaryHeap_ExtractTop_Random(t *testing.T) {
	maxSize, maxValue := 100, 10
	for i := 0; i < 1000; i++ {
		xs := helpers.RandomInts(maxSize, maxValue)
		t.Run(fmt.Sprintf("%v", xs), func(t *testing.T) {
			// initialized heap
			heap := NewBinaryHeap[int](append([]int{}, xs...), MinIntComparator)
			validateHeap(t, heap, xs)
		})
	}
}

func validateHeap(t *testing.T, heap *BinaryHeap[int], xs []int) {
	sorted := helpers.SortInts(xs)
	for i := 0; i < len(sorted); i++ {
		if heap.Size() != len(sorted)-i {
			t.Errorf("sorted:%v heap:%v Size()=%v != %v", sorted, heap.xs, heap.Size(), len(sorted)-i)
		}
		if top := heap.ExtractTop(); top != sorted[i] {
			t.Errorf("sorted:%v heap:%v ExtractTop()=%v != sorted[%v]=%v", sorted, heap.xs, top, i, sorted[i])
		}
	}
}
