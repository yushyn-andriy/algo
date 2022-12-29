package strings

import "math"

func SmallestSubstringContaining(bigString, smallString string) string {
	targetCharCounts := getCharCounts(smallString)
	substringBounds := getSubstringBounds(bigString, targetCharCounts)
	return getStringFromBounds(bigString, substringBounds)
}

func getCharCounts(str string) map[byte]int {
	charCounts := map[byte]int{}
	for _, ch := range str {
		charCounts[byte(ch)] += 1
	}
	return charCounts
}

func getSubstringBounds(str string, targetCharCounts map[byte]int) []int {
	substringBounds := []int{0, math.MaxInt32}
	substringCharCounts := map[byte]int{}
	numUniqueChars := len(targetCharCounts)
	numUniqueCharsDone := 0
	leftIdx := 0
	rightIdx := 0

	for rightIdx < len(str) {
		rightChar := str[rightIdx]
		if _, found := targetCharCounts[rightChar]; !found {
			rightIdx++
			continue
		}

		substringCharCounts[rightChar]++
		if substringCharCounts[rightChar] == targetCharCounts[rightChar] {
			numUniqueCharsDone++
		}

		for numUniqueCharsDone == numUniqueChars && leftIdx <= rightIdx {
			substringBounds = getCloserBounds(leftIdx, rightIdx, substringBounds[0], substringBounds[1])
			leftChar := str[leftIdx]
			if _, found := targetCharCounts[leftChar]; !found {
				leftIdx++
				continue
			}
			if substringCharCounts[leftChar] == targetCharCounts[leftChar] {
				numUniqueCharsDone--
			}

			substringCharCounts[leftChar]--
			leftIdx++
		}
		rightIdx++
	}
	return substringBounds
}

func getCloserBounds(idx1, idx2, idx3, idx4 int) []int {
	if idx2-idx1 < idx4-idx3 {
		return []int{idx1, idx2}
	}
	return []int{idx3, idx4}
}

func getStringFromBounds(str string, bounds []int) string {
	start, end := bounds[0], bounds[1]
	if end == math.MaxInt32 {
		return ""
	}
	return str[start : end+1]
}
