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

func TreeMin(tree *Tree) *Tree {
	for tree.Left != nil {
		tree = tree.Left
	}
	return tree
}

func TreeMax(tree *Tree) *Tree {
	for tree.Left != nil {
		tree = tree.Left
	}
	return tree
}

// NOTE: Fix Algorithm
// Delete ..
func Delete(tree *Tree, key int) {
	node := Search(tree, key)
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
	} else if node.Left != nil && node.Right == nil {
		if node.Parent.Left == node {
			node.Parent.Left = node.Left
		} else {
			node.Parent.Right = node.Left
		}
	} else if node.Left == nil && node.Right != nil {
		if node.Parent.Left == node {
			node.Parent.Left = node.Right
		} else {
			node.Parent.Right = node.Right
		}
	} else if node.Left != nil && node.Right != nil {
		if node.Parent.Left == node {
			node.Parent.Left = node.Right
			Insert(node.Parent.Left, node.Left)
		} else {
			node.Parent.Right = node.Right
			Insert(node.Parent.Right, node.Left)
		}
	}

}
