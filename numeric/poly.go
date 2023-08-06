package numeric

import (
	"math"
)

// PolyMul multiplication of two polynomials based on Fast Fourier transform (FFT).
//
// PolyMul(a, b) = 1/n * FFT(FFT(a, w) * FFT(b, w), w^-1)
func PolyMul(a []int, b []int) []int {
	n := 1
	for n < int(math.Max(float64(len(a)), float64(len(b)))) {
		n *= 2
	}
	n *= 2
	ac := make([]complex128, n)
	bc := make([]complex128, n)
	for i, x := range a {
		ac[i] = complex(float64(x), float64(0))
	}
	for i, x := range b {
		bc[i] = complex(float64(x), float64(0))
	}
	angle := 2.0 * math.Pi / float64(n)
	w := complex(math.Cos(angle), math.Sin(angle))
	ac = FFT(ac, w)
	bc = FFT(bc, w)
	for i := 0; i < n; i++ {
		ac[i] *= bc[i]
	}
	cc := FFT(ac, 1/w)
	c := make([]int, n)
	acc := 0
	for i, x := range cc {
		d := int(real(x)/float64(n) + 0.5)
		c[i] = (acc + d) % 10
		acc = (acc + d) / 10
	}
	// leading zeros trim
	for i := len(c) - 1; i >= 0; i-- {
		if c[i] != 0 || i == 0 {
			c = c[:i+1]
			break
		}
	}
	return c
}
