package trees

func SameBsts(arrayOne, arrayTwo []int) bool {
	if len(arrayOne) != len(arrayTwo) {
		return false
	}
	if len(arrayOne) == 0 && len(arrayTwo) == 0 {
		return true
	}
	if arrayOne[0] != arrayTwo[0] {
		return false
	}

	leftOne := getSmaller(arrayOne)
	leftTwo := getSmaller(arrayTwo)
	rightOne := getBiggerOrEqual(arrayOne)
	rightTwo := getBiggerOrEqual(arrayTwo)

	return SameBsts(leftOne, leftTwo) && SameBsts(rightOne, rightTwo)
}

func getSmaller(array []int) []int {
	smaller := []int{}
	for i, v := range array {
		if i == 0 {
			continue
		}
		if v < array[0] {
			smaller = append(smaller, v)
		}
	}

	return smaller
}

func getBiggerOrEqual(array []int) []int {
	bigger := []int{}
	for i, v := range array {
		if i == 0 {
			continue
		}
		if v >= array[0] {
			bigger = append(bigger, v)
		}
	}

	return bigger
}
