package dynamic

import "math"

// ChainMatrixMultiplication find the lowest cost of matrices multiplication by finding the best order.
func ChainMatrixMultiplication(dimensions []int64) int64 {
	if len(dimensions)%2 != 0 {
		panic("invalid input")
	}
	n := len(dimensions) / 2
	if n <= 1 {
		return 0
	}
	cost := make([][]int64, n+1)
	for i := 0; i < n+1; i++ {
		cost[i] = make([]int64, n+1)
		for j := 0; j < n+1; j++ {
			if i != j {
				cost[i][j] = math.MaxInt
			}
		}
	}
	for step := 1; step < n; step++ {
		for i := 1; i+step <= n; i++ {
			j := i + step
			for k := i; k < j; k++ {
				mult := dimensions[2*i-2] * dimensions[2*k-1] * dimensions[2*j-1]
				cost[i][j] = min(cost[i][j], cost[i][k]+cost[k+1][j]+mult)
			}
		}
	}
	return cost[1][n]
}
