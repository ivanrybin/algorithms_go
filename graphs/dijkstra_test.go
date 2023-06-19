package graphs

import (
	"reflect"
	"testing"
)

func TestDijkstra_Oriented_Weight_One(t *testing.T) {
	for _, tt := range []struct {
		name     string
		graph    OrGraph
		weights  map[[2]int]int
		start    int
		wantDist map[int]int
	}{
		{
			name: "one vertex",
			graph: OrGraph{
				1: {},
			},
			weights: map[[2]int]int{},
			start:   1,
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
			weights: map[[2]int]int{
				[2]int{1, 2}: 42,
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: 42,
			},
		},
		{
			name: "bamboo 1 backward last edge",
			graph: OrGraph{
				1: {},
				2: {1},
			},
			weights: map[[2]int]int{
				[2]int{2, 1}: 42,
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
			weights: map[[2]int]int{
				[2]int{1, 2}: 1,
				[2]int{2, 3}: 1,
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
			weights: map[[2]int]int{
				[2]int{1, 2}: 1,
				[2]int{3, 2}: 42,
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
			weights: map[[2]int]int{
				[2]int{1, 2}: 1,
				[2]int{2, 3}: 1,
				[2]int{4, 3}: 42,
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
			weights: map[[2]int]int{
				[2]int{1, 2}: 1,
				[2]int{2, 3}: 1,
				[2]int{3, 1}: 1,
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
			weights: map[[2]int]int{
				[2]int{1, 2}: 1,
				[2]int{2, 3}: 1,
				[2]int{3, 4}: 1,
				[2]int{4, 1}: 1,
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
			weights: map[[2]int]int{
				[2]int{1, 2}: 1,
				[2]int{2, 3}: 1,
				[2]int{3, 4}: 1,
				[2]int{4, 5}: 1,
				[2]int{5, 1}: 1,
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
			if gotDist, _ := Dijkstra(tt.graph, tt.weights, tt.start); !reflect.DeepEqual(gotDist, tt.wantDist) {
				t.Errorf("Dijkstra() got = %v, want %v", gotDist, tt.wantDist)
			}
		})
	}
}
