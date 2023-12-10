package strings

import (
	"reflect"
	"testing"
)

var suffixArrayTests = []struct {
	s    string
	want []int
}{
	{
		"",
		[]int{},
	},
	{
		"a",
		[]int{0},
	},
	{
		"ab",
		[]int{0, 1},
	},
	{
		"abc",
		[]int{0, 1, 2},
	},
	{
		"aaa",
		[]int{2, 1, 0},
	},
}

func TestSuffixArraySlowAngUgly(t *testing.T) {
	for _, tt := range suffixArrayTests {
		t.Run(tt.s, func(t *testing.T) {
			if got := SuffixArraySlowAngUgly(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SuffixArraySlowAngUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
