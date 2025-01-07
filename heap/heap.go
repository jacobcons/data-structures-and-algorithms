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

}
