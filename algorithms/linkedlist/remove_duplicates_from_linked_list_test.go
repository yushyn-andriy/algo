package linkedlist

import "testing"

func TestRemoveDuplicatesFromLinkedList(t *testing.T) {
	tests := []struct {
		in  *LinkedListAlgo
		out *LinkedListAlgo
	}{
		{
			&LinkedListAlgo{
				Value: 1,
				Next: &LinkedListAlgo{
					Value: 1,
					Next: &LinkedListAlgo{
						Value: 1,
					},
				},
			},
			&LinkedListAlgo{
				Value: 1,
				Next:  nil,
			},
		},
	}

	for i, test := range tests {
		out := RemoveDuplicatesFromLinkedList(test.in)
		if test.out.Next != nil {
			t.Errorf("test(%d) expected %v got %v", i, test.out, out)
		}
	}
}
