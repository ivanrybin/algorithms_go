package bitset

import (
	"fmt"
	"testing"
)

func TestBitSet_Len(t *testing.T) {
	for _, tt := range []struct {
		n    uint
		want int
	}{
		{
			n:    0,
			want: 0,
		},
		{
			n:    15,
			want: 15,
		},
	} {
		t.Run(fmt.Sprintf("%v %v", tt.n, tt.want), func(t *testing.T) {
			if got := NewBitSet(tt.n).Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitSet_Cap(t *testing.T) {
	for _, tt := range []struct {
		n    uint
		want int
	}{
		{
			n:    0,
			want: 0,
		},
		{
			n:    2,
			want: 2,
		},
		{
			n:    7,
			want: 8,
		},
		{
			n:    8,
			want: 8,
		},
		{
			n:    9,
			want: 16,
		},
		{
			n:    15,
			want: 16,
		},
		{
			n:    16,
			want: 16,
		},
	} {
		t.Run(fmt.Sprintf("%v %v", tt.n, tt.want), func(t *testing.T) {
			if got := NewBitSet(tt.n).Cap(); got != tt.want {
				t.Errorf("Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}
