package binarytrees

import "math"

type maxPathTreeInfo struct {
	currSum int
	maxSum  int
}

func MaxPathSum(tree *BinaryTree) int {
	return maxPathSum(tree).maxSum
}

func maxPathSum(tree *BinaryTree) *maxPathTreeInfo {
	if tree == nil {
		return &maxPathTreeInfo{0, math.MinInt32}
	}

	leftInfo := maxPathSum(tree.Left)
	rightInfo := maxPathSum(tree.Right)

	maxSum := max(leftInfo.maxSum, rightInfo.maxSum)
	sumThroughtRoot := tree.Value + leftInfo.currSum + rightInfo.currSum

	currSum := tree.Value + max(leftInfo.currSum, rightInfo.currSum)
	maxPathSum := max(max(sumThroughtRoot, currSum), maxSum)

	return &maxPathTreeInfo{currSum, maxPathSum}
}
