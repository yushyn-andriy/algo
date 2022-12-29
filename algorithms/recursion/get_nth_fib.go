package recursion

import "fmt"

func GetNthFib(n int) int {
	return getNthFib(n, map[int]int{1: 0, 2: 1})
}

func getNthFib(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		fmt.Println(n, "ok")
		return v
	}

	v := getNthFib(n-1, cache) + getNthFib(n-2, cache)
	cache[n] = v

	return v

}

func GetNthFibIterative(n int) int {
	x, y := 0, 1

	if n < 2 {
		return x
	}

	for i := 2; i < n; i++ {
		x, y = y, x+y
	}

	return y
}
