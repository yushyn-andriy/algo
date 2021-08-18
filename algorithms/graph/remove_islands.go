package graph

func RemoveIslands(matrix [][]int) [][]int {
	onesConnectedToBorder := make([][]bool, len(matrix))
	for i := range matrix {
		onesConnectedToBorder[i] = make([]bool, len(matrix[i]))
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			rowIsBorder := row == 0 || row == len(matrix)-1
			colIsBorder := col == 0 || col == len(matrix[row])-1
			isBorder := rowIsBorder || colIsBorder
			if !isBorder {
				continue
			}

			if matrix[row][col] != 1 {
				continue
			}

			findOnesConnectedToBorder(matrix, row, col, onesConnectedToBorder)
		}
	}

	for row := 0; row < len(matrix)-1; row++ {
		for col := 0; col < len(matrix[row])-1; col++ {
			if onesConnectedToBorder[row][col] {
				continue
			}
			matrix[row][col] = 0
		}
	}

	return matrix
}

func findOnesConnectedToBorder(matrix [][]int, startRow, startCol int, onesConnectedToBorder [][]bool) {
	stack := [][]int{{startRow, startCol}}

	var currentPosition []int
	for len(stack) > 0 {
		currentPosition, stack = stack[len(stack)-1], stack[:len(stack)-1]
		currentRow, currentCol := currentPosition[0], currentPosition[1]

		alreadyVisited := onesConnectedToBorder[currentRow][currentCol]
		if alreadyVisited {
			continue
		}

		onesConnectedToBorder[currentRow][currentCol] = true

		neighbors := getNeighbors(matrix, currentRow, currentCol)
		for _, neighbor := range neighbors {
			row, col := neighbor[0], neighbor[1]

			if matrix[row][col] != 1 {
				continue
			}
			stack = append(stack, neighbor)
		}
	}
}

func getNeighbors(matrix [][]int, row, col int) [][]int {
	neighbors := make([][]int, 0)
	numRows := len(matrix)
	numCols := len(matrix[row])

	if row-1 >= 0 {
		neighbors = append(neighbors, []int{row - 1, col})
	}
	if row+1 < numRows {
		neighbors = append(neighbors, []int{row + 1, col})
	}
	if col-1 >= 0 {
		neighbors = append(neighbors, []int{row, col - 1})
	}
	if col+1 < numCols {
		neighbors = append(neighbors, []int{row, col + 1})
	}

	return neighbors
}
