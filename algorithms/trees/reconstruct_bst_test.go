package trees_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestReconstructBst(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{10, 4, 2, 1, 5, 17, 19, 18},
			[]int{1, 2, 4, 5, 10, 17, 18, 19},
		},
	}

	for i, test := range tests {
		bst := trees.ReconstructBst(test.in)
		out := bst.InOrderTraverse([]int{})
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}

	}
}
