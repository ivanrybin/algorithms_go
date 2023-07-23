package helpers

import (
	"math/rand"
	"time"
)

func Shuffle[T any](xs []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(xs), func(i, j int) { xs[i], xs[j] = xs[j], xs[i] })
}

func RandomInts(maxSize int, maxValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := r.Intn(maxSize)
	xs := make([]int, 0, size)
	for i := 0; i < size; i++ {
		xs = append(xs, r.Intn(maxValue))
	}
	return xs
}

func RandomIntsExactSize(size int, maxValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	xs := make([]int, 0, size)
	for i := 0; i < size; i++ {
		xs = append(xs, r.Intn(maxValue))
	}
	return xs
}

func RandomPairsIntInt(maxSize int, maxValue int) []Pair[int, int] {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := r.Intn(maxSize)
	xs := make([]Pair[int, int], 0, size)
	for i := 0; i < size; i++ {
		xs = append(xs, Pair[int, int]{
			L: r.Intn(maxValue),
			R: r.Intn(maxValue),
		})
	}
	return xs
}
