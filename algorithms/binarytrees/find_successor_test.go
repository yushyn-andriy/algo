package binarytrees

import "testing"

func TestFindSuccessor(t *testing.T) {
	tests := []struct {
		rootValue              int
		treeArray              []int
		nodeValue              int
		expectedSuccessorValue int
	}{
		{
			10,
			[]int{5, 7},
			5,
			7,
		},
		{
			10,
			[]int{5},
			5,
			10,
		},
		{
			10,
			[]int{5, 4},
			4,
			5,
		},
		{
			10,
			[]int{5, 4, 15, 13, 17, 16},
			16,
			17,
		},
	}

	for i, test := range tests {
		root := BinaryTreeSuccessor{Value: test.rootValue}
		for _, value := range test.treeArray {
			root.ConstructBST(value)
		}
		node := root.Find(test.nodeValue)
		if node == nil {
			t.Errorf("test(%d) expected not nil node", i)
			continue
		}

		successor := FindSuccessor(&root, node)
		if successor == nil {
			t.Errorf("test(%d) expected not nil successor", i)
			continue
		}
		if successor.Value != test.expectedSuccessorValue {
			t.Errorf("test(%d) expected %v got %v", i, test.expectedSuccessorValue, node.Value)
		}
	}

}
