package trees

func FindKthLargestValueInBst(tree *BST, k int) int {
	return tree.kMax(k)
}

func (tree *BST) kMax(k int) int {
	sortedKLargest := make([]int, 0)
	tree.findKMax(&sortedKLargest, k)
	return sortedKLargest[k-1]
}

func (tree *BST) findKMax(sortedKLargest *[]int, k int) {
	if tree == nil || len(*sortedKLargest) == k {
		return
	}

	tree.Right.findKMax(sortedKLargest, k)
	if len(*sortedKLargest) <= k {
		*sortedKLargest = append(*sortedKLargest, tree.Value)
	}
	tree.Left.findKMax(sortedKLargest, k)
}
