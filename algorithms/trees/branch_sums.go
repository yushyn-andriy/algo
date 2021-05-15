package trees

// Insert ...
func Insert(node *Tree, t *Tree) {
	if node.Key > t.Key {
		if node.Left == nil {
			node.Left = t
			t.Parent = node
		} else {
			Insert(node.Left, t)
		}
	} else {
		if node.Right == nil {
			node.Right = t
			t.Parent = node
		} else {
			Insert(node.Right, t)
		}
	}
}

// Traverse ...
func Traverse(tree *Tree, f func(tree *Tree)) {
	if tree == nil {
		return
	}
	Traverse(tree.Left, f)
	f(tree)
	Traverse(tree.Right, f)
}

// BranchSum ...
func BranchSum(tree *Tree, accumulated int, s *[]int) *[]int {
	if tree == nil {
		return nil
	}

	if tree.Left == nil && tree.Right == nil {
		*s = append(*s, accumulated+tree.Key)
		return nil
	}

	BranchSum(tree.Left, accumulated+tree.Key, s)
	BranchSum(tree.Right, accumulated+tree.Key, s)

	return s
}
