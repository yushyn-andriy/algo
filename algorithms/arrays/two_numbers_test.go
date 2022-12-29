package arrays_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestTwoNumbers(t *testing.T) {
	tests := []struct {
		in     []int
		out    []int
		target int
	}{
		{
			[]int{3, 5, -4, 8, 11, 1, -1},
			[]int{11, -1},
			10,
		},
		{
			[]int{3, 2, 8},
			[]int{2, 8},
			10,
		}, {
			[]int{3, 4, 3, 7},
			[]int{3, 7},
			10,
		},
		{
			[]int{-18, 4, 3, 7, 28},
			[]int{-18, 28},
			10,
		},
		{
			[]int{4, 4},
			[]int{},
			10,
		},
	}

	for i, test := range tests {
		out := arrays.TwoNumberSum(test.in, test.target)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}

	}
}

func BenchmarkTwoNumbers(b *testing.B) {
	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	for i := 0; i < b.N; i++ {
		arrays.TwoNumbers(arr, targetSum)
	}
}
