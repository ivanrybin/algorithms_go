package helpers

func IncreasingInts(n uint) []int {
	xs := make([]int, 0, n)
	for i := 0; i < int(n); i++ {
		xs = append(xs, i)
	}
	return xs
}

func DecreasingInts(n uint) []int {
	xs := make([]int, 0, n)
	for i := int(n) - 1; i >= 0; i-- {
		xs = append(xs, i)
	}
	return xs
}
