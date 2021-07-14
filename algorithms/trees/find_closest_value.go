package trees

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {

	return findClosestValue(tree, target, math.MaxInt32)
}

func findClosestValue(node *BST, target int, closest int) int {
	if node == nil {
		return closest
	}

	if abs(target-node.Value) < abs(target-closest) {
		closest = node.Value
	}
	if target >= node.Value {
		return findClosestValue(node.Right, target, closest)
	} else {
		return findClosestValue(node.Left, target, closest)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
