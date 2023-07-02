package graphs

import (
	"container/heap"

	"github.com/ivanrybin/algorithms_go/helpers"
)

// Dijkstra time depends on priority queue implementation and density of a graph.
//
// Operations:
//
//	|V| * heap.Pop  - main cycle
//	|V| * heap.Push - inner cycle, initialization of priority queue
//	|E| * heap.Fix  - inner cycle, paths relaxation
//
//	Priority queue on array: |V| * |V|    	    + |V| * C        + |E| * C        = O(V^2)
//	Binary priority queue:   |V| * log|V| 	    + |V| * log|V|   + |E| * log|V|   = O((V + E) * log V)
//	K-priority queue:        |V| * k * log_k|V| + |V| * log_k|V| + |E| * log_k|V| = O((Vk + E) * log_kV)
func Dijkstra(graph Graph, weights map[[2]int]int, s int) (map[int]int, map[int]int) {
	dist := make(map[int]int, len(graph))
	prev := make(map[int]int, len(graph))
	type dst struct {
		v, d int
	}
	// map for access priority queue elements
	known := make(map[int]*helpers.Elem[dst], len(graph))
	queue := helpers.NewPriorityQueue[dst](func(l, r dst) bool { return l.d < r.d })
	dist[s], prev[s], known[s] = 0, -1, &helpers.Elem[dst]{Value: dst{v: s, d: 0}}
	heap.Push(queue, known[s])
	for queue.Len() > 0 {
		e := heap.Pop(queue).(*helpers.Elem[dst])
		for _, u := range graph[e.Value.v] {
			vu := [2]int{e.Value.v, u}
			// vertex is visited for the first time
			if du, ok := dist[u]; !ok {
				dist[u] = dist[e.Value.v] + weights[vu]
				prev[u] = e.Value.v
				known[u] = &helpers.Elem[dst]{Value: dst{v: u, d: dist[u]}}
				heap.Push(queue, known[u])
			} else if dvu := dist[e.Value.v] + weights[vu]; du > dvu {
				dist[u] = dvu
				prev[u] = e.Value.v
				known[u].Value.d = dvu
				heap.Fix(queue, known[u].Value.v)
			}
		}
	}
	return dist, prev
}
