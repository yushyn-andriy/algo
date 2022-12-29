package arrays_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestArrayOfProducts(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{5, 1, 4, 2},
			[]int{8, 40, 10, 20},
		},
		{
			[]int{0, 0, 0, 0},
			[]int{0, 0, 0, 0},
		},
		{
			[]int{0, 1, 2, 3},
			[]int{6, 0, 0, 0},
		},
	}

	for i, test := range tests {
		out := arrays.ArrayOfProducts(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
			return
		}
	}
}
