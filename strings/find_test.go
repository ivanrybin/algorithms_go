package strings

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	for _, tt := range findPatternTestCases {
		t.Run(fmt.Sprintf("%s %s", tt.s, tt.p), func(t *testing.T) {
			if got := KMP(tt.s, tt.p); got != tt.want {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindByPrefixFunction(t *testing.T) {
	for _, tt := range findPatternTestCases {
		t.Run(fmt.Sprintf("%s %s", tt.s, tt.p), func(t *testing.T) {
			if got := FindByPrefixFunction(tt.s, tt.p); got != tt.want {
				t.Errorf("FindByPrefixFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindByZFunction(t *testing.T) {
	for _, tt := range findPatternTestCases {
		t.Run(fmt.Sprintf("%s %s", tt.s, tt.p), func(t *testing.T) {
			if got := FindByZFunction(tt.s, tt.p); got != tt.want {
				t.Errorf("FindByZFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

var findPatternTestCases = []struct {
	s    string
	p    string
	want int
}{
	{
		s:    "a",
		p:    "a",
		want: 0,
	},
	{
		s:    "a",
		p:    "b",
		want: -1,
	},
	{
		s:    "ab",
		p:    "a",
		want: 0,
	},
	{
		s:    "ab",
		p:    "b",
		want: 1,
	},
	{
		s:    "ba",
		p:    "a",
		want: 1,
	},
	{
		s:    "bba",
		p:    "a",
		want: 2,
	},
	{
		s:    "bbba",
		p:    "a",
		want: 3,
	},
	{
		s:    "bbbba",
		p:    "a",
		want: 4,
	},
	{
		s:    "bbbbb",
		p:    "b",
		want: 0,
	},
	{
		s:    "papapatternpapapa",
		p:    "pattern",
		want: 4,
	},
	{
		s:    "ppapatpattpattepatterpatternpatternpattern",
		p:    "pattern",
		want: 21,
	},
}
