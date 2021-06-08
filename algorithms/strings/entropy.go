package strings

import (
	"math"
)

func Entropy(s string) float64 {
	var e float64
	for _, pi := range CalculateRunesProbability(CountRunes(s), len(s)) {
		e += -pi * math.Log2(pi)
	}

	return math.Round(e*100) / 100
}

func CountRunes(s string) map[rune]int {
	c := make(map[rune]int)
	for _, r := range s {
		c[r] += 1
	}
	return c
}

func CalculateRunesProbability(c map[rune]int, l int) map[rune]float64 {
	p := make(map[rune]float64, len(c))
	for k, v := range c {
		p[k] = float64(v) / float64(l)
	}

	return p
}
