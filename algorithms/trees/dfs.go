package trees

// DFSSearch ...
func DFSSearch(tree *Tree, key int) *Tree {
	if tree == nil {
		return nil
	}

	if tree.Key == key {
		return tree
	}

	node := DFSSearch(tree.Left, key)
	if node != nil {
		return node
	}
	node = DFSSearch(tree.Right, key)
	if node != nil {
		return node
	}

	return nil
}
