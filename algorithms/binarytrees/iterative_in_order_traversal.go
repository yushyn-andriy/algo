package binarytrees

type BinaryTreeI struct {
	Value int

	Left   *BinaryTreeI
	Right  *BinaryTreeI
	Parent *BinaryTreeI
}

func (tree *BinaryTreeI) IterativeInOrderTraversal(callback func(int)) {
	var previous, next *BinaryTreeI
	current := tree
	for current != nil {
		if previous == nil || previous == current.Parent {
			if current.Left != nil {
				next = current.Left
			} else {
				callback(current.Value)
				if current.Right != nil {
					next = current.Right
				} else {
					next = current.Parent
				}
			}
		} else if previous == current.Left {
			callback(current.Value)
			if current.Right != nil {
				next = current.Right
			} else {
				next = current.Parent
			}
		} else {
			next = current.Parent
		}
		previous, current = current, next
	}
}
