package arrays

import (
	"reflect"
	"testing"
)

func TestFourNumberSum(t *testing.T) {
	tests := []struct {
		array  []int
		target int
		out    [][]int
	}{
		{
			[]int{7, 6, 4, -1, 1, 2},
			16,
			[][]int{
				{7, 6, 4, -1},
				{7, 6, 1, 2},
			},
		},
	}

	for i, test := range tests {
		out := FourNumberSum(test.array, test.target)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
			return
		}
	}
}
