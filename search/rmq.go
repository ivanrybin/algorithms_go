package search

import (
	"fmt"
	"math"

	"github.com/ivanrybin/algorithms_go/helpers"
)

type RMQ[T any] struct {
	f  func(l, r T) T
	n  int
	vs []T
}

// NewRMQ builds RMQ or any other RMQ-like-idempotent structure.
//
// Array RMQ is of a size of the closest full binary tree with size bigger than provided.
//
//	len(xs) -> 2 * 2 ** ceil(log_2(len(xs))) - 1
//	2 -> 2
//	3 -> 4
//	4 -> 4
//	5 -> 8
//	6 -> 8
//	 ...
func NewRMQ[T any](xs []T, f func(l, r T) T) *RMQ[T] {
	size := 2*int(math.Pow(2.0, math.Ceil(math.Log2(float64(len(xs)))))) - 1
	rmq := make([]T, size)
	buildRMQ(rmq, xs, f, 0, 0, len(xs)-1)
	return &RMQ[T]{
		f:  f,
		n:  len(xs),
		vs: rmq,
	}
}

func buildRMQ[T any](rmq, xs []T, f func(l, r T) T, i, l, r int) {
	if l == r {
		rmq[i] = xs[l]
		return
	}
	ll, lr := leftSegment(l, r)
	rl, rr := rightSegment(l, r)
	buildRMQ[T](rmq, xs, f, lIdx(i), ll, lr)
	buildRMQ[T](rmq, xs, f, rIdx(i), rl, rr)
	if lIdx(i) < len(rmq) && rIdx(i) < len(rmq) {
		rmq[i] = f(rmq[lIdx(i)], rmq[rIdx(i)])
	}
}

func (rmq *RMQ[T]) F() func(l, r T) T {
	return rmq.f
}

func (rmq *RMQ[T]) Range() (int, int) {
	return 0, rmq.n - 1
}

func (rmq *RMQ[T]) Get(sl, sr int) T {
	il, ir, ok := intersect(0, rmq.n-1, sl, sr)
	if !ok {
		panic(fmt.Sprintf("no intersection of RMQ range [%v:%v] with [%v:%v]", 0, rmq.n, sl, sr))
	}
	return rmq.get(0, 0, rmq.n-1, il, ir)
}

func (rmq *RMQ[T]) get(i, l, r, sl, sr int) T {
	if l == sl && r == sr {
		return rmq.vs[i]
	}
	ll, lr := leftSegment(l, r)
	rl, rr := rightSegment(l, r)
	ill, ilr, okLeft := intersect(ll, lr, sl, sr)
	irl, irr, okRight := intersect(rl, rr, sl, sr)
	switch {
	case okLeft && okRight:
		return rmq.f(rmq.get(lIdx(i), ll, lr, ill, ilr), rmq.get(rIdx(i), rl, rr, irl, irr))
	case okLeft:
		return rmq.get(lIdx(i), ll, lr, ill, ilr)
	case okRight:
		return rmq.get(rIdx(i), rl, rr, irl, irr)
	default:
		panic("impossible")
	}
}

func (rmq *RMQ[T]) Update(j int, v T) {
	if !belongs(0, rmq.n-1, j) {
		panic(fmt.Sprintf("%v doesn't belong to RMQ range [%v:%v]", j, 0, rmq.n))
	}
	rmq.update(0, 0, rmq.n-1, j, v)
}

func (rmq *RMQ[T]) update(i, l, r, j int, v T) {
	for {
		if l == j && r == j {
			rmq.vs[i] = v
			break
		}
		ll, lr := leftSegment(l, r)
		rl, rr := rightSegment(l, r)
		switch {
		case belongs(ll, lr, j):
			i = lIdx(i)
			l, r = ll, lr
		case belongs(rl, rr, j):
			i = rIdx(i)
			l, r = rl, rr
		default:
			panic("impossible")
		}
	}
	for i = pIdx(i); i != -1; i = pIdx(i) {
		rmq.vs[i] = rmq.f(rmq.lV(i), rmq.rV(i))
	}
}

func (rmq *RMQ[T]) lV(i int) T {
	return rmq.vs[lIdx(i)]
}

func (rmq *RMQ[T]) rV(i int) T {
	return rmq.vs[rIdx(i)]
}

func lIdx(i int) int {
	return 2*i + 1
}

func rIdx(i int) int {
	return 2*i + 2
}

func pIdx(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func leftSegment(l, r int) (int, int) {
	m := (r - l + 1) / 2
	return l, l + m - 1
}

func rightSegment(l, r int) (int, int) {
	m := (r - l + 1) / 2
	return l + m, r
}

func intersect(l, r, sl, sr int) (int, int, bool) {
	if sl > r || sr < l {
		return 0, 0, false
	}
	return helpers.MaxInt(l, sl), helpers.MinInt(r, sr), true
}

func belongs(l, r, i int) bool {
	return i >= l && i <= r
}
