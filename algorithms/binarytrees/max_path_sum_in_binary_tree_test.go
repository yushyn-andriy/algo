package binarytrees

import "testing"

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		in         *BinaryTree
		maxPathSum int
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
			18,
		},
		{
			&BinaryTree{
				Value: -2,
				Left: &BinaryTree{
					Value: -1,
				},
			},
			-1,
		},
	}

	for i, test := range tests {
		maxPathSum := MaxPathSum(test.in)
		if maxPathSum != test.maxPathSum {
			t.Errorf("test(%d) expected %v got %v", i, test.maxPathSum, maxPathSum)
		}
	}

}
