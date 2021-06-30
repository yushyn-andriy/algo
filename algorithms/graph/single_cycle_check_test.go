package graph_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/graph"
)

func TestSingleCycleCheck(t *testing.T) {
	tests := []struct {
		x        []int
		expected bool
	}{
		{
			[]int{},
			true,
		},
		{
			[]int{2},
			true,
		},
		{
			[]int{2, 2, 2},
			true,
		},
		{
			[]int{2, 2, 1},
			false,
		},
		{
			[]int{2, 2, 3},
			false,
		},
		{
			[]int{2, 3, 1, -4, -4, 2},
			true,
		},
		{
			[]int{2, 4, 2, -2, -1, 1},
			true,
		},
		{
			[]int{2, 4, 2, -2, -1, -1},
			false,
		},
	}

	for i, test := range tests {
		result := graph.SingleCycleCheck(test.x)
		if result != test.expected {
			t.Errorf("test(%d) Expected %v got %v", i, test.expected, result)
		}
	}
}
