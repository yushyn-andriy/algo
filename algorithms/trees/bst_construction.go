package trees

import "fmt"

func (tree *BST) Insert(value int) *BST {
	insert(tree, value)
	return tree
}

func insert(node *BST, value int) {
	if node.Value > value {
		if node.Left == nil {
			leaf := &BST{
				Value:  value,
				Left:   nil,
				Right:  nil,
				Parent: node,
			}
			node.Left = leaf
		} else {
			insert(node.Left, value)
		}
	} else {
		if node.Right == nil {
			leaf := &BST{
				Value:  value,
				Left:   nil,
				Right:  nil,
				Parent: node,
			}
			node.Right = leaf
		} else {
			insert(node.Right, value)
		}
	}
}

func (tree *BST) Contains(value int) bool {
	return contains(tree, value)
}

func contains(node *BST, value int) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		return true
	}

	if node.Value >= value {
		return contains(node.Left, value)
	} else {
		return contains(node.Right, value)
	}
}
func (tree *BST) Remove(value int) *BST {
	remove(tree, value)
	fmt.Println("remove", tree)
	return tree
}

// Delete ..
func remove(tree *BST, value int) {
	node := search(tree, value)
	if node == nil {
		return
	}

	parent := node.Parent
	if node.Left == nil && node.Right == nil {
		if parent == nil {
			return
		}
		if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
	} else if node.Left == nil && node.Right != nil {
		if parent != nil {
			if parent.Left == node {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		} else {
			node.Value = node.Right.Value
			node.Left = node.Right.Left
			node.Right = node.Right.Right
		}
	} else if node.Left != nil && node.Right == nil {
		if parent != nil {
			if parent.Left == node {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			node.Value = node.Left.Value
			node.Right = node.Left.Right
			node.Left = node.Left.Left
		}
	} else if node.Left != nil && node.Right != nil {
		if node.Right.Left == nil {
			node.Right.Left = node.Left
			node.Value = node.Right.Value
			node.Left = node.Right.Left
			node.Right = node.Right.Right
			if parent != nil {
				if parent.Left == node {
					parent.Left = node.Right
				} else {
					parent.Right = node.Right
				}
			}

		} else {
			next := treeSuccessor(node)
			next.Parent.Left = next.Right
			next.Right = node.Right
			next.Left = node.Left
			node.Value = next.Value
			if parent != nil {
				if parent.Left == node {
					parent.Left = next
				} else {
					parent.Right = next
				}
			}
		}
	}
}

func search(node *BST, value int) *BST {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if node.Value >= value {
		return search(node.Left, value)
	} else {
		return search(node.Right, value)
	}
}

// TreeMin ...
func treeMin(tree *BST) *BST {
	for tree.Left != nil {
		tree = tree.Left
	}
	return tree
}

// TreeSuccessor ...
func treeSuccessor(tree *BST) *BST {
	if tree.Right != nil {
		return treeMin(tree.Right)
	}
	y := tree.Parent
	for y != nil && tree == y.Right {
		tree = y
		y = y.Parent
	}

	return y
}
