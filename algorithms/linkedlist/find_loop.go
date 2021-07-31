package linkedlist

type LinkedListLoop struct {
	Value   int
	Next    *LinkedListLoop
	visited bool
}

// o(n) time | o(1) space
func FindLoop(head *LinkedListLoop) *LinkedListLoop {
	for node := head; node != nil; node = node.Next {
		if node.visited {
			return node
		}
		node.visited = true
	}
	return nil
}

// o(n) time | o(1) space
// T = 2d + p
// F -> x -> d + p
// S -> 2x -> 2d + 2p
// R = t  - p - d
// r = 2d + p - p - d = d
func FindLoopAlgoSolution(head *LinkedListLoop) *LinkedListLoop {
	first := head.Next
	second := first.Next
	for first != second {
		first = first.Next
		second = second.Next
	}
	first = head
	for first != second {
		first = first.Next
		second = second.Next
	}
	return first
}
