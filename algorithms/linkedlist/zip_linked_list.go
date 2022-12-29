package linkedlist

func ZipLinkedList(linkedList *LinkedList) *LinkedList {
	if linkedList.Next == nil || linkedList.Next.Next == nil {
		return linkedList
	}

	firstHalfHead := linkedList
	secondHalfHead := splitLinkedList(linkedList)
	reversedSecondHalfHead := reverseLinkedList(secondHalfHead)

	return interweaveLinkedLists(firstHalfHead, reversedSecondHalfHead)
}

func splitLinkedList(linkedList *LinkedList) *LinkedList {
	slowIterator, fastIterator := linkedList, linkedList
	for fastIterator != nil && fastIterator.Next != nil {
		slowIterator, fastIterator = slowIterator.Next, fastIterator.Next.Next
	}
	secondHalfHead := slowIterator.Next
	slowIterator.Next = nil
	return secondHalfHead
}

func reverseLinkedList(head *LinkedList) *LinkedList {
	if head.Next == nil {
		return head
	}
	var first, second, third *LinkedList = nil, head, head.Next
	for {
		second.Next = first
		first = second
		second = third
		third = third.Next
		if third == nil {
			second.Next = first
			return second
		}
	}
}

/*
func reverseLinkedList(head *LinkedList) *LinkedList {
	var previousNode *LinkedList
	currentNode := head
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}*/

func interweaveLinkedLists(linkedList1, linkedList2 *LinkedList) *LinkedList {
	linkedList1Iterator, linkedList2Iterator := linkedList1, linkedList2
	for linkedList1Iterator != nil && linkedList2Iterator != nil {
		linkedList1IteratorNext := linkedList1Iterator.Next
		linkedList2IteratorNext := linkedList2Iterator.Next

		linkedList1Iterator.Next = linkedList2Iterator
		linkedList2Iterator.Next = linkedList1IteratorNext

		linkedList1Iterator = linkedList1IteratorNext
		linkedList2Iterator = linkedList2IteratorNext
	}

	return linkedList1
}
