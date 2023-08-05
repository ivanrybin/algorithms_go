package numeric

import "fmt"

// PlusMod by modulo m.
func PlusMod(a, b, m int64) int64 {
	return (a + b) % m
}

// MulMod by modulo m.
func MulMod(a, b, m int64) int64 {
	return (a * b) % m
}

// DivMod by modulo m.
func DivMod(a, b, m int64) (int64, error) {
	d, x, _ := EGCD(b, m)
	if d != 1 {
		return 0, fmt.Errorf("GCD(b, m) = %v != 1", d)
	}
	if x < 0 {
		x += m
	}
	return (a * x) % m, nil
}

// PowerMod by modulo m.
func PowerMod(a int64, b uint, m int64) int64 {
	if b == 0 {
		return 1 % m
	}
	x := PowerMod(a, b>>1, m)
	x = x * x
	if b%2 != 0 {
		x *= a
	}
	return x % m
}

// PowerModNonRecursive by modulo m, achieves b by powers of two, see comment for PowerNonRecursive.
func PowerModNonRecursive(a int64, b uint, m int64) int64 {
	res := int64(1)
	for b > 0 {
		x := a
		p := uint(2)
		for ; p<<1 <= b; p <<= 1 {
			x = (x * x) % m
		}
		b -= p >> 1
		res = (res * x) % m
	}
	return res % m
}
