package trees_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestSameBsts(t *testing.T) {
	tests := []struct {
		arrayOne []int
		arrayTwo []int
		out      bool
	}{
		{
			[]int{10, 15, 8, 12, 94, 81, 5, 2},
			[]int{10, 8, 5, 15, 2, 12, 94, 81},
			true,
		},
		{
			[]int{10, 15, 8, 12, 94, 81, 5, 2},
			[]int{11, 8, 5, 15, 2, 12, 94, 81},
			false,
		},
	}

	for i, test := range tests {
		out := trees.SameBsts(test.arrayOne, test.arrayTwo)
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}

	}
}
