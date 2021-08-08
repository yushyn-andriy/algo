package strings

type substring struct {
	left  int
	right int
}

func (ss substring) length() int { return ss.right - ss.left }

func LongestSubstringWithoutDuplication(str string) string {
	lastSeen := map[rune]int{}
	longest := substring{0, 1}

	startIndex := 0
	for i, ch := range str {
		if seenIndex, found := lastSeen[ch]; found && startIndex < seenIndex+1 {
			startIndex = seenIndex + 1
		}
		if longest.length() < i+1-startIndex {
			longest = substring{startIndex, i + 1}
		}
		lastSeen[ch] = i
	}
	return str[longest.left:longest.right]
}
