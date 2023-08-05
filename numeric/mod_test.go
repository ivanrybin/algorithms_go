package numeric

import (
	"math"
	"testing"
)

func TestPowerMod_MathPow(t *testing.T) {
	for a := int64(-10); a <= 10; a++ {
		for b := uint(0); b <= 10; b++ {
			for m := int64(1); m < 10; m++ {
				got, want := PowerMod(a, b, m), int64(math.Pow(float64(a), float64(b)))%m
				if got != want {
					t.Errorf("PowerMod(%v, %v, %v) = %v != %v", a, b, m, got, want)
				}
			}
		}
	}
}

func TestPowerModNonRecursive_MathPow(t *testing.T) {
	for a := int64(-10); a <= 10; a++ {
		for b := uint(0); b <= 10; b++ {
			for m := int64(1); m < 10; m++ {
				got, want := PowerModNonRecursive(a, b, m), int64(math.Pow(float64(a), float64(b)))%m
				if got != want {
					t.Errorf("PowerModNonRecursive(%v, %v, %v) = %v != %v", a, b, m, got, want)
				}
			}
		}
	}
}
