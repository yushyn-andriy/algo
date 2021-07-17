package trees

// O(n^2) time | O(n) space
func RightSmallerThan1(array []int) []int {
	smaller := make([]int, len(array))
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				smaller[i] += 1
			}
		}
	}

	return smaller
}
