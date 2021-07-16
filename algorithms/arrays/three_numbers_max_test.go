package arrays_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestMaxThreeNumbers(t *testing.T) {
	tests := []struct {
		arr      []int
		expected arrays.ThreeMax
	}{
		{
			[]int{6, 1, 2, 3},
			arrays.ThreeMax{
				{true, 2},
				{true, 3},
				{true, 6},
			},
		},
		{
			[]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7},
			arrays.ThreeMax{
				{true, 18},
				{true, 141},
				{true, 541},
			},
		},
	}

	for i, test := range tests {
		threeMax := arrays.ThreeNumbersMax(test.arr)
		if threeMax[arrays.FIRST].Value != test.expected[arrays.FIRST].Value {
			t.Errorf("test(%v) error expected %+v got %+v", i, test.expected[arrays.FIRST], threeMax[arrays.FIRST])
		}
		if threeMax[arrays.SECOND].Value != test.expected[arrays.SECOND].Value {
			t.Errorf("test(%v) error expected %+v got %+v", i, test.expected[arrays.SECOND], threeMax[arrays.SECOND])
		}
		if threeMax[arrays.THIRD].Value != test.expected[arrays.THIRD].Value {
			t.Errorf("test(%v) error expected %+v got %+v", i, test.expected[arrays.THIRD], threeMax[arrays.THIRD])
		}
	}
}

func TestFindThreeLargestNumbersAlgoSoulution1(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			[]int{6, 1, 2, 3},
			[]int{2, 3, 6},
		},
		{
			[]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7},
			[]int{18, 141, 541},
		},
	}

	for i, test := range tests {
		out := arrays.FindThreeLargestNumbersAlgoSoulution1(test.arr)
		if !reflect.DeepEqual(out, test.expected) {
			t.Errorf("test(%v) error expected %+v got %+v", i, test.expected, out)
		}
	}
}
