package helpers

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AbsInt(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func MinInts(x int, xs ...int) int {
	min := x
	for _, n := range xs {
		if n < min {
			min = n
		}
	}
	return min
}

func MinInArray(xs []int) int {
	if len(xs) == 0 {
		return -1
	}
	pos := 0
	for i := 0; i < len(xs); i++ {
		if xs[i] < xs[pos] {
			pos = i
		}
	}
	return pos
}
