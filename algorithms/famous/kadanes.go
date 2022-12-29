package famous

func KadanesAlgorithm(x []int) (best int) {
	curr := x[0]
	best = x[0]
	for i := 1; i < len(x); i++ {
		v := x[i]
		curr = max(v, curr+v)
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
