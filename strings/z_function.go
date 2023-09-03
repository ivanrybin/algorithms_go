package strings

import "github.com/ivanrybin/algorithms_go/helpers"

// ZFunctionTrivial O(n) / O(n^2) (mem / time).
func ZFunctionTrivial(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i := 0; i < n; i++ {
		for i+z[i] < n && s[i+z[i]] == s[z[i]] {
			z[i]++
		}
	}
	return z
}

// ZFunction O(n) / O(n) (mem / time).
func ZFunction(s string) []int {
	l, r, n := 0, 0, len(s)
	z := make([]int, n)
	if n != 0 {
		z[0] = n
	}
	for i := 1; i < n; i++ {
		// current prefix is inside the last prefix
		// [..l..i..r..]
		if i <= r {
			z[i] = helpers.MinInt(z[i-l], r-i+1)
		}
		// found match with prefix
		for i+z[i] < n && s[i+z[i]] == s[z[i]] {
			z[i]++
		}
		// current prefix is outside the last prefix
		// [..l..r..i..i+z[i]-1..]
		// [..l..i..r..i+z[i]-1..]
		if r < i+z[i]-1 {
			l, r = i, i+z[i]-1
		}
	}
	return z
}
