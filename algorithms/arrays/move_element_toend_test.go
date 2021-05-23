package arrays_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestMoveElementToEnd(t *testing.T) {
	tests := []struct {
		arr, expectedElemnts []int
		element              int
	}{
		{
			[]int{4, 1, 2, 2, 2, 2, 3, 2, 2},
			[]int{4, 1, 3, 2, 2, 2, 2, 2, 2},
			2,
		},
		{
			[]int{3, 3, 1, 2},
			[]int{2, 1, 3, 3},
			3,
		},
		{
			[]int{5, 5, 1, 6, 5, 5, 5},
			[]int{6, 1, 5, 5, 5, 5, 5},
			5,
		},
		{
			[]int{5, 5, 1, 6, 5, 5, 5},
			[]int{5, 5, 1, 6, 5, 5, 5},
			-100,
		},
		{
			[]int{5, 5, 1, 6, 5, 5, 5},
			[]int{5, 5, 5, 6, 5, 5, 1},
			1,
		},
	}

	for i, test := range tests {
		arr := arrays.MoveElementToEnd(test.arr, test.element)
		if !reflect.DeepEqual(arr, test.expectedElemnts) {
			t.Errorf("test(%d) expected %v got %v", i, test.expectedElemnts, arr)
			return
		}
	}
}
