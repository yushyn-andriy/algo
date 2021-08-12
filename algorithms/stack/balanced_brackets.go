package stack

type stack struct {
	array []rune
}

func BalancedBrackets(s string) bool {
	return IsBalanced(s)
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
		if r == '(' || r == '[' || r == '{' {
			s.Push(r)
		} else if r == ')' || r == ']' || r == '}' {
			last, ok := s.Pop()
			if !ok {
				return false
			}
			if r == ')' && last != '(' {
				return false
			}
			if r == ']' && last != '[' {
				return false
			}
			if r == '}' && last != '{' {
				return false
			}
		}
	}

	return s.Length() == 0
}
