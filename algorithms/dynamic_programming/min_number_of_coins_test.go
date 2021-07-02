package dynamic_programming_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/dynamic_programming"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{
			[]int{1, 2, 3},
			6,
			2,
		},
		{
			[]int{1, 1, 1},
			6,
			6,
		},
		{
			[]int{1, 2},
			6,
			3,
		},
		{
			[]int{3},
			4,
			0,
		},
	}
	for i, test := range tests {
		r := dynamic_programming.MinNumberOfCoins(test.amount, test.coins)
		if r != test.expected {
			t.Errorf("test(%d) Expected %v, got %v", i, test.expected, r)
		}
	}
}
