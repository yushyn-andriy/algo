package sorting


func InsertionSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		j := i + 1

		c := arr[j]
		for ; j > 0 && c < arr[j - 1]; j-- {
			arr[j] = arr[j - 1]
		}
		arr[j] = c
	}
}

func RecursiveInsertionSort(arr []int) {
	if len(arr) > 2 {
		RecursiveInsertionSort(arr[:len(arr) - 1])
	}

	j := len(arr) - 1
	c := arr[j]
	for ; j > 0 && c < arr[j - 1]; j-- {
		arr[j] = arr[j - 1]
	}
	arr[j] = c
}