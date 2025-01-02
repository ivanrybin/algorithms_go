package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	for _, tt := range []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{0},
			want: []int{0},
		},
		{
			in:   []int{0, 1},
			want: []int{0, 1},
		},
		{
			in:   []int{0, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			in:   []int{0, 1, 2, 3},
			want: []int{0, 1, 2, 3},
		},
		{
			in:   []int{0, 1, 2, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			in:   []int{1, 0, 1},
			want: []int{1, 2},
		},
		{
			in:   []int{2, 1, 0, 1, 2},
			want: []int{2, 3, 4},
		},
		{
			in:   []int{3, 2, 1, 0, 1, 2, 3},
			want: []int{3, 4, 5, 6},
		},
		{
			in:   []int{5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5},
			want: []int{5, 6, 7, 8, 9, 10},
		},
		{
			in:   []int{1, 0},
			want: []int{0},
		},
		{
			in:   []int{2, 1, 0},
			want: []int{0},
		},
		{
			in:   []int{3, 2, 1, 0},
			want: []int{0},
		},
		{
			in:   []int{4, 3, 2, 1, 0},
			want: []int{0},
		},
		{
			in:   []int{0, 0, 0, 0},
			want: []int{0},
		},
		{
			in:   []int{0, 0, 0, 1},
			want: []int{0, 3},
		},
		{
			in:   []int{0, 0, 0, 1, 0, 0, 2},
			want: []int{0, 3, 6},
		},
		{
			in:   []int{0, 10, 2, 15, 4, 20, 6, 25, 8, 35, 10, 40},
			want: []int{0, 1, 3, 5, 7, 9, 11},
		},
		{
			in:   []int{0, 1, 2, 0, 1, 2, 0, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			in:   []int{0, 1, 2, 0, 1, 2, 3, 0, 1, 2},
			want: []int{0, 1, 2, 6},
		},
		{
			in:   []int{0, 1, 2, 0, 1, 2, 0, 1, 2, 3},
			want: []int{0, 1, 2, 9},
		},
		{
			in:   []int{8, 3, 4, 6, 5, 2, 0, 7, 9, 1},
			want: []int{1, 2, 3, 7, 8},
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got := LongestIncreasingSubsequence(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LongestIncreasingSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestIncreasingSubsequenceModified(t *testing.T) {
	for _, tt := range []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{0},
			want: []int{0},
		},
		{
			in:   []int{0, 1},
			want: []int{0, 1},
		},
		{
			in:   []int{0, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			in:   []int{0, 1, 2, 3},
			want: []int{0, 1, 2, 3},
		},
		{
			in:   []int{0, 1, 2, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			in:   []int{1, 0, 1},
			want: []int{1, 2},
		},
		{
			in:   []int{2, 1, 0, 1, 2},
			want: []int{2, 3, 4},
		},
		{
			in:   []int{3, 2, 1, 0, 1, 2, 3},
			want: []int{3, 4, 5, 6},
		},
		{
			in:   []int{5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5},
			want: []int{5, 6, 7, 8, 9, 10},
		},
		{
			in:   []int{1, 0},
			want: []int{1},
		},
		{
			in:   []int{2, 1, 0},
			want: []int{2},
		},
		{
			in:   []int{3, 2, 1, 0},
			want: []int{3},
		},
		{
			in:   []int{4, 3, 2, 1, 0},
			want: []int{4},
		},
		{
			in:   []int{0, 0, 0, 0},
			want: []int{0},
		},
		{
			in:   []int{0, 0, 0, 1},
			want: []int{0, 3},
		},
		{
			in:   []int{0, 0, 0, 1, 0, 0, 2},
			want: []int{0, 3, 6},
		},
		{
			in:   []int{0, 10, 2, 15, 4, 20, 6, 25, 8, 35, 10, 40},
			want: []int{0, 2, 4, 6, 8, 10, 11},
		},
		{
			in:   []int{0, 1, 2, 0, 1, 2, 0, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			in:   []int{0, 1, 2, 0, 1, 2, 3, 0, 1, 2},
			want: []int{0, 1, 2, 6},
		},
		{
			in:   []int{0, 1, 2, 0, 1, 2, 0, 1, 2, 3},
			want: []int{0, 1, 2, 9},
		},
		{
			in:   []int{8, 3, 4, 6, 5, 2, 0, 7, 9, 1},
			want: []int{1, 2, 4, 7, 8},
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got := LongestIncreasingSubsequenceModified(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LongestIncreasingSubsequenceModified() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestIncreasingSubsequenceModifiedFast(t *testing.T) {
	for _, tt := range []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{0},
			want: []int{0},
		},
		{
			in:   []int{0, 1},
			want: []int{0, 1},
		},
		{
			in:   []int{0, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			in:   []int{0, 1, 2, 3},
			want: []int{0, 1, 2, 3},
		},
		{
			in:   []int{0, 1, 2, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			in:   []int{1, 0, 1},
			want: []int{1, 2},
		},
		{
			in:   []int{2, 1, 0, 1, 2},
			want: []int{2, 3, 4},
		},
		{
			in:   []int{3, 2, 1, 0, 1, 2, 3},
			want: []int{3, 4, 5, 6},
		},
		{
			in:   []int{5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5},
			want: []int{5, 6, 7, 8, 9, 10},
		},
		{
			in:   []int{1, 0},
			want: []int{1},
		},
		{
			in:   []int{2, 1, 0},
			want: []int{2},
		},
		{
			in:   []int{3, 2, 1, 0},
			want: []int{3},
		},
		{
			in:   []int{4, 3, 2, 1, 0},
			want: []int{4},
		},
		{
			in:   []int{0, 0, 0, 0},
			want: []int{0},
		},
		{
			in:   []int{0, 0, 0, 1},
			want: []int{0, 3},
		},
		{
			in:   []int{0, 0, 0, 1, 0, 0, 2},
			want: []int{0, 3, 6},
		},
		{
			in:   []int{0, 10, 2, 15, 4, 20, 6, 25, 8, 35, 10, 40},
			want: []int{0, 2, 4, 6, 8, 10, 11},
		},
		{
			in:   []int{0, 1, 2, 0, 1, 2, 0, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			in:   []int{0, 1, 2, 0, 1, 2, 3, 0, 1, 2},
			want: []int{0, 1, 2, 6},
		},
		{
			in:   []int{0, 1, 2, 0, 1, 2, 0, 1, 2, 3},
			want: []int{0, 1, 2, 9},
		},
		{
			in:   []int{8, 3, 4, 6, 5, 2, 0, 7, 9, 1},
			want: []int{1, 2, 4, 7, 8},
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got := LongestIncreasingSubsequenceModifiedFast(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LongestIncreasingSubsequenceModifiedFast() = %v, want %v", got, tt.want)
			}
		})
	}
}
