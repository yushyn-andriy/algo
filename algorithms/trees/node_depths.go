package trees


func NodeDepths(root *Tree) int {
	st := NewTreeStack()
	visited := make(map[*Tree]bool)
	depths := make(map[*Tree]int)

	st.PushTree(root)
	for {
		node, ok := st.PopTree()
		if !ok {
			break
		}

		if v, ok := visited[node]; ok && v {
			continue
		}
		visited[node] = true

		if node.Left != nil {
			st.PushTree(node.Left)
		}
		if node.Right != nil {
			st.PushTree(node.Right)
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
