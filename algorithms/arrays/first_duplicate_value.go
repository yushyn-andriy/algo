package arrays

func FirstDuplicateValue(array []int) int {
	for idx := range array {
		value := Abs(array[idx])
		if array[value-1] < 0 {
			return value
		} else {
			array[value-1] *= -1
		}
	}

	return -1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
