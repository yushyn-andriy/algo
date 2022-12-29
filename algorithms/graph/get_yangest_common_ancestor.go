package graph

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	depthOne, depthTwo := getDescendantDepth(descendantOne, top), getDescendantDepth(descendantTwo, top)
	if depthOne > depthTwo {
		return backtrackAncestralTree(descendantOne, descendantTwo, depthOne-depthTwo)
	}
	return backtrackAncestralTree(descendantTwo, descendantOne, depthTwo-depthOne)
}

func getDescendantDepth(descendant, topAncestor *AncestralTree) int {
	depth := 0
	for descendant != topAncestor {
		depth += 1
		descendant = descendant.Ancestor
	}
	return depth
}

func backtrackAncestralTree(lowerDescendant, higherDescendant *AncestralTree, diff int) *AncestralTree {
	for diff > 0 {
		lowerDescendant = lowerDescendant.Ancestor
		diff--
	}
	for lowerDescendant != higherDescendant {
		lowerDescendant = lowerDescendant.Ancestor
		higherDescendant = higherDescendant.Ancestor
	}
	return lowerDescendant
}
