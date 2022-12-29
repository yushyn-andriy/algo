package binarytrees

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

type TreeInfo struct {
	diameter int
	height   int
}

// O(N) time | O(h) space(worst O(N) space)
func BinaryTreeDiameter(tree *BinaryTree) int {
	return getTreeInfo(tree).diameter
}

func getTreeInfo(tree *BinaryTree) TreeInfo {
	if tree == nil {
		return TreeInfo{0, 0}
	}

	leftTreeInfo := getTreeInfo(tree.Left)
	rigthTreeInfo := getTreeInfo(tree.Right)

	longestPathThroughRoot := leftTreeInfo.height + rigthTreeInfo.height
	currentDiameter := max(longestPathThroughRoot, max(leftTreeInfo.diameter, rigthTreeInfo.diameter))
	currentHeight := 1 + max(leftTreeInfo.height, rigthTreeInfo.height)

	return TreeInfo{currentDiameter, currentHeight}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
