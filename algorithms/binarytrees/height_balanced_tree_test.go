package binarytrees

import "testing"

func TestHeightBalancedBinaryTree(t *testing.T) {
	tests := []struct {
		in         *BinaryTree
		isBalanced bool
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
						Left: &BinaryTree{
							Value: 7,
						},
						Right: &BinaryTree{
							Value: 8,
						},
					},
				},
				Right: &BinaryTree{
					Value: 3,
					Right: &BinaryTree{
						Value: 6,
					},
				},
			},
			true,
		},
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
						Left: &BinaryTree{
							Value: 7,
						},
						Right: &BinaryTree{
							Value: 8,
						},
					},
				},
				Right: &BinaryTree{
					Value: 3,
					Right: &BinaryTree{
						Value: 6,
						Right: &BinaryTree{
							Value: 6,
						},
					},
				},
			},
			false,
		},
	}

	for i, test := range tests {
		isBalanced := HeightBalancedBinaryTree(test.in)
		if isBalanced != test.isBalanced {
			t.Errorf("test(%d) expected %v got %v", i, test.isBalanced, isBalanced)
		}
	}
}
