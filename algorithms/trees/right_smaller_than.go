package trees

// O(n^2) time | O(n) space
func RightSmallerThan1(array []int) []int {
	smaller := make([]int, len(array))
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				smaller[i] += 1
			}
		}
	}

	return smaller
}

func RightSmallerThan2(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	smaller := make([]int, len(array))
	bst := &BSTree{Value: array[len(array)-1], IndexInArray: len(array) - 1}
	for i := len(array) - 2; i >= 0; i-- {
		bst.Insert(array[i], 0, i)
	}

	return bst.SmallerCounterInOrderTraverse(smaller)
}

type BSTree struct {
	Value          int
	Left           *BSTree
	Right          *BSTree
	LeftCounter    int
	SmallerCounter int
	IndexInArray   int
}

func (tree *BSTree) Insert(value int, smallerCount, indexInArray int) {
	if tree.Value > value {
		tree.LeftCounter += 1
		if tree.Left == nil {
			tree.Left = &BSTree{
				Value:          value,
				SmallerCounter: smallerCount,
				IndexInArray:   indexInArray,
			}
		} else {
			tree.Left.Insert(value, smallerCount, indexInArray)
		}
	} else {
		additionalNodeSmallerCount := 0
		if tree.Value < value {
			additionalNodeSmallerCount = 1
		}
		if tree.Right == nil {

			tree.Right = &BSTree{
				Value:          value,
				SmallerCounter: tree.LeftCounter + smallerCount + additionalNodeSmallerCount,
				IndexInArray:   indexInArray,
			}
		} else {
			tree.Right.Insert(value, smallerCount+tree.LeftCounter+additionalNodeSmallerCount, indexInArray)
		}
	}
}

func (tree *BSTree) SmallerCounterInOrderTraverse(array []int) []int {
	smallerCounterTraverse(tree, &array)
	return array
}

func smallerCounterTraverse(tree *BSTree, array *[]int) {
	if tree == nil {
		return
	}

	(*array)[tree.IndexInArray] = tree.SmallerCounter
	smallerCounterTraverse(tree.Left, array)
	smallerCounterTraverse(tree.Right, array)
}
