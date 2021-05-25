package arrays

func SpiralTraverse(x [][]int) []int {
	if len(x) == 0 {
		return []int{}
	}
	if len(x) == 1 {
		return x[0]
	}

	sequance := []int{}

	leftIndex, rightIndex := 0, len(x[0])-1
	upIndex, downIndex := 0, len(x)-1
	for leftIndex <= rightIndex && upIndex <= downIndex {
		right(x, &sequance, leftIndex, rightIndex, upIndex)
		upIndex++

		down(x, &sequance, upIndex, downIndex, rightIndex)
		rightIndex--

		left(x, &sequance, leftIndex, rightIndex, downIndex)
		downIndex--

		up(x, &sequance, upIndex, downIndex, leftIndex)
		leftIndex++
	}

	return sequance
}

func right(x [][]int, sequance *[]int, leftIndex, rightIndex, rowIndex int) {
	for ; leftIndex <= rightIndex; leftIndex++ {
		*sequance = append(*sequance, x[rowIndex][leftIndex])
	}

}
func down(x [][]int, sequance *[]int, upIndex, downIndex, colIndex int) {
	for ; upIndex <= downIndex; upIndex++ {
		*sequance = append(*sequance, x[upIndex][colIndex])
	}
}
func left(x [][]int, sequance *[]int, leftIndex, rightIndex, rowIndex int) {
	for ; leftIndex <= rightIndex; rightIndex-- {
		*sequance = append(*sequance, x[rowIndex][rightIndex])
	}
}
func up(x [][]int, sequance *[]int, upIndex, downIndex, colIndex int) {
	for ; upIndex <= downIndex; downIndex-- {
		*sequance = append(*sequance, x[downIndex][colIndex])
	}
}
