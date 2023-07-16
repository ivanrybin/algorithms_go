package helpers

func LessInt(a, b int) bool {
	return a < b
}

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

func GenSegments(l, r int) [][]int {
	segments := make([][]int, 0, (r-l)*(r-l)/2)
	for i := l; i <= r; i++ {
		for j := i; j <= r; j++ {
			segments = append(segments, []int{i, j})
		}
	}
	return segments
}

func FindIdempotentOnSegment[T any](xs []T, l, r int, f func(l, r T) T) T {
	v := xs[l]
	for _, x := range xs[l+1 : r+1] {
		v = f(v, x)
	}
	return v
}

func FindMaxOnSequentialSegment[T any](xs []T, _, r int, _ func(l, r T) T) T {
	return xs[r]
}

func FindMinOnSequentialSegment[T any](xs []T, l, _ int, _ func(l, r T) T) T {
	return xs[l]
}
