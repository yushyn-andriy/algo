package linkedlist

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	dl := NewDoublyLinkedList()

	head := &Node{Value: 1}
	dl.SetHead(head)
	dl.InsertAtPosition(2, &Node{Value: 2})
	dl.InsertAtPosition(3, &Node{Value: 3})
	dl.InsertAtPosition(1, &Node{Value: 4})
	dl.Remove(dl.Head)

	t.Errorf("%v", dl.ContainsNodeWithValue(1))
	t.Errorf("%v", dl.ContainsNodeWithValue(122))

	dl.Traverse(func(value int) {
		fmt.Println(value)
	})
	t.Errorf("___")

}
