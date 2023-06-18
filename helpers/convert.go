package helpers

func MapSlicesToSliceSlices(m map[int][]int) [][]int {
	ss := make([][]int, 0, len(m))
	for _, s := range m {
		ss = append(ss, s)
	}
	return ss
}

func SliceSlicesToSliceMaps(ss [][]int) []map[int]struct{} {
	sm := make([]map[int]struct{}, len(ss))
	for _, s := range ss {
		m := make(map[int]struct{}, len(s))
		for _, v := range s {
			m[v] = struct{}{}
		}
		sm = append(sm, m)
	}
	return sm
}
