package arrays

import (
	"math"
	"sort"
)

func SmallestDifference(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)

	pair := []int{}
	smallestDiff := math.Inf(1)
	currentDiff := math.Inf(1)

	for i, j := 0, 0; i < len(a) && j < len(b); {
		firstNum, secondNum := a[i], b[j]
		if firstNum < secondNum {
			currentDiff = float64(secondNum - firstNum)
			i++
		} else if firstNum > secondNum {
			currentDiff = float64(firstNum - secondNum)
			j++
		} else {
			return []int{firstNum, secondNum}
		}
		if currentDiff < smallestDiff {
			smallestDiff = currentDiff
			pair = []int{firstNum, secondNum}
		}
	}

	return pair
}
