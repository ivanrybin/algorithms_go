package list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDoublyLinkedList_PushBack(t *testing.T) {
	for size := 0; size < 10; size++ {
		want := make([]int, 0, size)
		l := DoublyLinkedList[int]{}
		for j := 0; j < size; j++ {
			want = append(want, j)
			l.PushBack(j)
		}
		if l.Size() != len(want) {
			t.Errorf("l.size=%v != len(want)=%v", l.size, len(want))
		}
		if got := l.GetAll(); !reflect.DeepEqual(got, want) {
			t.Errorf("got=%v != want=%v", got, want)
		}
	}
}

func TestDoublyLinkedList_PushFront(t *testing.T) {
	for size := 0; size < 10; size++ {
		want := make([]int, 0, size)
		l := DoublyLinkedList[int]{}
		for j := 0; j < size; j++ {
			want = append([]int{j}, want...)
			l.PushFront(j)
		}
		if l.Size() != len(want) {
			t.Errorf("l.size=%v != len(want)=%v", l.size, len(want))
		}
		if got := l.GetAll(); !reflect.DeepEqual(got, want) {
			t.Errorf("got=%v != want=%v", got, want)
		}
	}
}

func TestDoublyLinkedList_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		want := make([]int, 0, i)
		l := DoublyLinkedList[int]{}
		for j := 0; j < i; j++ {
			want = append(want, j)
			l.PushBack(j)
		}
		for j := 0; j < i; j++ {
			if got := l.Get(j); got != want[j] {
				t.Errorf("l.Get(%v)=%v != want[%v]=%v", j, got, j, want[j])
			}
		}
	}
}

func TestDoublyLinkedList_Invert(t *testing.T) {
	for _, tt := range []struct {
		data []int
		want []int
	}{
		{
			data: []int{},
			want: []int{},
		},
		{
			data: []int{1},
			want: []int{1},
		},
		{
			data: []int{1, 2},
			want: []int{2, 1},
		},
		{
			data: []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			data: []int{1, 2, 3, 4},
			want: []int{4, 3, 2, 1},
		},
		{
			data: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.data), func(t *testing.T) {
			l := SinglyLinkedList[int]{}
			for _, v := range tt.data {
				l.PushBack(v)
			}
			l.Invert()
			if got := l.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got=%v != want=%v", got, tt.want)
			}
			l.Invert()
			if got := l.GetAll(); !reflect.DeepEqual(got, tt.data) {
				t.Errorf("got=%v != data=%v", got, tt.data)
			}
			l.Invert()
			if got := l.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got=%v != want=%v", got, tt.want)
			}
		})
	}
}
