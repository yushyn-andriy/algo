package dynamic_programming_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/dynamic_programming"
)

func TestMaxSubsetSumNoAdjacent(t *testing.T) {
	tests := []struct {
		x []int // input array
		m int   // expected max ubset sum no adjacent
	}{
		{
			[]int{7, 10, 12, 7, 9, 14},
			33,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{9, 2, 3, 4, 5, 13},
			26,
		},
	}

	for i, test := range tests {
		m := dynamic_programming.MaxSubsetSumNoAdjacent(test.x)
		if m != test.m {
			t.Errorf("test(%d): Wrong subset sum,  expected %d but got %d", i, test.m, m)
		}
	}
}
