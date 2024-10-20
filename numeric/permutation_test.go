package numeric

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAllPermutations(t *testing.T) {
	type test struct {
		xs   []int
		want [][]int
	}
	for _, tt := range []test{
		{
			xs: []int{1},
			want: [][]int{
				{1},
			},
		},
		{
			xs: []int{1, 1},
			want: [][]int{
				{1, 1},
			},
		},
		{
			xs: []int{1, 2},
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			xs: []int{1, 1, 1},
			want: [][]int{
				{1, 1, 1},
			},
		},
		{
			xs: []int{1, 1, 2},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			xs: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			xs: []int{1, 2, 3, 4},
			want: [][]int{
				{1, 2, 3, 4},
				{1, 2, 4, 3},
				{1, 3, 2, 4},
				{1, 3, 4, 2},
				{1, 4, 2, 3},
				{1, 4, 3, 2},
				{2, 1, 3, 4},
				{2, 1, 4, 3},
				{2, 3, 1, 4},
				{2, 3, 4, 1},
				{2, 4, 1, 3},
				{2, 4, 3, 1},
				{3, 1, 2, 4},
				{3, 1, 4, 2},
				{3, 2, 1, 4},
				{3, 2, 4, 1},
				{3, 4, 1, 2},
				{3, 4, 2, 1},
				{4, 1, 2, 3},
				{4, 1, 3, 2},
				{4, 2, 1, 3},
				{4, 2, 3, 1},
				{4, 3, 1, 2},
				{4, 3, 2, 1},
			},
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.xs), func(t *testing.T) {
			if got := AllPermutations(tt.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextPermutation(t *testing.T) {
	type test struct {
		xs   []int
		want []int
	}
	for _, tt := range []test{
		{
			xs:   []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			xs:   []int{1, 3, 2},
			want: []int{2, 1, 3},
		},
		{
			xs:   []int{2, 1, 3},
			want: []int{2, 3, 1},
		},
		{
			xs:   []int{2, 3, 1},
			want: []int{3, 1, 2},
		},
		{
			xs:   []int{3, 1, 2},
			want: []int{3, 2, 1},
		},
		{
			xs:   []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.xs), func(t *testing.T) {
			if got := NextPermutation(tt.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
