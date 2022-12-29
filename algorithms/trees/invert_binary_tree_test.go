package trees_test

import (
	"fmt"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestInvertBinaryTree(t *testing.T) {
	tests := []struct {
		root                   *trees.Tree
		treeArr                []int
		expectedTraverseString string
	}{
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			"6:5:4:3:2:1:0:",
		},
		{
			&trees.Tree{Key: 2},
			[]int{1, 5, 10, 5, 6, 0},
			"10:6:5:5:2:1:0:",
		},
	}

	for i, test := range tests {
		root := test.root
		for _, key := range test.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}
		trees.InvertBinaryTree(root)

		s := ""
		trees.InOrderWalk(root, func(tree *trees.Tree) {
			s += fmt.Sprintf("%d:", tree.Key)
		})
		if s != test.expectedTraverseString {
			t.Errorf("test(%d) Expected traverse string %v, got %v", i, test.expectedTraverseString, s)
			return
		}
	}
}
