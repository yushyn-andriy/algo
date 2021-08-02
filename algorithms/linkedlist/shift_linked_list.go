package linkedlist

func ShiftLinkedList(head *LinkedList, k int) *LinkedList {
	// 1 detect list length
	listLength := 1
	listTail := head
	for listTail.Next != nil {
		listTail = listTail.Next
		listLength += 1
	}

	// because circullar
	offset := abs(k) % listLength
	if offset == 0 {
		return head
	}

	newTailPosition := listLength - offset
	if k <= 0 {
		newTailPosition = offset
	}

	newTail := head
	for i := 1; i < newTailPosition; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil
	listTail.Next = head

	return newHead
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
