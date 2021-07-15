package trees_test

import (
	"reflect"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestFindKthLargestValueInBst(t *testing.T) {
	tests := []struct {
		node trees.BST
		k    int
		out  int
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
			1,
			22,
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
			2,
			15,
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
			3,
			14,
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
			4,
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
			10,
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
			6,
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
			7,
			5,
		},
	}

	for i, test := range tests {
		out := trees.FindKthLargestValueInBst(&test.node, test.k)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("test(%d) expected %v, got %v", i, test.out, out)
		}
	}
}
