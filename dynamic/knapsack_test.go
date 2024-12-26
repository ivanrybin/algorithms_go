package dynamic

import (
	"fmt"
	"testing"
)

func Test_KnapsackWithRepetition(t *testing.T) {
	type test struct {
		items []item
		W     int
		want  float64
	}
	for _, tt := range []test{
		{
			items: []item{},
			W:     10,
			want:  0,
		},
		{
			items: []item{
				{v: 100, w: 10},
			},
			W:    10,
			want: 100,
		},
		{
			items: []item{
				{v: 100, w: 11},
			},
			W:    10,
			want: 0,
		},
		{
			items: []item{
				{v: 30, w: 6},
				{v: 16, w: 4},
				{v: 14, w: 3},
				{v: 9, w: 2},
			},
			W:    10,
			want: 48,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tt.items), func(t *testing.T) {
			if got := KnapsackWithRepetition(tt.items, tt.W); got != tt.want {
				t.Errorf("KnapsackWithRepetition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_KnapsackWithoutRepetition(t *testing.T) {
	type test struct {
		items []item
		W     int
		want  float64
	}
	for _, tt := range []test{
		{
			items: []item{},
			W:     10,
			want:  0,
		},
		{
			items: []item{
				{v: 100, w: 10},
			},
			W:    10,
			want: 100,
		},
		{
			items: []item{
				{v: 100, w: 11},
			},
			W:    10,
			want: 0,
		},
		{
			items: []item{
				{v: 30, w: 6},
				{v: 16, w: 4},
				{v: 14, w: 3},
				{v: 9, w: 2},
			},
			W:    10,
			want: 46,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tt.items), func(t *testing.T) {
			if got := KnapsackWithoutRepetition(tt.items, tt.W); got != tt.want {
				t.Errorf("KnapsackWithoutRepetition() = %v, want %v", got, tt.want)
			}
		})
	}
}
