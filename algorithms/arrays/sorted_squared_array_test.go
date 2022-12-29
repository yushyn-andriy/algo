package arrays_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestSortedSquaredArray(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{

		{
			[]int{-4, -2, 2, 4},
			[]int{4, 4, 16, 16},
		},
	}

	for i, test := range tests {
		output := arrays.SortedSquaredArray(test.input)
		if !reflect.DeepEqual(output, test.output) {
			t.Errorf("test(%d) Expected %v, got %v", i, test.output, output)
		}
	}
}
