package stack

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		array, out []int
	}{
		{
			[]int{3, 6, 2, -1, 5, 10},
			[]int{6, 10, 5, 5, 10, -1},
		},
	}

	for i, test := range tests {
		out := NextGreaterElement(test.array)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
