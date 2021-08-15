package graph

type Node1 struct {
	Name     string
	Children []*Node1
}

func (n *Node1) BreadthFirstSearch(array []string) []string {
	result := []string{}
	visited := make(map[*Node1]bool)
	stack := []*Node1{n}
	for len(stack) > 0 {
		node := Front(&stack)
		if visited[node] {
			continue
		}
		visited[node] = true
		stack = append(stack, node.Children...)
		result = append(result, node.Name)
	}

	return result
}

func Front(x *[]*Node1) *Node1 {
	el := (*x)[0]
	*x = (*x)[1:]
	return el
}
