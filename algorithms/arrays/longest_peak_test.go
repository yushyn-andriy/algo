package arrays_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestLongestPeak(t *testing.T) {
	tests := []struct {
		x                   []int
		expectedLongestPeak int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1},
			0,
		},
		{
			[]int{1, 2},
			0,
		},
		{
			[]int{1, 2, 3},
			0,
		},
		{
			[]int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3},
			6,
		},
		{
			[]int{1, 2, 9, 7, 4, 3, 2, 6, 5, -1, -3, 2, 3},
			7,
		},
		{
			[]int{1, 2, 3, 2, 1, 8, 7, 6, 5, 12, 2, 2, 3},
			5,
		},
	}

	for i, test := range tests {
		peak := arrays.LongestPeak(test.x)
		if peak != test.expectedLongestPeak {
			t.Errorf("test(%d) expecfted %d, got %d", i, test.expectedLongestPeak, peak)
			return
		}
	}
}
