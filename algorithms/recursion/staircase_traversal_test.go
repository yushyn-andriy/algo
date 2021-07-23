package recursion

import (
	"testing"
)

func TestStaircaseTraversal(t *testing.T) {
	tests := []struct {
		height, maxSteps int
		out              int
	}{
		{
			4, 2, 5,
		},
	}

	for i, test := range tests {
		out := StaircaseTraversal(test.height, test.maxSteps)
		if out != test.out {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}

func TestStaircaseTraversalSolution2(t *testing.T) {
	tests := []struct {
		height, maxSteps int
		out              int
	}{
		{
			4, 2, 5,
		},
	}

	for i, test := range tests {
		out := StaircaseTraversalSolution2(test.height, test.maxSteps)
		if out != test.out {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
