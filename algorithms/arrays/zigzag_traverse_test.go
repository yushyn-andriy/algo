package arrays

import (
	"reflect"
	"testing"
)

func TestZigzagTraverse(t *testing.T) {
	tests := []struct {
		in  [][]int
		out []int
	}{
		{
			[][]int{
				{1, 3, 4, 10},
				{2, 5, 9, 11},
				{6, 8, 12, 15},
				{7, 13, 14, 16},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
	}

	for i, test := range tests {
		out := ZigzagTraverse(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
		}
	}
}
