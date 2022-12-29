package recursion

import (
	"reflect"
	"testing"
)

func TestPowerset(t *testing.T) {
	tests := []struct {
		in  []int
		out [][]int
	}{
		{
			[]int{1, 2},
			[][]int{
				{1, 2},
				{1},
				{2},
				{},
			},
		},
		{
			[]int{},
			[][]int{{}},
		},
	}

	for i, test := range tests {
		out := Powerset(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
