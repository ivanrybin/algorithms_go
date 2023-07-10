package search

type AVLTree[T comparable] struct {
	root *AVLNode[T]
	less func(l, r T) bool
	eq   func(l, r T) bool
}

// RotateSmallLeft
//
//			x					y
//
//		y 		(γ)	  --> 	(α) 	x
//
//	(α)	   (β)				  	(β)    (γ)
func RotateSmallLeft[T comparable](x *AVLNode[T]) *AVLNode[T] {
	y := x.l
	x.l, x.p, y.r, y.p = y.r, y, x, x.p
	x.UpdateHeight()
	y.UpdateHeight()
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
	y := x.r
	x.p, x.r, y.p, y.l = y, y.l, x.p, x
	x.UpdateHeight()
	y.UpdateHeight()
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
	y := x.l
	RotateSmallRight(y)
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
	y := x.r
	RotateSmallLeft(y)
	return RotateSmallRight(x)
}
