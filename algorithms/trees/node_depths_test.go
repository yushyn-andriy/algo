package trees_test

import (
	"reflect"
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

func TestNodeDepthsMax(t *testing.T) {
	testCases := []struct {
		root                  *trees.Tree
		treeArr               []int
		expectedNodeDepthsMax int
	}{
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			3,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			5,
		},
		{
			&trees.Tree{Key: 4},
			[]int{1, 15, 14, 20, 6, 17, 11, 19, 10},
			5,
		},
	}

	for _, testCase := range testCases {
		root := testCase.root
		for _, key := range testCase.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}

		nodeDepths := trees.NodeDepthsMax(root)
		if nodeDepths != testCase.expectedNodeDepthsMax {
			t.Errorf("Expected depthsMax=%v, got=%v", testCase.expectedNodeDepthsMax, nodeDepths)
		}
	}
}

func TestMinHeightBST(t *testing.T) {
	tests := []struct {
		x                  []int
		expectedBuildedArr []int
		expectedDepthsMax  int
	}{
		{
			[]int{0, 1, 3, 4, 5, 6, 10},
			[]int{4, 1, 0, 3, 6, 5, 10},
			3,
		},
		{
			[]int{2, 3, 5, 6, 8, 10, 12, 13, 14, 17},
			[]int{10, 5, 2, 3, 6, 8, 14, 12, 13, 17},
			4,
		},
		{
			[]int{1, 6, 10, 11, 14, 15, 17, 19, 20},
			[]int{14, 10, 1, 6, 11, 19, 15, 17, 20},
			4,
		},
		{
			[]int{1, 2, 5, 7, 10, 13, 14, 15, 22},
			[]int{10, 5, 1, 2, 7, 15, 13, 14, 22},
			4,
		},
	}

	for i, test := range tests {
		root := trees.MinHeightBST(test.x)
		nodeDepthsMax := trees.NodeDepthsMax(root)
		if nodeDepthsMax != test.expectedDepthsMax {
			t.Errorf("test(%d) Expected depthsMax=%v, got=%v", i, test.expectedDepthsMax, nodeDepthsMax)
		}
		buildedArr := trees.BuildArr(test.x)
		if !reflect.DeepEqual(buildedArr, test.expectedBuildedArr) {
			t.Errorf("test(%d) Expected buildedArr=%v, got=%v", i, test.expectedBuildedArr, buildedArr)
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
