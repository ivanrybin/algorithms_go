package graphs

func DFS(g Graph, pre, post func(...any)) {
	visited := make(map[int]struct{}, len(g))
	for v := range g {
		if _, ok := visited[v]; !ok {
			exploreDFS(v, g, visited, pre, post)
		}
	}
}

func exploreDFS(v int, g Graph, visited map[int]struct{}, pre, post func(...any)) {
	visited[v] = struct{}{}
	if pre != nil {
		pre(v, g, visited)
	}
	for _, u := range g[v] {
		if _, ok := visited[u]; !ok {
			exploreDFS(u, g, visited, pre, post)
		}
	}
	if post != nil {
		post(v, g, visited)
	}
}
