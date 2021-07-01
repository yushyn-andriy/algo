package graph_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/graph"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		nodes    map[int][]int
		source   int
		expected []int
	}{
		{
			map[int][]int{
				1: {2, 3, 4},
				2: {1, 5, 6},
				3: {1, 7},
				4: {1, 8},
				5: {2},
				6: {2},
				7: {3},
				8: {3, 4},
			},
			1,
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			map[int][]int{
				0: {1, 2, 3},
				1: {0, 3, 4},
				2: {0, 4},
				3: {0, 1, 4},
				4: {1, 2, 3},
			},
			0,
			[]int{0, 1, 2, 3, 4},
		},
	}
	for i, test := range tests {
		g := graph.Graph{Nodes: test.nodes}
		r := g.BFS(test.source)
		if !reflect.DeepEqual(test.expected, r) {
			t.Errorf("test(%d) Expected %v, got %v", i, test.expected, r)
		}
	}
}
