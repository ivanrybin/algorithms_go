package search

import "math/rand"

// OrderStatistic
//
//	O(n * log n) best-case
//	O(n * log n) on average
//	O(n^2) worst-case
func OrderStatistic(xs []int, k int) int {
	if len(xs) == 0 || k <= 0 || k > len(xs) {
		panic("wrong arguments")
	}
	k = k - 1
	for l, r := partition3(rand.Intn(len(xs)), xs, 0, len(xs)-1); l > k || r < k; {
		if l > k {
			l, r = partition3(rand.Intn(l), xs, 0, l-1)
		} else if r < k {
			l, r = partition3(r+1+rand.Intn(len(xs)-r-1), xs, r+1, len(xs)-1)
		}
	}
	return xs[k]
}

// partition3 splits xs in place on [l, r] in: [ <x ][ == x][ >x ] where x = xs[p] on the function start.
func partition3(p int, xs []int, l, r int) (int, int) {
	x := xs[p]
	xs[p], xs[r] = xs[r], xs[p]
	// left border [x,...]
	lx := l
	for i := l; i < r; i++ {
		if xs[i] < x {
			xs[lx], xs[i] = xs[i], xs[lx]
			lx++
		}
	}
	xs[lx], xs[r] = xs[r], xs[lx]
	// right border of  [...,x]
	rx := lx
	for i := rx; i <= r; i++ {
		if xs[i] == x {
			xs[i], xs[rx] = xs[rx], xs[i]
			rx++
		}
	}
	return lx, rx - 1
}
