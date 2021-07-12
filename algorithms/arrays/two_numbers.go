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

// TwoNumberSum algoexpert submitted code
func TwoNumberSum(array []int, target int) []int {
	m := make(map[int]bool)
	for _, n := range array {
		m[n] = true
	}

	for _, n := range array {
		diff := target - n
		m[n] = false
		if val, ok := m[diff]; ok && val {
			return []int{n, diff}
		}
		m[n] = true
	}

	return []int{}
}
