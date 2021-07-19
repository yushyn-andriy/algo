package binarytrees

type treeInfo struct {
	height     int
	heightDiff int
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	return getBalancedTreeInfo(tree).heightDiff <= 1
}

func getBalancedTreeInfo(tree *BinaryTree) treeInfo {
	if tree == nil {
		return treeInfo{0, 0}
	}

	leftTreeInfo := getBalancedTreeInfo(tree.Left)
	rightTreeInfo := getBalancedTreeInfo(tree.Right)

	currentHeight := 1 + max(leftTreeInfo.height, rightTreeInfo.height)
	currentDiff := abs(leftTreeInfo.height - rightTreeInfo.height)
	maxDiff := max(currentDiff, max(leftTreeInfo.heightDiff, rightTreeInfo.heightDiff))

	return treeInfo{currentHeight, maxDiff}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
