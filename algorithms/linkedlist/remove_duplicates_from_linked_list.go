package linkedlist

type LinkedListAlgo struct {
	Value int
	Next  *LinkedListAlgo
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedListAlgo) *LinkedListAlgo {
	root := linkedList
	for currNode := root; currNode.Next != nil; {
		if currNode.Value == currNode.Next.Value {
			currNode.DeleteNext()
		} else {
			currNode = currNode.Next
		}
	}
	return root
}

func (ll *LinkedListAlgo) DeleteNext() {
	if ll != nil {
		ll.Next = ll.Next.Next
	}
}
