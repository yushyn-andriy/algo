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
