package helpers

import (
	"math/rand"
	"time"
)

func RandomInts(maxSize int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(maxSize)
	xs := make([]int, 0, size)
	for i := 0; i < size; i++ {
		xs = append(xs, rand.Intn(maxValue))
	}
	return xs
}
