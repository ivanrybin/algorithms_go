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
