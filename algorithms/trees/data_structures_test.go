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
		{
			&trees.Tree{Key: 15},
			[]int{6, 4, 3, 5, 9, 7, 8},
			6,
			"3:4:5:7:8:9:15:",
		},
		{
			&trees.Tree{Key: 15},
			[]int{6, 4, 3, 5, 9, 7, 8, 12, 10},
			9,
			"3:4:5:6:7:8:10:12:15:",
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

func TestTreeMin(t *testing.T) {
	tests := []struct {
		root           *trees.Tree
		treeArr        []int
		expectedMinKey int
	}{
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			0,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			2,
		},
	}

	for i, test := range tests {
		root := test.root
		for _, key := range test.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}
		minNone := trees.TreeMin(root)
		if minNone.Key != test.expectedMinKey {
			t.Errorf("test(%d) Expected key  %v, got %v", i, test.expectedMinKey, minNone.Key)
			return
		}
	}
}

func TestTreeMax(t *testing.T) {
	tests := []struct {
		root           *trees.Tree
		treeArr        []int
		expectedMaxKey int
	}{
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			6,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			17,
		},
	}

	for i, test := range tests {
		root := test.root
		for _, key := range test.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}
		maxNode := trees.TreeMax(root)
		if maxNode.Key != test.expectedMaxKey {
			t.Errorf("test(%d) Expected key  %v, got %v", i, test.expectedMaxKey, maxNode.Key)
			return
		}
	}
}

func TestTreeSuccessor(t *testing.T) {
	tests := []struct {
		root        *trees.Tree
		treeArr     []int
		successorOf int
		expectedKey int
	}{
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			5,
			6,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			6,
			8,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			13,
			14,
		},
	}

	for i, test := range tests {
		root := test.root
		for _, key := range test.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}
		successorOfNode := trees.Search(root, test.successorOf)

		node := trees.TreeSuccessor(successorOfNode)
		if node.Key != test.expectedKey {
			t.Errorf("test(%d) Expected key  %v, got %v", i, test.expectedKey, node.Key)
			return
		}
	}
}

func TestTreePredecessor(t *testing.T) {
	tests := []struct {
		root          *trees.Tree
		treeArr       []int
		predecessorOf int
		expectedKey   int
	}{
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			5,
			4,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			6,
			5,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			13,
			12,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			13,
			12,
		},
		{
			&trees.Tree{Key: 5},
			[]int{8, 7, 6},
			6,
			5,
		},
	}

	for i, test := range tests {
		root := test.root
		for _, key := range test.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}
		successorOfNode := trees.Search(root, test.predecessorOf)

		node := trees.TreePredecessor(successorOfNode)
		if node.Key != test.expectedKey {
			t.Errorf("test(%d) Expected key  %v, got %v", i, test.expectedKey, node.Key)
			return
		}
	}
}

func TestIsBST(t *testing.T) {
	tests := []struct {
		root          *trees.Tree
		treeArr       []int
		expectedIsBST bool
	}{
		{
			&trees.Tree{Key: 2},
			[]int{1, 4, 3, 5, 6, 0},
			true,
		},
		{
			&trees.Tree{Key: 14},
			[]int{5, 17, 2, 6, 3, 13, 10, 12, 8},
			true,
		},
	}

	for i, test := range tests {
		root := test.root
		for _, key := range test.treeArr {
			trees.Insert(root, &trees.Tree{Key: key})
		}
		isBST := trees.IsBST(root)
		if isBST != test.expectedIsBST {
			t.Errorf("test(%d) Expected %v, got %v", i, test.expectedIsBST, isBST)
			return
		}
	}
}

func TestNotBST(t *testing.T) {
	root1 := &trees.Tree{Key: 14}
	root1.Left = &trees.Tree{Key: 15}

	root2 := &trees.Tree{Key: 14}
	root2.Right = &trees.Tree{Key: 13}

	root3 := &trees.Tree{Key: 14}
	root3.Left = &trees.Tree{Key: 13}
	root3.Right = &trees.Tree{Key: 15}
	root3.Right.Left = &trees.Tree{Key: 16}

	tests := []struct {
		root          *trees.Tree
		expectedIsBST bool
	}{
		{
			root1,
			false,
		},
		{
			root2,
			false,
		},
		{
			root3,
			false,
		},
	}

	for i, test := range tests {
		root := test.root
		isBST := trees.IsBST(root)
		if isBST != test.expectedIsBST {
			t.Errorf("test(%d) Expected %v, got %v", i, test.expectedIsBST, isBST)
			return
		}
	}
}
