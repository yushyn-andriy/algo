package graph

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	n.depthsFirstSearch(&array)
	return array
}

func (n *Node) depthsFirstSearch(array *[]string) {
	*array = append(*array, n.Name)
	for _, node := range n.Children {
		node.depthsFirstSearch(array)
	}
}
