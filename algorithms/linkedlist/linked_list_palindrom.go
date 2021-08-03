package linkedlist

func LinkedListPalindrome(head *LinkedList) bool {
	isPalindrom, _ := linkedListPalindrome(head, head.Next)
	return isPalindrom
}

func linkedListPalindrome(left, right *LinkedList) (bool, *LinkedList) {
	if right == nil {
		return true, left
	}
	isPalindrom, nextLeft := linkedListPalindrome(left, right.Next)
	if nextLeft.Value != right.Value {
		return false, nextLeft.Next
	}
	return isPalindrom, nextLeft.Next
}
