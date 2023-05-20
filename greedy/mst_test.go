package greedy

import (
	"fmt"
	"testing"
)

func TestMSTPrim(t *testing.T) {
	for _, tt := range mstTestCases() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// duplicate weights by edges inverse
			weights := duplicateInverseEdges(tt.weights)
			mst := MSTPrim(tt.graph, weights)
			if got := sumMST(mst, weights); got != tt.wantSum {
				t.Errorf("MST sum: got=%v != want=%v", got, tt.wantSum)
			}
			if !checkEdges(mst, tt.wantMST) {
				t.Errorf("MST edges: got=%v != want=%v", mst, tt.wantMST)
			}
		})
	}
}

func TestMSTKruskal(t *testing.T) {
	for _, tt := range mstTestCases() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mst := MSTKruskal(tt.graph, tt.weights)
			if gotSum := sumMST(mst, tt.weights); gotSum != tt.wantSum {
				t.Errorf("MST sum: gotSum=%v != want=%v", gotSum, tt.wantSum)
			}
		})
	}
}

type mstTest struct {
	name    string
	graph   map[int][]int
	weights map[edge]int
	wantMST []edge
	wantSum int
}

func mstTestCases() []mstTest {
	return []mstTest{
		{
			name:    "empty",
			wantSum: 0,
		},
		{
			name: "one edge 42 | sum = 42",
			graph: map[int][]int{
				0: {1},
				1: {0},
			},
			weights: map[edge]int{
				edge{0, 1}: 42,
			},
			wantSum: 42,
			wantMST: []edge{
				{0, 1},
			},
		},
		{
			name: "tree 1-1-1-1 | sum = 42",
			graph: map[int][]int{
				0: {1},
				1: {2},
				2: {3},
				3: {4},
				4: {},
			},
			weights: map[edge]int{
				{0, 1}: 10,
				{1, 2}: 10,
				{2, 3}: 10,
				{3, 4}: 12,
			},
			wantSum: 42,
			wantMST: []edge{
				{0, 1},
				{1, 2},
				{2, 3},
				{3, 4},
			},
		},
		{
			name: "triangle 1-1-1 | sum = 2",
			graph: map[int][]int{
				0: {1, 2},
				1: {0, 2},
				2: {0, 1},
			},
			weights: map[edge]int{
				edge{0, 1}: 1,
				edge{1, 2}: 1,
				edge{2, 0}: 1,
			},
			wantSum: 2,
			wantMST: []edge{
				{0, 1},
				{0, 2},
			},
		},
		{
			name: "triangle 42-1-1 | sum = 2",
			graph: map[int][]int{
				0: {1, 2},
				1: {0, 2},
				2: {0, 1},
			},
			weights: map[edge]int{
				edge{0, 1}: 42,
				edge{1, 2}: 1,
				edge{2, 0}: 1,
			},
			wantSum: 2,
			wantMST: []edge{
				{1, 2},
				{2, 0},
			},
		},
		{
			name: "triangle 1-42-1 | sum = 2",
			graph: map[int][]int{
				0: {1, 2},
				1: {0, 2},
				2: {0, 1},
			},
			weights: map[edge]int{
				edge{0, 1}: 1,
				edge{1, 2}: 42,
				edge{2, 0}: 1,
			},
			wantSum: 2,
			wantMST: []edge{
				{0, 1},
				{2, 0},
			},
		},
		{
			name: "triangle 1-1-42 | sum = 2",
			graph: map[int][]int{
				0: {1, 2},
				1: {0, 2},
				2: {0, 1},
			},
			weights: map[edge]int{
				edge{0, 1}: 1,
				edge{1, 2}: 1,
				edge{2, 0}: 42,
			},
			wantSum: 2,
			wantMST: []edge{
				{0, 1},
				{1, 2},
			},
		},
		{
			name: "square 1-1-1-1 | sum = 3",
			graph: map[int][]int{
				0: {1, 3},
				1: {0, 2},
				2: {1, 3},
				3: {0, 1},
			},
			weights: map[edge]int{
				edge{0, 1}: 1,
				edge{0, 3}: 1,
				edge{1, 2}: 1,
				edge{2, 3}: 1,
			},
			wantSum: 3,
			wantMST: []edge{
				{0, 1},
				{1, 2},
				{0, 3},
			},
		},
		{
			name: "full square 1-1-1-1-1-1 | sum = 3",
			graph: map[int][]int{
				0: {1, 2, 3},
				1: {0, 2, 3},
				2: {0, 1, 3},
				3: {0, 1, 2},
			},
			weights: map[edge]int{
				edge{0, 1}: 1,
				edge{0, 2}: 1,
				edge{0, 3}: 1,
				edge{1, 2}: 1,
				edge{1, 3}: 1,
				edge{2, 3}: 1,
			},
			wantSum: 3,
			wantMST: []edge{
				{0, 1},
				{0, 2},
				{0, 3},
			},
		},
		{
			name: "full square 3-3-3-1-1-1 | sum = 5",
			graph: map[int][]int{
				0: {1, 2, 3},
				1: {0, 2, 3},
				2: {0, 1, 3},
				3: {0, 1, 2},
			},
			weights: map[edge]int{
				edge{0, 1}: 3,
				edge{0, 2}: 3,
				edge{0, 3}: 3,
				edge{1, 2}: 1,
				edge{1, 3}: 1,
				edge{2, 3}: 1,
			},
			wantSum: 5,
			wantMST: []edge{
				{0, 1},
				{1, 2},
				{1, 3},
			},
		},
		{
			name: "full square 3-3-1-3-1-1 | sum = 3",
			graph: map[int][]int{
				0: {1, 2, 3},
				1: {0, 2, 3},
				2: {0, 1, 3},
				3: {0, 1, 2},
			},
			weights: map[edge]int{
				edge{0, 1}: 3,
				edge{0, 2}: 3,
				edge{0, 3}: 1,
				edge{1, 2}: 3,
				edge{1, 3}: 1,
				edge{2, 3}: 1,
			},
			wantSum: 3,
			wantMST: []edge{
				{0, 3},
				{1, 3},
				{2, 3},
			},
		},
	}
}

func duplicateInverseEdges(weights map[edge]int) map[edge]int {
	ws := make(map[edge]int, len(weights))
	for e, w := range weights {
		ws[e] = w
		ws[edge{v: e.u, u: e.v}] = w
	}
	return ws
}

func sumMST(mst []edge, weights map[edge]int) int {
	sum := 0
	for _, e := range mst {
		if w, ok := weights[e]; !ok {
			panic(fmt.Sprintf("weights[%+v] doesn't exist", e))
		} else {
			sum += w
		}
	}
	return sum
}

func checkEdges(got, want []edge) bool {
	if len(want) != len(got) {
		return false
	}
	wm := map[edge]struct{}{}
	for _, e := range want {
		wm[e] = struct{}{}
		wm[edge{v: e.u, u: e.v}] = struct{}{}
	}
	for _, e := range got {
		if _, ok := wm[e]; !ok {
			return false
		}
	}
	return true
}
