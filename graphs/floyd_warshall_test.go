package graphs

import (
	"math"
	"reflect"
	"testing"
)

func TestFloydWarshall_Positive_Weight_One(t *testing.T) {
	for _, tt := range []struct {
		name              string
		graph             OrGraph
		weights           map[[2]int]int
		wantDist          map[int]map[int]int
		wantNegativeCycle bool
	}{
		{
			name: "one vertex",
			graph: OrGraph{
				1: {},
			},
			weights: map[[2]int]int{},
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
				},
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: 42,
				},
				2: {
					1: math.MaxInt,
					2: 0,
				},
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: math.MaxInt,
				},
				2: {
					1: 42,
					2: 0,
				},
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: 1,
					3: 2,
				},
				2: {
					1: math.MaxInt,
					2: 0,
					3: 1,
				},
				3: {
					1: math.MaxInt,
					2: math.MaxInt,
					3: 0,
				},
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: 1,
					3: math.MaxInt,
				},
				2: {
					1: math.MaxInt,
					2: 0,
					3: math.MaxInt,
				},
				3: {
					1: math.MaxInt,
					2: 42,
					3: 0,
				},
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: 1,
					3: 2,
					4: math.MaxInt,
				},
				2: {
					1: math.MaxInt,
					2: 0,
					3: 1,
					4: math.MaxInt,
				},
				3: {
					1: math.MaxInt,
					2: math.MaxInt,
					3: 0,
					4: math.MaxInt,
				},
				4: {
					1: math.MaxInt,
					2: math.MaxInt,
					3: 42,
					4: 0,
				},
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: 1,
					3: 2,
				},
				2: {
					1: 2,
					2: 0,
					3: 1,
				},
				3: {
					1: 1,
					2: 2,
					3: 0,
				},
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: 1,
					3: 2,
					4: 3,
				},
				2: {
					1: 3,
					2: 0,
					3: 1,
					4: 2,
				},
				3: {
					1: 2,
					2: 3,
					3: 0,
					4: 1,
				},
				4: {
					1: 1,
					2: 2,
					3: 3,
					4: 0,
				},
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: 1,
					3: 2,
					4: 3,
					5: 4,
				},
				2: {
					1: 4,
					2: 0,
					3: 1,
					4: 2,
					5: 3,
				},
				3: {
					1: 3,
					2: 4,
					3: 0,
					4: 1,
					5: 2,
				},
				4: {
					1: 2,
					2: 3,
					3: 4,
					4: 0,
					5: 1,
				},
				5: {
					1: 1,
					2: 2,
					3: 3,
					4: 4,
					5: 0,
				},
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotDist, _, gotCycle := FloydWarshall(tt.graph, tt.weights)
			if !reflect.DeepEqual(gotDist, tt.wantDist) {
				t.Errorf("FloydWarshall()\n gotDist = %v \nwantDist = %v", gotDist, tt.wantDist)
			}
			if gotCycle != tt.wantNegativeCycle {
				t.Errorf("FloydWarshall() gotCycle = %v, wantNegativeCycle %v", gotCycle, tt.wantNegativeCycle)
			}
		})
	}
}

func TestFloydWarshall_Negative_Weight_One(t *testing.T) {
	for _, tt := range []struct {
		name              string
		graph             OrGraph
		weights           map[[2]int]int
		wantDist          map[int]map[int]int
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: -2,
					3: -1,
				},
				2: {
					1: 2,
					2: 0,
					3: 1,
				},
				3: {
					1: 1,
					2: -1,
					3: 0,
				},
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
			wantDist: map[int]map[int]int{
				1: {
					1: 0,
					2: -3,
					3: -2,
					4: -1,
				},
				2: {
					1: 3,
					2: 0,
					3: 1,
					4: 2,
				},
				3: {
					1: 2,
					2: -1,
					3: 0,
					4: 1,
				},
				4: {
					1: 1,
					2: -2,
					3: -1,
					4: 0,
				},
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
			wantDist:          nil,
			wantNegativeCycle: true,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotDist, _, gotCycle := FloydWarshall(tt.graph, tt.weights)
			if !reflect.DeepEqual(gotDist, tt.wantDist) {
				t.Errorf("FloydWarshall()\n  gotDist = %v\n wantDist = %v", gotDist, tt.wantDist)
			}
			if gotCycle != tt.wantNegativeCycle {
				t.Errorf("FloydWarshall() gotCycle = %v, wantNegativeCycle %v", gotCycle, tt.wantNegativeCycle)
			}
		})
	}
}
