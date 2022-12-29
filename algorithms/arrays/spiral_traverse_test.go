package arrays_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/arrays"
)

func TestSpiralTraverse(t *testing.T) {
	tests := []struct {
		matrix                   [][]int
		expectedTraverseSequance []int
	}{
		{
			[][]int{
				{1, 2, 3, 4},
				{6, 7, 8, 9},
				{11, 12, 13, 14},
				{15, 16, 17, 18},
			},
			[]int{1, 2, 3, 4, 9, 14, 18, 17, 16, 15, 11, 6, 7, 8, 13, 12},
		},
		{
			[][]int{
				{1, 2, 3},
				{6, 7, 8},
				{11, 12, 13},
			},
			[]int{1, 2, 3, 8, 13, 12, 11, 6, 7},
		},
		{
			[][]int{
				{},
			},
			[]int{},
		},
		{
			[][]int{
				{1, 2, 3},
				{6, 7, 8},
			},
			[]int{1, 2, 3, 8, 7, 6},
		},
		{
			[][]int{
				{1, 2},
				{10, 3},
				{9, 4},
				{8, 5},
				{7, 6},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{10, 11, 12, 5},
				{9, 8, 7, 6},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			[][]int{
				{1},
				{3},
				{2},
				{5},
				{4},
				{7},
				{6},
			},
			[]int{1, 3, 2, 5, 4, 7, 6},
		},
		{
			[][]int{
				{1, 2, 3},
				{12, 13, 4},
				{11, 14, 5},
				{10, 15, 6},
				{9, 8, 7},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
	}

	for i, test := range tests {
		sequance := arrays.SpiralTraverse(test.matrix)
		if !reflect.DeepEqual(sequance, test.expectedTraverseSequance) {
			t.Errorf("test(%d) expected sequance=%v, got %v", i, test.expectedTraverseSequance, sequance)
			return
		}
	}
}
