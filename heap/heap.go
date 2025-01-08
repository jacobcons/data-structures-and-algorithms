package heap

import (
	"cmp"
)

type Ordered = cmp.Ordered

type Heap[T Ordered] struct {
	Items []T
}

func NewHeap[T Ordered]() *Heap[T] {
	var zeroValue T
	return &Heap[T]{[]T{zeroValue}}
}

func (h *Heap[T]) Add(item T) {
	h.Items = append(h.Items, item)
	currentIndex := len(h.Items) - 1
	var parentIndex int
	for currentIndex != 1 {
		parentIndex = currentIndex / 2
		if h.Items[parentIndex] <= h.Items[currentIndex] {
			break
		}

		h.Items[currentIndex], h.Items[parentIndex] = h.Items[parentIndex], h.Items[currentIndex]
		currentIndex = parentIndex
	}
}

func (h *Heap[T]) RemoveMin() (T, bool) {
	if h.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	// move rightmost leaf to root
	lastItemIndex := len(h.Items) - 1
	minItem := h.Items[1]
	h.Items[1] = h.Items[lastItemIndex]
	h.Items = h.Items[:lastItemIndex]

	// sink new root
	currentIndex := 1
	// stop when there's no left child i.e. sunk all the way to leaf node
	for currentIndex*2 <= len(h.Items)-1 {
		leftChildIndex := currentIndex * 2
		rightChildIndex := leftChildIndex + 1
		isRightChild := false
		if rightChildIndex < len(h.Items) {
			isRightChild = true
		}
		// stop if current node is less than or equal to both children
		if (h.Items[currentIndex] <= h.Items[leftChildIndex] && isRightChild == false) || (h.Items[currentIndex] <= h.Items[leftChildIndex] && h.Items[currentIndex] <= h.Items[rightChildIndex]) {
			break
		}

		// get index of smaller child
		var minItemIndex int
		if isRightChild == false || (h.Items[leftChildIndex] < h.Items[rightChildIndex]) {
			minItemIndex = leftChildIndex
		} else {
			minItemIndex = rightChildIndex
		}
		// swap current node with smaller child
		h.Items[currentIndex], h.Items[minItemIndex] = h.Items[minItemIndex], h.Items[currentIndex]

		currentIndex = minItemIndex
	}

	return minItem, true
}

func (h *Heap[T]) GetMin() (T, bool) {
	if h.Size() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return h.Items[1], true
}

func (h *Heap[T]) Size() int {
	return len(h.Items) - 1
}
