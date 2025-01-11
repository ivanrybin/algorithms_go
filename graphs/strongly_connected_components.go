package graphs

import (
	"slices"
)

// SCC O(V+E)/O(V+E) (time/mem).
func SCC(g Graph) map[int][]int {
	// 1. topological sort of G to get vertices in the postorder
	postOrder := make([]int, 0, len(g))
	post := func(v int, g Graph) {
		postOrder = append(postOrder, v)
	}
	DFS(g, nil, post)
	// 2. transpose G to get G^T
	gT := Transpose(g)
	// 3. run DFS on G^T in the reverse order of postorder of G
	slices.Reverse(postOrder)
	scc, sccN := map[int][]int{}, 0
	visited := map[int]struct{}{}
	preGT := func(v int, g Graph) {
		scc[sccN] = append(scc[sccN], v)
	}
	for _, v := range postOrder {
		if _, ok := visited[v]; !ok {
			dfs(v, gT, visited, preGT, nil)
			sccN++
		}
	}
	return scc
}
