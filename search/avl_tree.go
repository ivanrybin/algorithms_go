package search

import "fmt"

type AVLTree[T comparable] struct {
	size int
	root *AVLNode[T]
	less func(l, r T) bool
}

// NewAVLTree creates AVL tree.
func NewAVLTree[T comparable](less func(l, r T) bool) *AVLTree[T] {
	return &AVLTree[T]{
		size: 0,
		root: nil,
		less: less,
	}
}

// BuildAVLTree builds AVL tree based on provided values and comparison function.
func BuildAVLTree[T comparable](xs []T, less func(l, r T) bool) *AVLTree[T] {
	t := NewAVLTree(less)
	for _, x := range xs {
		t.Insert(x)
	}
	return t
}

// Insert value to a tree.
func (t *AVLTree[T]) Insert(x T) *AVLNode[T] {
	if t.root == nil {
		t.root = &AVLNode[T]{v: x, h: 1}
		t.size++
		return t.root
	}
	n := t.root.Insert(x, t.less)
	t.fix(n)
	t.size++
	return n
}

// Delete value from a tree if it's present.
func (t *AVLTree[T]) Delete(x T) bool {
	n := t.root.Find(x, t.less)
	if n == nil || n.Value() != x {
		return false
	}
	n = n.Delete()
	t.fix(n)
	t.size--
	return true
}

// Find the first node with this value otherwise nil.
func (t *AVLTree[T]) Find(x T) *AVLNode[T] {
	if t.root == nil {
		return nil
	}
	if n := t.root.Find(x, t.less); n.Value() == x {
		return n
	}
	return nil
}

// Size of a tree.
func (t *AVLTree[T]) Size() int {
	return t.size
}

// Less comparison function.
func (t *AVLTree[T]) Less() func(l, r T) bool {
	return t.less
}

// Height of a tree.
func (t *AVLTree[T]) Height() int {
	if t.root == nil {
		return 0
	}
	return t.root.Height()
}

// fix a tree with rotations starting from specified node along the path to the root.
func (t *AVLTree[T]) fix(n *AVLNode[T]) {
	for ; n != nil && n.Parent() != nil; n = n.Parent() {
		p := n.Parent()
		if p.Left() == n {
			p.l = FixNode(p.Left())
		} else if p.Right() == n {
			p.r = FixNode(p.Right())
		}
		p.UpdateHeight()
	}
	if n == nil {
		t.root = nil
	} else {
		t.root = FixNode(n)
		t.root.UpdateHeight()
	}
}

// FixNode subtree of a specified node with AVL rotations.
func FixNode[T comparable](n *AVLNode[T]) *AVLNode[T] {
	switch rt := n.DetectRotationType(); rt {
	case NoRotation:
		return n
	case SmallLeft:
		return RotateSmallLeft(n)
	case SmallRight:
		return RotateSmallRight(n)
	case BigLeft:
		return RotateBigLeft(n)
	case BigRight:
		return RotateBigRight(n)
	default:
		panic(fmt.Sprintf("unknown rotation type: %v", rt))
	}
}

// RotateSmallLeft
//
//			x					y
//
//		y 		(γ)	  --> 	(α) 	x
//
//	(α)	   (β)				  	(β)    (γ)
func RotateSmallLeft[T comparable](x *AVLNode[T]) *AVLNode[T] {
	y := x.Left()
	x.l, x.p, y.p, y.r = y.Right(), y, x.Parent(), x
	if x.Left() != nil {
		x.Left().p = x
	}
	x.UpdateHeight()
	y.UpdateHeight()
	y.Parent().UpdateHeight()
	return y
}

// RotateSmallRight
//
//		x                     y
//
//	(α) 	y	   -->     x     (γ)
//
//		(β)    (γ)     (α)   (β)
func RotateSmallRight[T comparable](x *AVLNode[T]) *AVLNode[T] {
	y := x.Right()
	x.p, x.r, y.p, y.l = y, y.Left(), x.Parent(), x
	if x.Right() != nil {
		x.Right().p = x
	}
	x.UpdateHeight()
	y.UpdateHeight()
	y.Parent().UpdateHeight()
	return y
}

// RotateBigLeft
//
//			x					     x                     z
//
//		y 		(δ)	  --> 	     z      (δ)  -->     y           x
//
//	(α)	    z			     y     (γ)           (α)   (β)   (γ)    (δ)
//
//	  	(β)   (γ)        (α)   (β)
func RotateBigLeft[T comparable](x *AVLNode[T]) *AVLNode[T] {
	x.l = RotateSmallRight(x.Left())
	x.Left().p = x
	x.UpdateHeight()
	return RotateSmallLeft(x)
}

// RotateBigRight
//
//			x				     x                         z
//
//	  (α) 		y	   --> 	 (α)     z       -->     y           x
//
//		     z	  (δ)		      (β)    y       (α)   (β)   (γ)    (δ)
//
//		(β)   (γ)                    (γ)   (δ)
func RotateBigRight[T comparable](x *AVLNode[T]) *AVLNode[T] {
	x.r = RotateSmallLeft(x.Right())
	x.Right().p = x
	x.UpdateHeight()
	return RotateSmallRight(x)
}
