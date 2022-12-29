package stack

type Stack interface {
	Push(value int)
	Pop() *int
	Size() int
	Peek() *int
}

type stack struct {
	items []int
}

func New() Stack {
	return &stack{}
}

func (s *stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *stack) Pop() *int {
	if len(s.items) > 0 {
		value := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return &value
	}
	return nil
}

func (s *stack) Peek() *int {
	if len(s.items) > 0 {
		return &s.items[len(s.items)-1]
	}
	return nil
}

func (s *stack) Size() int {
	return len(s.items)
}
