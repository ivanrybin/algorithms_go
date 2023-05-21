package disjoint_set

// DisjointSetLog faster implementation.
//   - MakeSet O(1)
//   - Union O(log(n)) - worst case
//   - Find O(log(n)) - worst case
type DisjointSetLog struct {
	parent []int
	rank   []int
}

func NewDisjointSetLog(n int) *DisjointSetLog {
	parent := make([]int, 0, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	return &DisjointSetLog{
		parent: parent,
		rank:   make([]int, 0, n),
	}
}

// MakeSet O(1).
func (d *DisjointSetLog) MakeSet(x int) {
	if d.parent[x] == -1 {
		d.parent[x] = x
	}
}

// Union O(log(n)) because of rank heuristic.
func (d *DisjointSetLog) Union(x, y int) {
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

// Find O(log(n)) because of rank heuristic in Union.
func (d *DisjointSetLog) Find(x int) int {
	for d.parent[x] != x && d.parent[x] != -1 {
		x = d.parent[x]
	}
	return x
}
