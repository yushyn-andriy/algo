package linkedlist

import (
	"strconv"
)

func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	firstNum, secondNum, multiplier := 0, 0, 1
	currFirstNode, currSecondNode := linkedListOne, linkedListTwo
	for {
		if currFirstNode == nil && currSecondNode == nil {
			break
		}
		if currFirstNode != nil {
			firstNum += currFirstNode.Value * multiplier
			currFirstNode = currFirstNode.Next
		}
		if currSecondNode != nil {
			secondNum += currSecondNode.Value * multiplier
			currSecondNode = currSecondNode.Next
		}
		multiplier *= 10
	}

	var head *LinkedList
	sum := strconv.Itoa(firstNum + secondNum)
	for _, r := range sum {
		value, _ := strconv.Atoi(string(r))
		newHead := &LinkedList{Value: value, Next: head}
		head = newHead
	}

	return head
}
