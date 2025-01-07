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
		if h.Items[parentIndex] > h.Items[currentIndex] {
			h.Items[currentIndex], h.Items[parentIndex] = h.Items[parentIndex], h.Items[currentIndex]
		} else {
			break
		}
		currentIndex = parentIndex
	}
}

func (h *Heap[T]) RemoveMin() T {
	return h.Items[0]
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
