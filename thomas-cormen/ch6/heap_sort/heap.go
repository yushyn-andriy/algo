package heapsort

func Left(i int) int {
	return i * 2
}

func Right(i int) int {
	return i*2 + 1
}

func Parent(i int) int {
	return i / 2
}

type Heap []int

func (h Heap) Size() int {
	return len(h) - 1
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func NewHeap(arr []int) *Heap {
	h := Heap(arr)
	return &h
}

func MaxHeapify(h Heap, i int) {
	left, right := Left(i), Right(i)
	largest := i
	if left < h.Size() && h[left] > h[i] {
		largest = left
	}
	if right < h.Size() && h[right] > h[largest] {
		largest = right
	}
	if largest != i {
		h.Swap(largest, i)
		MaxHeapify(h, largest)
	}
}

func BuildMaxHeap(h Heap) {
	for i := h.Size() / 2; i >= 1; i-- {
		MaxHeapify(h, i)
	}
}
