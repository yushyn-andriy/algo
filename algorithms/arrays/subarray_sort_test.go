package arrays

import (
	"reflect"
	"testing"
)

func TestSubarraySort(t *testing.T) {
	tests := []struct {
		array, out []int
	}{
		{
			[]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19},
			[]int{3, 9},
		},
	}

	for i, test := range tests {
		out := SubarraySort(test.array)
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
