package linkedlist

import "testing"

func TestMergeLinkedLists(t *testing.T) {
	listOne := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 2,
			Next: &LinkedList{
				Value: 3,
			},
		},
	}
	listTwo := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 1,
			Next: &LinkedList{
				Value: 5,
			},
		},
	}
	MergeLinkedLists(listOne, listTwo)
}
