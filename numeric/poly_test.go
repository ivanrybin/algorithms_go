package numeric

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestPolyMul(t *testing.T) {
	for a := uint64(0); a < 1000; a++ {
		for b := a; b < 1000; b++ {
			want := toPoly(a * b)
			if got := PolyMul(toPoly(a), toPoly(b)); !reflect.DeepEqual(got, want) {
				t.Errorf("PolyMul(%v, %v) = %v, want %v", a, b, got, want)
			}
		}
	}
}

func TestPolyMul_Random(t *testing.T) {
	for a := uint64(0); a < 10_000; a += uint64(rand.Intn(100)) {
		for b := uint64(0); b < 10_000; b += uint64(rand.Intn(100)) {
			want := toPoly(a * b)
			if got := PolyMul(toPoly(a), toPoly(b)); !reflect.DeepEqual(got, want) {
				t.Errorf("PolyMul(%v, %v) = %v, want %v", a, b, got, want)
			}
		}
	}
}

func toPoly(a uint64) []int {
	if a == 0 {
		return []int{0}
	}
	n := 0
	for a_ := a; a_ != 0; a_ /= 10 {
		n++
	}
	p := make([]int, 0, n+1)
	for ; a != 0; a /= 10 {
		p = append(p, int(a%10))
	}
	return p
}
