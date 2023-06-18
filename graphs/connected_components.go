package graphs

func CC(g Graph) map[int][]int {
	visited := make(map[int]struct{}, len(g))
	cc, ccN := make(map[int][]int, len(g)), 0
	for v := range g {
		if _, ok := visited[v]; !ok {
			exploreCC(v, g, visited, cc, ccN)
			ccN++
		}
	}
	return cc
}

func exploreCC(v int, g Graph, visited map[int]struct{}, cc map[int][]int, ccN int) {
	visited[v] = struct{}{}
	cc[ccN] = append(cc[ccN], v)
	for _, u := range g[v] {
		if _, ok := visited[u]; !ok {
			exploreCC(u, g, visited, cc, ccN)
		}
	}
}
