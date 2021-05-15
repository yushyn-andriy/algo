package trees_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestNodeDepths(t *testing.T) {
	testCases := []struct {
		root               *trees.Tree
		treeArr            []int
		expectedNodeDepths int
	}{
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			11,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			26,
		},
		{
			&trees.Tree{Key: 4},
			[]int{1, 15, 14, 20, 6, 17, 11, 19, 10},
			25,
		},
	}

	for _, testCase := range testCases {
		root := testCase.root
		for _, key := range testCase.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}

		nodeDepths := trees.NodeDepths(root)
		if nodeDepths != testCase.expectedNodeDepths {
			t.Errorf("Expected depths=%v, got=%v", testCase.expectedNodeDepths, nodeDepths)
		}
	}
}

func BenchmarkNodeDepths(b *testing.B) {
	root := &trees.Tree{Key: 4}
	treeArr := []int{1, 15, 14, 20, 6, 17, 11, 19, 10}
	for _, key := range treeArr {
		trees.Insert(root, &trees.Tree{Key: key})
	}

	for i := 0; i < b.N; i++ {
		trees.NodeDepths(root)
	}

}
