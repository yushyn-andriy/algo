package arrays

import "math"

type Block map[string]bool

func ApartmentHunting(blocks []Block, reqs []string) int {
	maxDistancesAtBlocks := make([]int, len(blocks))
	for i := range blocks {
		maxDistancesAtBlocks[i] = -1
		for _, req := range reqs {
			closestReqDistance := math.MaxInt32
			for j := range blocks {
				if blocks[j][req] {
					closestReqDistance = min1(closestReqDistance, distanceBetween(i, j))
				}
			}
			maxDistancesAtBlocks[i] = max1(maxDistancesAtBlocks[i], closestReqDistance)
		}
	}

	var optimalBlockIdx int
	smallestMaxDistance := math.MaxInt32
	for i, currentDistance := range maxDistancesAtBlocks {
		if currentDistance < smallestMaxDistance {
			smallestMaxDistance = currentDistance
			optimalBlockIdx = i
		}
	}
	return optimalBlockIdx
}

func distanceBetween(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
