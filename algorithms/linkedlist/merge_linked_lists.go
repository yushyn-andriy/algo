package linkedlist

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	var head  = &LinkedList{}
	currHead := head
	for {
		if headOne == nil && headTwo == nil {
			break
		}

		if headOne != nil && headTwo != nil {
			if headOne.Value < headTwo.Value {
				currHead =currHead.InsertAndGetNext(&LinkedList{Value: headTwo.Value})
				headOne = headOne.Next
			} else if headOne.Value > headTwo.Value {
				currHead = currHead.InsertAndGetNext(&LinkedList{Value: headTwo.Value})
				headTwo = headTwo.Next
			} else {
				currHead = currHead.InsertAndGetNext(&LinkedList{Value: headOne.Value})
				currHead = currHead.InsertAndGetNext(&LinkedList{Value: headOne.Value})
				headOne = headOne.Next
				headTwo = headTwo.Next
			}
		} else if headOne == nil {
			for ; headOne != nil; headOne = headOne.Next {
				currHead = currHead.InsertAndGetNext(&LinkedList{Value: headOne.Value})
			}
		} else if headTwo == nil {
			for ; headTwo != nil; headTwo = headTwo.Next {
				currHead.InsertAndGetNext(&LinkedList{Value: headTwo.Value})
				currHead = currHead.Next
			}
		}
	}

	return head.Next
}

func (ll *LinkedList) InsertAndGetNext(node *LinkedList) *LinkedList {
	ll.Next = node
	return ll.Next
}
