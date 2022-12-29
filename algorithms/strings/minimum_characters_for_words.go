package strings

func MinimumCharactersForWords(words []string) []string {
	minimumCharacterSet := make(map[string]int)
	for _, word := range words {
		currentCharacterSet := make(map[string]int)
		for _, ch := range word {
			currentCharacterSet[string(ch)] += 1
		}
		for key, count := range currentCharacterSet {
			minimumCharacterSet[key] = max(minimumCharacterSet[key], count)
		}
	}

	uniqueCharacterSet := []string{}
	for ch, count := range minimumCharacterSet {
		for i := 0; i < count; i++ {
			uniqueCharacterSet = append(uniqueCharacterSet, ch)
		}
	}

	return uniqueCharacterSet
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
