package greedy

import "sort"

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)

	if !fastest {
		reverseArrayInPlace(redShirtSpeeds)
	}

	totalSpeed := 0
	for idx := range redShirtSpeeds {
		rider1 := redShirtSpeeds[idx]
		rider2 := blueShirtSpeeds[len(blueShirtSpeeds)-idx-1]
		totalSpeed += max(rider1, rider2)
	}

	return totalSpeed
}

func reverseArrayInPlace(array []int) {
	var start = 0
	var end = len(array) - 1
	for start < end {
		temp := array[start]
		array[start] = array[end]
		array[end] = temp
		start += 1
		end -= 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
