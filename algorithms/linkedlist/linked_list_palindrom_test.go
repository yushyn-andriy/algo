package linkedlist

import "testing"

func TestLinkedListPalindrome(t *testing.T) {
	tests := []struct {
		in  *LinkedList
		out bool
	}{
		{
			&LinkedList{
				Value: 1,
				Next: &LinkedList{
					Value: 2,
					Next: &LinkedList{
						Value: 3,
						Next: &LinkedList{
							Value: 4,
						},
					},
				},
			},
			false,
		},
		{
			&LinkedList{
				Value: 1,
				Next: &LinkedList{
					Value: 2,
					Next: &LinkedList{
						Value: 2,
						Next: &LinkedList{
							Value: 1,
						},
					},
				},
			},
			true,
		},
		{
			&LinkedList{
				Value: 1,
			},
			true,
		},
		{
			&LinkedList{
				Value: 1,
				Next: &LinkedList{
					Value: 2,
					Next: &LinkedList{
						Value: 2,
						Next: &LinkedList{
							Value: 2,
						},
					},
				},
			},
			false,
		},
		{
			&LinkedList{
				Value: 1,
				Next: &LinkedList{
					Value: 2,
				},
			},
			false,
		},
	}

	for i, test := range tests {
		if LinkedListPalindrome(test.in) != test.out {
			t.Errorf("test(%d) expected %v got %v", i, test.out, !test.out)
		}
	}

}
