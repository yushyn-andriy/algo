package reverse

import (
	"math"
	"strconv"
)

// function, which takes a string as
// argument and return the reverse of string.
func reverseStr(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func reverse(x int) int {
	str := strconv.Itoa(abs(x))
	res := reverseStr(str)
	out, _ := strconv.Atoi(res)
	if out > math.MaxInt32 {
		return 0
	}
	if x < 0 {
		out = out * -1
	}

	return out
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
