package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoNumbers(t *testing.T) {
	tests := []struct {
		in     []int
		out    []int
		target int
	}{
		{
			[]int{3, 3},
			[]int{0, 1},
			6,
		},
		{
			[]int{3, 5, -4, 8, 11, 1, -1},
			[]int{4, 6},
			10,
		},
		{
			[]int{3, 2, 8},
			[]int{1, 2},
			10,
		}, {
			[]int{3, 4, 3, 7},
			[]int{0, 3},
			10,
		},
		{
			[]int{-18, 4, 3, 7, 28},
			[]int{0, 4},
			10,
		},
		{
			[]int{4, 4},
			[]int{},
			10,
		},
		{
			[]int{2, 5, 5, 1},
			[]int{1, 2},
			10,
		},
	}

	for i, test := range tests {
		out := twoSum(test.in, test.target)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}

	}
}
