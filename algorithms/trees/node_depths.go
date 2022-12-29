package trees

import (
	"sort"
)

func NodeDepths(root *Tree) int {
	st := NewTreeStack()
	visited := make(map[*Tree]bool)
	depths := make(map[*Tree]int)

	st.PushTree(root)
	for {
		node, ok := st.PopTree()
		if !ok {
			break
		}

		if v, ok := visited[node]; ok && v {
			continue
		}
		visited[node] = true

		if node.Left != nil {
			st.PushTree(node.Left)
		}
		if node.Right != nil {
			st.PushTree(node.Right)
		}
		if node == root {
			depths[node] = 0
		} else {
			depths[node] = depths[node.Parent] + 1
		}
	}

	depthsSum := 0
	for _, value := range depths {
		depthsSum += value
	}

	return depthsSum
}

func NodeDepthsMax(root *Tree) int {
	st := NewTreeStack()
	visited := make(map[*Tree]bool)
	depths := make(map[*Tree]int)

	st.PushTree(root)
	for {
		node, ok := st.PopTree()
		if !ok {
			break
		}

		if v, ok := visited[node]; ok && v {
			continue
		}
		visited[node] = true

		if node.Left != nil {
			st.PushTree(node.Left)
		}
		if node.Right != nil {
			st.PushTree(node.Right)
		}
		if node == root {
			depths[node] = 0
		} else {
			depths[node] = depths[node.Parent] + 1
		}
	}

	depthsMax := 0
	for _, value := range depths {
		if value > depthsMax {
			depthsMax = value
		}
	}

	return depthsMax
}

// MinHeightBST
func MinHeightBST(x []int) *Tree {
	root := &Tree{Key: x[0]}
	if len(x) == 1 {
		return root
	}

	sort.Ints(x)
	arr := BuildArr(x)

	for _, key := range arr {
		Insert(root, &Tree{Key: key})
	}
	return root
}

func BuildArr(x []int) []int {
	if len(x) == 1 {
		return []int{x[0]}
	}
	if len(x) == 2 {
		return []int{x[0], x[1]}
	}
	if len(x) == 3 {
		return []int{x[1], x[0], x[2]}
	}

	getMiddle := func(left, right int) int {
		return (left + right) / 2
	}

	mid := getMiddle(0, len(x))
	arr := []int{x[mid]}

	arr = append(arr, BuildArr(x[0:mid])...)
	arr = append(arr, BuildArr(x[mid+1:])...)

	return arr
}
