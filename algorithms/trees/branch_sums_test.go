package trees_test

import (
	"testing"

	"github.com/baybaraandrey/algo/algorithms/trees"
)

func TestBranchSum(t *testing.T) {
	testCases := []struct {
		root     *trees.Tree
		treeArr  []int
		expected []int
	}{
		{
			&trees.Tree{Key: 5},
			[]int{3, 8, 1, 4, 6, 10},
			[]int{9, 12, 19, 23},
		},
		{
			&trees.Tree{Key: 8},
			[]int{7, 25, 5, 86, 32, 105, 1, 6, 3, 96},
			[]int{24, 26, 151, 320},
		},
		{
			&trees.Tree{Key: 2},
			[]int{-5, 4, -6, 1, 3, 6, -2},
			[]int{-9, -4, 9, 12},
		},
	}
	for _, testCase := range testCases {
		root := testCase.root
		for _, key := range testCase.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}
		result := *trees.BranchSum(root, 0, &[]int{})
		if len(testCase.expected) != len(result) {
			t.Errorf("Expected result=%v, got %v", testCase.expected, result)
			return
		}
		for i := range testCase.expected {
			if result[i] != testCase.expected[i] {
				t.Errorf("Expected result=%v, got %v", testCase.expected, result)
			}
		}
	}
}

func BenchmarkBranchSum(b *testing.B) {
	root := &trees.Tree{Key: 5}

	for _, key := range []int{3, 8, 1, 4, 6, 10} {
		trees.Insert(root, &trees.Tree{Key: key})
	}

	for i := 0; i < b.N; i++ {
		trees.BranchSum(root, 0, &[]int{})
	}
}
