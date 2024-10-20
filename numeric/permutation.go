package numeric

import (
	"reflect"
	"slices"
)

// AllPermutations returns all possible permutations for the provided array.
// FYI: can be done better if you calculate # of permutations instead of comparing current permutation with the original one to stop the loop.
func AllPermutations(xs []int) [][]int {
	perms := [][]int{append([]int{}, xs...)}
	for curr := NextPermutation(xs); !reflect.DeepEqual(perms[0], curr); curr = NextPermutation(curr) {
		perms = append(perms, append([]int{}, curr...))
	}
	return perms
}

// NextPermutation (O(1) mem, O(n) time) in place performs next permutation for arrays i.e. the next permutation in ascending order of all sorted permutations of that array.
func NextPermutation(xs []int) []int {
	isNonAscending := true
	for i := 1; i < len(xs); i++ {
		if xs[i-1] < xs[i] {
			isNonAscending = false
			break
		}
	}
	if isNonAscending {
		slices.Reverse(xs)
		return xs
	}
	// find a breakpoint of a non-descending sequence starting from the end of the array
	for i := len(xs) - 2; i >= 0; i-- {
		if xs[i] < xs[i+1] {
			m := i + 1
			for j := len(xs) - 1; j >= i+1; j-- {
				if xs[j] > xs[i] {
					m = j
					break
				}
			}
			xs[i], xs[m] = xs[m], xs[i]
			// xs[i+1:] is a non-descending sequence i.e. we have to reverse it
			slices.Reverse(xs[i+1:])
			return xs
		}
	}
	return xs
}
