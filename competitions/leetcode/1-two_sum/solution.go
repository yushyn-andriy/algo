package two_sum

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for _, n := range nums {
		m[n] += 1
	}

	for i, n := range nums {
		diff := target - n

		m[n] -= 1
		if val, ok := m[diff]; ok && val > 0 {
			if diff == n {
				return []int{i, findNextAfterIndex(nums, i, n)}
			}
			return []int{i, findNextAfterIndex(nums, 0, diff)}
		}
		m[n] += 1
	}

	return []int{}
}

func findNextAfterIndex(nums []int, idx, target int) int {
	for i := idx + 1; i < len(nums); i++ {
		if nums[i] == target && idx != i {
			return i
		}
	}
	return -1
}
