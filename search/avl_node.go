package search

import "github.com/ivanrybin/algorithms_go/helpers"

type AVLNode[T comparable] struct {
	h    int
	v    T
	p    *AVLNode[T]
	l, r *AVLNode[T]
}

// IsLeaf [TESTED].
func (n *AVLNode[T]) IsLeaf() bool {
	return n.l == nil && n.r == nil
}

// IsLeft [TESTED].
func (n *AVLNode[T]) IsLeft() bool {
	return n.l != nil
}

// IsRight [TESTED].
func (n *AVLNode[T]) IsRight() bool {
	return n.r != nil
}

// IsFull [TESTED].
func (n *AVLNode[T]) IsFull() bool {
	return n.l != nil && n.r != nil
}

// Height [TESTED].
func (n *AVLNode[T]) Height() int {
	if n == nil {
		return 0
	}
	return n.h
}

func (n *AVLNode[T]) UpdateHeight() {
	if n != nil {
		n.h = helpers.MaxInt(n.l.Height(), n.r.Height()) + 1
	}
}

// IsBroken [TESTED].
func (n *AVLNode[T]) IsBroken() bool {
	if n == nil {
		return false
	}
	return helpers.AbsInt(n.l.Height()-n.r.Height()) > 1
}

// Find finds first exact match or a node which can be a parent for provided value.
// [TESTED].
func (n *AVLNode[T]) Find(v T, less func(l, r T) bool) *AVLNode[T] {
	if n == nil {
		return nil
	}
	curr := n
	for curr.v != v {
		if less(v, curr.v) {
			if curr.l == nil {
				return curr
			}
			curr = curr.l
		} else {
			if curr.r == nil {
				return curr
			}
			curr = curr.r
		}
	}
	return curr
}

// Insert [TESTED].
func (n *AVLNode[T]) Insert(v T, less func(l, r T) bool) *AVLNode[T] {
	node := &AVLNode[T]{h: 1, v: v, p: n}
	if less(v, n.v) {
		if n.l != nil {
			panic("broken invariant")
		}
		n.l = node
	} else {
		if n.r != nil {
			panic("broken invariant")
		}
		n.r = node
	}
	return node
}

type RotationType = int

const (
	NoRotation RotationType = iota
	SmallLeft
	SmallRight
	BigLeft
	BigRight
)

// DetectRotationType [TESTED].
func (n *AVLNode[T]) DetectRotationType() RotationType {
	switch {
	case !n.IsBroken():
		return NoRotation
	case n.l.Height() > n.r.Height():
		if n.l.l.Height() >= n.l.r.Height() {
			return SmallLeft
		}
		return BigLeft
	case n.l.Height() < n.r.Height():
		if n.r.l.Height() <= n.r.r.Height() {
			return SmallRight
		}
		return BigRight
	default:
		panic("broken invariant")
	}
}

func (n *AVLNode[T]) Detach() {
	if n.p != nil {
		if n.p.l == n {
			n.p.l = nil
		} else if n.p.r == n {
			n.p.r = nil
		}
	}
}
