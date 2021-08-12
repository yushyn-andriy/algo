package stack

func SunsetViews(buildings []int, direction string) []int {
	buildingsWithSunsetViews := make([]int, 0)

	startIdx := len(buildings) - 1
	step := -1
	if direction == "WEST" {
		startIdx = 0
		step = 1
	}

	idx := startIdx
	runnningMaxHeight := 0
	for idx >= 0 && idx < len(buildings) {
		buildingHeight := buildings[idx]

		if buildingHeight > runnningMaxHeight {
			buildingsWithSunsetViews = append(buildingsWithSunsetViews, idx)
		}

		runnningMaxHeight = max(runnningMaxHeight, buildingHeight)

		idx += step
	}

	if direction == "EAST" {
		reverse(buildingsWithSunsetViews)
	}

	return buildingsWithSunsetViews
}

func reverse(array []int) {
	l := len(array) - 1
	for i := 0; i < len(array)/2; i++ {
		array[i], array[l-i] = array[l-i], array[i]
	}
}
