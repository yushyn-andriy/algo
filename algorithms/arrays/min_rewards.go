package arrays

// o(n^2)
func MinRewards(scores []int) int {
	rewards := make([]int, len(scores))
	fill(rewards, 1)
	for i := 1; i < len(scores); i++ {
		j := i - 1
		if scores[i] > scores[j] {
			rewards[i] = rewards[j] + 1
			continue
		}
		for j >= 0 && scores[j] > scores[j+1] {
			rewards[j] = max(rewards[j], rewards[j+1]+1)
			j--
		}
	}

	return sum(rewards)
}

func fill(arr []int, val int) {
	for i := range arr {
		arr[i] = val
	}
}

func sum(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}
