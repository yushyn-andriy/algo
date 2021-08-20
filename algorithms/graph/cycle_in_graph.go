package graph

func CycleInGraph1(edges [][]int) bool {
	numberOfNodes := len(edges)
	visited := make([]bool, len(edges))
	currentlyInStack := make([]bool, len(edges))

	for node := 0; node < numberOfNodes; node++ {
		if visited[node] {
			continue
		}

		containsCycle := isNodeInCycle(node, edges, visited, currentlyInStack)
		if containsCycle {
			return true
		}
	}

	return false
}

func isNodeInCycle(node int, edges [][]int, visited []bool, currentlyInStack []bool) bool {
	visited[node] = true
	currentlyInStack[node] = true

	neighbors := edges[node]
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			containsInCycle := isNodeInCycle(neighbor, edges, visited, currentlyInStack)
			if containsInCycle {
				return true
			}
		} else if currentlyInStack[neighbor] {
			return true
		}
	}

	currentlyInStack[node] = false

	return false
}

type Color int

const (
	White Color = 0
	Grey  Color = 1
	Black Color = 2
)

func CycleInGraph2(edges [][]int) bool {
	numberOfNodes := len(edges)
	colors := make([]Color, len(edges))

	for node := 0; node < numberOfNodes; node++ {
		if colors[node] != White {
			continue
		}

		containsCycle := traverseAndColorNodes(node, edges, colors)
		if containsCycle {
			return true
		}
	}

	return false
}

func traverseAndColorNodes(node int, edges [][]int, colors []Color) bool {
	colors[node] = Grey

	neighbors := edges[node]
	for _, neighbor := range neighbors {
		neighborColor := colors[neighbor]

		if neighborColor == Grey {
			return true
		}

		if neighborColor == Black {
			continue
		}

		containsCycle := traverseAndColorNodes(neighbor, edges, colors)
		if containsCycle {
			return true
		}
	}

	colors[node] = Black
	return false
}
