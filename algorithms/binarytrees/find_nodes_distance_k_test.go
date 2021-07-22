package binarytrees

import (
	"reflect"
	"testing"
)

func TestFindNodesDistanceK(t *testing.T) {
	tests := []struct {
		in        *BinaryTree
		target, k int
		array     []int
	}{
		{
			&BinaryTree{
				Value: 1,
				Left: &BinaryTree{
					Value: 2,
					Left: &BinaryTree{
						Value: 4,
					}, 
					Right: &BinaryTree{
						Value: 5,
					},
				},
				Right: &BinaryTree{
					Value: 3,
					Left: &BinaryTree{
						Value: 6,
					},
					Right: &BinaryTree{
						Value: 7,
					},
				},
			},
			4, 2,
			[]int{5, 1},
		},
		{
			&BinaryTree{
				Value: -2,
				Left: &BinaryTree{
					Value: -1,
				},
			},
			-1, 1,
			[]int{-2},
		},
	}

	for i, test := range tests {
		out := FindNodesDistanceK(test.in, test.target, test.k)
		if !reflect.DeepEqual(out, test.array) {
			t.Errorf("test(%d) expected %v got %v", i, test.array, out)
		}
	}

}
