package graphs

import (
	"reflect"
	"testing"
)

func TestBFS_Unordered(t *testing.T) {
	for _, tt := range []struct {
		name     string
		graph    Graph
		start    int
		wantDist map[int]int
	}{
		{
			name: "one vertex",
			graph: Graph{
				1: {},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
			},
		},
		{
			name: "bamboo 1",
			graph: Graph{
				1: {2},
				2: {1},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
			},
		},
		{
			name: "bamboo 2",
			graph: Graph{
				1: {2},
				2: {1, 3},
				3: {2},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 2,
			},
		},
		{

			name: "bamboo 3",
			graph: Graph{
				1: {2},
				2: {1, 3},
				3: {2, 4},
				4: {3},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 2,
				4: 3,
			},
		},
		{
			name: "triangle",
			graph: Graph{
				1: {2, 3},
				2: {1, 3},
				3: {1, 2},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 1,
			},
		},
		{
			name: "square",
			graph: Graph{
				1: {2, 3},
				2: {1, 4},
				3: {1, 4},
				4: {2, 3},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 1,
				4: 2,
			},
		},
		{
			name: "full square",
			graph: Graph{
				1: {2, 3, 4},
				2: {1, 3, 4},
				3: {1, 2, 4},
				4: {1, 2, 3},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 1,
				4: 1,
			},
		},
		{
			name: "pentagon",
			graph: Graph{
				1: {5, 2},
				2: {1, 3},
				3: {2, 4},
				4: {3, 5},
				5: {4, 1},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 2,
				4: 2,
				5: 1,
			},
		},
		{
			name: "full pentagon",
			graph: Graph{
				1: {2, 3, 4, 5},
				2: {1, 3, 4, 5},
				3: {1, 2, 4, 5},
				4: {1, 2, 3, 5},
				5: {1, 2, 3, 4},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 1,
				4: 1,
				5: 1,
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if gotDist, _ := BFS(tt.graph, tt.start); !reflect.DeepEqual(gotDist, tt.wantDist) {
				t.Errorf("BFS() got = %v, want %v", gotDist, tt.wantDist)
			}
		})
	}
}

func TestBFS_Oriented(t *testing.T) {
	for _, tt := range []struct {
		name     string
		graph    OrGraph
		start    int
		wantDist map[int]int
	}{
		{
			name: "one vertex",
			graph: OrGraph{
				1: {},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
			},
		},
		{
			name: "bamboo 1",
			graph: OrGraph{
				1: {2},
				2: {},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
			},
		},
		{
			name: "bamboo 1 backward last edge",
			graph: OrGraph{
				1: {},
				2: {1},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
			},
		},
		{
			name: "bamboo 2",
			graph: OrGraph{
				1: {2},
				2: {3},
				3: {},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 2,
			},
		},
		{
			name: "bamboo 2 backward last edge",
			graph: OrGraph{
				1: {2},
				2: {},
				3: {2},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
			},
		},
		{
			name: "bamboo 3 backward last edge",
			graph: OrGraph{
				1: {2},
				2: {3},
				3: {},
				4: {3},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 2,
			},
		},
		{
			name: "triangle",
			graph: OrGraph{
				1: {2},
				2: {3},
				3: {1},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 2,
			},
		},
		{
			name: "square",
			graph: OrGraph{
				1: {2},
				2: {3},
				3: {4},
				4: {1},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 2,
				4: 3,
			},
		},
		{
			name: "pentagon",
			graph: OrGraph{
				1: {2},
				2: {3},
				3: {4},
				4: {5},
				5: {1},
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 1,
				3: 2,
				4: 3,
				5: 4,
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if gotDist, _ := BFS(tt.graph, tt.start); !reflect.DeepEqual(gotDist, tt.wantDist) {
				t.Errorf("BFS() got = %v, want %v", gotDist, tt.wantDist)
			}
		})
	}
}
