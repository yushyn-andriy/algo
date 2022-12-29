package arrays

func MoveElementToEnd(arr []int, a int) []int {
	l, r := 0, len(arr)-1
	for ; l < r; l++ {
		for ; arr[r] == a && l < r; r-- {
		}
		if arr[l] == a {
			swap(l, r, arr)
		}
	}
	return arr
}

func swap(i, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}
