package trees_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestValidateThreeNodes(t *testing.T) {
	nodeThree := &trees.BST{
		Value: 3,
	}
	nodeTwo := &trees.BST{
		Value: 2,
		Left: &trees.BST{
			Value: 1,
			Left: &trees.BST{
				Value: 0,
			},
		},
		Right: &trees.BST{
			Value: 4,
			Left:  nodeThree,
		},
	}
	nodeOne := &trees.BST{
		Value: 5,
		Left:  nodeTwo,
		Right: &trees.BST{
			Value: 7,
			Left: &trees.BST{
				Value: 6,
			},
			Right: &trees.BST{
				Value: 8,
			},
		},
	}

	out := trees.ValidateThreeNodes(nodeOne, nodeTwo, nodeThree)
	if out != true {
		t.Errorf("test(%d) expected %v got %v", 0, true, out)
	}

}
