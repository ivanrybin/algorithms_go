package strings

import "fmt"

// FindByZFunction finds entry of p in s O(|p+s|) /  O(|p+s|) (mem / time).
func FindByZFunction(s, p string) int {
	z := ZFunction(fmt.Sprintf("%s#%s", p, s))
	for i := len(p) + 1; i < len(z); i++ {
		if z[i] == len(p) {
			return i - len(p) - 1
		}
	}
	return -1
}
