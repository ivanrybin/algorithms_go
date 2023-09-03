package strings

// PrefixFunctionTrivial O(n) / O(n^3) (mem / time).
func PrefixFunctionTrivial(s string) []int {
	n := len(s)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		for k := 1; k <= i; k++ {
			if s[0:k] == s[i-k+1:i+1] {
				pi[i] = k
			}
		}
	}
	return pi
}

// PrefixFunction O(n) / O(n) (mem / time).
func PrefixFunction(s string) []int {
	n := len(s)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		k := pi[i-1] // # of matched characters
		for k > 0 && s[k] != s[i] {
			k = pi[k-1]
		}
		if s[k] == s[i] {
			k++
		}
		pi[i] = k
	}
	return pi
}
