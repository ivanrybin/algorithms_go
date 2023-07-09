package search

type AVLTree[T comparable] struct {
	root *AVLNode[T]
	less func(l, r T) bool
	eq   func(l, r T) bool
}
