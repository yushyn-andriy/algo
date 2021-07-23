package recursion

func StaircaseTraversal(height int, maxSteps int) int {
	// Write your code here.
	return staircaseTraversal(height, maxSteps, 0)
}

func staircaseTraversal(height int, maxSteps int, step int) int {
	diff := height - step
	if diff == 0 {
		return 1
	} else if diff < 0 {
		return -1
	}

	steps := 0
	for currStep := 1; currStep <= maxSteps; currStep++ {
		s := staircaseTraversal(height, maxSteps, currStep+step)
		if s > 0 {
			steps += s
		}
	}

	return steps
}

func StaircaseTraversalSolution2(height int, maxSteps int) int {
	return staircaseTraversalSolution2(height, maxSteps, map[int]int{0: 1, 1: 1})
}

func staircaseTraversalSolution2(height int, maxSteps int, memoize map[int]int) int {
	if ways, found := memoize[height]; found {
		return ways
	}

	numberOfWays := 0
	for step := 1; step < min(maxSteps, height)+1; step++ {
		numberOfWays += staircaseTraversalSolution2(height-step, maxSteps, memoize)
	}
	memoize[height] = numberOfWays

	return numberOfWays
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
