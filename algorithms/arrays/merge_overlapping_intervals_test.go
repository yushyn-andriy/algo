package arrays_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestMergeOverlappingIntervals(t *testing.T) {
	tests := []struct {
		in  [][]int
		out [][]int
	}{
		{
			[][]int{
				{1, 2},
				{3, 5},
				{4, 7},
				{6, 8},
				{9, 10},
			},
			[][]int{
				{1, 2},
				{3, 8},
				{9, 10},
			},
		},
	}

	for i, test := range tests {
		out := arrays.MergeOverlappingIntervals(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
			return
		}
	}
}
