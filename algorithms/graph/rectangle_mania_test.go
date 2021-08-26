package graph

import (
	"fmt"
	"testing"
)

func TestRectangleMania(t *testing.T) {
	tests := []struct {
		coords [][]int
		out    int
	}{
		{
			[][]int{
				{0, 0},
				{0, 1},
				{1, 1},
				{1, 0},
				{2, 1},
				{2, 0},
				{3, 1},
				{3, 0},
			},
			6,
		},
	}

	for i, test := range tests {
		out := RectangleMania(test.coords)
		if out != test.out {
			fmt.Printf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
