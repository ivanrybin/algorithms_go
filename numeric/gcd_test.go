package numeric

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	for _, tt := range []struct {
		a, b int64
		want int64
	}{
		{
			a: 10, b: 0, want: 10,
		},
		{
			a: 10, b: 1, want: 1,
		},
		{
			a: 10, b: 2, want: 2,
		},
		{
			a: 10, b: 3, want: 1,
		},
		{
			a: 10, b: 4, want: 2,
		},
		{
			a: 10, b: 5, want: 5,
		},
		{
			a: 10, b: 6, want: 2,
		},
		{
			a: 10, b: 7, want: 1,
		},
		{
			a: 10, b: 8, want: 2,
		},
		{
			a: 10, b: 9, want: 1,
		},
		{
			a: 10, b: 10, want: 10,
		},
		{
			a: 5, b: 10, want: 5,
		},
	} {
		t.Run(fmt.Sprintf("a=%v b=%v", tt.a, tt.b), func(t *testing.T) {
			if got := GCD(tt.a, tt.b); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEGCD(t *testing.T) {
	for _, tt := range []struct {
		a, b    int64
		wantGCD int64
		wantX   int64
		wantY   int64
	}{
		{
			a: 10, b: 0,
			wantGCD: 10,
			wantX:   1,
			wantY:   0,
		},
		{
			a: 15, b: 3,
			wantGCD: 3,
			wantX:   0,
			wantY:   1,
		},
		{
			a: 50, b: 40,
			wantGCD: 10,
			wantX:   1,
			wantY:   -1,
		},
		{
			a: 40, b: 50,
			wantGCD: 10,
			wantX:   -1,
			wantY:   1,
		},
		{
			a: -10, b: 0,
			wantGCD: -10,
			wantX:   1,
			wantY:   0,
		},
		{
			a: -15, b: 3,
			wantGCD: 3,
			wantX:   0,
			wantY:   1,
		},
		{
			a: 15, b: -3,
			wantGCD: -3,
			wantX:   0,
			wantY:   1,
		},
	} {
		t.Run(fmt.Sprintf("a=%v b=%v", tt.a, tt.b), func(t *testing.T) {
			d, x, y := EGCD(tt.a, tt.b)
			if tt.a*x+tt.b*y != d {
				t.Errorf("a * %v + b * %b = %v != GCD(a, b) = %v", x, y, tt.a*x+tt.b*y, d)
			}
			if d != tt.wantGCD {
				t.Errorf("gcd=%v != tt.wantGCD=%v", d, tt.wantGCD)
			}
			if x != tt.wantX {
				t.Errorf("x=%v != tt.wantX=%v", x, tt.wantX)
			}
			if y != tt.wantY {
				t.Errorf("y=%v != tt.wantY=%v", y, tt.wantY)
			}
		})
	}
}
