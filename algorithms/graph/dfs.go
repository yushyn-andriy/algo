package graph

type Graph struct {
	Nodes map[int][]int
}

func (g *Graph) Add(n int, conns []int) {
	for _, key := range conns {
		g.Nodes[key] = append(g.Nodes[key], n)
		g.Nodes[n] = append(g.Nodes[n], key)
	}
}

func (g *Graph) BFS(source int) []int {
	res := []int{}
	visited := make(map[int]bool)
	seq := []int{source}
	for len(seq) != 0 {
		el := front(&seq)
		if visited[el] {
			continue
		}
		visited[el] = true
		seq = append(seq, g.Nodes[el]...)
		res = append(res, el)
	}
	return res
}

func front(x *[]int) int {
	el := (*x)[0]
	*x = (*x)[1:]
	return el
}
