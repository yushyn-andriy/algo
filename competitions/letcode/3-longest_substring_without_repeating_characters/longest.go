package longestsubstring

// time O(n^2) | space O(n)
func LengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}

	var largest int = 1
	for i := 0; i < len(s)-1; i++ {
		m := make(map[byte]bool)
		m[s[i]] = true
		current := 1
		for j := i + 1; j < len(s); j++ {
			if _, exists := m[s[j]]; exists {
				break
			} else {
				m[s[j]] = true
				current += 1
			}
		}
		largest = max(largest, current)
	}

	return largest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type substring struct {
	left, right int
}

func (s *substring) length() int { return s.right - s.left }

// O(n) time | O(1) space
func lengthOfLongestSubstring(s string) int {
	lastSeen := map[byte]uint8{}
	longest := substring{0, 0}

	var startIndex int = 0
	for i := range s {
		if seenIndex, found := lastSeen[s[i]]; found && startIndex < int(seenIndex)+1 {
			startIndex = int(seenIndex) + 1
		}
		if longest.length() < i+1-startIndex {
			longest = substring{startIndex, i + 1}
		}
		lastSeen[s[i]] = uint8(i)
	}

	return longest.length()
}
