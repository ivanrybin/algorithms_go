package graphs

import (
	"cmp"
	"slices"
)

// SCC O(V+E)/O(V+E) (time/mem).
func SCC(g Graph) map[int][]int {
	// 1. calculate exit timestamps of vertices in G
	type vertex struct {
		v int
		t int
	}
	timestamp, exit := 0, make([]vertex, len(g))
	pre := func(v int, g Graph) {
		timestamp++
	}
	post := func(v int, g Graph) {
		exit[v] = vertex{v: v, t: timestamp}
		timestamp++
	}
	DFS(g, pre, post)
	// 2. transpose G and get G^T
	trans := Transpose(g)
	// 3. run DFS on G^T in the reverse order of exit timestamps of vertices of G
	slices.SortFunc(exit, func(a, b vertex) int {
		return cmp.Compare(a.t, b.t) * -1
	})
	scc, sccN := map[int][]int{}, 0
	visited := map[int]struct{}{}
	preTrans := func(v int, g Graph) {
		scc[sccN] = append(scc[sccN], v)
	}
	for _, v := range exit {
		if _, ok := visited[v.v]; !ok {
			dfs(v.v, trans, visited, preTrans, nil)
			sccN++
		}
	}
	return scc
}
