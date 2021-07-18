package binarytrees

import (
	"testing"
)

func TestBinaryTreeDiameter(t *testing.T) {
	tests := []struct {
		in  BinaryTree
		out int
	}{
		{
			BinaryTree{
				Value: 1,
				Left: &BinaryTree{
					Value: 3,
					Left: &BinaryTree{
						Value: 7,
						Left: &BinaryTree{
							Value: 8,
							Left: &BinaryTree{
								Value: 9,
							},
						},
					},
					Right: &BinaryTree{
						Value: 4,
						Right: &BinaryTree{
							Value: 5,
							Right: &BinaryTree{
								Value: 6,
							},
						},
					},
				},
				Right: &BinaryTree{
					Value: 2,
				},
			},
			6,
		},
	}

	for i, test := range tests {
		out := BinaryTreeDiameter(&test.in)
		if out != test.out {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
