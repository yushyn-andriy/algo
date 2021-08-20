package graph

import "testing"

func TestCycleInGraph1(t *testing.T) {
	tests := []struct {
		in  [][]int
		out bool
	}{
		{
			[][]int{
				{1, 3},
				{2, 3, 4},
				{0},
				{},
				{2, 5},
				{},
			},
			true,
		},
	}

	for i, test := range tests {
		out := CycleInGraph1(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}

func TestCycleInGraph2(t *testing.T) {
	tests := []struct {
		in  [][]int
		out bool
	}{
		{
			[][]int{
				{1, 3},
				{2, 3, 4},
				{0},
				{},
				{2, 5},
				{},
			},
			true,
		},
	}

	for i, test := range tests {
		out := CycleInGraph2(test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
