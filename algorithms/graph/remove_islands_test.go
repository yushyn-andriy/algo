package graph

import (
	"reflect"
	"testing"
)

func TestRemoveIslands(t *testing.T) {
	tests := []struct {
		matrix, out [][]int
	}{
		{
			[][]int{
				{1, 0, 0, 1},
				{1, 0, 1, 0},
				{0, 0, 1, 0},
				{1, 0, 0, 1},
			},
			[][]int{
				{1, 0, 0, 1},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 0, 0, 1},
			},
		},
	}

	for i, test := range tests {
		out := RemoveIslands(test.matrix)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}

}
