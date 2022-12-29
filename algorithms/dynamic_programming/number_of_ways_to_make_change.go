package dynamic_programming

func NumberOfWays(amount int, coins []int) int {
	if amount <= 0 {
		return 0
	}
	if len(coins) == 0 {
		return 1
	}

	ways := make([]int, amount+1)
	ways[0] = 1

	for _, coin := range coins {
		if coin <= amount {
			for a := coin; a < len(ways); a++ {
				ways[a] += ways[a-coin]
			}
		}
	}

	return ways[amount]
}
