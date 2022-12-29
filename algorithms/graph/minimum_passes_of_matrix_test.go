package graph

import "testing"

func TestMinimumPassesOfMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		passes int
	}{

		{
			[][]int{
				{0, -1},
				{-2, 2},
				{5, -1},
			},
			1,
		},
		{
			[][]int{
				{0, -1, -1},
				{-2, 2, -1},
				{5, -1, -3},
			},
			2,
		},
		{
			[][]int{
				{0, -1, -3, 2, 0},
				{1, -2, -5, -1, -3},
				{3, 0, 0, -4, -1},
			},
			3,
		},
	}

	for i, test := range tests {
		passes := MinimumPassesOfMatrix(test.matrix)
		if passes != test.passes {
			t.Errorf("test(%d) expected %v got %v", i, test.passes, passes)
		}
	}

}
