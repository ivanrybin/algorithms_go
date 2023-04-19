package heap

var MinIntComparator = func(x, y int) bool { return x < y }

var MaxIntComparator = func(x, y int) bool { return x > y }

type BinaryHeap[T any] struct {
	xs   []T
	comp func(x, y T) bool
}

func NewBinaryHeap[T any](xs []T, comp func(x, y T) bool) *BinaryHeap[T] {
	h := &BinaryHeap[T]{
		xs:   xs,
		comp: comp,
	}
	h.build()
	return h
}

func (h *BinaryHeap[T]) Top() T {
	return h.xs[0]
}

func (h *BinaryHeap[T]) ExtractTop() T {
	x := h.xs[0]
	h.swap(0, h.Size()-1)
	h.xs = h.xs[:h.Size()-1]
	h.siftDown(0)
	return x
}

func (h *BinaryHeap[T]) Insert(x T) {
	h.xs = append(h.xs, x)
	h.siftUp(h.Size() - 1)
}

func (h *BinaryHeap[T]) Size() int {
	return len(h.xs)
}

func (h *BinaryHeap[T]) build() {
	if h.Size() <= 1 {
		return
	}
	for i := len(h.xs) / 2; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *BinaryHeap[T]) siftDown(i int) {
	next := i
	if h.left(i) < h.Size() && h.comp(h.xs[h.left(i)], h.xs[i]) {
		next = h.left(i)
	}
	if h.right(i) < h.Size() && h.comp(h.xs[h.right(i)], h.xs[next]) {
		next = h.right(i)
	}
	if i < next {
		h.swap(i, next)
		h.siftDown(next)
	}
}

func (h *BinaryHeap[T]) siftUp(i int) {
	for p := h.parent(i); p >= 0 && h.comp(h.xs[i], h.xs[p]); p = h.parent(p) {
		h.swap(i, p)
		i = p
	}
}

func (h *BinaryHeap[T]) swap(i, j int) {
	h.xs[i], h.xs[j] = h.xs[j], h.xs[i]
}

func (h *BinaryHeap[T]) left(i int) int {
	return 2*i + 1
}

func (h *BinaryHeap[T]) right(i int) int {
	return 2*i + 2
}

func (h *BinaryHeap[T]) parent(i int) int {
	return (i - 1) / 2
}
