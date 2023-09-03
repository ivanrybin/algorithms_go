package strings

import "fmt"

// KMP Knuth–Morris–Pratt algorithm finds entry of p in s for O(|p+s|) /  O(|p+s|) (mem / time).
var KMP = FindByPrefixFunction

// FindByPrefixFunction finds entry of p in s for O(|p+s|) /  O(|p+s|) (mem / time).
func FindByPrefixFunction(s, p string) int {
	pi := PrefixFunction(fmt.Sprintf("%s#%s", p, s))
	for i := len(p) + 1; i < len(pi); i++ {
		if pi[i] == len(p) {
			return i - len(p) - pi[i]
		}
	}
	return -1
}

// FindByZFunction finds entry of p in s for O(|p+s|) /  O(|p+s|) (mem / time).
func FindByZFunction(s, p string) int {
	z := ZFunction(fmt.Sprintf("%s#%s", p, s))
	for i := len(p) + 1; i < len(z); i++ {
		if z[i] == len(p) {
			return i - len(p) - 1
		}
	}
	return -1
}
