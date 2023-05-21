package disjoint_set

// DisjointSetLogStar faster implementation.
//   - MakeSet O(1)
//   - Union amortized O(α(n))
//   - Find amortized O(α(n))
type DisjointSetLogStar struct {
	parent []int
	rank   []int
}

func NewDisjointSetLogStar(n int) *DisjointSetLogStar {
	parent := make([]int, 0, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	return &DisjointSetLogStar{
		parent: parent,
		rank:   make([]int, 0, n),
	}
}

// MakeSet O(1).
func (d *DisjointSetLogStar) MakeSet(x int) {
	if d.parent[x] == -1 {
		d.parent[x] = x
	}
}

// Union amortized O(α(n)) because of path compression in Find and rank heuristic in Union
func (d *DisjointSetLogStar) Union(x, y int) {
	if d.parent[x] == -1 {
		d.MakeSet(x)
	}
	if d.parent[y] == -1 {
		d.MakeSet(y)
	}
	x, y = d.Find(x), d.Find(y)
	switch {
	case x == y:
		return
	case d.rank[x] < d.rank[y]:
		d.parent[x] = y
	case d.rank[x] > d.rank[y]:
		d.parent[y] = x
	default:
		d.parent[x] = y
		d.rank[y]++
	}
}

// Find amortized O(α(n)) because of path compression in Find and rank heuristic in Union.
func (d *DisjointSetLogStar) Find(x int) int {
	for d.parent[x] != x && d.parent[x] != -1 {
		d.parent[x] = d.parent[d.parent[x]]
	}
	return d.parent[x]
}
