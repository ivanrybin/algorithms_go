package helpers

func UniqSequentialInts(n uint) []int {
	xs := make([]int, 0, n)
	for i := 0; i < int(n); i++ {
		xs = append(xs, i)
	}
	return xs
}
