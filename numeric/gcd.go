package numeric

// GCD finds GCD of two numbers.
func GCD(a uint64, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// EGCD finds EGCD of two numbers.
//
//	q = a // b
//	r = a - bq = a % b
//	ax  +  by  = GCD(a,b) = GCD(b, r)
//	bx' +  ry' = GCD(b, r)
//	bx' +  (a - bq) * y' = GCD(b, r) = d
//	ay' +  b * (x' - qy') = GCD(b, r) = d
//	x = y'
//	y = x' - qy'
func EGCD(a int64, b int64) (d int64, x int64, y int64) {
	if b == 0 {
		return a, 1, 0
	}
	q, r := a/b, a%b
	d, x_, y_ := EGCD(b, r)
	return d, y_, x_ - q*y_
}
