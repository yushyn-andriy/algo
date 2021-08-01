package linkedlist

import "testing"

func TestReverseLinkedList(t *testing.T) {
	linkedListOne := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 2,
			Next: &LinkedList{
				Value: 3,
				Next: &LinkedList{
					Value: 4,
					Next:  &LinkedList{Value: 5},
				},
			},
		},
	}

	ReverseLinkedList(linkedListOne)
}
