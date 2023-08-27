package strings

import (
	"fmt"
	"testing"
)

func TestFindByZFunction(t *testing.T) {
	for _, tt := range []struct {
		s    string
		p    string
		want int
	}{
		{s: "a", p: "a", want: 0},
		{s: "a", p: "b", want: -1},
		{s: "ab", p: "b", want: 1},
		{s: "papapatternpapapa", p: "pattern", want: 4},
	} {
		t.Run(fmt.Sprintf("%s %s", tt.s, tt.p), func(t *testing.T) {
			if got := FindByZFunction(tt.s, tt.p); got != tt.want {
				t.Errorf("FindByZFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
