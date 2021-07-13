package arrays_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestNonConstructibleChange(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{5, 7, 1, 1, 2, 3, 22},
			20,
		},
		{
			[]int{1, 1, 1, 1, 1, 1},
			7,
		},
		{
			[]int{87},
			1,
		},
		{
			[]int{2},
			1,
		},
		{
			[]int{1},
			2,
		},
	}

	for i, test := range tests {
		out := arrays.NonConstructibleChange(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
			return
		}
	}
}
