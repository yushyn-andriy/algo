package linkedlist

func NodeSwap(head *LinkedList) *LinkedList {
	return recursiveNodeSwap(head)
}

func recursiveNodeSwap(head *LinkedList) *LinkedList {
	if head == nil || head.Next == nil {
		return head
	}

	nextNode := head.Next
	head.Next = recursiveNodeSwap(head.Next.Next)
	nextNode.Next = head
	return nextNode
}
