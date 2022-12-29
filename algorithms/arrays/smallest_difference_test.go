package arrays_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestSmallestDifference(t *testing.T) {
	tests := []struct {
		a, b, expectedSmallestNums []int
	}{
		{
			[]int{-1, 5, 10, 20, 28, 3},
			[]int{26, 134, 135, 15, 17},
			[]int{28, 26},
		},
		{
			[]int{1, 2, 3},
			[]int{-4, -5, 0, 3},
			[]int{3, 3},
		},
		{
			[]int{1, 6, -5},
			[]int{3, -7, 2},
			[]int{1, 2},
		},
	}

	for i, test := range tests {
		smallestDiffNums := arrays.SmallestDifference(test.a, test.b)
		if !reflect.DeepEqual(smallestDiffNums, test.expectedSmallestNums) {
			t.Errorf("test(%d) expected %v got %v", i, test.expectedSmallestNums, smallestDiffNums)
			return
		}
	}
}
