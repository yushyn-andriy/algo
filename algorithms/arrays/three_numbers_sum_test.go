package arrays_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestThreeNumbers(t *testing.T) {
	tests := []struct {
		arr          []int
		targetSum    int
		expectedSums [][]int
	}{
		{
			[]int{12, 3, 1, 2, -6, 5, -8, 6},
			0,
			[][]int{
				{-8, 2, 6},
				{-8, 3, 5},
				{-6, 1, 5},
			},
		},
		{
			[]int{-10, -9, 8, 4, 6, 3, 5, -3, 7},
			1,
			[][]int{
				{-10, 3, 8},
				{-10, 4, 7},
				{-10, 5, 6},
				{-9, 3, 7},
				{-9, 4, 6},
			},
		},
	}
	for i, test := range tests {
		sum, err := arrays.ThreeNumbers(test.arr, test.targetSum)
		if err != nil {
			t.Errorf("test(%d) Got error %v", i, err)
			return
		}
		if !reflect.DeepEqual(test.expectedSums, sum) {
			t.Errorf("test(%d) expected %v got %v", i, test.expectedSums, sum)
			return
		}
	}

}
