package graph

// O(n) time | O(1) space
func SingleCycleCheck(x []int) bool {
	if len(x) == 0 {
		return true
	}
	var (
		visited   int
		currIndex int
		currSteps int = x[currIndex]
	)
	for visited < len(x) {
		if visited > 0 && currIndex == 0 {
			return false
		}

		visited += 1
		currIndex = getNextIndex(currIndex, currSteps, len(x))
		currSteps = x[currIndex]
	}

	return currIndex == 0
}

func getNextIndex(currIndex, currSteps, xLen int) int {
	currIndex = (currIndex + currSteps) % xLen
	if currIndex < 0 {
		currIndex = xLen + currIndex
	}
	return currIndex
}
