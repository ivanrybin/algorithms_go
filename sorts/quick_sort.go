package sorts

import "math/rand"

// QuickSort
//
//	O(n * log n) best-case
//	O(n * log n) on average
//	O(n^2) worst-case
func QuickSort(xs []int) []int {
	quickSort(xs, 0, len(xs)-1)
	return xs
}

// quickSort sort [l, r] part of xs.
func quickSort(xs []int, l, r int) {
	if l >= r {
		return
	}
	p := pivot(l, r)
	m := partition(xs, l, r, p)
	quickSort(xs, l, m-1)
	quickSort(xs, m+1, r)
}

// pivot randomly chooses pivot index.
func pivot(l, r int) int {
	return l + rand.Intn(r-l+1)
}

// partition by pivot xs[p] in [l, r] part of xs.
func partition(xs []int, l, r, p int) int {
	x := xs[p]
	xs[p], xs[r] = xs[r], xs[p]
	m := l
	for i := l; i < r; i++ {
		if xs[i] <= x {
			xs[m], xs[i] = xs[i], xs[m]
			m++
		}
	}
	xs[m], xs[r] = xs[r], xs[m]
	return m
}
