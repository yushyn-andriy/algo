package sorting

func BubbleSort(arr []int) []int {
	var isSorted bool
	indexOfLastUnsortedEl := len(arr)

Loop:
	for {
		isSorted = true
		indexOfLastUnsortedEl--
		for i := 0; i < indexOfLastUnsortedEl; i++ {
			j := i + 1
			if arr[i] > arr[j] {
				swap(i, j, arr)
				isSorted = false
			}
		}
		if isSorted {
			break Loop
		}
	}

	return arr
}

func swap(i, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}
