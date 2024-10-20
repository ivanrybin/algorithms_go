package graphs

// CC O(V+E) / O(V) (time / mem).
// * operations with map visited aren't counted.
func CC(g Graph) map[int][]int {
	visited := make(map[int]struct{}, len(g))
	cc, ccN := make(map[int][]int, len(g)), 0
	pre := func(v int, _ Graph) {
		cc[ccN] = append(cc[ccN], v)
	}
	for v := range g {
		if _, ok := visited[v]; !ok {
			dfs(v, g, visited, pre, nil)
			ccN++
		}
	}
	return cc
}
