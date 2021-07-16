package strings

func GenerateDocument(characters string, document string) bool {
	m := make(map[rune]int)
	for _, r := range characters {
		m[r] += 1
	}

	for _, r := range document {
		if countLeft, ok := m[r]; !ok || countLeft <= 0 {
			return false
		} else {
			m[r] -= 1
		}
	}

	return true
}
