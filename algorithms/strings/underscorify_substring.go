package strings

type interval struct {
	left, right int
}

func UnderscorifySubstring(str string, substring string) string {
	locations := []interval{}
	subLen := len(substring)
	for start := 0; start < len(str); start++ {
		end := start + subLen
		if end > len(str) {
			break
		}
		if str[start:end] == substring {
			locations = append(locations, interval{start, end})
		}
	}

	collapsedLocations := []*interval{}
	collapsedIdx := 0
	for i := 0; i < len(locations) && len(locations) > 0; i++ {
		currentInterval := locations[i]
		if len(collapsedLocations) == 0 {
			collapsedLocations = append(collapsedLocations, &currentInterval)
			continue
		}

		currentCollapsedInterval := collapsedLocations[collapsedIdx]
		if isOverlap(currentCollapsedInterval, &currentInterval) {
			currentCollapsedInterval.right = currentInterval.right
			continue
		}

		collapsedLocations = append(collapsedLocations, &currentInterval)
		collapsedIdx += 1
	}

	if len(locations) == 0 {
		return str
	}

	result := make([]rune, len(str)+2*len(collapsedLocations))
	resultIndex := 0
	locationIndex := 0
	for i, r := range str {
		location := collapsedLocations[locationIndex]
		if i == location.left {
			result[resultIndex] = '_'
			resultIndex += 1
		} else if i == location.right {
			result[resultIndex] = '_'
			resultIndex += 1
			if locationIndex+1 < len(collapsedLocations) {
				locationIndex += 1
			}
		}
		result[resultIndex] = r
		resultIndex += 1
	}

	if collapsedLocations[locationIndex].right == len(str) {
		result[len(result)-1] = '_'
	}

	return string(result)
}

func isOverlap(i1, i2 *interval) bool {
	return i1.right == i2.left+1 || i1.right == i2.left || i1.right+1 == i2.right
}
