package trees

// ========================================================
// Stack data structures                                  =
// ========================================================

// TreeStack ...
type TreeStack struct {
	data []*Tree
}

func (s *TreeStack) PushTree(v *Tree) {
	s.data = append(s.data, v)
}

func (s *TreeStack) PopTree() (*Tree, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, true
}

func NewTreeStack() *TreeStack {
	return &TreeStack{}
}

// ========================================================
// Tree data structures                                   =
// ========================================================

//  Tree ...
type Tree struct {
	Key, Value int
	Parent     *Tree
	Left       *Tree
	Right      *Tree
}

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

// InOrderWalk ...
func InOrderWalk(tree *Tree, f func(tree *Tree)) {
	if tree == nil {
		return
	}
	InOrderWalk(tree.Left, f)
	f(tree)
	InOrderWalk(tree.Right, f)
}

// PreOrderWalk ...
func PreOrderWalk(tree *Tree, f func(tree *Tree)) {
	if tree == nil {
		return
	}
	f(tree)
	PreOrderWalk(tree.Left, f)
	PreOrderWalk(tree.Right, f)
}

// PostOrderWalk ...
func PostOrderWalk(tree *Tree, f func(tree *Tree)) {
	if tree == nil {
		return
	}
	PostOrderWalk(tree.Left, f)
	PostOrderWalk(tree.Right, f)
	f(tree)
}

// Search ...
func Search(tree *Tree, key int) *Tree {
	if tree == nil {
		return nil
	}

	if tree.Key == key {
		return tree
	}

	if tree.Key >= key {
		return Search(tree.Left, key)
	} else {
		return Search(tree.Right, key)
	}
}

// TreeMin ...
func TreeMin(tree *Tree) *Tree {
	for tree.Left != nil {
		tree = tree.Left
	}
	return tree
}

// TreeMax ...
func TreeMax(tree *Tree) *Tree {
	for tree.Right != nil {
		tree = tree.Right
	}
	return tree
}

// TreeSuccessor ...
func TreeSuccessor(tree *Tree) *Tree {
	if tree.Right != nil {
		return TreeMin(tree.Right)
	}
	y := tree.Parent
	for y != nil && tree == y.Right {
		tree = y
		y = y.Parent
	}

	return y
}

// TreePredecessor ...
func TreePredecessor(tree *Tree) *Tree {
	if tree.Left != nil {
		return TreeMax(tree.Left)
	}

	y := tree.Parent
	for y != nil && tree == y.Left {
		tree = y
		y = y.Parent
	}

	return y
}

// Delete ..
func Delete(tree *Tree, key int) {
	node := Search(tree, key)
	if node == nil {
		return
	}

	parent := node.Parent
	if node.Left == nil && node.Right == nil {
		if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
	} else if node.Left == nil && node.Right != nil {
		if parent.Left == node {
			parent.Left = node.Right
		} else {
			parent.Right = node.Right
		}
	} else if node.Left != nil && node.Right == nil {
		if parent.Left == node {
			parent.Left = node.Left
		} else {
			parent.Right = node.Left
		}
	} else if node.Left != nil && node.Right != nil {
		if node.Right.Left == nil {
			if parent.Left == node {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
			node.Right.Left = node.Left
		} else {
			next := TreeSuccessor(node)
			next.Parent.Left = next.Right
			next.Right = node.Right
			next.Left = node.Left
			if parent.Left == node {
				parent.Left = next
			} else {
				parent.Right = next
			}
		}
	}
}

// IsBST ...
func IsBST(root *Tree) bool {
	stack := []*Tree{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)

		parent := node.Parent
		if parent != nil {
			if parent.Parent != nil {
				if isLeftNode(parent) && isRightNode(node) {
					if node.Key >= parent.Key && node.Key < parent.Parent.Key {
						continue
					} else {
						return false
					}
				} else if isRightNode(parent) && isLeftNode(node) {
					if node.Key < parent.Key && node.Key >= parent.Parent.Key {
						continue
					} else {
						return false
					}
				} else if isLeftNode(parent) && isLeftNode(node) {
					if node.Key < parent.Key && parent.Key < parent.Parent.Key {
						continue
					} else {
						return false
					}
				} else if isRightNode(parent) && isRightNode(node) {
					if node.Key >= parent.Key && parent.Key >= parent.Parent.Key {
						continue
					} else {
						return false
					}
				}
			}
		}

		if node.Right != nil && node.Key > node.Right.Key {
			return false
		}
		if node.Left != nil && node.Key <= node.Left.Key {
			return false
		}
	}

	return true
}

func isLeftNode(node *Tree) bool {
	if node.Parent.Left == node {
		return true
	}
	return false
}

func isRightNode(node *Tree) bool {
	if node.Parent.Right == node {
		return true
	}
	return false
}
