package strings

type stack struct {
	array []rune
}

func LongestBalancedSubstringCubic(str string) int {
	longestMax := 0
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j <= len(str); j++ {
			if IsBalanced(str[i:j]) {
				currentMax := j - i
				longestMax = max(currentMax, longestMax)
			}
		}
	}

	return longestMax
}

func (s *stack) Length() int {
	return len(s.array)
}

func (s *stack) Push(element rune) {
	s.array = append(s.array, element)
}

func (s *stack) Pop() (rune, bool) {
	if s.Length() <= 0 {
		return 0, false
	}
	last := s.array[s.Length()-1]
	s.array = s.array[:s.Length()-1]
	return last, true
}

func IsBalanced(str string) bool {
	s := stack{}
	for _, r := range str {
		if r == '(' {
			s.Push(r)
		} else {
			if _, ok := s.Pop(); !ok {
				return false
			}
		}
	}

	return s.Length() == 0
}
