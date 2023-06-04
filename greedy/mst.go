package greedy

import (
	"container/heap"
	"math"
	"sort"

	"github.com/ivanrybin/algorithms_go/data_structures/disjoint_set"
	"github.com/ivanrybin/algorithms_go/helpers"
)

type edge struct {
	v, u int
}

// MSTPrim O((V+E) * log(V)).
func MSTPrim(graph map[int][]int, weights map[edge]int) []edge {
	if len(graph) == 0 {
		return nil
	}
	n := len(graph)
	type weight struct {
		v, w int
	}
	// init priority queue
	// v: 0  -> w: 0
	// v: 1+ -> w: +inf
	prev, known := make([]int, n), map[int]*helpers.Elem[weight]{}
	queue := helpers.NewPriorityQueue[weight](func(l, r weight) bool { return l.w < r.w })
	prev[0], known[0] = -1, &helpers.Elem[weight]{Value: weight{v: 0, w: 0}}
	heap.Push(queue, known[0])
	for v := 1; v < n; v++ {
		prev[v], known[v] = -1, &helpers.Elem[weight]{Value: weight{v: v, w: math.MaxInt}}
		heap.Push(queue, known[v])
	}
	mst := make([]edge, 0, len(graph))
	for queue.Len() > 0 {
		// v: 0, 1, 2, ..., n-1
		pe := heap.Pop(queue).(*helpers.Elem[weight])
		for _, u := range graph[pe.Value.v] {
			e := edge{v: pe.Value.v, u: u}
			if elem, ok := known[u]; ok && weights[e] < elem.Value.w {
				elem.Value.w = weights[e]
				prev[u] = pe.Value.v
				heap.Fix(queue, elem.Index())
			}
		}
		// MST gathering
		if pe.Value.v != 0 {
			mst = append(mst, edge{v: prev[pe.Value.v], u: pe.Value.v})
		}
	}
	return mst
}

// MSTKruskal depends on edges sort algorithm:
// - any linear sort - O(E * α(n)) because of DisjointSet with path compression and rank heuristic
// - sort.Slice - O(E * log(E)) + O(E * α(n))
func MSTKruskal(graph map[int][]int, weights map[edge]int) []edge {
	n := len(graph)
	if n == 0 {
		return nil
	}
	// init disjoint-set
	set := disjoint_set.NewDisjointSetLogStar(n)
	for v := 0; v < n; v++ {
		set.MakeSet(v)
	}
	// sort edges by weight ascending
	edges := make([]edge, 0, len(weights))
	for e := range weights {
		edges = append(edges, e)
	}
	sort.Slice(edges, func(i, j int) bool {
		return weights[edges[i]] < weights[edges[j]]
	})
	// build MST
	mst := make([]edge, 0)
	for _, e := range edges {
		if set.Find(e.u) != set.Find(e.v) {
			set.Union(e.u, e.v)
			mst = append(mst, e)
		}
	}
	return mst
}
