package strings

// o(n^3) time | o(n) space
func LongestPalindromicSubstring(str string) string {
	longest := ""
	for i := range str {
		for j := i; j < len(str); j++ {
			substring := str[i : j+1]
			if len(substring) > len(longest) && isPalindrom(substring) {
				longest = substring
			}
		}
	}

	return longest
}

func isPalindrom(str string) bool {
	for i := range str {
		j := len(str) - i - 1
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
