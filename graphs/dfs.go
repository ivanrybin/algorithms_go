package graphs

// DFS O(V+E) / O(V) (time / mem).
// * operations with map visited aren't counted.
func DFS(g Graph, pre, post func(v int, g Graph)) {
	visited := make(map[int]struct{}, len(g))
	for v := range g {
		if _, ok := visited[v]; !ok {
			dfs(v, g, visited, pre, post)
		}
	}
}

func dfs(v int, g Graph, visited map[int]struct{}, pre, post func(v int, g Graph)) {
	visited[v] = struct{}{}
	if pre != nil {
		pre(v, g)
	}
	for _, u := range g[v] {
		if _, ok := visited[u]; !ok {
			dfs(u, g, visited, pre, post)
		}
	}
	if post != nil {
		post(v, g)
	}
}
