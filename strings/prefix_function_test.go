package strings

import (
	"reflect"
	"testing"
)

func TestPrefixFunctionTrivial(t *testing.T) {
	for _, tt := range prefixFunctionTestCases {
		t.Run(tt.s, func(t *testing.T) {
			if got := PrefixFunctionTrivial(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrefixFunctionTrivial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrefixFunction(t *testing.T) {
	for _, tt := range prefixFunctionTestCases {
		t.Run(tt.s, func(t *testing.T) {
			if got := PrefixFunction(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrefixFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

var prefixFunctionTestCases = []struct {
	s    string
	want []int
}{
	{
		s:    "",
		want: []int{},
	},
	{
		s:    "a",
		want: []int{0},
	},
	{
		s:    "aa",
		want: []int{0, 1},
	},
	{
		s:    "aaa",
		want: []int{0, 1, 2},
	},
	{
		s:    "aaaa",
		want: []int{0, 1, 2, 3},
	},
	{
		s:    "ababab",
		want: []int{0, 0, 1, 2, 3, 4},
	},
	{
		s:    "abacabab",
		want: []int{0, 0, 1, 0, 1, 2, 3, 2},
	},
	{
		s:    "abcabacababc",
		want: []int{0, 0, 0, 1, 2, 1, 0, 1, 2, 1, 2, 3},
	},
}
