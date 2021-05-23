package arrays

func MoveElementToEnd(arr []int, a int) []int {
	l, r := 0, len(arr)-1

	for ; l < r; l++ {
		leftElement := arr[l]
		rightElement := arr[r]
		if leftElement == a {
			for rightElement == a {
				r--
				rightElement = arr[r]
			}
			if l < r {
				swap(l, r, arr)
			}
		}
	}

	return arr
}

func swap(i, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}
