package trees

func (tree *BST) InOrderTraverse(array []int) []int {
	tree.inOrderTraverse(&array)
	return array
}

func (tree *BST) inOrderTraverse(x *[]int) {
	if tree == nil {
		return
	}
	tree.Left.inOrderTraverse(x)
	*x = append(*x, tree.Value)
	tree.Right.inOrderTraverse(x)
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	tree.preOrderTraverse(&array)
	return array
}

func (tree *BST) preOrderTraverse(x *[]int) {
	if tree == nil {
		return
	}
	*x = append(*x, tree.Value)
	tree.Left.preOrderTraverse(x)
	tree.Right.preOrderTraverse(x)
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	tree.postOrderTraverse(&array)
	return array
}

func (tree *BST) postOrderTraverse(x *[]int) {
	if tree == nil {
		return
	}
	tree.Left.postOrderTraverse(x)
	tree.Right.postOrderTraverse(x)
	*x = append(*x, tree.Value)
}
