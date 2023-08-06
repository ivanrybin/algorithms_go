package numeric

// Power just an example if a and b were bits arrays.
func Power(a int64, b uint) int64 {
	if b == 0 {
		return 1
	}
	x := Power(a, b>>1)
	x = x * x
	if b%2 != 0 {
		x *= a
	}
	return x
}

// PowerNonRecursive achieves b by powers of two.
//
// Input: b = 30.
//
// iteration 1:
//
//	p = [2, 16]
//	b = b - 16 = 14
//
// iteration 2:
//
//	p = [2, 8]
//	b = b - 8 = 6
//
// iteration 3:
//
//	p = [2, 4]
//	b = b - 4 = 2
//
// iteration 4:
//
//	p = [2, 2]
//	b = b - 2 = 0 -> done
func PowerNonRecursive(a int64, b uint) int64 {
	res := int64(1)
	for b > 0 {
		x := a
		p := uint(2)
		for ; p<<1 <= b; p <<= 1 {
			x *= x
		}
		b -= p >> 1
		res *= x
	}
	return res
}
