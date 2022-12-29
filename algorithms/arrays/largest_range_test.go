package arrays


import (
	"reflect"
	"testing"
)

func TestLargestRange(t *testing.T) {
	tests := []struct {
		in, out  []int
	}{
		{
			[]int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6},
			[]int{0, 7},
		},
	}

	for i, test := range tests {
		out := LargestRange(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
			return
		}
	}
}
