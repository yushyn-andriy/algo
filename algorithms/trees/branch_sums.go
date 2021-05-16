package trees


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
