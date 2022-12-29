package palindrom

func IsPalindrom(s string) bool {
	for i := 0; i < len(s); i++ {
		j := len(s) - i - 1
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func getLongestPalindrom(s string, startIdx, endIdx int) substring {
	for startIdx >= 0 && endIdx < len(s) {
		if s[startIdx] != s[endIdx] {
			break
		}
		startIdx -= 1
		endIdx += 1
	}
	return substring{startIdx + 1, endIdx}
}

type substring struct {
	left, rigth int
}

func (s *substring) length() int {
	return s.rigth - s.left
}

// O(n^2) time | O(1) space
func longestPalindrome1(s string) string {
	result := substring{0, 0}
	for i := 0; i < len(s); i++ {
		odd := getLongestPalindrom(s, i-1, i+1)
		even := getLongestPalindrom(s, i-1, i)
		longest := even
		if odd.length() > even.length() {
			longest = odd
		}
		if longest.length() > result.length() {
			result = longest
		}
	}
	return s[result.left:result.rigth]
}

// O(n^3) time | O(1) space
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	longest := string(s[0])
	for i := 0; i < len(s)-1; i++ {
		for j := i; j <= len(s); j++ {
			if IsPalindrom(s[i:j]) && len(s[i:j]) > len(longest) {
				longest = s[i:j]
			}
		}
	}
	return longest
}
