package disjoint_set

// DisjointSetLinear naive implementation.
//   - MakeSet O(1)
//   - Union O(n) - worst case
//   - Find O(n) - worst case
type DisjointSetLinear struct {
	parent []int
}

func NewDisjointSetLinear(n int) *DisjointSetLinear {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	return &DisjointSetLinear{
		parent: parent,
	}
}

// MakeSet O(1).
func (d *DisjointSetLinear) MakeSet(x int) {
	if d.parent[x] == -1 {
		d.parent[x] = x
	}
}

// Union O(n).
func (d *DisjointSetLinear) Union(x, y int) {
	if d.parent[x] == -1 {
		d.MakeSet(x)
	}
	if d.parent[y] == -1 {
		d.MakeSet(y)
	}
	x, y = d.Find(x), d.Find(y)
	d.parent[x] = y
}

// Find O(n).
func (d *DisjointSetLinear) Find(x int) int {
	for d.parent[x] != x && d.parent[x] != -1 {
		x = d.parent[x]
	}
	return x
}
