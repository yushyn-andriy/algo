package stack

func NextGreaterElement(array []int) []int {
	result := make([]int, 0)
	for range array {
		// {-1, -1, -1, -1, -1}
		result = append(result, -1)
	}
	stack := make([]int, 0)

	for idx := 0; idx < 2*len(array); idx++ {
		circularIdx := idx % len(array)

		for len(stack) > 0 && array[stack[len(stack)-1]] < array[circularIdx] {
			var top int
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			result[top] = array[circularIdx]
		}

		stack = append(stack, circularIdx)
	}

	return result
}
