package dynamic_programming

import "math"

func LevenshteinDistance(s1, s2 string) int {
	return levenshteinDist(s1, s2, len(s1)-1, len(s2)-1)
}

func levenshteinDist(s1, s2 string, i, j int) int {
	if i == 0 && j == 0 {
		return 0
	}
	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}

	if s1[i] == s2[j] {
		return levenshteinDist(s1, s2, i-1, j-1)
	}

	return 1 + minThree(
		levenshteinDist(s1, s2, i-1, j),
		levenshteinDist(s1, s2, i, j-1),
		levenshteinDist(s1, s2, i-1, j-1),
	)
}

func minThree(a, b, c int) int {
	return int(math.Min(
		math.Min(float64(a), float64(b)),
		float64(c),
	))
}
