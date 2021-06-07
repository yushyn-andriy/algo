package dynamic_programming

import (
	"math"
)

func MaxSubsetSumNoAdjacent(x []int) int {
	if len(x) < 1 {
		return 0
	}
	if len(x) == 1 {
		return x[0]
	}
	if len(x) == 2 {
		return int(math.Max(float64(x[0]), float64(x[1])))
	}

	return maxSubsetSum(x)
}

func maxSubsetSum(x []int) int {
	subsetSumArr := make([]float64, len(x))
	subsetSumArr[0] = float64(x[0])
	subsetSumArr[1] = math.Max(float64(x[0]), float64(x[1]))
	for i := 2; i < len(subsetSumArr); i++ {
		subsetSumArr[i] = math.Inf(-1)
	}

	for i := 2; i < len(subsetSumArr); i++ {
		subsetSumArr[i] = subsetSumsElement(x, subsetSumArr, i)
	}

	m := math.Inf(-1)
	for _, val := range subsetSumArr {
		if val > m {
			m = val
		}
	}
	return int(m)
}

func subsetSumsElement(x []int, s []float64, i int) float64 {
	if s[i] > math.Inf(-1) {
		return s[i]
	}
	return math.Max(
		subsetSumsElement(x, s, i-1),
		subsetSumsElement(x, s, i-2)+float64(x[i]),
	)
}
