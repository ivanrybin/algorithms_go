package graphs

type Graph = map[int][]int

type OrGraph = Graph

// Transpose reverses all edges of oriented graph O(V+E).
func Transpose(g Graph) Graph {
	trans := make(Graph, len(g))
	for v, neighbours := range g {
		for _, u := range neighbours {
			trans[u] = append(trans[u], v)
		}
	}
	return trans
}
