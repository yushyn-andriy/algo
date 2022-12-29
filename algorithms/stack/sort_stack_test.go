package stack

import (
	"reflect"
	"testing"
)

func TestSortStack(t *testing.T) {
	tests := []struct {
		in, out []int
	}{
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
	}

	for i, test := range tests {
		out := SortStack(test.in)
		if reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
