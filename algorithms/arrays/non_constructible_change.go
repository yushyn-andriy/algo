package arrays

import "sort"

func NonConstructibleChange(coins []int) int {
	if len(coins) == 0 {
		return 1
	}
	sort.Ints(coins)

	change := coins[0]
	for i, coin := range coins {
		if i == 0 {
			if change > 1 {
				return 1
			}
			continue
		}

		if coin > change+1 {
			return change + 1
		} else {
			change += coin
		}
	}

	return change + 1
}
