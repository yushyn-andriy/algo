package arrays

func LongestPeak(x []int) int {
	largestPeak := 0
	if len(x) < 3 {
		return 0
	}
	for i := 1; i < len(x)-1; i++ {
		if x[i-1] < x[i] && x[i] > x[i+1] {
			currentPeak := getLeftAvailableSteps(x, i) + getRightAvailableSteps(x, i) + 1
			if currentPeak > largestPeak {
				largestPeak = currentPeak
			}
		}
	}
	return largestPeak
}

func getLeftAvailableSteps(x []int, peakIndex int) int {
	steps := 0
	for i := peakIndex; i > 0; i-- {
		if x[i] > x[i-1] {
			steps += 1
		} else {
			break
		}
	}
	return steps
}

func getRightAvailableSteps(x []int, peakIndex int) int {
	steps := 0
	for i := peakIndex; i < len(x)-1; i++ {
		if x[i+1] < x[i] {
			steps += 1
		} else {
			break
		}
	}
	return steps
}
