package recursion

import (
	"reflect"
	"testing"
)

func TestGetPermutations(t *testing.T) {
	tests := []struct {
		in  []int
		out [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			[]int{},
			[][]int{},
		},
	}

	for i, test := range tests {
		out := GetPermutations(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
