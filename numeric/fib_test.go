package numeric

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	for n, want := range fibonacciSequence {
		t.Run(fmt.Sprintf("%v", n), func(t *testing.T) {
			if got := Fib(uint(n)); got != want {
				t.Errorf("Fib() = %v, want %v", got, want)
			}
		})
	}
}

var fibonacciSequence = []uint64{
	0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229,
}
