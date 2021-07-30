package linkedlist

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	nodesCount := 0
	for node := head; node != nil; node = node.Next {
		nodesCount += 1
	}

	position := nodesCount - k
	if position == 0 {
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return
	}
	for node := head; node != nil; node = node.Next {
		if position == 1 {
			nodeToDelete := node.Next
			node.Next = node.Next.Next
			nodeToDelete.Next = nil
			break
		}
		position--
	}
}
