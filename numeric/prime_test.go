package numeric

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	for _, tt := range []struct {
		x    uint64
		want bool
	}{
		{
			x: 0, want: false,
		},
		{
			x: 1, want: false,
		},
		{
			x: 2, want: true,
		},
		{
			x: 3, want: true,
		},
		{
			x: 4, want: false,
		},
		{
			x: 5, want: true,
		},
		{
			x: 6, want: false,
		},
		{
			x: 1005, want: false,
		},
		{
			x: 2147483647, want: true,
		},
	} {
		t.Run(fmt.Sprintf("x=%v", tt.x), func(t *testing.T) {
			if got, divisor := IsPrime(tt.x); got != tt.want {
				t.Errorf("IsPrime() = %v (divisor=%v), want %v", got, divisor, tt.want)
			}
		})
	}
}
