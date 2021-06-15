package famous

func Kadanes(x []int) (best int) {
	curr := 0
	for _, v := range x {
		curr = max(0, curr+v)
		best = max(best, curr)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return x
	}
	return y
}
