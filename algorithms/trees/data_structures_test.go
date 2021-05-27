package trees_test

import (
	"fmt"
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestDelete(t *testing.T) {
	tests := []struct {
		root                   *trees.Tree
		treeArr                []int
		keyNodeToDelete        int
		expectedTraverseString string
	}{
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			3,
			"0:1:2:4:5:6:",
		},
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			5,
			"0:1:2:3:4:6:",
		},
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			4,
			"0:1:2:3:5:6:",
		},
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			1,
			"0:2:3:4:5:6:",
		},
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			0,
			"1:2:3:4:5:6:",
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			3,
			"2:5:6:8:10:12:13:14:17:",
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			10,
			"2:3:5:6:8:12:13:14:17:",
		},
		{
			&trees.Tree{Key: 10},
			[]int{12, 11},
			12,
			"10:11:",
		},
		{
			&trees.Tree{Key: 10},
			[]int{6, 7},
			6,
			"7:10:",
		},
		{
			&trees.Tree{Key: 10},
			[]int{6, 7},
			600,
			"6:7:10:",
		},
	}

	for i, test := range tests {
		root := test.root
		for _, key := range test.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}
		trees.Delete(root, test.keyNodeToDelete)

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
