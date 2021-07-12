package arrays

func SortedSquaredArray1(array []int) []int {
	out := []int{}

	negative, positive := []int{}, []int{}
	for _, x := range array {
		sqr := x * x
		if x < 0 {
			negative = append(negative, sqr)
		} else {
			positive = append(positive, sqr)
		}
	}

	i, j := len(negative)-1, 0
Loop:
	for {
		if i >= 0 && j < len(positive) {
			n, p := negative[i], positive[j]
			if n < p {
				out = append(out, n)
				i--
			} else if n > p {
				out = append(out, p)
				j++
			} else {
				out = append(out, p)
				out = append(out, n)
				i--
				j++
			}
		} else {
			break Loop
		}
	}
	if i >= 0 {
		for ; i >= 0; i-- {
			out = append(out, negative[i])
		}
	}
	if j < len(positive) {
		out = append(out, positive[j:]...)
	}

	return out
}

func SortedSquaredArray(array []int) []int {
	sortedSquares := make([]int, len(array))

	smallerValueIdx := 0
	largerValueIdx := len(array) - 1

	for idx := len(array) - 1; idx >= 0; idx-- {
		smallerValue := array[smallerValueIdx]
		largerValue := array[largerValueIdx]

		if abs(smallerValue) > abs(largerValue) {
			sortedSquares[idx] = smallerValue * smallerValue
			smallerValueIdx += 1
		} else {
			sortedSquares[idx] = largerValue * largerValue
			largerValueIdx -= 1
		}
	}

	return sortedSquares
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
