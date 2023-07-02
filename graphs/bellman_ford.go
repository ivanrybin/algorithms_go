package graphs

import "math"

// BellmanFord O(V*E) / O(V) (time/mem), works with negative edges.
func BellmanFord(graph Graph, weights map[[2]int]int, start int) (map[int]int, map[int]int, bool) {
	dist := make(map[int]int, len(graph))
	prev := make(map[int]int, len(graph))
	for v := range graph {
		dist[v] = math.MaxInt
	}
	dist[start] = 0
	// |V| - 1 iterations: each edge relaxation
	for i := 0; i+1 < len(graph); i++ {
		for u, list := range graph {
			for _, v := range list {
				w := weights[[2]int{u, v}]
				if dist[u] != math.MaxInt && dist[v] > dist[u]+w {
					dist[v] = dist[u] + w
					prev[v] = u
				}
			}
		}
	}
	// |V|th iteration: check for negative cycle
	for u, list := range graph {
		for _, v := range list {
			if dist[u] != math.MaxInt && dist[v] > dist[u]+weights[[2]int{u, v}] {
				return nil, nil, true
			}
		}
	}
	return dist, prev, false
}
