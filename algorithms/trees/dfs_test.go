package trees_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestDFSSearch(t *testing.T) {
	testCases := []struct {
		root    *trees.Tree
		treeArr []int
		key     int
	}{
		{
			&trees.Tree{Key: 5},
			[]int{3, 8, 1, 4, 6, 10},
			10,
		},
		{
			&trees.Tree{Key: 8},
			[]int{7, 25, 5, 86, 32, 105, 1, 6, 3, 96},
			25,
		},
		{
			&trees.Tree{Key: 2},
			[]int{-5, 4, -6, 1, 3, 6, -2},
			-5,
		},
		{
			&trees.Tree{Key: 3},
			[]int{3, 8, 1, 4, 6, 10, 32, 37, 45, 83,
				20, 19, 80, 27, 73, 60, 25, 88, 23, 33, 100, 26, 70, 56, 63, 92},
			56,
		},
	}
	for _, testCase := range testCases {
		root := testCase.root
		for _, key := range testCase.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}

		node := trees.DFSSearch(root, testCase.key)
		if node == nil {
			t.Errorf("Expected node not nil")
		} else {
			if node.Key != testCase.key {
				t.Errorf("Expected Key=%v got key=%v", testCase.key, node.Key)
			}
		}
	}
}

func BenchmarkDFSSearch(b *testing.B) {
	root := &trees.Tree{Key: 5}

	for _, key := range []int{3, 8, 1, 4, 6, 10, 32, 37, 45, 83,
		20, 19, 80, 27, 73, 60, 25, 88, 23, 33, 100, 26, 70, 56, 63, 92} {
		trees.Insert(root, &trees.Tree{Key: key})
	}

	for i := 0; i < b.N; i++ {
		trees.DFSSearch(root, 6)
	}
}
