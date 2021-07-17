package trees

func ValidateThreeNodes(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	if checkNext(nodeOne, nodeTwo, nodeThree) {
		return checkNext(nodeTwo, nodeThree, nodeOne)
	}
	if checkNext(nodeThree, nodeTwo, nodeOne) {
		return checkNext(nodeTwo, nodeOne, nodeThree)
	}

	return false
}

func checkNext(node, potentialNext, notNext *BST) bool {
	if node == nil {
		return false
	}
	if node.Value == notNext.Value {
		return false
	}
	if node.Value == potentialNext.Value {
		return true
	}
	if potentialNext.Value < node.Value {
		return checkNext(node.Left, potentialNext, notNext)
	}
	return checkNext(node.Right, potentialNext, notNext)
}
