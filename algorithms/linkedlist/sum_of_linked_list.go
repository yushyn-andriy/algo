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

func SumOfLinkedLists2(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	newLinkedListHeadPointer := &LinkedList{Value: 0}
	currentNode := newLinkedListHeadPointer

	carry := 0
	currFirstNode, currSecondNode := linkedListOne, linkedListTwo
	for currFirstNode != nil || currSecondNode != nil || carry != 0 {
		var valueOne, valueTwo int
		if currFirstNode != nil {
			valueOne = currFirstNode.Value
		}
		if currSecondNode != nil {
			valueTwo = currSecondNode.Value
		}

		sumOfValues := valueOne + valueTwo + carry
		newValue := sumOfValues % 10
		newNode := &LinkedList{Value: newValue}
		currentNode.Next = newNode
		currentNode = newNode

		carry = sumOfValues / 10
		if currFirstNode != nil {
			currFirstNode = currFirstNode.Next
		}
		if currSecondNode != nil {
			currSecondNode = currSecondNode.Next
		}
	}

	return newLinkedListHeadPointer.Next
}
