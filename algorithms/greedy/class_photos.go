package greedy

import "sort"


func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Slice(redShirtHeights, func(i, j int) bool {
		return redShirtHeights[i] > redShirtHeights[j]
	})
	sort.Slice(blueShirtHeights, func(i, j int) bool {
		return blueShirtHeights[i] > blueShirtHeights[j]
	})

	shirtColorInFirstRow := "BLUE"
	if redShirtHeights[0] < blueShirtHeights[0] {
		shirtColorInFirstRow ="RED"
	}

	for idx := range redShirtHeights {
		redShirtHeigh := redShirtHeights[idx]
		blueShirtHeight := blueShirtHeights[idx]

		if shirtColorInFirstRow == "RED" {
			if redShirtHeigh >= blueShirtHeight {
				return false
			}
		} else {
			if blueShirtHeight >= redShirtHeigh {
				return false
			}
		}

	}

	return true
}
