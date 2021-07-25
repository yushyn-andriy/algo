package strings

// O(n) time | O(1) space - because hash table will never have more than 26 character frequance
func FirstNonRepeatingCharacterSolution1(str string) int {
	m := make(map[rune]int)
	for _, v := range str {
		m[v] += 1
	}
	for i, v := range str {
		count := m[v]
		if count < 2 {
			return i
		}
	}

	return -1
}
