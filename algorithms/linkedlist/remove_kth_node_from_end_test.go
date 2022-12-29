package linkedlist

import "testing"

func TestSumOfLinkedLists(t *testing.T) {
	linkedListOne := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 2,
			Next:  &LinkedList{Value: 3},
		},
	}
	linkedListTwo := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 2,
			Next:  &LinkedList{Value: 3},
		},
	}

	res := SumOfLinkedLists(linkedListOne, linkedListTwo)

	t.Errorf("%v", res)
}
