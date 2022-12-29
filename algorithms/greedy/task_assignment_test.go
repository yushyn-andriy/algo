package greedy

import (
	"reflect"
	"testing"
)

func TestTaskAssignment(t *testing.T) {
	tests := []struct {
		tasks []int
		k     int
		out   [][]int
	}{
		{
			[]int{1, 3, 5, 3, 1, 4},
			3,
			[][]int{
				{0, 2},
				{4, 5},
				{1, 3},
			},
		},
	}

	for i, test := range tests {
		out := TaskAssignment(test.k, test.tasks)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
		}
	}
}
