package trees

import "math"

func ReconstructBst(preOrderTraversalValues []int) *BST {
	// Write your code here.
	return reconstruct(&treeInfo{0}, preOrderTraversalValues, math.MinInt32, math.MaxInt32)
}

func reconstruct(currentSubtreeInfo *treeInfo, preOrderTraversalValues []int, min, max int) *BST {
	if currentSubtreeInfo.rootIdx >= len(preOrderTraversalValues) {
		return nil
	}

	rootValue := preOrderTraversalValues[currentSubtreeInfo.rootIdx]
	if rootValue < min || rootValue >= max {
		return nil
	}

	currentSubtreeInfo.rootIdx += 1

	leftTree := reconstruct(currentSubtreeInfo, preOrderTraversalValues, min, rootValue)
	rightTree := reconstruct(currentSubtreeInfo, preOrderTraversalValues, rootValue, max)

	return &BST{Value: rootValue, Left: leftTree, Right: rightTree}
}

type treeInfo struct {
	rootIdx int
}
