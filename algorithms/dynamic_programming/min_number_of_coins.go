package dynamic_programming

import "math"

func MinNumberOfCoins(amount int, coins []int) int {
	r := int(minNumberOfCoins(amount, coins))
	if r < 0 {
		return 0
	}
	return r
}

func minNumberOfCoins(amount int, coins []int) float64 {
	if amount < 0 {
		return math.Inf(1)
	} else if amount == 0 {
		return 0
	}

	m := math.Inf(1)
	currWay := math.Inf(1)
	for _, coin := range coins {
		currWay = math.Min(
			currWay,
			math.Min(
				m,
				minNumberOfCoins(amount-coin, coins)+1,
			))
	}

	return currWay
}

func min(array []float64) float64 {
	m := math.Inf(1)
	for _, v := range array {
		if v < m {
			m = v
		}

	}
	return m
}
