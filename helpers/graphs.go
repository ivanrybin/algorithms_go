package helpers

import "reflect"

func CCEqual(got [][]int, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}
	gotM := SliceSlicesToSliceMaps(got)
	wantM := SliceSlicesToSliceMaps(want)
	for _, w := range wantM {
		ok := false
		for _, g := range gotM {
			if reflect.DeepEqual(w, g) {
				ok = true
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
