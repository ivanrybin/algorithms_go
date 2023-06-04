package dynamic

import hs "github.com/ivanrybin/algorithms_go/helpers"

// EditDistance O(|s1| * |s2|) / O(|s1| * |s2|) (mem / time).
// !!! could be improved in memory to O(|s2|).
func EditDistance(s1, s2 string) int {
	r1, r2 := []rune(s1), []rune(s2)
	ed := make([][]int, 0, len(r1)+1)
	for i := 0; i < len(r1)+1; i++ {
		ed = append(ed, make([]int, len(r2)+1))
	}
	// i - corresponds to s1 (r1)
	// j - corresponds to s2 (r2)
	// E[i, j] = E(s1[:i], s2[:j])
	// E[i, 0] = i - insertions
	// E[0, j] = j - deletions
	// E[i, j] = min {
	//		    	E[i-1, j-1] + [s1[i] != s2[j]] // substitution
	//				E[i-1, j] + 1 				   // deletion
	//				E[i, j-1] + 1 				   // insertion
	//			 }
	for i := 0; i <= len(r1); i++ {
		ed[i][0] = i
	}
	for j := 0; j <= len(r2); j++ {
		ed[0][j] = j
	}
	for i := 1; i <= len(r1); i++ {
		for j := 1; j <= len(r2); j++ {
			diff := 0
			if r1[i-1] != r2[j-1] { // i-1 and j-1 because ed was built with +1 in indexing
				diff = 1
			}
			ed[i][j] = hs.MinInts(
				ed[i-1][j-1]+diff,
				ed[i-1][j]+1,
				ed[i][j-1]+1,
			)
		}
	}
	return ed[len(r1)][len(r2)]
}

// EditDistanceOptimized O(min{|s1| * |s2|}) / O(|s1| * |s2|) (mem / time).
func EditDistanceOptimized(s1, s2 string) int {
	r1, r2 := []rune(s1), []rune(s2)
	// i - corresponds to s1 (r1)
	// j - corresponds to s2 (r2)
	// E[i, j] = E(s1[:i], s2[:j])
	// E[i, 0] = i - insertions
	// E[0, j] = j - deletions
	// E[i, j] = min {
	//		    	E[i-1, j-1] + [s1[i] != s2[j]] // substitution
	//				E[i-1, j] + 1 				   // deletion
	//				E[i, j-1] + 1 				   // insertion
	//			 }
	// swap for memory reduce
	if len(r1) > len(r2) {
		r1, r2 = r2, r1
	}
	ed := [][]int{
		make([]int, len(r1)+1), // corresponds to E[:,j-1]
		make([]int, len(r1)+1), // corresponds to E[:,j]
	}
	for i := 0; i <= len(r1); i++ {
		ed[0][i] = i
	}
	for j := 1; j <= len(r2); j++ {
		ed[1][0] = j
		for i := 1; i <= len(r1); i++ {
			diff := 0
			if r1[i-1] != r2[j-1] { // i-1 and j-1 because edCurr was built with +1 in indexing
				diff = 1
			}
			ed[1][i] = hs.MinInts(
				ed[0][i-1]+diff,
				ed[1][i-1]+1,
				ed[0][i]+1,
			)
		}
		ed[0], ed[1] = ed[1], ed[0]
	}
	return ed[0][len(r1)]
}
