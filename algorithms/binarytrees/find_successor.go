package binarytrees

type BinaryTreeSuccessor struct {
	Value int

	Left   *BinaryTreeSuccessor
	Right  *BinaryTreeSuccessor
	Parent *BinaryTreeSuccessor
}

func (tree *BinaryTreeSuccessor) Find(value int) *BinaryTreeSuccessor {
	if tree == nil {
		return nil
	}

	if value == tree.Value {
		return tree
	}

	if value >= tree.Value {
		return tree.Right.Find(value)
	}
	return tree.Left.Find(value)
}

func (tree *BinaryTreeSuccessor) ConstructBST(value int) {
	if value >= tree.Value {
		if tree.Right != nil {
			tree.Right.ConstructBST(value)
		} else {
			tree.Right = &BinaryTreeSuccessor{
				Value:  value,
				Parent: tree,
			}
		}
	} else {
		if tree.Left != nil {
			tree.Left.ConstructBST(value)
		} else {
			tree.Left = &BinaryTreeSuccessor{
				Value:  value,
				Parent: tree,
			}
		}
	}
}

func FindSuccessor(tree *BinaryTreeSuccessor, node *BinaryTreeSuccessor) *BinaryTreeSuccessor {
	if node.Parent == nil && node.Right == nil {
		return nil
	}

	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}

	if node.Parent.Left == node {
		return node.Parent
	}

	if node.Parent.Right == node {
		parent := node.Parent
		for parent != nil && parent.Right == node {
			node = parent
			parent = parent.Parent
		}

		return parent
	}

	return node
}
