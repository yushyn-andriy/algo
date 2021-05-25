package arrays

import "fmt"

func LongestPeak(x []int) int {
	largestPeak := 0
	if len(x) < 3 {
		return 0
	}

	peaks := findPeakIndexes(x)
	fmt.Println(peaks)
	for _, peak := range peaks {
		currentPeak := getLeftAvailableSteps(x, peak) + getRightAvailableSteps(x, peak) + 1
		if currentPeak > largestPeak {
			largestPeak = currentPeak
		}
	}

	return largestPeak
}

func findPeakIndexes(x []int) []int {
	peaks := []int{}
	for i := 1; i < len(x)-1; i++ {
		if x[i-1] < x[i] && x[i] > x[i+1] {
			peaks = append(peaks, i)
		}
	}
	return peaks
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
	for i := peakIndex; i < len(x); i++ {
		if x[i+1] < x[i] {
			steps += 1
		} else {
			break
		}
	}
	return steps
}
