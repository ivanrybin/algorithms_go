package graphs

import "math"

// FloydWarshall O(V^3) / O(V^2) (time/mem).
// * map operations with dist and prev aren't counted.
func FloydWarshall(graph Graph, weights map[[2]int]int) (map[int]map[int]int, map[int]map[int]int, bool) {
	dist := make(map[int]map[int]int, len(graph))
	prev := make(map[int]map[int]int, len(graph))
	// initialization
	for u, list := range graph {
		dist[u] = make(map[int]int, len(graph))
		prev[u] = make(map[int]int, len(graph))
		for v := range graph {
			dist[u][v] = math.MaxInt
			prev[u][v] = -1
		}
		dist[u][u] = 0
		for _, v := range list {
			dist[u][v] = weights[[2]int{u, v}]
			prev[u][v] = u
		}
	}
	// algorithm
	for k := range graph {
		for u := range graph {
			for v := range graph {
				if dist[u][k] != math.MaxInt && dist[k][v] != math.MaxInt && dist[u][v] > dist[u][k]+dist[k][v] {
					dist[u][v] = dist[u][k] + dist[k][v]
					prev[u][v] = prev[k][v]
				}
			}
		}
	}
	// negative cycle check
	for v := range graph {
		if dist[v][v] < 0 {
			return nil, nil, true
		}
	}
	return dist, prev, false
}
