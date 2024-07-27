package main

import "fmt"

type bst struct {
	value int
	left  *bst
	right *bst
}

func newBst(value int) *bst {
	return &bst{value, nil, nil}
}

func (t *bst) insert(value int) *bst {
	if t == nil {
		return newBst(value)
	}
	if value < t.value {
		t.left = t.left.insert(value)
	} else if value > t.value {
		t.right = t.right.insert(value)
	}
	return t
}

func (t *bst) search(value int) bool {
	if t == nil {
		return false
	}
	if value < t.value {
		return t.left.search(value)
	} else if value > t.value {
		return t.right.search(value)
	}
	return true
}

func (t *bst) preOrder() {
	if t == nil {
		return
	}
	fmt.Println(t.value)
	t.left.preOrder()
	t.right.preOrder()
}

func main() {
	t := newBst(3)
	t.preOrder()
	t.insert(2)
	t.insert(5)
	t.insert(1)
	t.insert(4)
	t.insert(6)

	fmt.Println()
	t.preOrder()

	fmt.Println(t.search(2))
	fmt.Println(t.search(9))
}
