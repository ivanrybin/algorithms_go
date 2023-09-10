package strings

import (
	"sort"

	hs "github.com/ivanrybin/algorithms_go/helpers"
)

// SuffixArraySlowAngUgly O(n * log^2).
func SuffixArraySlowAngUgly(s string) []int {
	s += "$"
	n := len(s)
	sa, c := make([]int, n), make([]int, n)
	// k = 0
	{
		a := make([]hs.Pair[uint8, int], n)
		for i := 0; i < n; i++ {
			a[i] = hs.Pair[uint8, int]{L: s[i], R: i}
		}
		sort.SliceStable(a, func(i, j int) bool {
			return a[i].L < a[j].L
		})
		for i := 0; i < n; i++ {
			sa[i] = a[i].R
		}

		c[sa[0]] = 0
		for i := 1; i < n; i++ {
			if a[i].L == a[i-1].L {
				c[sa[i]] = c[sa[i-1]]
			} else {
				c[sa[i]] = c[sa[i-1]] + 1
			}
		}
	}
	type IIP = hs.Pair[int, int]
	type PI = hs.Pair[IIP, int]
	for k := 0; (1 << k) < n; k++ {
		a := make([]PI, n)
		for i := 0; i < n; i++ {
			a[i] = PI{L: IIP{L: c[i], R: c[(i+(1<<k))%n]}, R: i}
		}
		sort.SliceStable(a, func(i, j int) bool {
			if a[i].L.L == a[j].L.L {
				return a[i].L.R < a[j].L.R
			}
			return a[i].L.L < a[j].L.L
		})
		for i := 0; i < n; i++ {
			sa[i] = a[i].R
		}

		c[sa[0]] = 0
		for i := 1; i < n; i++ {
			if a[i].L == a[i-1].L {
				c[sa[i]] = c[sa[i-1]]
			} else {
				c[sa[i]] = c[sa[i-1]] + 1
			}
		}
	}
	return sa
}
