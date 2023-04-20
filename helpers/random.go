package helpers

import "math/rand"

func RandomInts(maxSize int, maxValue int) []int {
	size := rand.Intn(maxSize)
	xs := make([]int, 0, size)
	for i := 0; i < size; i++ {
		xs = append(xs, rand.Intn(maxValue))
	}
	return xs
}

func RandomPairsIntInt(maxSize int, maxValue int) []Pair[int, int] {
	size := rand.Intn(maxSize)
	xs := make([]Pair[int, int], 0, size)
	for i := 0; i < size; i++ {
		xs = append(xs, Pair[int, int]{
			L: rand.Intn(maxValue),
			R: rand.Intn(maxValue),
		})
	}
	return xs
}
