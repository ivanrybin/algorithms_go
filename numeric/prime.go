package numeric

// IsPrime checks whether a number is a prime number.
func IsPrime(x uint64) (bool, uint64) {
	if x < 2 {
		return false, x
	}
	if x == 2 {
		return true, x
	}
	if x%2 == 0 {
		return false, 2
	}
	for i := uint64(3); i*i <= x; i += 2 {
		if x%i == 0 {
			return false, i
		}
	}
	return true, 1
}
