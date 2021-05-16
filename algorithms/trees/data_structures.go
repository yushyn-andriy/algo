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

// InOrderWalk ...
func PreOrderWalk(tree *Tree, f func(tree *Tree)) {
	if tree == nil {
		return
	}
	f(tree)
	PreOrderWalk(tree.Left, f)
	PreOrderWalk(tree.Right, f)
}

// InOrderWalk ...
func PostOrderWalk(tree *Tree, f func(tree *Tree)) {
	if tree == nil {
		return
	}
	PostOrderWalk(tree.Right, f)
	f(tree)
	PostOrderWalk(tree.Left, f)
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
