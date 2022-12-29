package trees

func InvertBinaryTree(node *Tree) {
	if node == nil {
		return
	}

	swapNodes(node)
	InvertBinaryTree(node.Left)
	InvertBinaryTree(node.Right)
}

func swapNodes(node *Tree) {
	node.Left, node.Right = node.Right, node.Left
}
