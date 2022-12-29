package dynamic_programming

import (
	"math"
)

func LevenshteinDistance(s1, s2 string) int {
	m := [][]int{}
	for j := 0; j < len(s2)+1; j++ {
		m = append(m, make([]int, len(s1)+1))
	}
	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			if i*j == 0 {
				m[j][i] = i + j
			}
		}
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s2[j-1] {
				m[j][i] = m[j-1][i-1]
			} else {
				m[j][i] = 1 + minThree(m[j-1][i], m[j][i-1], m[j-1][i-1])
			}
		}
	}

	return m[len(s2)][len(s1)]
}

func minThree(a, b, c int) int {
	return int(math.Min(
		math.Min(float64(a), float64(b)),
		float64(c),
	))
}
