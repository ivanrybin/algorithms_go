package dynamic

type item struct {
	v float64
	w int
}

// KnapsackWithRepetition O(nW) time, O(W) memory.
//
// K(w) = max { K(w - w_i) + v_i } for i: w_i <= w
func KnapsackWithRepetition(items []item, W int) float64 {
	knapsack := make([]float64, W+1)
	for w := 1; w <= W; w++ {
		for _, item := range items {
			if w >= item.w && knapsack[w] < knapsack[w-item.w]+item.v {
				knapsack[w] = knapsack[w-item.w] + item.v
			}
		}
	}
	return knapsack[W]
}

// KnapsackWithoutRepetition O(nW) time, O(nW) memory.
//
// K(w, i) = max { K(w - w_i, i - 1) + v_i, K(w, i-1) }
func KnapsackWithoutRepetition(items []item, W int) float64 {
	knapsack := make([][]float64, W+1)
	for w := 0; w < len(knapsack); w++ {
		knapsack[w] = make([]float64, len(items)+1)
	}
	for i := 1; i <= len(items); i++ {
		item := items[i-1]
		for w := 0; w <= W; w++ {
			if w < item.w {
				knapsack[w][i] = knapsack[w][i-1]
			} else {
				knapsack[w][i] = max(knapsack[w-item.w][i-1]+item.v, knapsack[w][i-1])
			}
		}
	}
	return knapsack[W][len(items)]
}
