package dynamic_programming_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/dynamic_programming"
)

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		amount int
		coins  []int
		n      int // expected number of ways
	}{
		{
			10,
			[]int{1, 5, 10, 25},
			4,
		}, // 10 * 1, 5 * 1 + 1 * 5, 1 * 10, 5 * 2
		{
			10,
			[]int{1, 5},
			3, // 1 * 10, 5 * 2, 5 * 1 + 5 * 1
		},
		{
			10,
			[]int{1},
			1, // 1 * 10
		},
		{
			7,
			[]int{2, 5},
			1, // 2 * 1 + 5*1
		},
		{
			9,
			[]int{2},
			0,
		},
	}

	for i, test := range tests {
		n := dynamic_programming.NumberOfWays(test.amount, test.coins)
		if n != test.n {
			t.Errorf("test(%d) Expected number of ways %d but got %d", i, test.n, n)
		}
	}
}
