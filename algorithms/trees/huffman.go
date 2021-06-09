package trees

import (
	"container/heap"
	"fmt"
)

type HuffmanNode struct {
	Key    int
	Value  string
	Freq   float64
	index  int // The index of the item in the heap.
	Parent *HuffmanNode
	Left   *HuffmanNode
	Right  *HuffmanNode
}

func (h *HuffmanNode) String() string {
	return fmt.Sprintf("HuffmanNode(Value='%v', Freq='%v')", h.Value, h.Freq)
}

// InOrderWalk ...
func HuffmanInOrderWalk(tree *HuffmanNode, f func(tree *HuffmanNode)) {
	if tree == nil {
		return
	}
	HuffmanInOrderWalk(tree.Left, f)
	f(tree)
	HuffmanInOrderWalk(tree.Right, f)
}

// InOrderWalk ...
func huffmanInOrderWalkToGetCodes(tree *HuffmanNode, s string, codes map[string]string) {
	if tree == nil {
		return
	}

	if tree.Left == nil && tree.Right == nil {
		codes[tree.Value] = s
	}
	huffmanInOrderWalkToGetCodes(tree.Left, s+"0", codes)
	huffmanInOrderWalkToGetCodes(tree.Right, s+"1", codes)
}

func HuffmanCodesFromTree(node *HuffmanNode) map[string]string {
	m := make(map[string]string)
	huffmanInOrderWalkToGetCodes(node, "", m)
	return m
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Freq < pq[j].Freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*HuffmanNode)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func BuildHuffmanTree(runesFreq map[rune]float64) *HuffmanNode {
	pq := make(PriorityQueue, len(runesFreq))

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	i := 0
	for value, priority := range runesFreq {
		pq[i] = &HuffmanNode{
			Value: string(value),
			Freq:  priority,
			index: i,
		}
		i++
	}
	heap.Init(&pq)

	n := pq.Len()
	for i := 0; i < n-1; i++ {
		newNode := &HuffmanNode{}
		x := heap.Pop(&pq).(*HuffmanNode)
		y := heap.Pop(&pq).(*HuffmanNode)
		newNode.Left = x
		newNode.Right = y
		newNode.Value = x.Value + y.Value
		newNode.Freq = x.Freq + y.Freq
		heap.Push(&pq, newNode)
	}

	return heap.Pop(&pq).(*HuffmanNode)
}
