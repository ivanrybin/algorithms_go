package dynamic

import (
	"fmt"
	"testing"
)

func editDistanceTestCases() []struct {
	s1, s2 string
	want   int
} {
	return []struct {
		s1, s2 string
		want   int
	}{
		{
			s1:   "", // eq
			s2:   "",
			want: 0,
		},
		{
			s1:   "", // ins
			s2:   "abc",
			want: 3,
		},
		{
			s1:   "abc", // del
			s2:   "",
			want: 3,
		},
		{
			s1:   "abc", // eq
			s2:   "abc",
			want: 0,
		},
		{
			s1:   "ab", // sub
			s2:   "ax",
			want: 1,
		},
		{
			s1:   "ab", // del
			s2:   "a",
			want: 1,
		},
		{
			s1:   "ab", // ins
			s2:   "abx",
			want: 1,
		},
		{
			s1:   "ab", // sub
			s2:   "xb",
			want: 1,
		},
		{
			s1:   "xab", // del
			s2:   "ab",
			want: 1,
		},
		{
			s1:   "abab",
			s2:   "baba",
			want: 2,
		},
		{
			s2:   "sunny", // sunny -> s_nny (del) -> s_noy (sub) -> s_nowy (ins) -> snowy
			s1:   "snowy",
			want: 3,
		},
		{
			s1:   "exponential",
			s2:   "polynomial",
			want: 6,
		},
		{
			s1:   "aaaaaaaaaaaaaaaaaa",
			s2:   "bbbbbbbbbbbbbbbbbb",
			want: 18,
		},
	}
}

func TestEditDistance(t *testing.T) {
	for _, tt := range editDistanceTestCases() {
		tt := tt
		t.Run(fmt.Sprintf("(%v,%v)_%v", tt.s1, tt.s2, tt.want), func(t *testing.T) {
			if got := EditDistance(tt.s1, tt.s2); got != tt.want {
				t.Errorf("got=%v != want=%v", got, tt.want)
			}
		})
	}
}

func TestEditDistanceOptimized(t *testing.T) {
	for _, tt := range editDistanceTestCases() {
		tt := tt
		t.Run(fmt.Sprintf("(%v,%v)_%v", tt.s1, tt.s2, tt.want), func(t *testing.T) {
			if got := EditDistanceOptimized(tt.s1, tt.s2); got != tt.want {
				t.Errorf("got=%v != want=%v", got, tt.want)
			}
		})
	}
}
