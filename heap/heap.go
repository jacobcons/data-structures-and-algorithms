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

func (h *Heap[T]) add(item T) {
	h.Items = append(h.Items, item)
	currentIndex := len(h.Items) - 1
	var parentIndex int
	for true {
		// if reached root
		if currentIndex == 1 {
			break
		}

		parentIndex = currentIndex / 2
		if h.Items[parentIndex] > h.Items[currentIndex] {
			h.Items[currentIndex], h.Items[parentIndex] = h.Items[parentIndex], h.Items[currentIndex]
		} else {
			break
		}
		currentIndex = parentIndex
	}
}
