package trees_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestFindClosestValue(t *testing.T) {
	tests := []struct {
		node   trees.BST
		target int
		out    int
	}{
		{
			trees.BST{
				Value: 10,
				Left: &trees.BST{
					Value: 5,
					Left: &trees.BST{
						Value: 2,
						Left: &trees.BST{
							Value: 1,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &trees.BST{
						Value: 5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &trees.BST{
					Value: 15,
					Left: &trees.BST{
						Value: 13,
						Left:  nil,
						Right: &trees.BST{
							Value: 14,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &trees.BST{
						Value: 22,
						Left:  nil,
						Right: nil,
					},
				},
			},
			12,
			13,
		},
		{
			trees.BST{
				Value: 10,
				Left: &trees.BST{
					Value: 5,
					Left: &trees.BST{
						Value: 2,
						Left: &trees.BST{
							Value: 1,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &trees.BST{
						Value: 5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &trees.BST{
					Value: 15,
					Left: &trees.BST{
						Value: 13,
						Left:  nil,
						Right: &trees.BST{
							Value: 14,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &trees.BST{
						Value: 22,
						Left:  nil,
						Right: nil,
					},
				},
			},
			5,
			5,
		},
		{
			trees.BST{
				Value: 10,
				Left: &trees.BST{
					Value: 5,
					Left: &trees.BST{
						Value: 2,
						Left: &trees.BST{
							Value: 1,
							Left:  nil,
							Right: &trees.BST{
								Value: 1,
								Left:  nil,
								Right: &trees.BST{
									Value: 1,
									Left:  nil,
									Right: nil,
								},
							},
						},
						Right: nil,
					},
					Right: &trees.BST{
						Value: 5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &trees.BST{
					Value: 15,
					Left: &trees.BST{
						Value: 13,
						Left:  nil,
						Right: &trees.BST{
							Value: 14,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &trees.BST{
						Value: 22,
						Left:  nil,
						Right: nil,
					},
				},
			},
			1,
			2,
		},
	}

	for i, test := range tests {
		out := test.node.FindClosestValue(test.target)
		if out != test.out {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
			return
		}
	}
}
