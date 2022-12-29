package stack

import "testing"

func TestLargestRectangleUnderSkyline(t *testing.T) {
	tests := []struct {
		buildings []int
		maxArea   int
	}{
		{
			[]int{1, 2, 4, 3, 1, 1, 1},
			7,
		},
	}

	for i, test := range tests {
		maxArea := LargestRectangleUnderSkyline(test.buildings)
		if test.maxArea != maxArea {
			t.Errorf("test(%d) expected %v got %v", i, test.maxArea, maxArea)
		}
	}
}
