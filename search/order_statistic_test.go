package search

import (
	"fmt"
	"reflect"
	"testing"

	hs "github.com/ivanrybin/algorithms_go/helpers"
)

func TestOrderStatistic_Uniq_Ordered(t *testing.T) {
	for size := 1; size < 50; size++ {
		xs := hs.UniqSequentialInts(uint(size))
		t.Run(fmt.Sprintf("%v", xs), func(t *testing.T) {
			for k, want := range xs {
				if got := OrderStatistic(append([]int{}, xs...), k+1); got != want {
					t.Errorf("OrderStatistic(%v, %v)=%v != %v", xs, k+1, got, want)
				}
			}
		})
	}
}

func TestOrderStatistic(t *testing.T) {
	for _, tt := range []struct {
		in   []int
		k    int
		want int
	}{
		{
			in:   []int{42},
			k:    1,
			want: 42,
		},
		{
			in:   []int{42, 42},
			k:    1,
			want: 42,
		},
		{
			in:   []int{0, 42},
			k:    1,
			want: 0,
		},
		{
			in:   []int{42, 0},
			k:    2,
			want: 42,
		},
		{
			in:   []int{42, 0},
			k:    1,
			want: 0,
		},
		{
			in:   []int{0, 1337},
			k:    2,
			want: 1337,
		},
		{
			in:   []int{1, 0, 0, 0, 0, 0, 0},
			k:    7,
			want: 1,
		},
		{
			in:   []int{0, 2, 0, 42, 0, 4, 0},
			k:    7,
			want: 42,
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got := OrderStatistic(append([]int{}, tt.in...), tt.k); got != tt.want {
				t.Errorf("OrderStatistic(%v, %v)=%v != %v", tt.in, tt.k, got, tt.want)
			}
		})
	}
}

func TestPartition3(t *testing.T) {
	for _, tt := range []struct {
		in     []int
		p      int
		wantLR []int
		wantIn []int
	}{
		{
			in:     []int{0},
			p:      0,
			wantLR: []int{0, 0},
			wantIn: []int{0},
		},
		{
			in:     []int{0, 1},
			p:      0,
			wantLR: []int{0, 0},
			wantIn: []int{0, 1},
		},
		{
			in:     []int{0, 1},
			p:      1,
			wantLR: []int{1, 1},
			wantIn: []int{0, 1},
		},
		{
			in:     []int{1, 0},
			p:      0,
			wantLR: []int{1, 1},
			wantIn: []int{0, 1},
		},
		{
			in:     []int{1, 0, 0},
			p:      0,
			wantLR: []int{2, 2},
			wantIn: []int{0, 0, 1},
		},
		{
			in:     []int{0, 0, 0},
			p:      2,
			wantLR: []int{0, 2},
			wantIn: []int{0, 0, 0},
		},
		{
			in:     []int{1, 0, 2},
			p:      0,
			wantLR: []int{1, 1},
			wantIn: []int{0, 1, 2},
		},
		{
			in:     []int{0, 1, 2},
			p:      1,
			wantLR: []int{1, 1},
			wantIn: []int{0, 1, 2},
		},
		{
			in:     []int{2, 1, 0},
			p:      1,
			wantLR: []int{1, 1},
			wantIn: []int{0, 1, 2},
		},
		{
			in:     []int{1, 2, 0, 0, 2, 1},
			p:      0,
			wantLR: []int{2, 3},
			wantIn: []int{0, 0, 1, 1, 2, 2},
		},
		{
			in:     []int{2, 2, 2, 2, 2, 1, 0},
			p:      5,
			wantLR: []int{1, 1},
			wantIn: []int{0, 1, 2, 2, 2, 2, 2},
		},
		{
			in:     []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			p:      0,
			wantLR: []int{0, 4},
			wantIn: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		},
		{
			in:     []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			p:      1,
			wantLR: []int{5, 9},
			wantIn: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			lx, rx := partition3(tt.p, tt.in, 0, len(tt.in)-1)
			if !reflect.DeepEqual(tt.in, tt.wantIn) {
				t.Errorf("p=%v: got=%v != wantIn=%v lx=%v rx=%v", tt.p, tt.in, tt.wantIn, lx, rx)
			}
			if !reflect.DeepEqual([]int{lx, rx}, tt.wantLR) {
				t.Errorf("p=%v: got=[%v, %v] != wantLR=%v", tt.p, lx, rx, tt.wantLR)
			}
		})
	}
}
