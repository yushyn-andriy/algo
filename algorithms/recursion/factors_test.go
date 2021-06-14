package recursion_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/recursion"
)

func TestFactors(t *testing.T) {
	tests := []struct {
		fn []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
	}

	for i, test := range tests {
		factors := recursion.Factors(test.fn)
		for _, l := range factors {
			t.Errorf("test(%d) %v", i, l)
		}
	}
}
