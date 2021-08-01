package arrays

// o(n) time | o(1) space
func FindPeakLinear(array []int) int {
	for i := 1; i < len(array); i++ {
		if i == len(array)-1 {
			if array[i] >= array[i-1] {
				return i
			}
		} else {
			if array[i] >= array[i-1] && array[i] >= array[i+1] {
				return i
			}
		}
	}

	return -1
}
