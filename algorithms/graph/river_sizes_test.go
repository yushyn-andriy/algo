package graph

import (
	"reflect"
	"testing"
)

func TestRiverSizes(t *testing.T) {
	tests := []struct {
		matrix [][]int
		out    []int
	}{
		{
			[][]int{
				{1, 0, 0, 0},
				{0, 1, 1, 0},
				{1, 0, 0, 0},
				{1, 1, 0, 0},
			},
			[]int{1, 2, 3},
		},
	}

	for i, test := range tests {
		out := RiverSizes(test.matrix)
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
