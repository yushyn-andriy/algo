package searching

import (
	"reflect"
	"testing"
)

func TestSearchInSortedMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		out    []int
	}{
		{
			[][]int{
				{1, 4, 7, 12, 15, 1000},
				{2, 5, 19, 31, 32, 1001},
				{3, 8, 24, 33, 35, 1002},
				{40, 41, 42, 44, 45, 1003},
				{99, 100, 103, 106, 128, 1004},
			},
			44,
			[]int{3, 3},
		},
	}

	for i, test := range tests {
		out := SearchInSortedMatrix(test.matrix, test.target)
		if reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
