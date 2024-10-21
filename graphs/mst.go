package graphs

import (
	"container/heap"
	"math"
	"sort"

	"github.com/ivanrybin/algorithms_go/data_structures/disjoint_set"
	hs "github.com/ivanrybin/algorithms_go/helpers"
)

type edge struct {
	v, u int
}

// MSTPrim O((V+E) * log(V)).
func MSTPrim(g Graph, weights map[edge]int) []edge {
	if len(g) == 0 {
		return nil
	}
	n := len(g)
	type weight struct {
		v, w int
	}
	// init priority queue
	// v: 0  -> w: 0
	// v: 1+ -> w: +inf
	prev, elems := make([]int, n), map[int]*hs.Elem[weight]{}
	queue := hs.NewPriorityQueue[weight](func(l, r weight) bool { return l.w < r.w })
	prev[0], elems[0] = -1, &hs.Elem[weight]{Value: weight{v: 0, w: 0}}
	heap.Push(queue, elems[0])
	for v := 1; v < n; v++ {
		prev[v], elems[v] = -1, &hs.Elem[weight]{Value: weight{v: v, w: math.MaxInt}}
		heap.Push(queue, elems[v])
	}
	mst := make([]edge, 0, len(g))
	for queue.Len() > 0 {
		// v: 0, 1, 2, ..., n-1
		v := heap.Pop(queue).(*hs.Elem[weight])
		for _, u := range g[v.Value.v] {
			vu := edge{v: v.Value.v, u: u}
			if elem, ok := elems[u]; ok && weights[vu] < elem.Value.w {
				elem.Value.w = weights[vu]
				prev[u] = v.Value.v
				heap.Fix(queue, elem.Index())
			}
		}
		// MST gathering
		if v.Value.v != 0 {
			mst = append(mst, edge{v: prev[v.Value.v], u: v.Value.v})
		}
	}
	return mst
}

// MSTKruskal O(E * log(V)).
func MSTKruskal(g Graph, weights map[edge]int) []edge {
	n := len(g)
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
