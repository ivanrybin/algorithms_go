package graphs

import (
	"math"
	"reflect"
	"testing"
)

func TestBellmanFord_Positive_Weight_One(t *testing.T) {
	for _, tt := range []struct {
		name              string
		graph             OrGraph
		weights           map[[2]int]int
		start             int
		wantDist          map[int]int
		wantNegativeCycle bool
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
				2: math.MaxInt,
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
				3: math.MaxInt,
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
				4: math.MaxInt,
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
			gotDist, _, gotCycle := BellmanFord(tt.graph, tt.weights, tt.start)
			if !reflect.DeepEqual(gotDist, tt.wantDist) {
				t.Errorf("BellmanFord() gotDist = %v, wantDist %v", gotDist, tt.wantDist)
			} else if gotCycle != tt.wantNegativeCycle {
				t.Errorf("BellmanFord() gotCycle = %v, wantNegativeCycle %v", gotCycle, tt.wantNegativeCycle)
			}
		})
	}
}

func TestBellmanFord_Negative_Weight_One(t *testing.T) {
	for _, tt := range []struct {
		name              string
		graph             OrGraph
		weights           map[[2]int]int
		start             int
		wantDist          map[int]int
		wantNegativeCycle bool
	}{
		{
			name: "triangle negative",
			graph: OrGraph{
				1: {2},
				2: {3},
				3: {1},
			},
			weights: map[[2]int]int{
				[2]int{1, 2}: -2,
				[2]int{2, 3}: 1,
				[2]int{3, 1}: 1,
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: -2,
				3: -1,
			},
		},
		{
			name: "triangle negative cycle",
			graph: OrGraph{
				1: {2},
				2: {3},
				3: {1},
			},
			weights: map[[2]int]int{
				[2]int{1, 2}: -3,
				[2]int{2, 3}: 1,
				[2]int{3, 1}: 1,
			},
			start:             1,
			wantDist:          nil,
			wantNegativeCycle: true,
		},
		{
			name: "square negative",
			graph: OrGraph{
				1: {2},
				2: {3},
				3: {4},
				4: {1},
			},
			weights: map[[2]int]int{
				[2]int{1, 2}: -3,
				[2]int{2, 3}: 1,
				[2]int{3, 4}: 1,
				[2]int{4, 1}: 1,
			},
			start: 1,
			wantDist: map[int]int{
				1: 0,
				2: -3,
				3: -2,
				4: -1,
			},
		},
		{
			name: "square negative cycle",
			graph: OrGraph{
				1: {2},
				2: {3},
				3: {4},
				4: {1},
			},
			weights: map[[2]int]int{
				[2]int{1, 2}: -4,
				[2]int{2, 3}: 1,
				[2]int{3, 4}: 1,
				[2]int{4, 1}: 1,
			},
			start:             1,
			wantDist:          nil,
			wantNegativeCycle: true,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotDist, _, gotCycle := BellmanFord(tt.graph, tt.weights, tt.start)
			if !reflect.DeepEqual(gotDist, tt.wantDist) {
				t.Errorf("BellmanFord() gotDist = %v, wantDist %v", gotDist, tt.wantDist)
			} else if gotCycle != tt.wantNegativeCycle {
				t.Errorf("BellmanFord() gotCycle = %v, wantNegativeCycle %v", gotCycle, tt.wantNegativeCycle)
			}
		})
	}
}
