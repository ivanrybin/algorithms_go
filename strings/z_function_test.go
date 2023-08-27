package strings

import (
	"reflect"
	"testing"
)

func TestZFunctionTrivial(t *testing.T) {
	for _, tt := range zFunctionTestCases {
		t.Run(tt.s, func(t *testing.T) {
			if got := ZFunctionTrivial(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZFunctionTrivial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZFunction(t *testing.T) {
	for _, tt := range zFunctionTestCases {
		t.Run(tt.s, func(t *testing.T) {
			if got := ZFunction(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

var zFunctionTestCases = []struct {
	s    string
	want []int
}{
	{
		s:    "",
		want: []int{},
	},
	{
		s:    "a",
		want: []int{1},
	},
	{
		s:    "aa",
		want: []int{2, 1},
	},
	{
		s:    "aaa",
		want: []int{3, 2, 1},
	},
	{
		s:    "aaaa",
		want: []int{4, 3, 2, 1},
	},
	{
		s:    "ababab",
		want: []int{6, 0, 4, 0, 2, 0},
	},
	{
		s:    "abacabab",
		want: []int{8, 0, 1, 0, 3, 0, 2, 0},
	},
}
