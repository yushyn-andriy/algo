package linkedlist


func ReverseLinkedList(head *LinkedList) *LinkedList {
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
