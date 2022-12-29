package arrays_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestIsMonotonic(t *testing.T) {
	tests := []struct {
		arr                 []int
		expectedIsMonotonic bool
	}{
		{
			[]int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001},
			true,
		},
		{
			[]int{1, 2},
			true,
		},
		{
			[]int{2, 1},
			true,
		},
		{
			[]int{1, 3, 2},
			false,
		},
		{
			[]int{4, 3, 2, -1, 0},
			false,
		},
	}

	for i, test := range tests {
		isMonotonic := arrays.IsMonotonic(test.arr)
		if isMonotonic != test.expectedIsMonotonic {
			t.Errorf("test(%d) expected isMonotonic=%v, got %v", i, test.expectedIsMonotonic, isMonotonic)
			return
		}
	}
}
