package arrays_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestFirstDuplicateValue(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{5, 1, 4, 2, 5},
			5,
		},
		{
			[]int{1, 2, 1, 3},
			1,
		},
		{
			[]int{2, 3, 4, 2},
			2,
		},
		{
			[]int{2, 1, 5, 3, 3, 2, 4},
			3,
		},
	}

	for i, test := range tests {
		out := arrays.FirstDuplicateValue(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
			return
		}
	}
}
