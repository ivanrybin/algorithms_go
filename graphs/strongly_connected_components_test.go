package graphs

import (
	"fmt"
	"testing"

	"github.com/ivanrybin/algorithms_go/helpers"
)

func TestSCC(t *testing.T) {
	for _, tt := range []struct {
		name string
		in   Graph
		want [][]int
	}{
		{
			in: Graph{
				0: {0},
			},
			want: [][]int{
				{0},
			},
		},
		{
			in: Graph{
				0: {1},
				1: {0},
				2: {3},
				3: {2},
			},
			want: [][]int{
				{0, 1},
				{2, 3},
			},
		},
		{
			in: Graph{
				0: {1},
				1: {2},
				2: {0},
			},
			want: [][]int{
				{0, 1, 2},
			},
		},
		{
			in: Graph{
				0: {1},
				1: {0},
				2: {2},
			},
			want: [][]int{
				{0, 1},
				{2},
			},
		},
		{
			in: Graph{
				0: {1},
				1: {2, 3, 6},
				2: {0, 3},
				3: {4},
				4: {3, 5},
				5: {5},
				6: {4, 7},
				7: {5, 6},
			},
			want: [][]int{
				{0, 1, 2},
				{3, 4},
				{5},
				{6, 7},
			},
		},
	} {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got := helpers.MapSlicesToSliceSlices(SCC(tt.in)); !helpers.CCEqual(got, tt.want) {
				t.Errorf("SCC() = %v, want %v", got, tt.want)
			}
		})
	}
}
