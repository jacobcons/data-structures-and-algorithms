package heap

import (
	"testing"
)

func hasMinHeapProperty[T Ordered](h *Heap[T]) bool {
	for i := 1; i < len(h.Items); i++ {
		current := h.Items[i]

		leftChildIndex := i * 2
		if leftChildIndex < len(h.Items) {
			leftChild := h.Items[leftChildIndex]
			if current > leftChild {
				return false
			}
		}

		rightChildIndex := i*2 + 1
		if rightChildIndex < len(h.Items) {
			rightChild := h.Items[rightChildIndex]
			if current > rightChild {
				return false
			}
		}
	}

	return true
}

func TestAdd(t *testing.T) {
	h := NewHeap[int]()
	h.add(5)
	h.add(1)
	h.add(3)
	h.add(2)
	t.Log(h.Items)

	if hasMinHeapProperty(h) == false {
		t.Errorf("Adding some values doesn't maintain min-heap property")
	}
}
