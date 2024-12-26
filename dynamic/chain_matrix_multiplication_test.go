package dynamic

import "testing"

func Test_ChainMatrixMultiplication(t *testing.T) {
	type test struct {
		name       string
		dimensions []int64
		want       int64
	}
	for _, tt := range []test{
		{
			name:       "20x10",
			dimensions: []int64{20, 10},
			want:       0,
		},
		{
			name:       "10x10 10x10",
			dimensions: []int64{10, 10, 10, 10},
			want:       1000,
		},
		{
			name:       "30x10 10x20 20x40",
			dimensions: []int64{30, 10, 10, 20, 20, 40},
			want:       20000,
		},
		{
			name:       "50x20 20x1 1x10 10x100",
			dimensions: []int64{50, 20, 20, 1, 1, 10, 10, 100},
			want:       7000,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChainMatrixMultiplication(tt.dimensions); got != tt.want {
				t.Errorf("ChainMatrixMultiplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
