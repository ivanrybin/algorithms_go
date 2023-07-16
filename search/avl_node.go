package search

import (
	"fmt"

	"github.com/ivanrybin/algorithms_go/helpers"
)

type AVLNode[T comparable] struct {
	h    int
	v    T
	p    *AVLNode[T]
	l, r *AVLNode[T]
}

// Min [WASN'T TESTED].
func (n *AVLNode[T]) Min() *AVLNode[T] {
	v := n
	for v != nil && v.IsLeft() {
		v = v.Left()
	}
	return v
}

// Max [WASN'T TESTED].
func (n *AVLNode[T]) Max() *AVLNode[T] {
	v := n
	for v != nil && v.IsRight() {
		v = v.Right()
	}
	return v
}

// Succ finds successor of a node by value [WASN'T TESTED].
func (n *AVLNode[T]) Succ() *AVLNode[T] {
	if n.IsRight() {
		return n.Right().Min()
	}
	v := n
	for v.Parent() != nil && v.Parent().Left() != v {
		v = v.Parent()
	}
	return v.Parent()
}

// Pred finds predecessor of a node by value [WASN'T TESTED].
func (n *AVLNode[T]) Pred() *AVLNode[T] {
	if n.IsLeft() {
		return n.Left().Max()
	}
	v := n
	for v.Parent() != nil && v.Parent().Right() != v {
		v = v.Parent()
	}
	return v.Parent()
}

// Find finds first exact match or a node which can be a parent for provided value [TESTED].
func (n *AVLNode[T]) Find(x T, less func(l, r T) bool) *AVLNode[T] {
	if n == nil {
		return nil
	}
	curr := n
	for curr.Value() != x {
		if less(x, curr.Value()) {
			if curr.Left() == nil {
				return curr
			}
			curr = curr.Left()
		} else {
			if curr.Right() == nil {
				return curr
			}
			curr = curr.Right()
		}
	}
	return curr
}

// Insert value to a subtree of a node [WASN'T TESTED].
func (n *AVLNode[T]) Insert(v T, less func(l, r T) bool) *AVLNode[T] {
	curr := n
	for {
		if less(v, curr.Value()) {
			if curr.Left() == nil {
				return curr.Attach(v, less)
			}
			curr = curr.Left()
		} else {
			if curr.Right() == nil {
				return curr.Attach(v, less)
			}
			curr = curr.Right()
		}
	}
}

// Attach child to this node [TESTED].
func (n *AVLNode[T]) Attach(v T, less func(l, r T) bool) *AVLNode[T] {
	node := &AVLNode[T]{h: 1, v: v, p: n}
	if less(v, n.Value()) {
		if n.Left() != nil {
			panic("broken invariant")
		}
		n.l = node
	} else {
		if n.Right() != nil {
			panic("broken invariant")
		}
		n.r = node
	}
	n.UpdateHeight()
	return node
}

// Delete node and get parent node of deleted node [WASN'T TESTED].
func (n *AVLNode[T]) Delete() *AVLNode[T] {
	switch {
	case n.IsLeaf():
		return n.deleteAsLeaf()
	case n.IsFull():
		return n.deleteAsFull()
	case n.IsLeft():
		return n.deleteAsWithOnlyLeftChild()
	case n.IsRight():
		return n.deleteAsWithOnlyRightChild()
	default:
		panic("broken invariant")
	}
}

// deleteAsLeaf deletes node as a leaf.
func (n *AVLNode[T]) deleteAsLeaf() *AVLNode[T] {
	n.Detach()
	n.Parent().UpdateHeight()
	return n.Parent()
}

// deleteAsLeaf deletes node as it has both children by finding predecessor of a node, swapping values and deleting predecessor.
func (n *AVLNode[T]) deleteAsFull() *AVLNode[T] {
	pred := n.Pred()
	pred.v, n.v = n.Value(), pred.Value()
	switch {
	case pred.IsLeaf():
		return pred.deleteAsLeaf()
	case pred.IsFull():
		panic("broken invariant")
	case pred.IsLeft():
		return pred.deleteAsWithOnlyLeftChild()
	case pred.IsRight():
		return pred.deleteAsWithOnlyRightChild()
	default:
		panic("broken invariant")
	}
}

// deleteAsWithOnlyLeftChild deletes node by squeezing as it has only the left child.
func (n *AVLNode[T]) deleteAsWithOnlyLeftChild() *AVLNode[T] {
	if n.Parent() != nil {
		if n.Parent().l == n {
			n.Parent().l = n.Left()
		} else if n.Parent().r == n {
			n.Parent().r = n.Left()
		}
	}
	n.Left().p = n.Parent()
	n.Parent().UpdateHeight()
	return n.Left()
}

// deleteAsWithOnlyRightChild deletes node by squeezing as it has only the right child.
func (n *AVLNode[T]) deleteAsWithOnlyRightChild() *AVLNode[T] {
	if n.Parent() != nil {
		if n.Parent().Left() == n {
			n.Parent().l = n.Right()
		} else if n.Parent().Right() == n {
			n.Parent().r = n.Right()
		}
	}
	n.Right().p = n.Parent()
	n.Parent().UpdateHeight()
	return n.Right()
}

// Detach node from a parent [WASN'T TESTED].
func (n *AVLNode[T]) Detach() {
	if n.Parent() != nil {
		if n.Parent().Left() == n {
			n.Parent().l = nil
		}
		if n.Parent().Right() == n {
			n.Parent().r = nil
		}
		n.Parent().UpdateHeight()
	}
}

// InorderTraverse collects nodes in non-decreasing.
func (n *AVLNode[T]) InorderTraverse() []*AVLNode[T] {
	if n == nil {
		return []*AVLNode[T]{}
	}
	return append(append(n.Left().InorderTraverse(), n), n.Right().InorderTraverse()...)
}

func (n *AVLNode[T]) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("v=%v, h=%v l=[%v] r=[%v]", n.Value(), n.Height(), n.Left().String(), n.Right().String())
}

// Value of a node.
func (n *AVLNode[T]) Value() T {
	return n.v
}

// Left child of a node.
func (n *AVLNode[T]) Left() *AVLNode[T] {
	return n.l
}

// Right child of a node.
func (n *AVLNode[T]) Right() *AVLNode[T] {
	return n.r
}

// Parent of a node.
func (n *AVLNode[T]) Parent() *AVLNode[T] {
	return n.p
}

// IsLeaf [TESTED].
func (n *AVLNode[T]) IsLeaf() bool {
	return n.Left() == nil && n.Right() == nil
}

// IsLeft child present [TESTED].
func (n *AVLNode[T]) IsLeft() bool {
	return n.Left() != nil
}

// IsRight [TESTED].
func (n *AVLNode[T]) IsRight() bool {
	return n.Right() != nil
}

// IsFull [TESTED].
func (n *AVLNode[T]) IsFull() bool {
	return n.Left() != nil && n.Right() != nil
}

// Height [TESTED].
func (n *AVLNode[T]) Height() int {
	if n == nil {
		return 0
	}
	return n.h
}

// HeightDiff returns heights difference of children of a node.
func (n *AVLNode[T]) HeightDiff() int {
	if n == nil {
		return 0
	}
	return helpers.AbsInt(n.Left().Height() - n.Right().Height())
}

// UpdateHeight [WASN'T TESTED].
func (n *AVLNode[T]) UpdateHeight() {
	if n != nil {
		n.h = helpers.MaxInt(n.Left().Height(), n.Right().Height()) + 1
	}
}

// IsBroken [TESTED].
func (n *AVLNode[T]) IsBroken() bool {
	if n == nil {
		return false
	}
	return helpers.AbsInt(n.Left().Height()-n.Right().Height()) > 1
}

// AVLRotationType type of the AVL tree rotation according to the tree definition.
type AVLRotationType = int

const (
	NoRotation AVLRotationType = iota
	SmallLeft
	SmallRight
	BigLeft
	BigRight
)

// DetectRotationType according to the AVL tree definition [TESTED].
func (n *AVLNode[T]) DetectRotationType() AVLRotationType {
	switch {
	case !n.IsBroken():
		return NoRotation
	case n.Left().Height() > n.Right().Height():
		if n.Left().Left().Height() >= n.Left().Right().Height() {
			return SmallLeft
		}
		return BigLeft
	case n.Left().Height() < n.Right().Height():
		if n.Right().Left().Height() <= n.Right().Right().Height() {
			return SmallRight
		}
		return BigRight
	default:
		panic("broken invariant")
	}
}
