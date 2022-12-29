package trees

import "math"

func (tree *BST) ValidateBst() bool {
	// Write your code here.
	return tree.validateBst(math.MinInt32, math.MaxInt32)
}

func (tree *BST) validateBst(min, max int) bool {
	if tree.Value < min || tree.Value >= max {
		return false
	}

	if tree.Left != nil && !tree.Left.validateBst(min, tree.Value) {
		return false
	}

	if tree.Right != nil && !tree.Right.validateBst(tree.Value, max) {
		return false
	}

	return true
}
