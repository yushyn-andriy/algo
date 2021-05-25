package arrays

import "fmt"

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
	fmt.Println(leftIndex, rightIndex, upIndex, downIndex)
	for leftIndex <= rightIndex && upIndex <= downIndex {
		fmt.Println(leftIndex, rightIndex, upIndex, downIndex)

		fmt.Println("right", leftIndex, rightIndex, upIndex, downIndex)
		right(x, &sequance, leftIndex, rightIndex, upIndex)
		upIndex++

		fmt.Println("down", leftIndex, rightIndex, upIndex, downIndex)
		down(x, &sequance, upIndex, downIndex, rightIndex)
		rightIndex--

		fmt.Println("left", leftIndex, rightIndex, upIndex, downIndex)
		left(x, &sequance, leftIndex, rightIndex, downIndex)
		downIndex--

		fmt.Println("up", leftIndex, rightIndex, upIndex, downIndex)
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
