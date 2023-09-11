package helpers

func MinInt[T IntegerType](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MinInts[T IntegerType](x T, xs ...T) T {
	m := x
	for _, n := range xs {
		if n < m {
			m = n
		}
	}
	return m
}

func MaxInt[T IntegerType](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxInts[T IntegerType](x T, xs ...T) T {
	m := x
	for _, n := range xs {
		if n > m {
			m = n
		}
	}
	return m
}

func AbsInt[T Int](a T) T {
	if a < 0 {
		return -1 * a
	}
	return a
}

func LessInt[T IntegerType](a, b T) bool {
	return a < b
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
