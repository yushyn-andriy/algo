package arrays

func IsMonotonic(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}
	direction := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if direction == 0 {
			direction = arr[i+1] - arr[i]
			continue
		}

		diff := arr[i+1] - arr[i]
		if diff > 0 && direction < 0 || diff < 0 && direction > 0 {
			return false
		}
	}
	return true
}
