package trees

type Stack struct {
	data []*Tree
}

func (s *Stack) Push(v *Tree) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (*Tree, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, true
}

func NewStack() *Stack {
	return &Stack{}
}

func NodeDepths(root *Tree) int {
	st := NewStack()
	visited := make(map[*Tree]bool)
	depths := make(map[*Tree]int)

	st.Push(root)
	for {
		node, ok := st.Pop()
		if !ok {
			break
		}

		if v, ok := visited[node]; ok && v {
			continue
		}
		visited[node] = true

		if node.Left != nil {
			st.Push(node.Left)
		}
		if node.Right != nil {
			st.Push(node.Right)
		}
		if node == root {
			depths[node] = 0
		} else {
			depths[node] = depths[node.Parent] + 1
		}
	}

	depthsSum := 0
	for _, value := range depths {
		depthsSum += value
	}

	return depthsSum
}
