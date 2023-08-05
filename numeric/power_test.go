package numeric

import (
	"math"
	"testing"
)

func TestPower_MathPow(t *testing.T) {
	for a := int64(-10); a <= 10; a++ {
		for b := uint(0); b <= 10; b++ {
			got, want := Power(a, b), int64(math.Pow(float64(a), float64(b)))
			if got != want {
				t.Errorf("Power(%v, %v) = %v != %v", a, b, got, want)
			}
		}
	}
}

func TestPowerNonRecursive_MathPow(t *testing.T) {
	for a := int64(-10); a <= 10; a++ {
		for b := uint(0); b <= 10; b++ {
			got, want := PowerNonRecursive(a, b), int64(math.Pow(float64(a), float64(b)))
			if got != want {
				t.Errorf("PowerNonRecursive(%v, %v) = %v != %v", a, b, got, want)
			}
		}
	}
}
