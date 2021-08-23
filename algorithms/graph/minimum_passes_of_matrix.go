package graph

func MinimumPassesOfMatrix(matrix [][]int) int {
	currentQueue, nextQueue := make([][]int, 0), make([][]int, 0)
	passes := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] > 0 {
				currentQueue = append(currentQueue, []int{i, j})
			}
		}
	}

	for len(currentQueue) > 0 || len(nextQueue) > 0 {
		flipped := false
		for len(currentQueue) > 0 {
			currentPosition := currentQueue[0]
			currentQueue = currentQueue[1:]
			positiveNeighbors := getOrCreatePositiveNeighbors(currentPosition[0], currentPosition[1], matrix)
			nextQueue = append(nextQueue, positiveNeighbors...)
			if len(positiveNeighbors) > 0 {
				flipped = true
			}
		}
		if flipped {
			passes += 1
		}

		currentQueue, nextQueue = nextQueue, currentQueue
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < 0 {
				return -1
			}
		}
	}

	return passes
}

func getOrCreatePositiveNeighbors(i, j int, matrix [][]int) [][]int {
	neighbors := make([][]int, 0)

	if i+1 < len(matrix) && matrix[i+1][j] < 0 {
		matrix[i+1][j] *= -1
		neighbors = append(neighbors, []int{i + 1, j})
	}
	if i-1 >= 0 && matrix[i-1][j] < 0 {
		matrix[i-1][j] *= -1
		neighbors = append(neighbors, []int{i - 1, j})
	}
	if j+1 < len(matrix[i]) && matrix[i][j+1] < 0 {
		matrix[i][j+1] *= -1
		neighbors = append(neighbors, []int{i, j + 1})
	}
	if j-1 >= 0 && matrix[i][j-1] < 0 {
		matrix[i][j-1] *= -1
		neighbors = append(neighbors, []int{i, j - 1})
	}

	return neighbors
}
