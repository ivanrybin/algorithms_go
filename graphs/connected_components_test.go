package graphs

import (
	"fmt"
	"testing"

	"github.com/ivanrybin/algorithms_go/helpers"
)

func TestCC(t *testing.T) {
	for _, tt := range []struct {
		name string
		in   Graph
		want [][]int
	}{
		{
			in: Graph{
				1: {},
			},
			want: [][]int{
				{1},
			},
		},
		{
			in: Graph{
				1: {},
				2: {},
			},
			want: [][]int{
				{1},
				{2},
			},
		},
		{
			in: Graph{
				1: {2},
				2: {1},
			},
			want: [][]int{
				{1, 2},
			},
		},
		{
			in: Graph{
				1: {},
				2: {},
				3: {},
			},
			want: [][]int{
				{1}, {2}, {3},
			},
		},
		{
			in: Graph{
				1: {2},
				2: {1},
				3: {4},
				4: {3},
			},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			in: Graph{
				1: {2, 3},
				2: {1},
				3: {1, 4},
				4: {3},
			},
			want: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			in: Graph{
				1: {2, 3},
				2: {1},
				3: {1, 4},
				4: {3},
				5: {6},
				6: {5, 7, 8},
				7: {6},
				8: {6},
				9: {},
			},
			want: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9},
			},
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got := helpers.MapSlicesToSliceSlices(CC(tt.in)); !helpers.CCEqual(got, tt.want) {
				t.Errorf("CC() = %v, want %v", got, tt.want)
			}
		})
	}
}
