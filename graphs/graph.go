package graphs

type Graph = map[int][]int

type OrGraph = Graph

func Reversed(g Graph) Graph {
	reversed := make(Graph, len(g))
	for v, list := range g {
		for _, u := range list {
			reversed[u] = append(reversed[u], v)
		}
	}
	return reversed
}
