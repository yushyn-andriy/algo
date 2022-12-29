package stack

func LargestRectangleUnderSkyline(buildings []int) int {
	maxArea := 0
	for pillarIdx := range buildings {
		currentHeight := buildings[pillarIdx]

		var furthestLeft = pillarIdx
		for furthestLeft > 0 && buildings[furthestLeft-1] >= currentHeight {
			furthestLeft -= 1
		}

		var furthestRight = pillarIdx
		for furthestRight < len(buildings)-1 && buildings[furthestRight+1] >= currentHeight {
			furthestRight += 1
		}

		areaWithCurrentBuilding := (furthestRight - furthestLeft + 1) * currentHeight
		maxArea = max(areaWithCurrentBuilding, maxArea)
	}

	return maxArea
}
