package search

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/ivanrybin/algorithms_go/helpers"
)

func Test_NewRMQ_Min(t *testing.T) {
	for _, tt := range []struct {
		xs   []int
		want []int
	}{
		{
			xs:   []int{1},
			want: []int{1},
		},
		{
			xs:   []int{1, 2},
			want: []int{1, 1, 2},
		},
		{
			xs:   []int{1, 2, 3},
			want: []int{1, 1, 2, 0, 0, 2, 3},
		},
		{
			xs:   []int{1, 2, 3, 4},
			want: []int{1, 1, 3, 1, 2, 3, 4},
		},
		{
			xs:   []int{1, 2, 3, 4, 5},
			want: []int{1, 1, 3, 1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 4, 5},
		},
		{
			xs:   []int{1, 2, 3, 4, 5, 6},
			want: []int{1, 1, 4, 1, 2, 4, 5, 0, 0, 2, 3, 0, 0, 5, 6},
		},
		{
			xs:   []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{1, 1, 4, 1, 2, 4, 6, 0, 0, 2, 3, 4, 5, 6, 7},
		},
		{
			xs:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: []int{1, 1, 5, 1, 3, 5, 7, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.xs), func(t *testing.T) {
			if got := NewRMQ(tt.xs, helpers.MinInt[int]); !reflect.DeepEqual(got.vs, tt.want) {
				t.Errorf("NewRMQ(%v)=%v != %v", tt.xs, got.vs, tt.want)
			}
		})
	}
}

func Test_NewRMQ_Max(t *testing.T) {
	for _, tt := range []struct {
		xs   []int
		want []int
	}{
		{
			xs:   []int{1},
			want: []int{1},
		},
		{
			xs:   []int{1, 2},
			want: []int{2, 1, 2},
		},
		{
			xs:   []int{1, 2, 3},
			want: []int{3, 1, 3, 0, 0, 2, 3},
		},
		{
			xs:   []int{1, 2, 3, 4},
			want: []int{4, 2, 4, 1, 2, 3, 4},
		},
		{
			xs:   []int{1, 2, 3, 4, 5},
			want: []int{5, 2, 5, 1, 2, 3, 5, 0, 0, 0, 0, 0, 0, 4, 5},
		},
		{
			xs:   []int{1, 2, 3, 4, 5, 6},
			want: []int{6, 3, 6, 1, 3, 4, 6, 0, 0, 2, 3, 0, 0, 5, 6},
		},
		{
			xs:   []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{7, 3, 7, 1, 3, 5, 7, 0, 0, 2, 3, 4, 5, 6, 7},
		},
		{
			xs:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: []int{8, 4, 8, 2, 4, 6, 8, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.xs), func(t *testing.T) {
			if got := NewRMQ(tt.xs, helpers.MaxInt[int]); !reflect.DeepEqual(got.vs, tt.want) {
				t.Errorf("NewRMQ(%v)=%v != %v", tt.xs, got.vs, tt.want)
			}
		})
	}
}

func TestRMQ_Get_Random_MinInt(t *testing.T) {
	repeats, maxValue, maxSize := 5, 100, 128
	for size := 1; size <= maxSize; size++ {
		for i := 0; i < repeats; i++ {
			xs := helpers.RandomIntsExactSize(size, maxValue)
			t.Run(fmt.Sprintf("%v", len(xs)), func(t *testing.T) {
				rmq := NewRMQ(xs, helpers.MinInt[int])
				segments := helpers.GenSegments(0, size-1)
				checkSegments(t, rmq, xs, segments, helpers.FindIdempotentOnSegment[int])
			})
		}
	}
}

func TestRMQ_Get_Random_MaxInt(t *testing.T) {
	repeats, maxValue, maxSize := 5, 100, 128
	for size := 1; size <= maxSize; size++ {
		for i := 0; i < repeats; i++ {
			xs := helpers.RandomIntsExactSize(size, maxValue)
			t.Run(fmt.Sprintf("%v", len(xs)), func(t *testing.T) {
				rmq := NewRMQ(xs, helpers.MaxInt[int])
				segments := helpers.GenSegments(0, size-1)
				checkSegments(t, rmq, xs, segments, helpers.FindIdempotentOnSegment[int])
			})
		}
	}
}

func TestRMQ_Get_Sequential_MinInt(t *testing.T) {
	sequence := helpers.IncreasingInts(uint(128))
	for size := 1; size <= len(sequence); size++ {
		xs := sequence[0:size]
		t.Run(fmt.Sprintf("%v", size), func(t *testing.T) {
			rmq := NewRMQ(xs, helpers.MinInt[int])
			segments := helpers.GenSegments(0, size-1)
			checkSegments(t, rmq, xs, segments, helpers.FindMinOnSequentialSegment[int])
		})
	}
}

func TestRMQ_Get_Sequential_MaxInt(t *testing.T) {
	sequence := helpers.IncreasingInts(uint(128))
	for size := 1; size <= len(sequence); size++ {
		xs := sequence[0:size]
		t.Run(fmt.Sprintf("%v", size), func(t *testing.T) {
			rmq := NewRMQ(xs, helpers.MaxInt[int])
			segments := helpers.GenSegments(0, size-1)
			checkSegments(t, rmq, xs, segments, helpers.FindMaxOnSequentialSegment[int])
		})
	}
}

func TestRMQ_Update_Random_MinInt(t *testing.T) {
	repeats, maxValue, maxSize := 5, 100, 128
	for size := 1; size <= maxSize; size++ {
		xs, segments := helpers.RandomIntsExactSize(size, maxValue), helpers.GenSegments(0, size-1)
		t.Run(fmt.Sprintf("%v", len(xs)), func(t *testing.T) {
			rmq := NewRMQ(xs, helpers.MinInt[int])
			for i := 0; i < repeats; i++ {
				j, v := rand.Intn(rmq.n), rand.Intn(maxValue)
				rmq.Update(j, v)
				xs[j] = v
				checkSegments(t, rmq, xs, segments, helpers.FindIdempotentOnSegment[int])
			}
		})
	}
}

func TestRMQ_Update_Random_MaxInt(t *testing.T) {
	repeats, maxValue, maxSize := 5, 100, 128
	for size := 1; size <= maxSize; size++ {
		xs, segments := helpers.RandomIntsExactSize(size, maxValue), helpers.GenSegments(0, size-1)
		t.Run(fmt.Sprintf("%v", len(xs)), func(t *testing.T) {
			rmq := NewRMQ(xs, helpers.MaxInt[int])
			for i := 0; i < repeats; i++ {
				j, v := rand.Intn(rmq.n), rand.Intn(maxValue)
				rmq.Update(j, v)
				xs[j] = v
				checkSegments(t, rmq, xs, segments, helpers.FindIdempotentOnSegment[int])
			}
		})
	}
}

func checkSegments[T comparable](t *testing.T, rmq *RMQ[T], xs []T, segments [][]int, idempotent func(xs []T, l, r int, f func(l, r T) T) T) {
	for _, s := range segments {
		want := idempotent(xs, s[0], s[1], rmq.F())
		if got := rmq.Get(s[0], s[1]); got != want {
			t.Errorf("Get(%v, %v)=%v != %v", s[0], s[1], got, want)
		}
	}
}
