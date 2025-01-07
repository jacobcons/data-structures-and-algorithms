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

func makePopulatedHeap() *Heap[int] {
	h := NewHeap[int]()
	h.Add(5)
	h.Add(1)
	h.Add(3)
	h.Add(2)
	return h
}

func TestAdd(t *testing.T) {
	h := makePopulatedHeap()

	if hasMinHeapProperty(h) == false {
		t.Errorf("Adding some values doesn't maintain min-heap property")
	}

	if h.Size() != 4 {
		t.Errorf("Size must be 4 after adding 4 values")
	}
}

func TestRemoveMin(t *testing.T) {
	h := makePopulatedHeap()
	expectedMinimums := []int{1, 2, 3, 5}
	expectedSize := len(expectedMinimums)
	for _, expectedMinimum := range expectedMinimums {
		actualMinimum := h.RemoveMin()
		expectedSize--
		if actualMinimum != expectedMinimum {
			t.Errorf("actualMinimum=%v, expectedMinimum=%v", actualMinimum, expectedMinimum)
		}

		if h.Size() != expectedSize {
			t.Errorf("h.Size()=%v, expectedSize=%v", h.Size(), expectedSize)
		}

		if hasMinHeapProperty(h) == false {
			t.Errorf("removing minimums doesn't maintain min-heap property")
		}
	}
}

func TestGetMin(t *testing.T) {
	h := makePopulatedHeap()
	initialSize := h.Size()
	minimum, _ := h.GetMin()
	if minimum != 1 {
		t.Errorf("minimum=%v, expectedMinimum=1", minimum)
	}
	if initialSize != h.Size() {
		t.Errorf("size should remain same after calling GetMin()")
	}

	h.RemoveMin()
	_, success := h.GetMin()
	if success == false {
		t.Errorf("expected GetMin() to not be successfull when heap is empty")
	}
}
