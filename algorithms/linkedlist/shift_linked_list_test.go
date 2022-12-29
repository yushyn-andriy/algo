package linkedlist

import "testing"

func TestShiftLinkedList(t *testing.T) {
	ShiftLinkedList(&LinkedList{
		Value: 0,
		Next: &LinkedList{
			Value: 1,
			Next: &LinkedList{
				Value: 2,
				Next: &LinkedList{
					Value: 3,
					Next: &LinkedList{
						Value: 4,
						Next: &LinkedList{
							Value: 5,
						},
					},
				},
			},
		},
	}, 2)
}
