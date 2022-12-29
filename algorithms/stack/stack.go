package stack

type Stack struct {
	data []int
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, true
}

func New() *Stack {
	return &Stack{}
}

type MinMaxStack struct {
	data                 map[int]int
	currentIdx, min, max int
	minMaxStack          []entry
}

type entry struct {
	min, max int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{
		data:        map[int]int{},
		currentIdx:  0,
		minMaxStack: []entry{},
	}
}

func (stack *MinMaxStack) Peek() int {
	v, _ := stack.data[stack.currentIdx-1]
	return v
}

func (stack *MinMaxStack) Pop() int {
	stack.minMaxStack = stack.minMaxStack[:len(stack.minMaxStack)-1]
	stack.currentIdx--
	v, found := stack.data[stack.currentIdx]
	if found {
		delete(stack.data, stack.currentIdx)
	}
	return v
}

func (stack *MinMaxStack) Push(number int) {
	newMinMax := entry{min: number, max: number}
	if len(stack.minMaxStack) > 0 {
		lastMinMax := stack.minMaxStack[len(stack.minMaxStack)-1]
		newMinMax.min = min(lastMinMax.min, number)
		newMinMax.max = max(lastMinMax.max, number)
	}
	stack.minMaxStack = append(stack.minMaxStack, newMinMax)
	stack.data[stack.currentIdx] = number
	stack.currentIdx += 1
}

func (stack *MinMaxStack) GetMin() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].min
}

func (stack *MinMaxStack) GetMax() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
