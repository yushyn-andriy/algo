package arrays

// TwoNUmbers ...
func TwoNumbers(arr []int, targetSum int) []int {
	hashMap := make(map[int]bool)

	for _, n := range arr {
		hashMap[n] = true
	}

	for _, n := range arr {
		diff := targetSum - n
		hashMap[n] = false
		if val, ok := hashMap[diff]; ok && val {
			return []int{n, diff}
		}
		hashMap[n] = true
	}
	return []int{}
}
