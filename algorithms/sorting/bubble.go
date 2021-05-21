package sorting

func BubbleSort(arr []int) []int {
	indexOfLastUnsortedEl := len(arr)

	for indexOfLastUnsortedEl >= 0 {
		indexOfLastUnsortedEl--
		for i := 0; i < indexOfLastUnsortedEl; i++ {
			j := i + 1
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
