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
