package trees

func (tree *BST) Insert(value int) *BST {
	tree.insert(value)
	return tree
}

func (tree *BST) insert(value int) {
	if tree.Value > value {
		if tree.Left == nil {
			leaf := &BST{
				Value: value,
				Left:  nil,
				Right: nil,
			}
			tree.Left = leaf
		} else {
			tree.Left.insert(value)
		}
	} else {
		if tree.Right == nil {
			leaf := &BST{
				Value: value,
				Left:  nil,
				Right: nil,
			}
			tree.Right = leaf
		} else {
			tree.Right.insert(value)
		}
	}
}

func (tree *BST) Contains(value int) bool {
	return tree.contains(value)
}

func (tree *BST) contains(value int) bool {
	if tree == nil {
		return false
	}

	if tree.Value == value {
		return true
	}

	if tree.Value >= value {
		return tree.Left.contains(value)
	} else {
		return tree.Right.contains(value)
	}
}
func (tree *BST) Remove(value int) *BST {
	tree.remove(value, nil)
	return tree
}

// Delete ..
func (tree *BST) remove(value int, parent *BST) {
	if value < tree.Value {
		if tree.Left != nil {
			tree.Left.remove(value, tree)
		}
	} else if value > tree.Value {
		if tree.Right != nil {
			tree.Right.remove(value, tree)
		}
	} else {
		if tree.Left != nil && tree.Right != nil {
			tree.Value = tree.Right.getMinValue()
			tree.Right.remove(tree.Value, tree)
		} else if parent == nil {
			if tree.Left != nil {
				tree.Value = tree.Left.Value
				tree.Right = tree.Left.Right
				tree.Left = tree.Left.Left
			} else if tree.Right != nil {
				tree.Value = tree.Right.Value
				tree.Right = tree.Right.Right
				tree.Left = tree.Right.Left
			} else {
				// do nothing
			}
		} else if parent.Left == tree {
			if tree.Left != nil {
				parent.Left = tree.Left
			} else {
				parent.Left = tree.Right
			}
		} else if parent.Right == tree {
			if tree.Left != nil {
				parent.Right = tree.Left
			} else {
				parent.Right = tree.Right
			}
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}

	return tree.Left.getMinValue()
}
